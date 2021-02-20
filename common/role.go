package common

import "database/sql/driver"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser Role= "user"
)

func (r *Role) Scan(value interface{}) error {
	*r = Role(value.([]byte))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) String() string{
	role := *r
	if role == RoleAdmin {
		return "admin"
	} else {
		return "user"
	}
}