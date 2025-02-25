package main

import (
	"fmt"
	"log"
	"os"

	tokenizer "github.com/ableinc/prompt-token-count/cmd/tokenizer"
)

func encodingForModel(model string) (*tokenizer.Encoding, error) {
	switch model {
	case "gpt-4", "gpt-4-turbo", "gpt-3.5-turbo", "text-embedding-ada-002":
		return tokenizer.GetEncoding("cl100k_base")
	case "gpt-4o", "gpt-4o-mini":
		return tokenizer.GetEncoding("o200k_base")
	case "text-davinci-002", "text-davinci-003", "code-davinci-002", "text-davinci-edit-001":
		return tokenizer.GetEncoding("p50k_base")
	case "code-cushman-001", "davinci", "curie", "babbage", "ada":
		return tokenizer.GetEncoding("r50k_base")
	case "text-ada-001", "text-babbage-001", "text-curie-001", "text-davinci-001", "code-davinci-001":
		return tokenizer.GetEncoding("gpt2")
	default:
		return nil, fmt.Errorf("unknown model: %s", model)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) != 2 {
		log.Fatal("You must provided a model name and prompt.")
	}
	var prompt tokenizer.TokenString
	modelEncoding, err := encodingForModel(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := os.Stat(args[1])
	if err == nil && !fileInfo.IsDir() {
		data, err := os.ReadFile(args[1])
		if err != nil {
			log.Fatal("Unable to open file at path: ", args[1])
		}
		prompt = tokenizer.TokenString(string(data))
	} else {
		prompt = tokenizer.TokenString(args[1])
	}
	tokens := modelEncoding.Encode(prompt)
	fmt.Println("Encoding: ", tokens)
	prompt = modelEncoding.Decode(tokens)
	fmt.Println("Decoding: ", prompt)
	fmt.Println("Number of tokens (raw text): ", prompt.CountTokens())
	fmt.Println("NUmber of tokens (encodings): ", tokens.CountTokens())
}
