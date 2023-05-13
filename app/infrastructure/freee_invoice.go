package infrastructure

import (
	"fmt"
	"github.com/colere-inc/seen-api/app/common/config"
	"io"
	"net/http"
	"net/url"
)

const freeeInvoiceApiEndpointUrl = "https://api.freee.co.jp/iv"

type FreeeInvoice struct {
	Client      *http.Client
	CompanyId   string
	AccessToken string
}

func NewFreeeInvoice() *FreeeInvoice {
	return &FreeeInvoice{
		Client:      &http.Client{},
		CompanyId:   config.FreeeCompanyId,
		AccessToken: config.GetFreeeAccessToken(),
	}
}

func (fa *FreeeInvoice) Do(method string, path string, values url.Values, body io.Reader) Response {
	// prepare request
	requestUrl := freeeInvoiceApiEndpointUrl + path
	if values != nil {
		requestUrl = fmt.Sprintf("%s?%s", requestUrl, values.Encode())
	}
	req, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		panic(fmt.Sprintf("failed to create new request: %v", err))
	}
	req.Header.Add("accept", "application/json")
	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", fa.AccessToken))

	// request
	client := fa.Client
	res, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("failed to send request: %v", err))
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
