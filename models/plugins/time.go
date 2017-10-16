package plugins

import "time"

func init() {
	p := plugin_time{}
	Regist("plugin_time", p)
}

type plugin_time struct {
}

func (this plugin_time) Keyword() []string {
	return []string{"时间","几点","时辰"}
}

func (this plugin_time) Flag() bool {
	return true
}

func (this plugin_time) Active() bool {
	return true
}

func (this plugin_time) Run(msg string) {
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))
}