package tokenizer

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	internal "github.com/ableinc/prompt-token-count/internal"
)

type Encoding struct {
	Name           string
	PatStr         string
	MergeableRanks map[string]int
	SpecialTokens  map[string]int
	Pattern        *regexp.Regexp
}

type TokenString string

var encodings = map[string]*Encoding{
	"cl100k_base": {
		Name:   "cl100k_base",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
	"p50k_base": {
		Name:   "p50k_base",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
	"r50k_base": {
		Name:   "r50k_base",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
	"o200k_base": {
		Name:   "o200k_base",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
	"p50k_edit": {
		Name:   "p50k_edit",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
	"gpt2": {
		Name:   "gpt2",
		PatStr: `'s|'t|'re|'ve|'m|'ll|'d| ?\p{L}+| ?\p{N}+| ?[^\s\p{L}\p{N}]+|\s+`,
	},
}

func GetEncoding(name string) (*Encoding, error) {
	enc, ok := encodings[name]
	if !ok {
		return nil, fmt.Errorf("unknown encoding: %s", name)
	}

	if enc.MergeableRanks == nil {
		mergeableRanks, err := internal.LoadMergeableRankFile()
		if err != nil {
			return nil, fmt.Errorf("failed to load mergeable ranks: %v", err)
		}
		enc.MergeableRanks = mergeableRanks
	}

	if enc.SpecialTokens == nil {
		specialTokens, err := internal.LoadSpecialTokensFile()
		if err != nil {
			return nil, fmt.Errorf("failed to open special tokens: %v", err)
		}
		enc.SpecialTokens = specialTokens
	}

	if enc.Pattern == nil {
		enc.Pattern = regexp.MustCompile(enc.PatStr)
	}

	return enc, nil
}

func (enc *Encoding) Encode(text TokenString) []int {
	tokens := enc.Pattern.FindAllString(text.ToString(), -1)
	ids := make([]int, 0, len(tokens))

	for _, token := range tokens {
		if id, ok := enc.SpecialTokens[token]; ok {
			ids = append(ids, id)
		} else {
			subTokens := enc.bytePairEncode(token)
			for _, subToken := range subTokens {
				ids = append(ids, enc.MergeableRanks[subToken])
			}
		}
	}
	return ids
}

func (enc *Encoding) DecodeRaw(tokens []int) string {
	decoded := make([]string, len(tokens))
	for i, token := range tokens {
		found := false
		for str, id := range enc.SpecialTokens {
			if id == token {
				decoded[i] = str
				found = true
				break
			}
		}
		if !found {
			for str, id := range enc.MergeableRanks {
				if id == token {
					decoded[i] = str
					break
				}
			}
		}
	}
	return strings.Join(decoded, "")
}

func (enc *Encoding) Decode(tokens []int) TokenString {
	decoded := make([]string, len(tokens))
	for i, token := range tokens {
		found := false
		for str, id := range enc.SpecialTokens {
			if id == token {
				decoded[i] = str
				found = true
				break
			}
		}
		if !found {
			for str, id := range enc.MergeableRanks {
				if id == token {
					// Remove leading "Ġ" (space indicator) if present
					decoded[i] = strings.TrimPrefix(str, "Ġ")
					break
				}
			}
		}
	}
	// Join tokens and replace "!" with a space
	return TokenString(strings.ReplaceAll(strings.Join(decoded, ""), "!", " "))
}

func (enc *Encoding) bytePairEncode(token string) []string {
	parts := strings.Split(token, "")
	for {
		bestPair := ""
		bestRank := -1

		for i := range len(parts) - 1 {
			pair := parts[i] + parts[i+1]
			if rank, ok := enc.MergeableRanks[pair]; ok {
				if bestRank == -1 || rank < bestRank {
					bestPair = pair
					bestRank = rank
				}
			}
		}

		if bestPair == "" {
			break
		}

		newParts := make([]string, 0, len(parts))
		i := 0
		for i < len(parts) {
			if i < len(parts)-1 && parts[i]+parts[i+1] == bestPair {
				newParts = append(newParts, bestPair)
				i += 2
			} else {
				newParts = append(newParts, parts[i])
				i++
			}
		}
		parts = newParts
	}
	return parts
}

func (text TokenString) CountTokens() int {
	// Rule of thumb: 1 token ~ 4 characters
	tokenLength := 4.0

	// Calculate the number of tokens
	numTokens := float64(len(text)) / tokenLength

	// Round up to the nearest integer
	return int(math.Ceil(numTokens))
}

func (text TokenString) ToString() string {
	return string(text)
}
