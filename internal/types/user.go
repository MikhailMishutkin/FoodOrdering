package types

import (
	"time"
)

type User struct {
	Uuid       int
	Name       string
	OfficeUuid int
	OfficeName string
	CreatedAt  time.Time
}
