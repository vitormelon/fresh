package runner

import (
	"fmt"
	"io"
	"os/exec"
)

func run() bool {
	runnerLog("RODANDOOOOOOO...")
	fmt.Println(buildPath())
	cmd := exec.Command("dlv", "debug", "--listen=:40000", "--headless=true", "--api-version=2", "--log")

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
