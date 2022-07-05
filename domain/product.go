package domain

import "context"

type Attribute struct {
	Label  string   `json:"label"`
	Values []string `json:"values"`
}

type Product struct {
	ID         int64
	UUID       string      `json:"uuid"`
	Rate       float32     `json:"rate"`
	Pictures   []string    `json:"pictures"`
	Name       string      `json:"name"`
	Detail     string      `json:"detail"`
	Favorite   bool        `json:"favorite"`
	Attributes []Attribute `json:"attributes"`
}

type ProductUseCase interface {
	Get(ctx context.Context, uuid string) (*Product, error)
}

type ProductRepository interface {
	GetByUUID(ctx context.Context, uuid string) (*Product, error)
}
