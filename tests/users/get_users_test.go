package users

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get users list"),
		axiom.WithCaseMeta(
			axiom.WithMetaTag("smoke"),
			axiom.WithMetaStory("list users"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetUsersClientFixture(cfg)

		result, resp, err := client.GetUsers(cfg.Context.Raw)
		require.NoError(cfg.SubT, err)

		assert.Equal(cfg.SubT, http.StatusOK, resp.StatusCode())
		assert.NotEmpty(cfg.SubT, result.Users)
		assert.Greater(cfg.SubT, result.Total, 0)
	})
}
