package utils

import (
	"bytes"
	"os/exec"
)

var byteCRLF = []byte{'\n'}

/**
运行一个带参数的命令并返回结果
*/
func RunCommand(cmd string, args ...string) (string, error) {
	var stdout bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stdout = &stdout
	c.Stderr = &stdout

	err := c.Run()
	if err != nil {
		return "", err
	} else {
		return stdout.String(), nil
	}
}
