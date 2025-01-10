package pkg

import (
	"database/sql/driver"
	"encoding/json"
)

type Meta map[string]any

// Scan implements the database/sql.Scanner interface.
func (m *Meta) Scan(value any) error {
	s, ok := value.(string)
	if !ok {
		return nil
	}

	return json.Unmarshal([]byte(s), m)
}

// Value implements the database/sql/driver.Valuer interface.
func (m Meta) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Meta Meta   `json:"meta"`
}

type Order struct {
	Id     int  `json:"id"`
	UserId int  `json:"user_id"`
	Meta   Meta `json:"meta"`
}

type CType string

const (
	Text    CType = "TEXT"
	Number  CType = "DOUBLE"
	Boolean CType = "BOOLEAN"
	Date    CType = "DATE"
)

// CustomColumns stores the columns that are set to meta
type CustomColumns struct {
	Id         int    `json:"id"`
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	ColumnType CType  `json:"column_type"`
	IsNullable bool   `json:"is_nullable"`
}
