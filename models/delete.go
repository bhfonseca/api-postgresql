package models

import "api-postgesql/db"

func Delete(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos where id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}