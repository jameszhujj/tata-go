package tata

import (
	"encoding/base64"
)

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) GetBasicAuthString() string {
	s := c.Username + ":" + c.Password
	e := base64.StdEncoding.EncodeToString([]byte(s))
	a := "Basic " + e

	return a
}
