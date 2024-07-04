package domain

import (
    "time"
)

type Gender uint8

const (
    Male   Gender = 1
    Female Gender = 2
    Other  Gender = 3
)

type User struct {
    Id        int
    FirstName string
    LastName  string
    Birthday  time.Time
    Gender    Gender
    Code      string
    CreatedAt time.Time
    UpdatedAt time.Time
}
