drpcli system
-------------

Access CLI commands relating to system

Synopsis
~~~~~~~~

Access CLI commands relating to system

Options
~~~~~~~

::

     -h, --help   help for system

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

     -c, --catalog string          The catalog file to use to get product information (default "https://repo.rackn.io")
     -d, --debug                   Whether the CLI should run in debug mode
     -D, --download-proxy string   HTTP Proxy to use for downloading catalog and content
     -E, --endpoint string         The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
     -f, --force                   When needed, attempt to force the operation - used on some update/patch calls
     -F, --format string           The serialization we expect for output.  Can be "json" or "yaml" or "text" or "table" (default "json")
     -H, --no-header               Should header be shown in "text" or "table" mode
     -x, --noToken                 Do not use token auth or token cache
     -P, --password string         password of the Digital Rebar Provision user (default "r0cketsk8ts")
     -J, --print-fields string     The fields of the object to display in "text" or "table" mode. Comma separated
     -r, --ref string              A reference object for update commands that can be a file name, yaml, or json blob
     -T, --token string            token of the Digital Rebar Provision access
     -t, --trace string            The log level API requests should be logged at on the server side
     -Z, --traceToken string       A token that individual traced requests should report in the server logs
     -j, --truncate-length int     Truncate columns at this length (default 40)
     -u, --url-proxy string        URL Proxy for passing actions through another DRP
     -U, --username string         Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
~~~~~~~~

-  `drpcli <drpcli.html>`__ - A CLI application for interacting with the
   DigitalRebar Provision API
-  `drpcli system action <drpcli_system_action.html>`__ - Display the
   action for this system
-  `drpcli system actions <drpcli_system_actions.html>`__ - Display
   actions for this system
-  `drpcli system active <drpcli_system_active.html>`__ - Switch DRP to
   HA Active State
-  `drpcli system certs <drpcli_system_certs.html>`__ - Access CLI
   commands to get and set the TLS cert the API uses
-  `drpcli system ha <drpcli_system_ha.html>`__ - Access CLI commands to
   get the state of high availability
-  `drpcli system passive <drpcli_system_passive.html>`__ - Switch DRP
   to HA Passive State
-  `drpcli system runaction <drpcli_system_runaction.html>`__ - Run
   action on object from plugin
-  `drpcli system signurl <drpcli_system_signurl.html>`__ - Generate a
   RackN Signed URL for download
-  `drpcli system upgrade <drpcli_system_upgrade.html>`__ - Upgrade DRP
   with the provided file
