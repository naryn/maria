package plugins

import "time"

func init() {
	p := plugin_weather{}
	Regist("plugin_weather", p)
}

type plugin_weather struct {
}

func (this plugin_weather) Keyword() []string {
	return []string{""}
}

func (this plugin_weather) Flag() bool {
	return true
}

func (this plugin_weather) Active() bool {
	return true
}

func (this plugin_weather) Run(msg string) {
	//Say("好的, 现在" + msg)
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))

}