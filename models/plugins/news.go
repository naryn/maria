package plugins

import "time"

func init() {
	p := plugin_news{}
	Regist("plugin_news", p)
}

type plugin_news struct {
}

func (this plugin_news) Keyword() []string {
	return []string{"古诗","诗词","吟诗"}
}

func (this plugin_news) Flag() bool {
	return true
}

func (this plugin_news) Active() bool {
	return true
}

func (this plugin_news) Run(msg string) {
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))
}