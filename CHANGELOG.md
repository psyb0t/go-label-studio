# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.1, 2026-10-17

- `make sec` no longer fails on the generated client's free-form JSON fields. Semgrep's unsafe-deserialization-interface rule is skipped for the generated client only. It still runs on handwritten code, and every other rule still runs on the generated client.

## v0.1.0, 2026-10-17

- Added `labelstudio`, a REST client generated with oapi-codegen from HumanSignal's Label Studio API document (OpenAPI 3.0.3; the document's `info.version` is empty). It covers all 376 paths, including projects, tasks, annotations, predictions, imports, exports, storage, users, organizations, workspaces, webhooks and ML backends.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec from `https://raw.githubusercontent.com/HumanSignal/label-studio-client-generator/master/fern/openapi/openapi.yaml` and regenerates the client. The spec is not committed.
