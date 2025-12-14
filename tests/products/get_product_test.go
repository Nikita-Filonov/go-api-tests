package products

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProductByID(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get product by id"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("get product"),
			axiom.WithMetaTag("regression"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetProductsClientFixture(cfg)

		product, resp, err := client.GetProduct(cfg.Context.Raw, 1)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusOK, resp.StatusCode())
		assert.Equal(cfg.SubT, 1, product.ID)
		assert.NotEmpty(cfg.SubT, product.Title)
		assert.Greater(cfg.SubT, product.Price, 0.0)
	})
}
