package database

import "fmt"

func GetMaxOrder(sql *SQL, tableName, field, value string) (int, error) {
	orderField := string(tableName[:len(tableName)-1]) + "_order"
	sqlStatement := fmt.Sprintf("SELECT COALESCE(MAX(%s), 0) FROM %s WHERE %s=\"%s\"", orderField, tableName, field, value)
	var result int
	if err := sql.Db.Get(&result, sqlStatement); err != nil {
		return 0, err
	}
	return result, nil
}
