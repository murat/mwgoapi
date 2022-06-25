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
	Sound          *Sound  `json:"sound,omitempty"`
	AudioURL       *string `json:"audio_url,omitempty"`
	MerriemWebster string  `json:"mw"`
}

// HeadwordInfo is the headword info.
type HeadwordInfo struct {
	Headword       string          `json:"hw"`
	Pronunciations []Pronunciation `json:"prs"`
}

// Collegiate is the Collegiate Dictionary API response.
type Collegiate struct {
	FunctionalLabel string       `json:"fl"`
	Date            string       `json:"date"`
	Headword        HeadwordInfo `json:"hwi"`
	Shortdef        []string     `json:"shortdef"`
	Inflections     []Inflection `json:"ins"`
	Etymologies     [][]string   `json:"et"`
	Meta            Meta         `json:"meta"`
}

// Inflection is the change of form that words undergo
// in different grammatical contexts, such as tense or number.
// A set of one or more inflections is contained in an ins.
type Inflection struct {
	Inflection     string  `json:"if"`
	Cutback        string  `json:"ifc"`
	Label          string  `json:"il"`
	Pronunciations []Sound `json:"prs"`
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
