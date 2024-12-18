package models

import "time"

type Client struct {
	ID        int
	Num1      string
	Num2      string
	Num3      string
	Subsidiary   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
