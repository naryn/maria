package conversation

import (
	"time"
	"strings"
	"log"
	"sync"
	"fmt"
	"github.com/naryn/maria/models/plugins"
	"github.com/naryn/maria/models/tts"
)

var Msg chan string

var isWakeUp = false
var keyword = ""

func init() {
	keyword = "上升"
	Msg = make(chan string, 1)
}
func Timer(waitGroup *sync.WaitGroup) {
	t1 := time.NewTimer(time.Second * 999)
	for {
		select {
		case msg := <-Msg:
			if (isWakeUp) {
				fmt.Println("isWakeup ing")

				analyseAction(msg)
			} else {
				wakeUp(msg)
				t1.Reset(time.Second * 5)
			}
		case <-t1.C:
			isWakeUp = false
			t1.Reset(time.Second * 999)
		}
	}
	waitGroup.Done()
}

func wakeUp(msg string) bool {
	if (strings.Contains(msg, keyword)) {
		fmt.Println("wakeup !!")
		isWakeUp = true
		tts.Say("我在")
	} else {
		fmt.Println("no wakeup !!" + msg + " <=>" + keyword)
		isWakeUp = false
	}
	return isWakeUp
}

func analyseAction(msg string) {
	//分词 语意分析
	log.Println("analyseAction " + msg)
	plugins.Each(msg)
}
