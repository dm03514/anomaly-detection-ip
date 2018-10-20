package blacklist

type StubBlacklist struct {
	Result Result
	Error  error
}

func (s *StubBlacklist) Test(ip string) (Result, error) {
	return s.Result, s.Error
}
