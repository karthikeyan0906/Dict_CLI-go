package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Structs to hold the API response structure
// sample json -> https://api.dictionaryapi.dev/api/v2/entries/en/pen
type Definition struct {
	Definition string `json:"definition"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type WordData struct { //custom data structure
	Word     string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Meanings []Meaning `json:"meanings"`
}

func main() {
	// Check if user provided a word for argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run dict.go <word>")
		return
	}

	word := os.Args[1]                                                             //declaring and initializing in one step and stores value of arg[1] as string
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word) //formats the string

	// Make HTTP GET request
	resp, err := http.Get(url) //https://pkg.go.dev/net/http

	if err != nil {
		fmt.Println("unable to fetch,check args")
		return
	}
	defer resp.Body.Close()
	// read response
	body, err := io.ReadAll(resp.Body)

	// Parse JSON into our structs
	var data []WordData               // data holds slice of worddata struct
	err = json.Unmarshal(body, &data) // storing unstructured json of body to pointer variable data,err stores any error

	// checking if meanings definitions exist
	if len(data) > 0 && len(data[0].Meanings) > 0 && len(data[0].Meanings[0].Definitions) > 0 {
		fmt.Printf("%s (%s): %s\n",
			data[0].Word,
			data[0].Meanings[0].PartOfSpeech,
			data[0].Meanings[0].Definitions[0].Definition)

		if data[0].Phonetic != "" {
			fmt.Println("Pronunciation:", data[0].Phonetic)
		} else {
			fmt.Println("Pronunciation: Not available")
		}
	} else {
		fmt.Println("No definitions found.")
	}
}
