# [Security Check](./README.md) / Contributing Guide

-   [Audit](#audit)
-   [Format and Lint](#format-and-lint)
-   [Pre-commit Hook](#pre-commit-hook)

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

Tests implemented with [Jest](https://jestjs.io), configured in `jest.config.js`, based on test files matching `./client/**/*.spec.ts`.

```bash
yarn test
```
