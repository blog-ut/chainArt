// Package util
// Time : 2023/8/23 14:10
// Author : PTJ
// File : str
// Software: GoLand
package util

import (
	"database/sql/driver"
	"strings"
)

type Strs []string

func (m *Strs) Scan(val interface{}) error {
	s := val.([]byte)
	ss := strings.Split(string(s), "|")
	*m = ss
	return nil
}

func (m Strs) Value() (driver.Value, error) {
	str := strings.Join(m, "|")
	return str, nil
}
