package plugins

import (
	"strings"
	"fmt"
	"github.com/naryn/maria/models/tts"
)

var Plugins map[string]Need

func init() {
	fmt.Println("init plugin")
	//Plugins = make(map[string]Need)
}

type Need interface {
	Flag() bool
	Active() bool
	Keyword() []string
	Run(msg string)
}

func Say(msg string)  {
	tts.Say(msg)
}

func Each(str string) {

	matching := false
	for _, plugin := range Plugins {

		for _, w := range plugin.Keyword() {
			if (strings.Contains(str, w)) {
				/* 查看插件是否是启用状态 */
				if plugin.Flag() {
					matching = true
					if plugin.Active(){
						go plugin.Run(str)
					}else{
						go plugin.Run(str)
					}

				}
			}
		}

	}

	if matching == false {
		go Plugins["plugins_robot"].Run(str)
	}
}

func Regist(name string, plugin Need) {
	if Plugins == nil {
		Plugins = make(map[string]Need)
	}
	Plugins[name] = plugin
}
