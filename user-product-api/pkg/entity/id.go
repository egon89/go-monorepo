package entity

import "github.com/google/uuid"

type Id = uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func ParseId(stringId string) (Id, error) {
	id, err := uuid.Parse(stringId)

	return Id(id), err
}
