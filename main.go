package GoUrban

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// An array of definitions
type Definitions []Definition

// A struct representing a definition definition.
//UpVotes are known as thumbs_up on the website, but UpVotes was shorter. Same with thumbs_down.
//current_vote was omitted because it never has a value.
type Definition struct {
	Def       string   `json:"definition"`
	Permalink string   `json:"permalink"`
	UpVotes   int      `json:"thumbs_up"`
	DownVotes int      `json:"thumbs_down"`
	SoundUrls []string `json:"sound_urls"`
	Author    string   `json:"author"`
	Word      string   `json:"word"`
	DefID     int      `json:"defid"`
	WrittenOn string   `json:"written_on"`
	Example   string   `json:"example"`
}

// Usage: Define(term). This will query the urban dictionary api for your term.
func Define(term string) Definitions {
	resp, err := http.Get("https://api.urbandictionary.com/v0/define?term=" + term)
	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error fetching definition: ", err)
	}
	var unmar map[string]Definitions
	json.Unmarshal(body, &unmar)
	return unmar["list"]

}
