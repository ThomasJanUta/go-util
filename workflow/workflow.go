// Package workflow consistently logs workflow status messages and feedback.
package workflow

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Workflow struct {
	name    string
	logger  *log.Logger
	start   time.Time
	options interface{}
}

func New(workflowFn interface{}, logger *log.Logger) *Workflow {
	fullName := runtime.FuncForPC(reflect.ValueOf(workflowFn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	shortName := parts[len(parts)-1]
	wf := Workflow{
		name:   shortName,
		logger: logger,
		start:  time.Now(),
	}
	return &wf
}

func (wf *Workflow) Start(args ...interface{}) *Workflow {
	err := wf.logger.Output(2, fmt.Sprintln("### WORKFLOW START:", wf.name, args))
	if err != nil {
		fmt.Println("log.Output errored during logging:", err)
		fmt.Sprintln(args...)
	}
	return wf
}
func (wf *Workflow) Finish(args ...interface{}) {
	duration := time.Since(wf.start).Minutes()
	err := wf.logger.Output(2, fmt.Sprintln("### WORKFLOW FINISH:", wf.name,
		"| duration in minutes:", fmt.Sprintf("%.3f", duration),
		args))
	if err != nil {
		fmt.Println("log.Output errored during logging:", err)
		fmt.Sprintln(args...)
	}
}
func (wf *Workflow) Step(args ...interface{}) {
	err := wf.logger.Output(2, fmt.Sprintln("#####", wf.name, args))
	if err != nil {
		fmt.Println("log.Output errored during logging:", err)
		fmt.Sprintln(args...)
	}
}
func (wf *Workflow) Result(args ...interface{}) {
	err := wf.logger.Output(2, fmt.Sprintln("### WORKFLOW RESULT:", wf.name, args))
	if err != nil {
		fmt.Println("log.Output errored during logging:", err)
		fmt.Sprintln(args...)
	}
}
