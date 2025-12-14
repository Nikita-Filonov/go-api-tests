package tests

import (
	"time"

	"go-api-tests/fixtures"
	"go-api-tests/hooks"

	"github.com/Nikita-Filonov/axiom"
	"github.com/Nikita-Filonov/axiom/plugins/testallure"
	"github.com/Nikita-Filonov/axiom/plugins/testlogger"
	"github.com/Nikita-Filonov/axiom/plugins/teststats"
	"github.com/Nikita-Filonov/axiom/plugins/testtags"
)

var TestStats = teststats.NewStats()

var BaseRunner = axiom.NewRunner(
	axiom.WithRunnerMeta(
		axiom.WithMetaEpic("api-tests"),
		axiom.WithMetaLayer("api"),
		axiom.WithMetaSeverity(axiom.SeverityNormal),
		axiom.WithMetaTag("dummyjson"),
	),

	axiom.WithRunnerFixture("config", fixtures.SetConfigFixture),
	axiom.WithRunnerFixture("logger", fixtures.SetLoggerFixture),
	axiom.WithRunnerFixture("users", fixtures.SetUsersClientFixture),
	axiom.WithRunnerFixture("products", fixtures.SetProductsClientFixture),

	axiom.WithRunnerHooks(
		axiom.WithBeforeAll(hooks.AllureBeforeAllHook),
	),

	axiom.WithRunnerPlugins(
		testtags.Plugin(),
		teststats.Plugin(TestStats),
		testlogger.Plugin(),
		testallure.Plugin(),
	),

	axiom.WithRunnerRetry(
		axiom.WithRetryTimes(3),
		axiom.WithRetryDelay(2*time.Second),
	),

	axiom.WithRunnerParallel(),
)
