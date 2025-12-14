package fixtures

import (
	"go-api-tests/clients/users"
	"go-api-tests/http"

	"github.com/Nikita-Filonov/axiom"
)

func SetUsersClientFixture(cfg *axiom.Config) (any, func(), error) {
	config := GetConfigFixture(cfg)
	logger := GetLoggerFixture(cfg)

	client := users.NewClient(
		config.HTTP,
		http.NewClientLogger(logger),
		cfg,
	)

	return client, nil, nil
}

func GetUsersClientFixture(cfg *axiom.Config) *users.Client {
	return axiom.GetFixture[*users.Client](cfg, "users")
}
