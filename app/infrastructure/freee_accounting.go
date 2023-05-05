package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/colere-inc/seen-api/app/common/config"
)

const freeeAccountingApiEndpointUrl = "https://api.freee.co.jp/api/1"

type FreeeAccounting struct {
	Client      *http.Client
	CompanyId   string
	AccessToken string
}

func NewFreeeAccounting() *FreeeAccounting {
	return &FreeeAccounting{
		Client:      &http.Client{},
		CompanyId:   config.FreeeCompanyId,
		AccessToken: config.FreeeAccessToken,
	}
}

func (fa *FreeeAccounting) Do(method string, path string, values url.Values, body io.Reader) Response {
	// prepare request
	requestUrl := freeeAccountingApiEndpointUrl + path
	if values != nil {
		requestUrl = fmt.Sprintf("%s?%s", requestUrl, values.Encode())
	}
	req, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", fa.AccessToken))

	// request
	client := fa.Client
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

	// convert to response
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return Response{StatusCode: res.StatusCode, ResBody: resBody}
}

type Response struct {
	StatusCode int
	ResBody    []byte
}
