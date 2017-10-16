package plugins

import "time"

func init() {
	p := plugin_robot{}
	Regist("plugin_weather", p)
}

type plugin_robot struct {
}

func (this plugin_robot) Keyword() []string {
	return []string{""}
}

func (this plugin_robot) Flag() bool {
	return true
}

func (this plugin_robot) Active() bool {
	return true
}

func (this plugin_robot) Run(msg string) {
	//Say("好的, 现在" + msg)
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))

}