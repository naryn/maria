package tts

import (
	//"os"
	//"os/signal"
	"fmt"
	"sync"
)

var say chan string

var Tts map[string]Need

func init() {
	fmt.Println("init tts")
	say = make(chan string, 1)
	//Tts = make(map[string]Need)
}

type Need interface {
	Say(msg string)
}

func Say(msg string) {
	fmt.Println("tts Say")
	say <- msg
}

func Run(ttsName string , waitGroup *sync.WaitGroup) {

	fmt.Println("tts run")
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, os.Interrupt, os.Kill)
	for {
		select {
		case  msg :=<- say:
			if call, ok := Tts[ttsName]; ok {
				fmt.Println("channel " + msg)
				call.Say(msg)
			} else {
				fmt.Println("Can't read tts " + ttsName)
				return
			}

		/*case <-sig:
			return*/
		}
	}
	fmt.Println("tts done")
	waitGroup.Done()
}

func Regist(name string, tts Need) {
	if Tts == nil {
		Tts = make(map[string]Need)
	}
	Tts[name] = tts
	fmt.Println(Tts)
}
