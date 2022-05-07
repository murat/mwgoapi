package mwgoapi

import (
	"encoding/json"
	"fmt"
)

// Meta is the metadata for the API.
type Meta struct {
	ID        string   `json:"id"`
	UUID      string   `json:"uuid"`
	Sort      string   `json:"sort"`
	Src       string   `json:"src"`
	Section   string   `json:"section"`
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}

// Sound ...
type Sound struct {
	Audio string `json:"audio"`
	Ref   string `json:"ref"`
	Stat  string `json:"stat"`
}

// Pronunciation ...
type Pronunciation struct {
	MerriemWebster string  `json:"mw"`
	Sound          *Sound  `json:"sound,omitempty"`
	AudioURL       *string `json:"audio_url,omitempty"`
}

// HeadwordInfo is the headword info.
type HeadwordInfo struct {
	Headword       string          `json:"hw"`
	Pronunciations []Pronunciation `json:"prs"`
}

// Collegiate is the Collegiate Dictionary API response.
type Collegiate struct {
	Meta            Meta                `json:"meta"`
	Headword        HeadwordInfo        `json:"hwi"`
	FunctionalLabel string              `json:"fl"`
	Inflections     []map[string]string `json:"ins"`
	Date            string              `json:"date"`
	Etymologies     [][]string          `json:"et"`
	Shortdef        []string            `json:"shortdef"`
}

func (c *Collegiate) UnmarshalJSON(data []byte) error {
	type CollegiateAlias Collegiate
	tmp := &struct {
		*CollegiateAlias
	}{
		(*CollegiateAlias)(c),
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("could not unmarshal Collegiate, %w", err)
	}

	for i := 0; i < len(tmp.Headword.Pronunciations); i++ {
		sound := tmp.Headword.Pronunciations[i].Sound
		if sound == nil {
			continue
		}
		audio := sound.Audio
		audioURL := fmt.Sprintf("%s/%s/%s/%s.%s", AudioBaseURL, AudioFormat, audio[0:1], audio, AudioFormat)
		c.Headword.Pronunciations[i].AudioURL = &audioURL
	}

	return nil
}
