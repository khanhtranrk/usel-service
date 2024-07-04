package domain

import (
    "time"
)

type Status string

const (
    Active   Status = "active"
    Inactive Status = "inactive"
    Banned   Status = "banned"
)

type Account struct {
    ID            int64
    Email         string
    PasswordDigest string
    Status        Status
    UserID        int64 
    Deleted       bool
    DeletedAt     *time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
