package mainchat

import (
	"flag"
	"fmt"
	"os"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/chat"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/common"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/config"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/util"
)

func init() {
	flag.BoolVar(&opt.Edit, "e", false, "Edit configuration file")
	flag.BoolVar(&opt.Edit, "edit", false, "Edit configuration file")

	flag.BoolVar(&opt.List, "l", false, "List all supported OpenAI model")
	flag.BoolVar(&opt.List, "list", false, "List all supported OpenAI model")

	flag.BoolVar(&opt.Remove, "rm", false, "Remove configuration file")

	flag.BoolVar(&opt.Version, "V", false, "Show current version")
	flag.BoolVar(&opt.Version, "version", false, "Show current version")

	openAiBaseUrl := os.Getenv("OPENAI_BASE_URL")

	if openAiBaseUrl == "" {
		os.Setenv("OPENAI_BASE_URL", "http://localhost:8082/v1")
	}

	flag.Usage = func() {
		showBanner()
		showUsage()
	}
	flag.Parse()

	switch {
	case opt.List:
		listAllModels()
	case opt.Remove:
		removeConfig()
	case opt.Version:
		showVersion()
	}

	// if opt.List {
	// 	listAllModels()
	// }

	// if opt.Remove {
	// 	removeConfig()
	// }

	// if opt.Version {
	// 	showVersion()
	// }
}

func main(sessionId string) {
	cfgPath := common.GetConfigPath()
fmt.Println("config path: ", cfgPath)
	cfg, err := config.Load(cfgPath)
	fmt.Println("err: ", err)
	if err == nil {
		cfg.SessionId = sessionId
		m = chat.New(cfg)

		if opt.Edit {
			m = config.New(sessionId, cfg)
		}
	} else { 
		m = config.New(sessionId)
	}

	util.RunProgram(m)
}

func Run(sessionId string) {
	main(sessionId)
}
