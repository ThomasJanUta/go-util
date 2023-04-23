package workflow

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_NewWorkflow(t *testing.T) {
	var buf bytes.Buffer
	infoLog := log.New(&buf, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	workflowFn := func() {}
	wf := New(workflowFn, infoLog)
	wf.Start()
	wf.Finish()
	assert.Contains(t, buf.String(), "WORKFLOW START:")
	assert.Contains(t, buf.String(), "WORKFLOW FINISH:")
}
