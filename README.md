# Prompt Token Count

Tokenize a prompt and/or get the number of tokens a prompt will take

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/ableinc/prompt-token-count/cmd/tokenizer"
)

func encodingForModel(model string) (*tokenizer.Encoding, error) {
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
	model, err := encodingForModel("gpt-4")
	if err != nil {
		log.Fatalf("incorrect model provided: %v", err)
	}
	prompt := tokenizer.TokenString("Give me Golang code to create a binary tree.")
	fmt.Println("Number of tokens: ", prompt.CountTokens())
	tokens := model.Encode(prompt)
	fmt.Println("Encoding: ", tokens)
	prompt = model.Decode(tokens)
	fmt.Println("Decoding: ", prompt)
}
```
