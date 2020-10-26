package table

type TableValue struct {
	Header []string
	Values []map[string]ColumnValue
}

type ValueType int

const (
	_ ValueType = iota
	Integer
)

func (v ValueType) String() string {
	switch v {
	case Integer:
		return "Integer"
	default:
		return "Unknown"
	}
}

type Value struct {
	Type     ValueType
	Integral int
}

type ColumnValue struct {
	Name  string
	Value Value
}
