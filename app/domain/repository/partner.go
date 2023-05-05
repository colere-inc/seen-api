package repository

import (
	"github.com/colere-inc/seen-api/app/domain/model"
)

type PartnerRepository interface {
	GetPartnerById(id int64) (*model.Partner, error)
	GetPartnerByName(name string) (*model.Partner, error)
}
