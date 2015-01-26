package models

import (
	"github.com/go-gorp/gorp"
	"time"
)

type User struct {
	Id      int64
	Created int64
	Updated int64
	Email   string
}

func (i *User) PreInsert(s gorp.SqlExecutor) error {
	i.Created = time.Now().UnixNano()
	i.Updated = i.Created
	return nil
}
