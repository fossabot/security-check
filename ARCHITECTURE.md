# [Security Check](./README.md) / Architecture

-   [Components](#components)
-   [Test Execution](#test-execution)

## Components

### `/api`

[GoLang](https://golang.org/) web API used by the client component. See [`api/openapi.yaml`](./api/openapi.yaml) for API specification.

### `/cli`

[GoLang](https://golang.org/) command line interface used to interact with the application core.

### `/client`

[ReactJS](https://reactjs.org/) and [TypeScript](https://www.typescriptlang.org/) single page application witch interacts with API component.

### `/core`

[GoLang](https://golang.org/) application core that is accessed by CLI and API components.

## Test Execution
