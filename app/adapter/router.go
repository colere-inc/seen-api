package adapter

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/colere-inc/seen-api/app/domain/model"

	"github.com/colere-inc/seen-api/app/common/config"

	"github.com/labstack/echo/v4"
)

const freeeAccountingApiEndpointUrl = "https://api.freee.co.jp/api/1"

func NewRouter(e *echo.Echo) {
	e.GET("/accounting/partners", ListPartners())
	e.POST("/accounting/partners", hello)
}

func ListPartners() echo.HandlerFunc {
	return func(c echo.Context) error {
		values := url.Values{}
		values.Set("company_id", config.FreeeCompanyId)
		req, err := http.NewRequest("GET", freeeAccountingApiEndpointUrl+"/partners?"+values.Encode(), nil)
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.FreeeAccessToken)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(res.Body)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		var partners model.Partners
		err = json.Unmarshal(resBody, &partners)
		if err != nil {
			panic(err)
		}
		return c.JSON(res.StatusCode, partners)
	}
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Text: "Hello, world!"})
}

type response struct {
	Text string `json:"text"`
}
