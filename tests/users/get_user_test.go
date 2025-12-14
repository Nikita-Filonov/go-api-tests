package users

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get user by id"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("get user"),
			axiom.WithMetaTag("regression"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetUsersClientFixture(cfg)

		user, resp, err := client.GetUser(cfg.Context.Raw, 1)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusOK, resp.StatusCode())
		assert.Equal(cfg.SubT, 1, user.ID)
		assert.NotEmpty(cfg.SubT, user.Username)
		assert.NotEmpty(cfg.SubT, user.Email)
	})
}
