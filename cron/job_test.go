package cron_test

import (
	"bytes"
	"io/ioutil"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/croner/cron"
)

var _ = Describe("NewJob", func() {
	var cmd, schedule string
	var args []string

	var job *cron.Job
	var err error

	JustBeforeEach(func() {
		log.SetOutput(ioutil.Discard)
		job, err = cron.NewJob(cmd, args, schedule)
	})

	AfterEach(func() {
		if job != nil {
			job.Stop()
		}
	})

	Context("created with valid parameters", func() {
		BeforeEach(func() {
			cmd = "sleep"
			schedule = "* * * * *"
			args = []string{"0.1"}
		})

		It("creates a new job", func() {
			Ω(err).ShouldNot(HaveOccurred())
			Ω(job).ShouldNot(BeNil())
			Ω(job.Cmd).Should(Equal(cmd))
			Ω(job.Schedule).Should(Equal(schedule))
			Ω(job.Args).Should(Equal(args))
		})
	})

	Context("with an invalid schedule", func() {
		BeforeEach(func() {
			cmd = "sleep"
			schedule = "foobar"
			args = []string{"0.1"}
		})

		It("returns an error", func() {
			Ω(err).Should(HaveOccurred())
			Ω(job).Should(BeNil())
		})
	})
})

var _ = Describe("Execute", func() {
	var cmd, schedule string
	var args []string

	var job *cron.Job
	var err error

	BeforeEach(func() {
		log.SetOutput(ioutil.Discard)
	})

	JustBeforeEach(func() {
		job, err = cron.NewJob(cmd, args, schedule)
	})

	AfterEach(func() {
		if job != nil {
			job.Stop()
		}
	})

	Context("a job created with valid parameters", func() {
		BeforeEach(func() {
			cmd = "sleep"
			schedule = "* * * * *"
			args = []string{"0.1"}
		})

		It("executes", func() {
			job.Execute()
			Ω(job.Last).ShouldNot(BeNil())
			Ω(job.Last.Pid).ShouldNot(BeZero())
		})
	})

	Context("a job created with an invalid command", func() {
		var b bytes.Buffer
		var exitStatus int

		BeforeEach(func() {
			cmd = "ping1234_Idontexist"
			schedule = "* * * * *"
			args = []string{"google.com"}
			log.SetOutput(&b)
			cron.OsExit = func(code int) { exitStatus = code }
		})

		It("logs an error", func() {
			job.Execute()
			Ω(job.Last).Should(BeNil())
			Ω(b.String()).Should(ContainSubstring("executable file not found"))
			Ω(exitStatus).Should(Equal(1))
		})
	})
})
