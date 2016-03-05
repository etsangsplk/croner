package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goadesign/goa"
	"github.com/goadesign/middleware"
	"github.com/rightscale/croner/app"
	"github.com/rightscale/croner/cron"
	"github.com/rightscale/croner/swagger"
)

func main() {

	var (
		help     = flag.Bool("h", false, "display usage")
		port     = flag.String("p", "18080", "bind service to a specific port, set to 0 to not open HTTP port at all")
		schedule = flag.String("s", "* * * * *", "schedule the task the cron style")
	)

	flagArgs, execArgs := splitArgs()
	os.Args = flagArgs

	flag.Parse()

	if len(execArgs) == 0 || *help {
		println("Usage of ", os.Args[0])
		println(os.Args[0], " [ OPTIONS ] -- [ COMMAND ]")
		flag.PrintDefaults()
		println(`\nExample:\n%s -p 8080 -s "0 0 * * *" -- echo hello`, os.Args[0])
		os.Exit(1)
	}

	job, err := cron.NewJob(execArgs[0], execArgs[1:len(execArgs)], *schedule)
	if err != nil {
		log.Fatalf("Failed to start job: %s", err)
		os.Exit(1)
	}
	log.Println(job.String())
	if *port != "0" {
		startService(*port, job)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	println(<-ch)
	job.Stop()
}

func startService(port string, job *cron.Job) {
	service := goa.New("croner")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())
	c := NewJobController(service, job)
	app.MountJobController(service, c)
	swagger.MountController(service)
	service.ListenAndServe(":" + port)
}

func splitArgs() (flagArgs []string, execArgs []string) {
	split := len(os.Args)
	for idx, e := range os.Args {
		if e == "--" {
			split = idx
			break
		}
	}
	flagArgs = os.Args[0:split]
	if split < len(os.Args) {
		execArgs = os.Args[split+1 : len(os.Args)]
	} else {
		execArgs = []string{}
	}
	return flagArgs, execArgs
}
