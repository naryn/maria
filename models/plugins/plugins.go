package plugins

import (
	"strings"
	"fmt"
	"github.com/naryn/maria/models/tts"
	"time"
)

const DEFAULT_PLUGIN = "plugins_robot"
const CMD_STOP = "STOP"
const CMD_PAUSE = "PAUSE"
const CMD_RESET = "RESET"

var Plugins map[string]Need

var PluginPause bool
var PluginStop bool
var PluginThread string

func init() {
	fmt.Println("init plugin")
	PluginPause = false
	PluginThread = ""
	//Plugins = make(map[string]Need)
}

type Need interface {
	Flag() bool
	Active() bool
	Keyword() []string
	Run(msg string)
	Control(msg string) string
	Say(msg string)
}

func RunBefore(msg string) {

}
func RunAfter(msg string) {

}

func threadRunPlugin(plugin Need, msg string) {
	RunBefore(msg)
	plugin.Run(msg)
	RunAfter(msg)
}

func Say(msg string) {
	tts.Say(msg)
}

func PluginControl() bool {
	for PluginPause {
		time.Sleep(0.8 * time.Second);
	}

	return !PluginStop
}

func SendCmd(cmd string) {

	switch cmd {
	case CMD_STOP:
		PluginStop = true
		PluginPause = false
	case CMD_PAUSE:
		PluginPause = true
	case CMD_RESET:
		PluginPause = false
		PluginStop = false
	}
	time.Sleep(1.3 * time.Second);
}

func Each(str string) {

	matching := false
	for key, plugin := range Plugins {

		for _, w := range plugin.Keyword() {
			if strings.Contains(str, w) && plugin.Flag() {
				/* 查看插件是否是启用状态 */
				matching = true

				if PluginPause {
					SendCmd(plugin.Control(str))
				} else if plugin.Active() {
					if key != PluginThread && PluginThread != "" {
						SendCmd(CMD_STOP)
						PluginThread = ""
					}
					if PluginThread == "" {
						PluginThread = key
						SendCmd(CMD_RESET)
						go threadRunPlugin(plugin, str)
						return
					} else {
						//todo
					}

				}
			}
		}

	}

	if matching == false {
		SendCmd(CMD_STOP)
		PluginThread = DEFAULT_PLUGIN
		go threadRunPlugin(Plugins[DEFAULT_PLUGIN], str)
	}

}

func Regist(name string, plugin Need) {
	if Plugins == nil {
		Plugins = make(map[string]Need)
	}
	Plugins[name] = plugin
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}