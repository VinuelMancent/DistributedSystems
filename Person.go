package main

import "github.com/google/uuid"

type Person struct {
	id            uuid.UUID
	isScrumMaster bool
}
