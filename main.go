package main

import (
	"fmt"
	"log"
	"os"

	tokenizer "github.com/ableinc/prompt-token-count/cmd/tokenizer"
)

func EncodingForModel(model string) (*tokenizer.Encoding, error) {
	switch model {
	case "gpt-4", "gpt-3.5-turbo", "text-embedding-ada-002":
		return tokenizer.GetEncoding("cl100k_base")
	case "text-davinci-002", "text-davinci-003":
		return tokenizer.GetEncoding("p50k_base")
	default:
		return nil, fmt.Errorf("unknown model: %s", model)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) != 2 {
		log.Fatal("You must provided a model name and prompt.")
	}
	modelEncoding, err := EncodingForModel(args[0])
	if err != nil {
		log.Fatal(err)
	}
	prompt := tokenizer.TokenString(args[1])
	tokens := modelEncoding.Encode(prompt)
	fmt.Println("Encoding: ", tokens)
	prompt = modelEncoding.Decode(tokens)
	fmt.Println("Decoding: ", prompt)
	fmt.Println("Number of tokens: ", prompt.CountTokens())
}
