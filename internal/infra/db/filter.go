package db

type ComparisonOperator byte

const (
	_ ComparisonOperator = iota
	Equal
	LesserThan
	LesserThanOrEqual
	GreaterThan
	GreaterThanOrEqual

	InValues
)

type Filter struct {
	FieldName string
	Value     interface{}
	Operator  ComparisonOperator
}
type LogicalOperator byte

const (
	_ LogicalOperator = iota
	And
	Or
	Not
)

func Compare(key string, value interface{}, operator ComparisonOperator) Filter {
	return Filter{
		FieldName: key,
		Value:     value,
		Operator:  operator,
	}
}
