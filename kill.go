package main

import (
	"os/exec"
	"runtime"
)

func KillPort(port string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		killCmd := "(Get-NetTCPConnection -LocalPort " + port + ").OwningProcess -Force"
		cmd = exec.Command("Stop-Process", "-Id", killCmd)
	} else {
		killCmd := "lsof -i tcp:" + port + " | grep LISTEN | awk '{print $2}' | xargs kill -9"
		cmd = exec.Command("bash", "-c", killCmd)
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
