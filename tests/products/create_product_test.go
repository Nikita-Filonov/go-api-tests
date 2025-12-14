package products

import (
	"net/http"
	"testing"

	"go-api-tests/clients/products"
	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("create product"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("regression"),
			axiom.WithMetaStory("create product"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetProductsClientFixture(cfg)

		req := products.CreateProductRequest{
			Title:       gofakeit.ProductName(),
			Description: gofakeit.Sentence(8),
			Price:       gofakeit.Price(10, 1000),
			Stock:       gofakeit.Number(1, 100),
			Category:    "electronics",
		}

		product, resp, err := client.CreateProduct(cfg.Context.Raw, req)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusCreated, resp.StatusCode())
		assert.Equal(cfg.SubT, req.Title, product.Title)
		assert.Equal(cfg.SubT, req.Price, product.Price)
		assert.Equal(cfg.SubT, req.Stock, product.Stock)
	})
}
