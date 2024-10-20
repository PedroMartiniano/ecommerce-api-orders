package vo

import "github.com/google/uuid"

type UUID struct {
	value string
}

func NewUUID(value string) UUID {
	if value == "" {
		id, _ := uuid.NewV7()
		value = id.String()
	}

	return UUID{
		value: value,
	}
}
