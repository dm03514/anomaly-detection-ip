package ipsets

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

var db *sql.DB

func init() {
	db, err := sql.Open("postgres", "postgresql://root:root@localhost/ipsets?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func BenchmarkQueryNonExistent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := db.Query("select * from ipsets where address >>= '1.1.1.1'")
		if err != nil {
			b.Error(err)
		}
	}
}
