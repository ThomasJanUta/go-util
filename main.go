package main

import (
	"github.com/ThomasJanUta/go-util/workflow"
	"log"
	"os"
	"time"
)

func main() {
	var infoLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	wf := workflow.New(main, infoLog).Start()
	wf.Step("step 1")
	wf.Result("1 + 2 =", 1+2)
	time.Sleep(1 * time.Second)
	wf.Finish()
}
