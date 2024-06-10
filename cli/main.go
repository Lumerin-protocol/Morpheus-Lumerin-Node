package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"bufio"
	"strings"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/api-gateway/client"

	"github.com/urfave/cli/v2"
)

const httpErrorMessage string = "internal error: %v; http status: %v"

func main() {
	api_host := "http://localhost:8082"
	file, err := os.Open(".env")
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			nv := strings.SplitN(scanner.Text(), "=", 2)
			if nv == nil || len(nv) != 2 {
				continue
			}
			n := strings.Trim(nv[0], " ")
			v := strings.Trim(nv[1], " ")
			if n == "API_HOST" {
				api_host = v
			}
		}
	}
	defer file.Close()

	actions := NewActions(client.NewApiGatewayClient(api_host, http.DefaultClient))
	app := &cli.App{
		Usage: "A client to call the Morpheus Lumerin API",
		Commands: []*cli.Command{
			{
				Name: "getAllowance",
				Aliases: []string{"ga"},
				Usage: "get allowance",
				Action: actions.getAllowance,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "spender",
						Required: true,
					},
				},
			},
			{
				Name: "approveAllowance",
				Aliases: []string{"aa"},
				Usage: "approve allowance",
				Action: actions.approveAllowance,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "spender",
						Required: true,
					},
					&cli.Uint64Flag{
						Name: "amount",
						Required: true,
					},
				},
			},
			{
				Name:    "healthcheck",
				Aliases: []string{"he"},
				Usage:   "check application health",
				Action:  actions.healthcheck,
			},
			{
				Name:    "proxyRouterConfig",
				Aliases: []string{"prc"},
				Usage:   "view proxy router config",
				Action:  actions.proxyRouterConfig,
			},
			{
				Name:    "proxyRouterFiles",
				Aliases: []string{"prf"},
				Usage:   "get the files associated with the proxy router pid",
				Action:  actions.proxyRouterFiles,
			},
			{
				Name:    "createChatCompletions",
				Aliases: []string{"ccc"},
				Usage:   "create chat completions by sending a prompt to the AI engine",
				Action:  actions.createChatCompletions,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "prompt",
						Required: true,
					},
					&cli.StringFlag{
						Name: "messages",
					},
				},
			},
			{
				Name:    "initiateProxySession",
				Aliases: []string{"ips"},
				Usage:   "initiate a proxy session",
				Action:  actions.initiateProxySession,
			},
			{
				Name:    "blockchainProviders",
				Aliases: []string{"bp"},
				Usage:   "list blockchain providers",
				Action:  actions.blockchainProviders,
			},
			{
				Name:    "blockchainProvidersCreate",
				Aliases: []string{"bpc"},
				Usage:   "create a blockchain provider",
				Action:  actions.createBlockchainProvider,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "address",
						Required: true,
					},
					&cli.Uint64Flag{
						Name: "stake",
						Required: true,
					},
					&cli.StringFlag{
						Name: "endpoint",
						Required: true,
					},
				},
			},
			{
				Name:    "blockchainProvidersBids",
				Aliases: []string{"bpb"},
				Usage:   "list provider bids",
				Action:  actions.blockchainProvidersBids,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "address",
						Required: true,
					},
					&cli.Int64Flag{
						Name: "offset",
					},
					&cli.UintFlag{
						Name: "limit",
					},
				},
			},
			{
				Name:    "blockchainModels",
				Aliases: []string{"bm"},
				Usage:   "list models",
				Action:  actions.blockchainModels,
			},
			{
				Name:    "openBlockchainSession",
				Aliases: []string{"obs"},
				Usage:   "open a blockchain session",
				Action:  actions.openBlockchainSession,
			},
			{
				Name:    "closeBlockchainSession",
				Aliases: []string{"cbs"},
				Usage:   "close a blockchain session",
				Action:  actions.closeBlockchainSession,
			},
			{
				Name:    "createAndStreamChatCompletions",
				Aliases: []string{"csc"},
				Usage:   "create and stream chat completions",
				Action:  actions.createAndStreamChatCompletions,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "prompt",
						Required: true,
					},
					&cli.StringFlag{
						Name: "messages",
					},
				},
			},
			{
				Name:    "createAndStreamSessionChatCompletions",
				Aliases: []string{"cssc"},
				Usage:   "create and stream session chat completions",
				Action:  actions.createChatCompletions,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "prompt",
						Required: true,
					},
					&cli.StringFlag{
						Name: "messages",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type actions struct {
	client *client.ApiGatewayClient
}

func NewActions(c *client.ApiGatewayClient) *actions {
	return &actions{client: c}
}

func (a *actions) getAllowance(cCtx *cli.Context) error {
	spender := cCtx.String("spender")
	res, err := a.client.GetAllowance(cCtx.Context, spender)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (a *actions) approveAllowance(cCtx *cli.Context) error {
	spender := cCtx.String("spender")
	amount := cCtx.Uint64("amount")
	res, err := a.client.ApproveAllowance(cCtx.Context, spender, amount)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (a *actions) healthcheck(cCtx *cli.Context) error {
	res, err := a.client.HealthCheck(cCtx.Context)
	if err != nil {
		return err
	}
	// Output the result of the healthcheck
	jsonData, err := json.Marshal(res)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) proxyRouterConfig(cCtx *cli.Context) error {
	config, err := a.client.GetProxyRouterConfig(cCtx.Context)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(config)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) proxyRouterFiles(cCtx *cli.Context) error {
	files, err := a.client.GetProxyRouterFiles(cCtx.Context)

	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(files)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) createChatCompletions(cCtx *cli.Context) error {
	prompt := cCtx.String("prompt")
	var messages []client.ChatCompletionMessage
	json.Unmarshal([]byte(cCtx.String("messages")), &messages)

	completion, err := a.client.Prompt(cCtx.Context, prompt, messages)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(completion)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) createAndStreamChatCompletions(cCtx *cli.Context) error {
	prompt := cCtx.String("prompt")
	var messages []*client.ChatCompletionMessage
	json.Unmarshal([]byte(cCtx.String("messages")), &messages)

	completion, err := a.client.PromptStream(cCtx.Context, prompt, messages, func(msg client.ChatCompletionStreamResponse) error {
		fmt.Println(msg)
		return nil
	})

	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(completion)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) createSessionChatCompletions(cCtx *cli.Context) error {
	prompt := cCtx.String("prompt")
	var messages []client.ChatCompletionMessage
	json.Unmarshal([]byte(cCtx.String("messages")), &messages)

	completion, err := a.client.SessionPrompt(cCtx.Context, prompt)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(completion)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) initiateProxySession(cCtx *cli.Context) error {
	session, err := a.client.InitiateSession(cCtx.Context)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(session)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) blockchainProviders(cCtx *cli.Context) error {
	providers, err := a.client.GetAllProviders(cCtx.Context)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(providers)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) createBlockchainProvider(cCtx *cli.Context) error {
	address := cCtx.String("address")
	stake := cCtx.Uint64("stake")
	endpoint := cCtx.String("endpoint")

	providers, err := a.client.CreateNewProvider(cCtx.Context, address, stake, endpoint)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(providers)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) blockchainProvidersBids(cCtx *cli.Context) error {
	address := cCtx.String("address")
	offset := cCtx.Int64("offset")
	limit := cCtx.Uint("limit")

	bids, err := a.client.GetBidsByProvider(cCtx.Context, address, big.NewInt(offset), uint8(limit))
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(bids)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) blockchainModels(cCtx *cli.Context) error {
	models, err := a.client.GetAllModels(cCtx.Context)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(models)
	fmt.Println(string(jsonData))
	return nil
}

func (a *actions) openBlockchainSession(cCtx *cli.Context) error {
	err := a.client.OpenSession(cCtx.Context)
	if err != nil {
		return err
	}
	// TODO: Output a message indicating the blockchain session was opened and showing relevant data for the session
	// jsonData, err := json.Marshal(bids)
	// fmt.Println(string(jsonData))
	return nil
}

func (a *actions) closeBlockchainSession(cCtx *cli.Context) error {
	err := a.client.CloseSession(cCtx.Context)
	if err != nil {
		return err
	}

	// TODO: Output a message indicating the blockchain session was closed and showing relevant data for the session
	// jsonData, err := json.Marshal(bids)
	// fmt.Println(string(jsonData))
	return nil
}
