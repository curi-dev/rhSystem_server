package valueobjects

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AccessKey struct {
	Id        uuid.UUID
	Value     string
	Candidate string
	CreatedAt time.Time
}

func (self *AccessKey) IsValid() bool {
	elapsedTime := time.Since(self.CreatedAt).Minutes()

	fmt.Println("elapsedTime: ", elapsedTime-(3*60))

	return elapsedTime-(3*60) <= 5 // timezone compensation
}
