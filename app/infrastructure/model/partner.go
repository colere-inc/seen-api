package model

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/colere-inc/seen-api/app/infrastructure"
)

type InfraPartnerRepository struct {
	DB              *infrastructure.DB
	FreeeAccounting *infrastructure.FreeeAccounting
}

func NewPartnerRepository(
	db *infrastructure.DB,
	freeeAccounting *infrastructure.FreeeAccounting,
) repository.PartnerRepository {
	return InfraPartnerRepository{
		DB:              db,
		FreeeAccounting: freeeAccounting,
	}
}

func (p InfraPartnerRepository) GetPartnerById(id int64) (*model.Partner, error) {
	// request
	values := url.Values{}
	values.Set("company_id", p.FreeeAccounting.CompanyId)
	res := p.FreeeAccounting.Do("GET", fmt.Sprintf("/partners/%d", id), values, nil)

	// unmarshal
	var partnerRes partnerResponse
	err := json.Unmarshal(res.ResBody, &partnerRes)
	if err != nil {
		panic(err)
	}
	return &partnerRes.Partner, err
}

func (p InfraPartnerRepository) GetPartnerByName(name string) (*model.Partner, error) {
	var id int64 = 0 // TODO
	return p.GetPartnerById(id)
}

type partnerResponse struct {
	Partner model.Partner `json:"partner"`
}
