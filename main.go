package GoUrban

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Definitions []Definition
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
