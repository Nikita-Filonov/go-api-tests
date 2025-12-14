package users

import (
	"net/http"
	"testing"

	"go-api-tests/fixtures"

	"github.com/Nikita-Filonov/axiom"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	c := axiom.NewCase(
		axiom.WithCaseName("get non-existing user"),
		axiom.WithCaseMeta(
			axiom.WithMetaStory("get user"),
			axiom.WithMetaTag("negative"),
		),
	)

	runner.RunCase(t, c, func(cfg *axiom.Config) {
		client := fixtures.GetUsersClientFixture(cfg)

		user, resp, err := client.GetUser(cfg.Context.Raw, 999999)

		assert.Error(cfg.SubT, err)
		assert.Nil(cfg.SubT, user)
		assert.Equal(cfg.SubT, http.StatusNotFound, resp.StatusCode())
	})
}
