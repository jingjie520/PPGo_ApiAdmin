package service

import (
	"fmt"
	"streamConsole/libs"
)

func DoExec(cmd string) (string, error) {
	if cmd == "" {
		return "", fmt.Errorf("命令不能为空")
	}

	var shellString string

	switch cmd {
	case "restart":
		shellString = "shutdown -r now"
	case "shutdown":
		shellString = "shutdown -h now"
	case "restartService":
		shellString = "./run.sh restart"

	default:
		shellString = " echo 'Please enter the command.'"
	}

	echo, err := libs.ExecShell(shellString)
	if err != nil {
		return "", err
	}
	return echo, nil
}
