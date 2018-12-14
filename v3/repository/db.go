package repository

import (
	"github.com/alamin-mahamud/rest-api-go/v3/entity"
)

// IDb is the base type of databse
type IDb interface {
	Open()
	Query() (entity.Account, error)
	Seed()
}
