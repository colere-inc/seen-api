package repository

import (
	"github.com/colere-inc/seen-api/app/domain/model"
)

type PartnerRepository interface {
	GetById(id int64) (*model.Partner, error)
	GetByName(name string) (*model.Partner, error)

	Add(name string) (*model.Partner, error)
}
