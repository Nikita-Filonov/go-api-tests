package products

import (
	"go-api-tests/tests"

	"github.com/Nikita-Filonov/axiom"
)

var runner = tests.BaseRunner.Join(
	axiom.NewRunner(
		axiom.WithRunnerMeta(
			axiom.WithMetaTag("products"),
			axiom.WithMetaFeature("products"),
		),
	),
)
