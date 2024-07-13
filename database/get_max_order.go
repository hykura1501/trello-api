package database

import "fmt"

func GetMaxOrder(sql *SQL, tableName, field, value string) (int, error) {
	sqlStatement := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s=\"%s\"", tableName, field, value)
	var result int
	if err := sql.Db.Get(&result, sqlStatement); err != nil {
		return 0, err
	}
	return result, nil
}
