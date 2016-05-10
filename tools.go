package main

type mapcolumn map[string]ColumnType

func (m mapcolumn) Keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}
