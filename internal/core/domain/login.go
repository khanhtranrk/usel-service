package domain

import (
    "time"
)

type Login struct {
    ID           int64
    RefreshToken string    
    RequestIP    string    
    Browser      int64     
    AccountID    int64     
    CreatedAt    time.Time 
    UpdatedAt    time.Time 
}
