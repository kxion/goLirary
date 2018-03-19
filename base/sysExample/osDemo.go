package sysExample

import (
	"os"
)

type OsDemo struct{}

func NewOsDemo() *OsDemo {
	return &OsDemo{}
}

func (o *OsDemo) ReadEnvByKey(key string) string {
	return os.Getenv(key)
}
