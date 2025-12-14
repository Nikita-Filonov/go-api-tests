package products

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

func TestGetProductNotFound(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get non-existing product"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("negative"),
			axiom.WithMetaStory("get product"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetProductsClientFixture(cfg)

		product, resp, err := client.GetProduct(cfg.Context.Raw, 999999)

		assert.Error(cfg.SubT, err)
		assert.Nil(cfg.SubT, product)
		assert.Equal(cfg.SubT, http.StatusNotFound, resp.StatusCode())
	})
}
