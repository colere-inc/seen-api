package model

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/colere-inc/seen-api/app/common/config"
	"google.golang.org/api/iterator"
	"net/url"
	"strconv"

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

func (p InfraPartnerRepository) GetById(id int64) (*model.Partner, error) {
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

func (p InfraPartnerRepository) Add(name string) (*model.Partner, error) {
	// request body
	body := postRequestBody{CompanyID: p.FreeeAccounting.CompanyId, Name: name}
	requestBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	// request
	res := p.FreeeAccounting.Do("POST", "/partners", nil, bytes.NewBuffer(requestBody))

	// unmarshal
	var partnerRes partnerResponse
	err = json.Unmarshal(res.ResBody, &partnerRes)
	if err != nil {
		panic(err)
	}
	return &partnerRes.Partner, err
}

func (p InfraPartnerRepository) GetByName(name string) (*model.Partner, error) {
	ctx := context.Background()
	partnerID := p.searchFirestoreByName(name, ctx)
	return p.GetById(partnerID)
}

func (p InfraPartnerRepository) searchFirestoreByName(name string, ctx context.Context) int64 {
	query := p.DB.Collection(config.FreeeCompaniesCollectionId).
		Doc(config.FreeeCompanyId).
		Collection(config.FreeePartnersSubCollectionId).
		Where("name", "==", name)
	var partnerID string
	it := query.Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("documents iterator: %v", err))
		}
		if partnerID != "" {
			panic(fmt.Sprintf("multiple documents are found (name = %s)", name))
		}
		partnerID = doc.Ref.ID
	}

	if partnerID == "" {
		panic(fmt.Sprintf("not found (name = %s)", name))
	}

	partnerIntID, err := strconv.ParseInt(partnerID, 10, 64)
	if err != nil {
		panic(err)
	}
	return partnerIntID
}

type postRequestBody struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
}

type partnerResponse struct {
	Partner model.Partner `json:"partner"`
}
