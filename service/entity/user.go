package entity

import "time"

type User struct {
    ID        uint64    `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Salt      string    `json:"salt"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Banned    bool      `json:"banned"`
}

