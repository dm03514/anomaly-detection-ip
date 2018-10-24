package blacklist

import "encoding/json"

type Provider struct {
	Name     string
	Metadata map[string]string
}

type Result struct {
	Found     bool
	CIDR      string
	Providers []Provider
}

func (r Result) Bytes() ([]byte, error) {
	return json.Marshal(r)
}
