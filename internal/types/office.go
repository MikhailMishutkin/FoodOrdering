package types

import "time"

type Office struct {
	Uuid      int
	Name      string
	Address   string
	CreatedAt time.Time
}
