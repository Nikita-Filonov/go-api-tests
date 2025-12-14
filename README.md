# Go API Tests (Axiom + Allure)

This repository is a reference implementation of API testing architecture in Go, built on top of the standard testing
package with [Axiom](https://github.com/Nikita-Filonov/axiom) as a test execution engine.

The project demonstrates how to:

- write clean, boring, readable tests,
- move all test infrastructure out of test cases,
- automatically produce steps, logs, artifacts, and Allure reports,
- without breaking Go’s testing philosophy or go test workflow.

Tests are executed against a public API: https://dummyjson.com

---

## 🧠 Core idea

Tests should not know about test infrastructure.

In this project:

- the HTTP client is integrated with the test runtime,
- steps and artifacts are created inside HTTP hooks,
- test cases contain only business assertions.

As a result:

- tests read like specifications,
- reports are generated automatically,
- infrastructure is centralized and reusable.

---

## 🧱 Project structure

```text
tests/
 ├─ runner.go              # base runner (test platform)
 ├─ users/                 # user domain tests
 │   ├─ runner.go
 │   ├─ get_user_test.go
 │   └─ ...
 └─ products/              # product domain tests

fixtures/
 ├─ config.go              # configuration fixture
 ├─ logger.go              # logger fixture
 ├─ users.go               # users client fixture
 └─ products.go            # products client fixture

clients/
 ├─ users/                 # domain client
 └─ products/

http/
 ├─ client.go              # HTTP client
 ├─ hooks.go               # resty hooks → Axiom steps
 └─ logger.go

hooks/
 └─ allure.go              # Allure bootstrap hook
```

---

## 🧪 Technologies

### Language & core tools

- Go 1.22+
- `testing` — Go standard testing package
- `testify` — assertions

### Test execution engine

- [Axiom](https://github.com/Nikita-Filonov/axiom)
    - fixtures with lifecycle management
    - retries for flaky tests
    - metadata (tags, stories, features, severity)
    - hooks and plugin system
    - fully compatible with go test

### HTTP

- resty
    - before / after / error hooks
    - deep integration with test runtime

### Reporting

- Allure
    - HTTP request / response steps
    - structured artifacts
    - execution history via GitHub Pages

### CI

- GitHub Actions
    - tests on every commit and pull request
    - Allure report generation
    - automatic deployment to GitHub Pages

## 🎯 Who this project is for

This project is useful if you:

- write API / integration / e2e tests in Go
- want structured reports without logging in tests
- use Allure for reporting
- are building a test platform, not just test cases
