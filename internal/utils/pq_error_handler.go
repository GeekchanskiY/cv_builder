package utils

import "github.com/lib/pq"

func PQErrorHandler(err error) (*pq.Error, bool) {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr, true
	}
	return nil, false
}
