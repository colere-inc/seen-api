package adapter

import (
	"io/ioutil"
	"net/http"
	"net/url"

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
		defer res.Body.Close()

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			// エラー処理
			panic(err)
		}
		return c.JSON(res.StatusCode, response{Text: string(resBody)})
	}
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Text: "Hello, world!"})
}

type response struct {
	Text string `json:"text"`
}
