package utils

import uuid "github.com/iris-contrib/go.uuid"

func Uuid() string {
	uid, _ := uuid.NewV4()
	rst := uid.String()
	return rst
}
