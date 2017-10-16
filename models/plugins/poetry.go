package plugins

import "time"

func init() {
	p := plugin_time{}
	Regist("plugin_poetry", p)
}

type plugin_poetry struct {
}

func (this plugin_poetry) Keyword() []string {
	return []string{"古诗","诗词","吟诗"}
}

func (this plugin_poetry) Flag() bool {
	return true
}

func (this plugin_poetry) Active() bool {
	return true
}

func (this plugin_poetry) Run(msg string) {
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))
}