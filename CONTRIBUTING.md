# [Security Check](./README.md) / Contributing Guide

-   [Audit](#audit)
-   [Format and Lint](#format-and-lint)
-   [Pre-commit Hook](#pre-commit-hook)
-   [Test](#test)
-   [API Documentation](#api-documentation)

## Audit

Package dependency audit provided by [Snyk](https://snyk.io).

```bash
yarn audit
```

## Format and Lint

Code auto formatting and linting provided by [Prettier](https://prettier.io/), configured in `.prettierrc.yml` and `.prettierignore`.

```bash
# Validate if code matches expected format.
yarn lint

# Modify source code to make it match expected format.
yarn format
```

## Pre-commit Hook

On commit, auto formatting is run by [Husky](https://github.com/typicode/husky), configured in `.huskyrc.js`.

## Test

### Client

Tests implemented with [Jest](https://jestjs.io), configured in `jest.config.js`, based on test files matching `./client/**/*.spec.ts`.

```bash
yarn test
```

### API, CLI, Core

```bash
go test ./{api,cli,core}/... -cover
```

## Dependencies

### API, CLI, Core

[Go Modules](https://github.com/golang/go/wiki/Modules) has been chosen to manage Go dependencies for this repository. Vendoring is used to install these modules into `./vendor` directory instead of into `$GOPATH` which helps to isolate this project from other Go projects being developed on the same workstation. Running `go mod vendor` to install dependencies and then providing `-mod=vendor` argument to the `go build` command enables this functionality.

## API Documentation

```bash
yarn docs:api
```
