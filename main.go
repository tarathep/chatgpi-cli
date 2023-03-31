package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

var messages []Message

type OpenaiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenaiResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usages  Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Messages     Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// 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
func main() {
	APIKEY := os.Getenv("GPT_CLI_APIKEY")
	MODEL := os.Getenv("GPT_CLI_MODEL")

	if MODEL == "" {
		MODEL = "gpt-3.5-turbo"
	}
	if APIKEY == "" {
		fmt.Println("Please set environment variable named GPT_CLI_APIKEY")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ask a question: ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)

		if question == "exit" {
			break
		}

		messages = append(messages, Message{
			Role:    "user",
			Content: question,
		})

		response := getOpenAIResponse(APIKEY, MODEL)
		fmt.Println(response.Choices[0].Messages.Content)
		print("\n")
	}
}

func getOpenAIResponse(apiKey string, model string) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    model,
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	messages = append(messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return response
}
