package main

import "fmt"

var company = map[string]any{
	"Engineering": map[string]any{
		"Backend": map[string]any{
			"Alice": "Developer",
			"Bob":   "Developer",
		},
		"Frontend": map[string]any{
			"John": "Developer",
		},
	},
	"HR": map[string]any{
		"Mary": "Manager",
	},
}

func isString(val any) bool {
	_, ok := val.(string)
	return ok
}

func printEvery(listing map[string]any) {
	for k, v := range listing {
		fmt.Println(k)
		if !isString(v) {
			if m, ok:=v.(map[string]any); ok {
				printEvery(m)
			}
		} else {
			fmt.Println(k)
		}
	}
}

func main() {
	printEvery(company)
}