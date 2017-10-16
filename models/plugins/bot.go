package plugins

import "time"

func init() {
	p := plugin_bot{}
	Regist("plugin_bot", p)
}

type plugin_bot struct {
}

func (this plugin_bot) Keyword() []string {
	return []string{""}
}

func (this plugin_bot) Flag() bool {
	return true
}

func (this plugin_bot) Active() bool {
	return true
}

func (this plugin_bot) Run(msg string) {
	//Say("好的, 现在" + msg)
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))

}