package entity


// Document represent single intro js tour, or other information
// contextually targeted to defined group of users
type Document struct {
 	Id string
	Payload string
	CreatedAt int64
	UpdatedAt int64

	RuleGroups []*RuleGroup
	CombineWithOR bool
}


type FieldType string

const (
	UserId FieldType = "user_id"
	URL FieldType = "url"
	TimeOnPage FieldType = "time_on_page"
	Custom FieldType = "custom"
)

type Operator string

const (
	EQ Operator = "eq"
	NEQ Operator = "neq"
	GT Operator = "gt"
	GTE Operator = "gte"
	LT Operator = "lt"
	LTE Operator = "lte"
	LIKE Operator = "like" // sql like function style
	BETWEEN Operator = "between"
	ANY Operator = "any"
)

// Rule single condition
// db represantation rule_id, fk_doc_id, field_name, op, values
type Rule struct {
	Field FieldType
	CustomFieldName string
	Operator Operator
	Values []string
}


type RuleGroup struct {
	Rules []*Rule
	CombineWithOR bool // if set true all rules combined with OR operator otherwise with AND
}

// UserData holds all context of user to be used for listing matching documents
type UserData struct {
	Id string
	URL string
	TimeOnPage int64 // in seconds
	Custom map[string]string
}
