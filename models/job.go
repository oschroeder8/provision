package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

// JobAction is something that job runner will need to do.
// If path is specified, then the runner will place the contents into that location.
// If path is not specified, then the runner will attempt to bash exec the contents.
//
// JobAction is generated by rendering a Template on a Task against a specific Machine.
// swagger:model
type JobAction struct {
	// Name is the name of this particular JobAction.  It is taken from
	// the name of the corresponding Template on the Task this Action was rendered from.
	// required: true
	Name string
	// Path is the location that Content should be written to on disk.
	//
	// If Path is absolute, it will be written in that location.
	//
	// If Path is relative, it will be written relative to the temporary direcory created
	// for running the Job in.
	//
	// If Path is empty, then Content is interpreted as a script to be run.
	// required: true
	Path string
	// Content is the rendered version of the Template on the Task corresponding to this
	// JobAction.
	// required: true
	Content string
	// Meta is a copt of the Meta field of the corresponding Template from the Task this Job was built from.
	// required: true
	Meta map[string]string
}

func (ja *JobAction) ValidForOS(target string) bool {
	if oses, ok := ja.Meta["OS"]; !ok || target == "" {
		// if the JobAction does not have an OS, it is valid for all of them.
		// Ditto if we are not actually looking for a target OS.
		return true
	} else {
		canPerform := false
		for _, v := range strings.Split(oses, ",") {
			testOS := strings.TrimSpace(v)
			if testOS == "any" || testOS == target {
				canPerform = true
				break
			}
		}
		return canPerform
	}
}

type JobActions []*JobAction

func (ja JobActions) FilterOS(forOS string) JobActions {
	if len(ja) == 0 {
		return ja
	}
	if _, ok := ja[0].Meta["OS"]; !ok {
		return ja
	}
	res := JobActions{}
	for i := range ja {
		action := ja[i]
		if action.ValidForOS(forOS) {
			res = append(res, action)
		}
	}
	return res
}

// Job contains information on a Job that is running for a specific
// Task on a Machine.
//
// swagger:model
type Job struct {
	Validation
	Access
	Meta
	Owned
	Bundled
	// The UUID of the job.  The primary key.
	// required: true
	// swagger:strfmt uuid
	Uuid uuid.UUID `index:",key"`
	// The UUID of the previous job to run on this machine.
	// swagger:strfmt uuid
	Previous uuid.UUID
	// The machine the job was created for.  This field must be the UUID of the machine.
	// required: true
	// swagger:strfmt uuid
	Machine uuid.UUID
	// The task the job was created for.  This will be the name of the task.
	// read only: true
	Task string
	// The stage that the task was created in.
	// read only: true
	Stage string
	// Context is the context the job was created to run in.
	Context string
	// The state the job is in.  Must be one of "created", "running", "failed", "finished", "incomplete"
	// required: true
	State string
	// The final disposition of the job.
	// Can be one of "reboot","poweroff","stop", or "complete"
	// Other substates may be added as time goes on
	ExitState string
	// The time the job started running.
	StartTime time.Time
	// The time the job failed or finished.
	EndTime time.Time
	// Archived indicates whether the complete log for the job can be
	// retrieved via the API.  If Archived is true, then the log cannot
	// be retrieved.
	//
	// required: true
	Archived bool
	// Whether the job is the "current one" for the machine or if it has been superceded.
	//
	// required: true
	Current bool
	// The current index is the machine CurrentTask that created this job.
	//
	// required: true
	// read only: true
	CurrentIndex int
	// The next task index that should be run when this job finishes.  It is used
	// in conjunction with the machine CurrentTask to implement the server side of the
	// machine agent state machine.
	//
	// required: true
	// read only: true
	NextIndex int
	// The workflow that the task was created in.
	// read only: true
	Workflow string
	// The bootenv that the task was created in.
	// read only: true
	BootEnv string
	// ExtraClaims is the expanded list of extra Claims that were added to the
	// default machine Claims via the ExtraRoles field on the Task that the Job
	// was created to run.
	ExtraClaims []*Claim `json:"ExtraClaims,omitempty"`
	// Token is the JWT token that should be used when running this Job.  If not
	// present or empty, the Agent running the Job will use its ambient Token
	// instead.  If set, the Token will only be valid for the current Job.
	Token string `json:"Token,omitempty"`
}

func (j *Job) GetMeta() Meta {
	return j.Meta
}

func (j *Job) SetMeta(d Meta) {
	j.Meta = d
}

func (j *Job) Validate() {
	if !strings.Contains(j.Task, ":") {
		j.AddError(ValidName("Invalid Task", j.Task))
	}
	j.AddError(ValidName("Invalid Stage", j.Stage))
	switch j.State {
	case "created", "running", "incomplete":
	case "failed", "finished":
	default:
		j.AddError(fmt.Errorf("Invalid State `%s`", j.State))
	}
	if j.ExitState != "" {
		switch j.ExitState {
		case "reboot", "poweroff", "stop", "complete", "failed":
		default:
			j.AddError(fmt.Errorf("Invalid ExitState `%s`", j.ExitState))
		}
	}
}

func (j *Job) Prefix() string {
	return "jobs"
}

func (j *Job) Key() string {
	return j.Uuid.String()
}

func (j *Job) KeyName() string {
	return "Uuid"
}

func (j *Job) Fill() {
	if j.Meta == nil {
		j.Meta = Meta{}
	}
	j.Validation.fill(j)
}

func (j *Job) AuthKey() string {
	return j.Machine.String()
}

func (b *Job) SliceOf() interface{} {
	s := []*Job{}
	return &s
}

func (b *Job) ToModels(obj interface{}) []Model {
	items := obj.(*[]*Job)
	res := make([]Model, len(*items))
	for i, item := range *items {
		res[i] = Model(item)
	}
	return res
}

func (b *Job) CanHaveActions() bool {
	return true
}
