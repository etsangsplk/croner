# croner

croner provides a simple cron like utility which runs a cron job and exposes a HTTP interface to
query the job status. croner should be run as a daemon. Example usage:

```
./croner -p 8080 -s "* * * * *" -- echo foo
```

The above runs the command `echo foo` using the cron schedule `"* * * * *"` and listens on port
8080. Once croner is running sending HTTP GET requests to the port it is listening to returns a
JSON document containing the information about the job:

```
http localhost:8080/job
HTTP/1.1 200 OK
Content-Length: 189
Content-Type: application/vnd.rightscale.croner.job+json
Date: Sat, 05 Mar 2016 00:49:00 GMT

{
    "cmd": "echo",
    "last": {
        "exit_status": 0,
        "finished_at": "2016-03-04T16:49:00.001280118-08:00",
        "pid": 6879,
        "started_at": "2016-03-04T16:49:00.000508062-08:00",
        "stderr": ""
    },
    "schedule": "* * * * *"
}
```
(note the `/job` path)
The complete API documentation is available on [swagger.goa.design](http://swagger.goa.design/?url=rightscale%2Fcroner%2Fdesign)
