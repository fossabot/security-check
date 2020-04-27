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
