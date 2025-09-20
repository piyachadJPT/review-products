package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"column:password_hash;not null" json:"-"`
	Name      *string
	Role      string `gorm:"default:user"`
	Avatar    *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SKU         *string   `gorm:"unique"`
	Name        string    `gorm:"not null"`
	Description *string   `gorm:"type:text"`
	Price       float64   `gorm:"type:numeric(12,2);default:0"`
	Stock       int       `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Images []ProductImage `gorm:"foreignKey:ProductID"`
	Review []Review       `gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductID string    `gorm:"not null"`
	URL       string
	Alt       *string
	Position  int
}

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductID string    `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	Title     *string
	Body      string `gorm:"type:text;not null"`
	Rating    int    `gorm:"not null;check:rating>=1 AND rating<=5"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
