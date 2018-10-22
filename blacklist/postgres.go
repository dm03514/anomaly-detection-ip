package blacklist

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Test(ip string) (Result, error) {
	q := `
SELECT address, provider FROM ipsets
`
	fmt.Println(q)
	rows, err := p.db.Query(q)
	if err != nil {
		return Result{}, nil
	}
	defer rows.Close()
	result := Result{}
	for rows.Next() {
		result.Found = true

		provider := Provider{}

		if err := rows.Scan(&result.CIDR, &provider.Name); err != nil {
			return result, err
		}

		result.Providers = append(result.Providers, provider)
		/*
			var (
				id   int64
				name string
			)
			fmt.Printf("id %d name is %s\n", id, name)
		*/
	}
	if err := rows.Err(); err != nil {
		return Result{}, err
	}

	return result, nil
}

func NewPostges(dbConnectionString string) (*Postgres, error) {
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}
