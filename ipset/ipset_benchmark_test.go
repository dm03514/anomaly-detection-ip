package ipset

import "testing"

/*jk
func BenchmarkIPSet_Test5(b *testing.B) {
	s, err := New(
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
)
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		_, err := s.Test("1.1.1.1")
		if err != nil {
			b.Error(err)
		}
	}
}
*/

func BenchmarkIPSet_Test24(b *testing.B) {
	s, err := New(
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
		"fullbogons",
		"dshield",
		"spamhaus_drop",
		"spamhaus_edrop",
	)
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		s.Test("1.1.1.1")
	}
}