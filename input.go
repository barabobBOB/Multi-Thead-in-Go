package main

import (
	"bufio"
	"os"
	"strings"
)

type Command struct {
	keyword string
	path    string
}

func InputCommend() Command {
	scanText := bufio.NewScanner(os.Stdin)
	scanText.Scan()
	commend := scanText.Text()
	commandList := strings.Split(commend, " ")

	if commandList[0] != "dgrep" {
		panic("존재하지 않는 명령어입니다.")
	}

	if len(commandList) >= 4 {
		panic("명령어가 올바르지 않습니다.")
	}

	return Command{commandList[1], commandList[2]}
}
