package npndatabase

import (
	"fmt"
	"strings"
)

const whereSpaces = " where "

// Creates a SQL insert statement, potentially for multiple rows
func SQLInsert(table string, columns []string, rows int) string {
	if rows <= 0 {
		rows = 1
	}
	colString := strings.Join(columns, ", ")
	var placeholders []string
	for i := 0; i < rows; i++ {
		var ph []string
		for idx := range columns {
			ph = append(ph, fmt.Sprintf("$%v", (i*len(columns))+idx+1))
		}
		placeholders = append(placeholders, "("+strings.Join(ph, ", ")+")")
	}
	return fmt.Sprintf("insert into %v (%v) values %v", table, colString, strings.Join(placeholders, ", "))
}

// Creates a SQL select statement
func SQLSelect(columns string, tables string, where string, orderBy string, limit int, offset int) string {
	if len(columns) == 0 {
		columns = "*"
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	orderByClause := ""
	if len(orderBy) > 0 {
		orderByClause = " order by " + orderBy
	}

	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" limit %v", limit)
	}

	offsetClause := ""
	if offset > 0 {
		offsetClause = fmt.Sprintf(" offset %v", offset)
	}

	return "select " + columns + " from " + tables + whereClause + orderByClause + limitClause + offsetClause
}

// Creates a SQL select statement with a simple where clause
func SQLSelectSimple(columns string, tables string, where ...string) string {
	return SQLSelect(columns, tables, strings.Join(where, " and "), "", 0, 0)
}

// Creates a SQL update statement
func SQLUpdate(table string, columns []string, where string) string {
	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	stmts := make([]string, 0, len(columns))
	for i, col := range columns {
		s := fmt.Sprintf("%v = $%v", col, i+1)
		stmts = append(stmts, s)
	}
	return fmt.Sprintf("update %v set %v%v", table, strings.Join(stmts, ", "), whereClause)
}

// Creates a SQL delete statement
func SQLDelete(table string, where string) string {
	if len(strings.TrimSpace(where)) == 0 {
		return "attempt to delete from [" + table + "] with empty where clause"
	}
	return "delete from " + table + whereSpaces + where
}
