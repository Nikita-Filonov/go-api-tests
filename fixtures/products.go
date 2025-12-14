package fixtures

import (
	"go-api-tests/clients/products"
	"go-api-tests/http"

	"github.com/Nikita-Filonov/axiom"
)

func SetProductsClientFixture(cfg *axiom.Config) (any, func(), error) {
	config := GetConfigFixture(cfg)
	logger := GetLoggerFixture(cfg)

	client := products.NewClient(
		config.HTTP,
		http.NewClientLogger(logger),
		cfg,
	)

	return client, nil, nil
}

func GetProductsClientFixture(cfg *axiom.Config) *products.Client {
	return axiom.GetFixture[*products.Client](cfg, "products")
}
