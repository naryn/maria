package plugins

import "time"

func init() {
	p := plugin_music{}
	Regist("plugin_music", p)
}

type plugin_music struct {
}

func (this plugin_music) Keyword() []string {
	return []string{"古诗","诗词","吟诗"}
}

func (this plugin_music) Flag() bool {
	return true
}

func (this plugin_music) Active() bool {
	return true
}

func (this plugin_music) Run(msg string) {
	Say("好的, 现在 " + time.Now().Format("2006-01-02 15:04"))
}