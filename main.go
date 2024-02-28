package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var CONVERSATIONAL bool
var VERBOSE bool

func parseFlags() {
	flag.BoolVar(&VERBOSE, "v", false, "verbose output")
	flag.BoolVar(&CONVERSATIONAL, "c", false, "enables conversational output")
	flag.Parse()
}

func gatherPrompt() (string, error) {
	fmt.Println("Ask your overlord(s)...")
	fmt.Print("> ")
	prompt, err := bufio.NewReader(os.Stdin).ReadString(byte('\n'))
	if err != nil {
		log.Fatal("Do not meddle with the breaking of bread.")
	}
	return prompt, nil
}

func main() {
	parseFlags()

	if CONVERSATIONAL {
		fmt.Println("CONVERSATIONAL MODE...")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("PANEM_NOSTRUM_GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-pro")
	prompt, err := gatherPrompt()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Candidates) > 0 {
		fmt.Print("Gemini: ", resp.Candidates[0].Content.Parts[0])
	}
}
