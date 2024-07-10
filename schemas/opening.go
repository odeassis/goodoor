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
}

type OpeningResponse struct {
	ID        uint           `json: "id"`
	Title     string         `json: "title"`
	Role      string         `json: "role"`
	Company   string         `json: "company"`
	Location  string         `json: "location"`
	Remote    bool           `json: "remote"`
	Link      string         `json: "link"`
	CreatedAt time.Time      `json: "createdAt"`
	UpdatedAt time.Time      `json: "updatedAt"`
	DeletedAt gorm.DeletedAt `json: "deletedAt, omitempty"`
}
