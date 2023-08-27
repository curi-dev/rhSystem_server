package valueobjects

import "github.com/google/uuid"

type ValidSlot struct {
	Id      uuid.UUID
	Weekday int
	Slot    int
}
