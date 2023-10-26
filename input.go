package main

import (
	"bufio"
	"os"
	"strings"
)

func InputCommend() (string, string) {
	scanText := bufio.NewScanner(os.Stdin)
	scanText.Scan()
	commend := scanText.Text()
	commandList := strings.Split(commend, " ")

	if commandList[0] != "dgrep" {
		panic("명령어가 올바르지 않습니다.")
	}

	if len(commandList) >= 4 {
		panic("명령어가 올바르지 않습니다.")
	}

	return commandList[1], commandList[2]
}
