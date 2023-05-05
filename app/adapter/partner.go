package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/colere-inc/seen-api/app/domain/model"

	"github.com/colere-inc/seen-api/app/infrastructure"
	"github.com/labstack/echo/v4"
)

func get(fa *infrastructure.FreeeAccounting) echo.HandlerFunc {
	return func(c echo.Context) error {
		var partnerID string
		echo.PathParamsBinder(c).String("partnerID", &partnerID)
		return getById(c, fa, partnerID)
	}
}

func getById(c echo.Context, fa *infrastructure.FreeeAccounting, partnerID string) error {
	// request
	values := url.Values{}
	values.Set("company_id", fa.CompanyId)
	res := fa.Do("GET", fmt.Sprintf("/partners/%s", partnerID), values, nil)

	// unmarshal
	var partnerRes partnerResponse
	err := json.Unmarshal(res.ResBody, &partnerRes)
	if err != nil {
		panic(err)
	}
	return c.JSON(res.StatusCode, partnerRes.Partner)
}

type partnerResponse struct {
	Partner model.Partner `json:"partner"`
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Text: "Hello, world!"})
}

type response struct {
	Text string `json:"text"`
}
