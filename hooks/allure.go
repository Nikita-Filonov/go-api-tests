package hooks

import (
	"os"

	"github.com/Nikita-Filonov/axiom"
)

func AllureBeforeAllHook(_ *axiom.Runner) {
	if _, ok := os.LookupEnv("ALLURE_RESULTS_PATH"); !ok {
		_ = os.Setenv("ALLURE_RESULTS_PATH", ".")
	}
}
