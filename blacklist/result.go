package blacklist

import "encoding/json"

type Result struct {
	Found    bool
	Metadata interface{}
}

func (r Result) Bytes() ([]byte, error) {
	return json.Marshal(r)
}
