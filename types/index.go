package types

import (
	"github.com/google/uuid"
)

type StatusNames string

const (
	TODO        StatusNames = "TODO"
	IN_PROGRESS StatusNames = "IN-PROGRESS"
	DONE        StatusNames = "DONE"
)

type TaskJSONFormat struct {
	ID         uuid.UUID   `json:"id"`
	Descripton string      `json:"description"`
	Status     StatusNames `json:"status"`
	CreatedAt  string      `json:"createdAt"`
	UpdatedAt  string      `json:"updatedAt"`
}

type TasksArrayFormat struct {
	Tasks []TaskJSONFormat `json:"tasks"`
}
