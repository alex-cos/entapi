package entapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	APIURL = "https://recherche-entreprises.api.gouv.fr"
)

type GeoAPI struct {
	client  *http.Client
	timeout time.Duration
	limit   int32
	perPage int32
}

func New() *GeoAPI {
	return NewWithClient(http.DefaultClient)
}

func NewWithClient(httpClient *http.Client) *GeoAPI {
	return NewWithClientTimeout(httpClient, httpClient.Timeout)
}

func NewWithTimeout(timeout time.Duration) *GeoAPI {
	return NewWithClientTimeout(http.DefaultClient, timeout)
}

func NewWithClientTimeout(httpClient *http.Client, timeout time.Duration) *GeoAPI {
	return &GeoAPI{
		client:  httpClient,
		timeout: timeout,
		limit:   10,
		perPage: 10,
	}
}

func (api *GeoAPI) SearchEntreprise(text string) (*EntreprisesResponse, error) {
	var entreprises *EntreprisesResponse

	params := url.Values{
		"q":                              {text},
		"limite_matching_etablissements": {strconv.Itoa(int(api.limit))},
		"minimal":                        {"true"},
		"include":                        {"siege,complements,dirigeants"},
		"page":                           {"1"},
		"per_page":                       {strconv.Itoa(int(api.perPage))},
	}
	reqURL := fmt.Sprintf("%s/search?%s", APIURL, params.Encode())

	body, err := api.request(reqURL)
	if err != nil {
		return entreprises, err
	}

	err = json.Unmarshal(body, &entreprises)
	if err != nil {
		return entreprises, err
	}

	return entreprises, nil
}

// Unexported functions

func (api *GeoAPI) request(reqURL string) ([]byte, error) {
	ctx, cancel := api.getContext()
	if cancel != nil {
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, ErrNewRequest(err)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, ErrDoRequest(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadBody(http.MethodGet, reqURL, err)
	}

	if resp.StatusCode >= 400 {
		return nil, ErrServerHTTPError(resp.StatusCode)
	}

	return body, nil
}

func (api *GeoAPI) getContext() (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc

	ctx := context.Background()
	if api.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), api.timeout)
	}

	return ctx, cancel
}
