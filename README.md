# Prompt Token Count

This library allows you to tokenize a prompt and/or get the number of tokens a prompt will take

## Usage

```go
package main

import (
	"fmt"
	"github.com/ableinc/prompt-token-count"
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
	var string modelName = "gpt-4"
	var string prompt = "Give me the Golang code to implement prompt-token-count library as a command line tool."
	modelEncoding, err := EncodingForModel(modelName)
	if err != nil {
		log.Fatal(err)
	}
	// Cast prompt to custome string type
	prompt = tokenizer.TokenString(prompt)
	// Get token count based on prompt size (note: you can do this without encoding/decode)
	fmt.Println("Number of tokens: ", prompt.CountTokens()) // output: 22
	// Tokenize
	tokens := modelEncoding.Encode(prompt)
	fmt.Println("Encoding: ", tokens) // output:  [38 72 85 68 0 76 68 0 83 71 68 0 38 78 75 64 77 70 0 66 78 67 68 0 83 78 0 72 76 79 75 68 76 68 77 83 0 79 81 78 76 79 83 12 83 78 74 68 77 12 66 78 84 77 83 0 75 72 65 81 64 81 88 0 64 82 0 64 0 66 78 76 76 64 77 67 0 75 72 77 68 0 83 78 78 75 13]
	// Decode tokens (tokens > original prompt)
	prompt = modelEncoding.Decode(tokens) // output: Give me the Golang code to implement prompt-token-count library as a command line tool.
	fmt.Println("Decoding: ", prompt)
```
