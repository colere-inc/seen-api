package adapter

import (
	"net/http"

	"github.com/colere-inc/seen-api/app/domain/repository"
	"github.com/labstack/echo/v4"
)

type PartnerController struct {
	PartnerRepository repository.PartnerRepository
}

func NewPartnerController(pr repository.PartnerRepository) *PartnerController {
	return &PartnerController{PartnerRepository: pr}
}

func (pc *PartnerController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		partner, err := pc.PartnerRepository.GetByName(name)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, partner)
	}
}

func (pc *PartnerController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var partnerID int64
		echo.PathParamsBinder(c).Int64("partnerID", &partnerID)
		partner, err := pc.PartnerRepository.GetById(partnerID)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, partner)
	}
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Text: "Hello, world!"})
}

type response struct {
	Text string `json:"text"`
}
