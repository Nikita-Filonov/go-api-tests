package products

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get products list"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("list products"),
			axiom.WithMetaTag("smoke"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetProductsClientFixture(cfg)

		result, resp, err := client.GetProducts(cfg.Context.Raw)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusOK, resp.StatusCode())
		assert.NotEmpty(cfg.SubT, result.Products)
		assert.Greater(cfg.SubT, result.Total, 0)
	})
}
