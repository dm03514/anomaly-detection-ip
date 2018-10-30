package ipset

import (
	"os/exec"
	"regexp"
	"fmt"
	"errors"
	"strings"
)

var (
	ipsetPath            string
	errIpsetNotFound     = errors.New("Ipset utility not found")
	errIpsetNotSupported = errors.New("Ipset utility version is not supported, requiring version >= 6.0")
)

// IPSet implements an Interface to an set.
type IPSet struct {
	Name       string
	HashType   string
	HashFamily string
	HashSize   int
	MaxElem    int
	Timeout    int

	path string
}

// Test is used to check whether the specified entry is in the set or not.
func (s *IPSet) Test(entry string) (bool, error) {
	out, err := exec.Command(
		"/bin/sh",
		"-c",
		strings.Join([]string{"sudo", s.path, "test", s.Name, entry}, " "),
	).CombinedOutput()

	if err == nil {
		reg, e := regexp.Compile("NOT")
		if e == nil && reg.MatchString(string(out)) {
			return false, nil
		} else if e == nil {
			return true, nil
		} else {
			return false, fmt.Errorf("error testing entry %s: %v", entry, e)
		}
	} else {
		return false, fmt.Errorf("error testing entry %s: %v (%s)", entry, err, out)
	}
}

type IPSets struct {
	sets []IPSet
}

func New(setName ...string) (IPSets, error) {

	path, err := exec.LookPath("ipset")
	if err != nil {
		return IPSets{}, errIpsetNotFound
	}

	s := IPSets{}
	for _, name := range setName {
		s.sets = append(s.sets, IPSet{
			Name: name,
			path: path,
		})
	}
	return s, nil
}

func (i IPSets) Test(entry string) (bool, error) {
	for _, s := range i.sets {
		s.Test(entry)
	}
	return false, nil
}
