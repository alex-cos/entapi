package entapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/alex-cos/restc"
)

const (
	APIURL = "https://recherche-entreprises.api.gouv.fr"
)

type GeoAPI struct {
	client  *restc.Client
	limit   int32
	perPage int32
}

func New() *GeoAPI {
	return NewWithClient(http.DefaultClient)
}

func NewWithClient(httpClient *http.Client) *GeoAPI {
	return NewWithClientTimeout(httpClient, restc.DefaultTimeout)
}

func NewWithTimeout(timeout time.Duration) *GeoAPI {
	return NewWithClientTimeout(http.DefaultClient, timeout)
}

func NewWithClientTimeout(httpClient *http.Client, timeout time.Duration) *GeoAPI {
	return &GeoAPI{
		client:  restc.NewWithClientTimeout(APIURL, httpClient, timeout),
		limit:   10,
		perPage: 10,
	}
}

func (api *GeoAPI) SearchEntreprise(text string) (*EntreprisesResponse, error) {
	params := url.Values{
		"q":                              {text},
		"limite_matching_etablissements": {strconv.Itoa(int(api.limit))},
		"minimal":                        {"true"},
		"include":                        {"siege,complements,dirigeants"},
		"page":                           {"1"},
		"per_page":                       {strconv.Itoa(int(api.perPage))},
	}

	req := restc.Get("search").
		SetHeader("Accept", "application/json").
		SetQueryParamsFromValues(params).
		SetResponseType(new(EntreprisesResponse))

	resp, err := api.client.Execute(req)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error server returned an http error: %s", resp.Status())
	}

	entreprises, ok := resp.Content().(*EntreprisesResponse)
	if !ok {
		return nil, errors.New("wrong expected type")
	}

	return entreprises, nil
}
