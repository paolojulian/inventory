package id

import (
	"github.com/google/uuid"
)

type UUID string

func NewUUID() string {
	return uuid.New().String()
}
