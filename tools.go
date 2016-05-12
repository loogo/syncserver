package main

type mapcolumn map[string]ColumnType

func (m *mapcolumn) Keys() (keys []string) {
	for key, val := range *m {
		if val.Alias != "" {
			keys = append(keys, key+" as "+val.Alias)
			continue
		}
		keys = append(keys, key)
	}
	return
}

func (m *mapcolumn) getByAlias(alias string) *ColumnType {
	for _, val := range *m {
		if val.Alias == alias {
			return &val
		}
	}
	return &ColumnType{}
}
