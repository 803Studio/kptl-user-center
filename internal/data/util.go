package data

import (
	"fmt"
	"github.com/803Studio/kptl-user-center/internal/interfaces"
	"strings"
)

type SelectByApi[V string | int, RES interfaces.SqlModel] func(limit int, value V) ([]RES, error)

func fmtSqlStr(table, key string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", table, key)
}

func createSelectByApi[V string | int, RES interfaces.SqlModel](table, key string, newFn func() RES) SelectByApi[V, RES] {
	sqlStrRaw := fmtSqlStr(table, key)

	return func(limit int, value V) ([]RES, error) {
		sqlStr := sqlStrRaw
		if limit != 0 {
			sqlStr = sqlStr + fmt.Sprintf(" LIMIT %d", limit)
		}

		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			return nil, err
		}

		rows, err := stmt.Query(value)
		if err != nil {
			return nil, err
		}

		result := make([]RES, 0)
		for rows.Next() {
			instance := newFn()
			err = rows.Scan(instance.PtrVec()...)
			if err != nil {
				return nil, err
			}
			result = append(result, instance)
		}

		err = rows.Close()
		if err != nil {
			return nil, err
		}

		err = stmt.Close()
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

type InsertApi[T interfaces.SqlModel] func(value T) (int64, error)

func fmtInsertSql(table string, value interfaces.SqlModel) string {
	keys := value.Keys()
	//result be like
	//INSERT INTO table(v1, v2, v3, ...) VALUES (?, ?, ?, ?, ...)
	return fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES (%s)",
		table,
		strings.Join(keys, ","),
		strings.Repeat("?,", len(keys))[:(len(keys)<<1)-1],
	)
}

func createInsertApi[T interfaces.SqlModel](table string) InsertApi[T] {
	var sqlCache string

	return func(value T) (int64, error) {
		tx, err := db.Begin()
		if err != nil {
			return 0, err
		}

		if sqlCache == "" {
			sqlCache = fmtInsertSql(table, value)
		}

		result, err := tx.Exec(sqlCache, value.Values()...)
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				return 0, rollBackErr
			}
			return 0, err
		}

		err = tx.Commit()
		if err != nil {
			return 0, err
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}

		return lastId, nil
	}
}
