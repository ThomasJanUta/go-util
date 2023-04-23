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
	infoLog *log.Logger
	start   time.Time
	options interface{}
}

func New(workflowFunction interface{}, infoLog *log.Logger) *Workflow {
	fullName := runtime.FuncForPC(reflect.ValueOf(workflowFunction).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	shortName := parts[len(parts)-1]
	wf := Workflow{
		name:    shortName,
		infoLog: infoLog,
		start:   time.Now(),
	}
	return &wf
}

func (wf *Workflow) Start(args ...interface{}) {
	wf.infoLog.Println("### WORKFLOW START:", wf.name, args)
}
func (wf *Workflow) Finish(args ...interface{}) {
	duration := time.Since(wf.start).Minutes()
	wf.infoLog.Println("### WORKFLOW FINISH:", wf.name,
		"| duration in minutes:", fmt.Sprintf("%.3f", duration),
		args)
}
func (wf *Workflow) Step(args ...interface{}) {
	wf.infoLog.Println("#####", wf.name, args)
}
func (wf *Workflow) Result(args ...interface{}) {
	wf.infoLog.Println("### WORKFLOW RESULT:", wf.name, args)
}
