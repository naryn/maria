package plugins

import "time"

func init() {
	p := plugin_alarmClock{}
	Regist("plugin_bot", p)
}

type plugin_alarmClock struct {
}

func (this plugin_alarmClock) Keyword() []string {
	return []string{""}
}

func (this plugin_alarmClock) Flag() bool {
	return true
}

func (this plugin_alarmClock) Active() bool {
	return true
}

func (this plugin_alarmClock) Run(msg string) {
	//Say("好的, 现在" + msg)
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))

}