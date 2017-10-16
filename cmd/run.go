package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"Maria/models/stt"
	"Maria/models/tts"
	"sync"
	"Maria/models/conversation"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Maria assistant Server",
	Long: `Run a Maria assistant server

Start a Maria server using a specific tts stt.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please select a tts stt (--help for available options)")

		fmt.Println(args)

		if len(args) < 2 {
			fmt.Println("Please params tts stt")
			log.Print("must  2 params,tts stt")
		} else {
			waitGroup := &sync.WaitGroup{}

			fmt.Println("stt: "+args[0])
			fmt.Println(stt.Stt)
			waitGroup.Add(1)
			go stt.Run(args[0], waitGroup)

			fmt.Println("tts: "+args[1])
			fmt.Println(tts.Tts)
			waitGroup.Add(1)
			go tts.Run(args[1], waitGroup)

			waitGroup.Add(1)
			go conversation.Timer(waitGroup)
			waitGroup.Wait()
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
