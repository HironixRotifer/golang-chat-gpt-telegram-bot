package openai

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

const (
	API_KEY = "" //
)

var (
	// type of openai engine
	EngineTypes = openai.GPT3Dot5Turbo16K0613
	client      *openai.Client
)

func SetupClientOpenAi() {
	os.Getenv(API_KEY) //TODO:
	client = openai.NewClient(API_KEY)
}

// func SetupAzureClient()  (error, *openai.Client) {
// 	azureKey := os.Getenv("AZURE_OPENAI_API_KEY")       // Your azure API key
// 	azureEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT") // Your azure OpenAI endpoint
// 	config := openai.DefaultAzureConfig(azureKey, azureEndpoint)

// 	// If you use a deployment name different from the model name, you can customize the AzureModelMapperFunc function
// 	config.AzureModelMapperFunc = func(model string) string {
// 		azureModelMapping := make(map[string]string, 1)
// 		azureModelMapping["gpt-3.5-turbo"] = "Rosetta"
// 		return azureModelMapping[model]
// 	}
// 	client := openai.NewClientWithConfig(config)

// 	return nil, client
// }
