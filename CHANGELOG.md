# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-17

- Added `labelstudio`, a REST client generated with oapi-codegen from HumanSignal's Label Studio API document (OpenAPI 3.0.3; the document's `info.version` is empty). It covers all 376 paths, including projects, tasks, annotations, predictions, imports, exports, storage, users, organizations, workspaces, webhooks and ML backends.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from `https://raw.githubusercontent.com/HumanSignal/label-studio-client-generator/master/fern/openapi/openapi.yaml` and regenerates the client. The spec is not committed.
