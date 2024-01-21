package entity

import "github.com/google/uuid"

type UniqueID string

func NewUniqueID() UniqueID {
	return UniqueID(uuid.New().String())
}

func UniqueIDFromValue(value string) UniqueID {
	id, err := uuid.Parse(value)
	if err != nil {
		return ""
	}

	return UniqueID(id.String())
}

func ValueFromUniqueID(value UniqueID) string {
	return string(value)
}
