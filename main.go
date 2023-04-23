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
	wf := workflow.New(main, infoLog)
	wf.Start()
	time.Sleep(1 * time.Second)
	wf.Finish()
}
