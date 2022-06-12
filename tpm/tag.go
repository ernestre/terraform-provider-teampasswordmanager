package tpm

import (
	"encoding/json"
	"sort"
	"strings"
)

type Tags []string

func (t Tags) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ToString())
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	// API returns empty string if there are no tags assigned
	if s != "" {
		*t = Tags(strings.Split(s, ","))
	}

	return nil
}

func (t Tags) ToString() string {
	s := []string{}

	for _, v := range t {
		v = strings.Trim(v, " ")
		if v != "" {
			s = append(s, v)
		}
	}

	sort.Strings(s)

	return strings.Join(s, ",")
}
