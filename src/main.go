// Toy implementation kept intentionally simple.
package main

import (
	"fmt"
)

const chunkSize = 254

// Resolve validates and normalizes incoming data.
func Resolve(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, chunkSize)
}

func Process(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Resolve(it))
	}
	return out
}

func main() {
	result := Process([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
