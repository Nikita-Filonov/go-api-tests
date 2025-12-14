package users

import (
	"net/http"
	"testing"

	"go-api-tests/clients/users"
	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("create user"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("create user"),
			axiom.WithMetaTag("regression"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetUsersClientFixture(cfg)

		req := users.CreateUserRequest{
			Email:     gofakeit.Email(),
			Username:  gofakeit.Username(),
			LastName:  gofakeit.LastName(),
			FirstName: gofakeit.FirstName(),
		}
		user, resp, err := client.CreateUser(cfg.Context.Raw, req)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusCreated, resp.StatusCode())
		assert.Equal(cfg.SubT, req.Email, user.Email)
		assert.Equal(cfg.SubT, req.Username, user.Username)
		assert.Equal(cfg.SubT, req.LastName, user.LastName)
		assert.Equal(cfg.SubT, req.FirstName, user.FirstName)
	})
}
