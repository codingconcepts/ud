package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/codingconcepts/ud/pkg/client"
	"github.com/codingconcepts/ud/pkg/model"
	"github.com/fatih/color"
)

func main() {
	args := os.Args[1:]
	resp, err := client.Request(strings.Join(args, " "))
	if err != nil {
		log.Fatal(err)
	}

	dump(resp)
}

func dump(resp *model.Response) {
	defer func() {
		color.Unset()
	}()

	if len(resp.Results) < 1 {
		color.Red("No results")
		return
	}

	// if multiple results, sort them by rank descending
	// to ensure the best result is the one we display
	if len(resp.Results) > 1 {
		sort.Slice(resp.Results, func(i int, j int) bool {
			return resp.Results[i].ThumbsUp > resp.Results[j].ThumbsUp
		})
	}

	example := resp.Results[0]
	if example.Definition != "" {
		fmt.Printf("%s: %s\n", color.RedString("Definition"), color.WhiteString(example.Definition))
	}
	if example.Example != "" {
		fmt.Printf("%s: %s\n", color.YellowString("Example"), color.WhiteString(example.Example))
	}
}
