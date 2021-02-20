package common

import "database/sql/driver"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser Role= "user"
)

func (e *Role) Scan(value interface{}) error {
	*e = Role(value.([]byte))
	return nil
}

func (e Role) Value() (driver.Value, error) {
	return string(e), nil
}