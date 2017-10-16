package stt

import (
	"fmt"
	"sync"
	"github.com/naryn/maria/models/tts"
)

var Stt map[string]Need

func init() {
	fmt.Println("init stt")
}

type Need interface {
	Start()
}

func Regist(name string, stt Need) {
	if Stt == nil {
		Stt = make(map[string]Need)
	}
	Stt[name] = stt
}

func Say(msg string)  {
	tts.Say(msg)
}
func Run(sttName string , waitGroup *sync.WaitGroup) {
	fmt.Println("stt run")

	if call, ok := Stt[sttName]; ok {
		call.Start()
	} else {
		fmt.Println("Can't read stt " + sttName)
	}
	fmt.Println("stt done")

	waitGroup.Done()
}
