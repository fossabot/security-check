# [Security Check](./README.md) / Architecture

-   [Conceptual Components](#conceptual-components)
-   [System Components](#system-components)

## Conceptual Components

### Checks

A Check is a foundational component of the system responsible for performing a single validation.

An example is a URL parameters reflective cross site scripting check that populates GET parameters with malicious content and looks at the rendered HTML for the content to be reflected.

```yaml
name: "Example Search Reflective XSS"
request_url: "https://example.com/search?q={value}"
response_reflection_selector: ".query"
```

### Suites

A suite is a collection of one or more checks along with configuration on how the checks will be executed.

### Runs

A run is a moment in time when one or more suites are executed and the results of running those suites are collated.

## System Components

### `/api`

[GoLang](https://golang.org/) web API used by the client component. See [`api/openapi.yaml`](./api/openapi.yaml) for API specification.

### `/cli`

[GoLang](https://golang.org/) command line interface used to interact with the application core.

### `/client`

[ReactJS](https://reactjs.org/) and [TypeScript](https://www.typescriptlang.org/) single page application witch interacts with API component.

### `/core`

[GoLang](https://golang.org/) application core that is accessed by CLI and API components.
