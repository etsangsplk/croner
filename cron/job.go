package cron

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/rightscale/croner/app"
	"github.com/robfig/cron"
)

// Job captures information about a running or completed job.
type Job struct {
	*sync.Mutex
	Cmd      string
	Args     []string
	Schedule string
	Running  map[int]*Execution // Executions indexed by pid
	Last     *Execution

	wg   *sync.WaitGroup
	cron *cron.Cron
}

// Execution represents a job execution.
type Execution struct {
	Pid        int
	StartedAt  time.Time
	FinishedAt time.Time
	StdErr     string
	ExitStatus int
}

// Human readable job description.
func (job *Job) String() string {
	return fmt.Sprintf("%s %s %s", job.Schedule, job.Cmd, strings.Join(job.Args, " "))
}

// NewJob creates a job and schedules it.
func NewJob(cmd string, args []string, schedule string) (*Job, error) {
	job := &Job{
		Mutex:    new(sync.Mutex),
		Cmd:      cmd,
		Args:     args,
		Schedule: schedule,
		Running:  make(map[int]*Execution),

		wg:   &sync.WaitGroup{},
		cron: cron.New(),
	}
	c := cron.New()
	err := c.AddFunc(schedule, func() {
		job.wg.Add(1)
		job.Execute()
		job.wg.Done()
	})
	if err != nil {
		return nil, err
	}
	job.cron = c
	log.Println("Starting", job.String())
	c.Start()
	return job, nil
}

// Stop stops the job schedule.
func (job *Job) Stop() {
	job.Lock()
	defer job.Unlock()
	log.Println("Stopping")
	job.cron.Stop()
	log.Println("Waiting")
	job.wg.Wait()
	log.Println("Exiting")
}

// OsExit is to be nice to tests
var OsExit = os.Exit

// Execute job
func (job *Job) Execute() {
	cmd := exec.Command(job.Cmd, job.Args...)
	cmd.Stdout = ioutil.Discard
	exe := &Execution{StartedAt: time.Now()}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logFatalf(err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		logFatalf("cmd.Start: %v", err)
		return
	}
	exe.Pid = cmd.Process.Pid
	job.Lock()
	job.Running[exe.Pid] = exe
	job.Unlock()

	log.Println(exe.Pid, "cmd:", job.Cmd, strings.Join(job.Args, " "))

	b, err := ioutil.ReadAll(stderr)
	if err != nil {
		logFatalf("reading from stderr: %s", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			// so set the error code to tremporary value
			exe.ExitStatus = 127
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exe.ExitStatus = status.ExitStatus()
				log.Printf("%d Exit Status: %d", exe.Pid, exe.ExitStatus)
			}
		} else {
			logFatalf("cmd.Wait: %v", err)
			return
		}
	}
	exe.FinishedAt = time.Now()
	exe.StdErr = string(b)
	job.Lock()
	defer job.Unlock()
	delete(job.Running, exe.Pid)
	job.Last = exe
}

// ToMediaType converts the job into a media type
func (job *Job) ToMediaType() *app.Job {
	j := &app.Job{
		Cmd:      job.Cmd,
		Schedule: job.Schedule,
	}
	if job.Last != nil {
		j.Last = job.Last.ToMediaType()
	}
	job.Lock()
	defer job.Unlock()
	j.Running = make(app.ExecutionCollection, len(job.Running))
	i := 0
	for _, e := range job.Running {
		j.Running[i] = e.ToMediaType()
		i++
	}
	return j
}

// ToMediaType converts the execution into a media type
func (exe *Execution) ToMediaType() *app.Execution {
	e := &app.Execution{
		Pid:       exe.Pid,
		StartedAt: exe.StartedAt,
	}
	var zeroTime time.Time
	if exe.FinishedAt != zeroTime {
		e.FinishedAt = &exe.FinishedAt
		e.ExitStatus = &exe.ExitStatus
		e.Stderr = &exe.StdErr
	}
	return e
}

// logFatalf is equivalent to log.Fatalf but uses OsExit instead of os.Exit
func logFatalf(format string, v ...interface{}) {
	log.Printf(format, v...)
	OsExit(1)
}
