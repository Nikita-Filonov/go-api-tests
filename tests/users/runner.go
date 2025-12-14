package users

import (
	"go-api-tests/tests"

	"github.com/Nikita-Filonov/axiom"
)

var runner = tests.BaseRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag("users"),
			axiom.WithMetaFeature("users"),
		),
	),
)
