package core

import (
	"fmt"
	"strings"
)

type DbQuery struct {
	QueryWheres []DbQueryWhere `json:"query_wheres"`
	PageSize    int            `json:"page_size"`
	PageNumber  int            `json:"page_number"`
}

func NewDbQuery(wheres []DbQueryWhere, ps int, pn int) *DbQuery {
	return &DbQuery{
		QueryWheres: wheres,
		PageSize:    ps,
		PageNumber:  pn,
	}
}

func (q *DbQuery) GetWhereClause() (whereClause string, values []interface{}) {
	whereClause = " 1=1 "
	if len(q.QueryWheres) < 1 {
		return whereClause, values
	}
	var sb strings.Builder

	sb.WriteString("1=1")
	for _, where := range q.QueryWheres {
		var subClause strings.Builder
		for _, filter := range where.QueryFilters {
			fieldName := filter.FieldName
			var op string
			switch filter.FilterOperation {
			case "EQ":
				op = " = ? "
				values = append(values, filter.FilterValues[0])
				break
			case "NEQ":
				op = " <> ? "
				values = append(values, filter.FilterValues[0])
				break
			case "LT":
				op = " < ? "
				values = append(values, filter.FilterValues[0])
				break
			case "LTE":
				op = " <= ? "
				values = append(values, filter.FilterValues[0])
				break
			case "GT":
				op = " > ? "
				values = append(values, filter.FilterValues[0])
				break
			case "GTE":
				op = " >= ? "
				values = append(values, filter.FilterValues[0])
				break
			case "LIKE":
				op = " LIKE ? "
				values = append(values, "%"+fmt.Sprint(filter.FilterValues[0])+"%")
				break
			case "IN":
				op = " IN ? "
				values = append(values, filter.FilterValues)
				break
			case "BTN":
				op = " BETWEEN ? AND ? "
				values = append(values, filter.FilterValues[0], filter.FilterValues[1])
				break
			}
			subClause.WriteString(fmt.Sprintf(" %s %s AND ", fieldName, op))
		}
		subWhere := subClause.String()
		subWhere = strings.Trim(subWhere, "AND ")
		sb.WriteString(subWhere)
		sb.WriteString(fmt.Sprintf(" %s ", where.Connecter))
	}
	whereClause = whereClause + sb.String()
	whereClause = strings.Trim(whereClause, "AND ")
	whereClause = strings.Trim(whereClause, "OR ")
	return whereClause, values
}