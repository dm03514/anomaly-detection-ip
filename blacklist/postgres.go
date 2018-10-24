package blacklist

import (
	"database/sql"
	"time"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Test(ip string) (Result, error) {
	q := `
SELECT address, provider FROM ipsets WHERE address >>= $1 
`
	// fmt.Println(q)
	rows, err := p.db.Query(q, ip)
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
	}

	if err := rows.Err(); err != nil {
		return Result{}, err
	}

	return result, nil
}

func NewPostges(dbConnectionString string) (*Postgres, error) {
	db, err := sql.Open("postgres", dbConnectionString)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(time.Millisecond * 500)
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
