package ipsets

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Metadata struct {
	Extra map[string]string
}

func (m Metadata) Name() string {
	return m.Extra["Name"]
}

func (m Metadata) JSON() (string, error) {
	bs, err := json.Marshal(m.Extra)
	return string(bs), err
}

type Netset struct {
	f io.Reader
}

func NewNetset(f io.Reader) (Netset, error) {
	return Netset{
		f: f,
	}, nil
}

func (n Netset) CIDRS() ([]string, error) {
	var buf bytes.Buffer
	var cidrs []string

	if _, err := io.Copy(&buf, n.f); err != nil {
		return nil, err
	}

	r := bufio.NewReader(&buf)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			fmt.Printf("CIDRS: end of file reached: %q\n", line)
			break
		}
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		// TODO add CIDR regex match!
		cidrs = append(cidrs, strings.TrimSuffix(line, "\n"))
	}
	return cidrs, nil
}

func (n Netset) Metadata() (Metadata, error) {
	return Metadata{
		Extra: map[string]string{
			"Name":            "fullbogons",
			"Type":            "ipv4 hash:net ipset",
			"Maintainer":      "Team Cymru",
			"MaintainerURL":   "http://www.team-cymru.org/",
			"ListSourceURL":   "http://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt",
			"SourceFileDate":  "Tue Oct 16 20:50:04 UTC 2018",
			"Category":        "unroutable",
			"Version":         "1",
			"ThisFileDate":    "Tue Oct 16 21:47:59 UTC 2018",
			"UpdateFrequency": "1 day",
			"Aggregation":     "none",
			"Entries":         "3059 subnets, 613341384 unique IPs",
		},
	}, nil
}
