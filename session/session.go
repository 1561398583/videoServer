package session

import "sync"

type SimpleSession struct {
	UserID string
	TTL int64
}

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}
