package plugins

import (
	"os"
	"os/signal"
	"github.com/bobertlo/go-mpg123/mpg123"
	"encoding/binary"
	"bytes"
	"github.com/gordonklaus/portaudio"
	"strings"
)

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

/**
  控制函数,返回控制指令。
 */
func (this plugin_music) Control(str string) string {
	if strings.Contains(str, "播放"){
		return CMD_RESET
	}

	if strings.Contains(str, "暂停"){
		return CMD_PAUSE
	}

	if strings.Contains(str, "停止") || strings.Contains(str, "关闭"){
		return CMD_STOP
	}

	return true
}

func (this plugin_music) Run(msg string) {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	// create mpg123 decoder instance
	decoder, err := mpg123.NewDecoder("")
	chk(err)

	fileName := "yuanzougaofei.mp3"
	chk(decoder.Open(fileName))
	defer decoder.Close()

	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()
	for PluginControl(){
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		chk(err)

		chk(binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out))
		chk(stream.Write())
		select {
		case <-sig:
			return
		default:
		}
	}
}
