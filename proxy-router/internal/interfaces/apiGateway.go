package interfaces

import (
	"context"
	"math/big"

	"github.com/gin-gonic/gin"
)

type ApiGateway interface {
	// Proxy Router Api
	  GetConfig(ctx context.Context) interface{}
	
	  GetFiles(ctx *context.Context) (int, interface{})
	
	  HealthCheck(ctx context.Context) interface{}
	
	  InitiateSession(ctx context.Context) (int, interface{})
	
	  SendPrompt(ctx context.Context) (bool, int, interface{})
	
	// AiEngine
	  Prompt(ctx context.Context, req interface{}) (interface{}, error)
	
	// AiEngine
	  PromptStream(ctx context.Context, req interface{}, flush interface{}) (interface{}, error)
	
	// RpcProxy
	  GetLatestBlock(ctx context.Context) (uint64, error) 

	  GetAllProviders(ctx context.Context) (int, gin.H) 

	  GetAllModels(ctx context.Context) (int, gin.H) 
	
	  GetBidsByProvider(ctx context.Context, providerAddr string, offset *big.Int, limit uint8) (int, gin.H) 
	
	  GetBidsByModelAgent(ctx context.Context, modelAgentId [32]byte, offset *big.Int, limit uint8) (int, gin.H)
	
	  OpenSession(ctx *gin.Context) (int, gin.H)
	
	  CloseSession(ctx *gin.Context) (int, gin.H)	
}