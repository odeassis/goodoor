package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Title    string
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint      `json: "id"`
	Title     string    `json: "title"`
	Role      string    `json: "role"`
	Company   string    `json: "company"`
	Location  string    `json: "location"`
	Remote    bool      `json: "remote"`
	Link      string    `json: "link"`
	Salary    int64     `json: "salary"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
	DeletedAt time.Time `json: "deletedAt, omitempty"`
}
