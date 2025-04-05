package util

import "github.com/google/uuid"

func NewUuidV7() string {
	uid, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return uid.String()
}
