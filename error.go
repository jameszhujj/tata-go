package tata

import (
	"errors"
	"fmt"
)

var (
	errEmptyResponse = errors.New("empty http response")
)

type Error struct {
	StatusCode int
}

func (err Error) Error() string {
	return fmt.Sprintf("Error calling Tata-Communications API\nstatus: %d", err.StatusCode)
}
