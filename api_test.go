package entapi_test

import (
	"fmt"
	"testing"

	"github.com/alex-cos/entapi"
	"github.com/stretchr/testify/assert"
)

func TestSearchEntreprise(t *testing.T) {
	t.Parallel()

	api := entapi.New()

	resp, err := api.SearchEntreprise("TF1")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	if !testing.Short() {
		fmt.Printf("resp = %+v\n", resp)
	}
}
