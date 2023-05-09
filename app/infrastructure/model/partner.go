package model

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/colere-inc/seen-api/app/common/config"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/colere-inc/seen-api/app/domain/model"
	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/colere-inc/seen-api/app/infrastructure"
)

type PartnerRepository struct {
	DB              *infrastructure.DB
	FreeeAccounting *infrastructure.FreeeAccounting
}

func NewPartnerRepository(
	db *infrastructure.DB,
	freeeAccounting *infrastructure.FreeeAccounting,
) repository.PartnerRepository {
	return PartnerRepository{
		DB:              db,
		FreeeAccounting: freeeAccounting,
	}
}

func (p PartnerRepository) GetById(id int64) (*model.Partner, error) {
	// request
	values := url.Values{}
	values.Set("company_id", p.FreeeAccounting.CompanyId)
	res := p.FreeeAccounting.Do(http.MethodGet, fmt.Sprintf("/partners/%d", id), values, nil)

	if res.StatusCode != http.StatusOK {
		log.Println("failed")
		panic(fmt.Sprintf("unexpected status: got %v, error: %s", res.StatusCode, string(res.ResBody)))
	}
	log.Println("success")

	// unmarshal
	var partnerRes partnerResponse
	err := json.Unmarshal(res.ResBody, &partnerRes)
	if err != nil {
		panic(err)
	}
	return &partnerRes.Partner, err
}

func (p PartnerRepository) Add(name string) (*model.Partner, error) {
	ctx := context.Background()

	// request body
	body := postRequestBody{CompanyID: p.FreeeAccounting.CompanyId, Name: name}
	requestBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	// request
	res := p.FreeeAccounting.Do(http.MethodPost, "/partners", nil, bytes.NewBuffer(requestBody))
	if res.StatusCode != http.StatusCreated {
		log.Println("failed")
		panic(fmt.Sprintf("unexpected status: got %v, error: %s", res.StatusCode, string(res.ResBody)))
	}
	log.Println("success")

	// unmarshal
	var partnerRes partnerResponse
	err = json.Unmarshal(res.ResBody, &partnerRes)
	if err != nil {
		panic(err)
	}

	// add to Firestore
	partner := partnerRes.Partner
	p.addToFirestore(ctx, partner.ID, partner.Name)

	return &partnerRes.Partner, err
}

func (p PartnerRepository) addToFirestore(
	ctx context.Context,
	partnerID int64,
	name string,
) *firestore.WriteResult {
	id := strconv.FormatInt(partnerID, 10)
	data := partnerDocData{Name: name}
	result, err := p.getCollection().Doc(id).Set(ctx, data)
	if err != nil {
		panic(err)
	}
	return result
}

func (p PartnerRepository) GetByName(name string) (*model.Partner, error) {
	ctx := context.Background()
	partnerID := p.searchFirestoreByName(ctx, name)
	return p.GetById(partnerID)
}

func (p PartnerRepository) searchFirestoreByName(ctx context.Context, name string) int64 {
	query := p.getCollection().Where("name", "==", name)
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
		panic(fmt.Sprintf("not found (name = %s, Collection = %s)", name, p.getCollection().Path))
	}

	partnerIntID, err := strconv.ParseInt(partnerID, 10, 64)
	if err != nil {
		panic(err)
	}
	return partnerIntID
}

func (p PartnerRepository) getCollection() *firestore.CollectionRef {
	return p.DB.Collection(config.FreeePartnersCollectionId)
}

type postRequestBody struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
}

// Firestore における partner のデータ
type partnerDocData struct {
	Name string `firestore:"name"`
}

type partnerResponse struct {
	Partner model.Partner `json:"partner"`
}
