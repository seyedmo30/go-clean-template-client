package api

import _ "embed"

// OpenAPISpec contains the entire OpenAPI JSON spec embedded in the binary.
// You can serve it over HTTP or use it for code generation at runtime.

//go:embed openapi.json
var OpenAPISpec []byte
