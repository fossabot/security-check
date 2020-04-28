# [Security Check](./README.md) / Architecture

-   [Conceptual Components](#conceptual-components)

## Conceptual Components

### Checks

A Check is a foundational component of the system responsible for performing a single validation.

An example is a URL parameters reflective cross site scripting check that populates GET parameters with malicious content and looks at the rendered HTML for the content to be reflected.

```yaml
name: "Example Search Reflective XSS"
request_url: "https://example.com/search?q={value}"
response_reflection_selector: ".query"
```
