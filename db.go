package truc

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (e *Engine) ConnectDb(url ...string) error {
	if len(url) == 0 {
		url = append(url, e.Params.Pg)
	}
	db, err := sql.Open("postgres", url[0])
	if err != nil {
		return err
	}

	e.Db = db
	var result int
	err = e.Db.QueryRow("SELECT 1 + 1 AS result").Scan(&result)

	return err
}
