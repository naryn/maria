package tts

import (
	"github.com/naryn/maria/utils"
	"fmt"
)

func init() {
	fmt.Println("init macOS")

	tts := tts_macOS{}
	Regist("macOS", tts)
}

type tts_macOS struct {
}

func (this tts_macOS) Say(msg string) {
	fmt.Println("say " + msg)
	command := "say"
	params := []string{msg}
	utils.ExecCommand(command, params)
}
