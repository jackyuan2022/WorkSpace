package core

type DbQueryWhere struct {
	QueryFilters []DbQueryFilter `json:"query_filters"`
	Connecter    string          `json:"connecter"`
}

func NewDbQueryWhere(filters []DbQueryFilter, connector string) DbQueryWhere {
	return DbQueryWhere{
		QueryFilters: filters,
		Connecter:    connector,
	}
}
