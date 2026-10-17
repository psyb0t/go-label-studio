# go-label-studio

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-label-studio.svg)](https://pkg.go.dev/github.com/psyb0t/go-label-studio)
[![CI](https://github.com/psyb0t/go-label-studio/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-label-studio/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-label-studio/badges/version.svg)](https://github.com/psyb0t/go-label-studio/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-label-studio/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-label-studio/badges/importers.svg)](https://github.com/psyb0t/go-label-studio/blob/badges/importers.md)

Unofficial Go client for the [Label Studio API](https://api.labelstud.io), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from HumanSignal's public OpenAPI document. I am not affiliated with HumanSignal.

The package is `labelstudio`, in the repository root. Everything in `client_*.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-label-studio
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	labelstudio "github.com/psyb0t/go-label-studio"
)

func main() {
	auth := labelstudio.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Token "+os.Getenv("LABEL_STUDIO_API_KEY"),
			)

			return nil
		},
	)

	client, err := labelstudio.NewClientWithResponses(
		os.Getenv("LABEL_STUDIO_URL"),
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.ApiProjectsRetrieveWithResponse(
		context.Background(),
		1,
		nil,
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

Method names come from the spec's operation IDs, such as `ApiProjectsRetrieveWithResponse` and `ApiTasksListWithResponse`. All 376 paths and 592 operations are generated. The spec includes Enterprise endpoints; availability depends on the server edition and your permissions.

## Auth

The spec's `Token` security scheme is an API key in the `Authorization` header with the `Token ` prefix. The request editor above supplies it. Read the token from `LABEL_STUDIO_API_KEY`; the client does not read environment variables itself. Find the user token on the User Account page in Label Studio.

Set `LABEL_STUDIO_URL` to your Label Studio server's origin, such as `http://localhost:8000` for a local installation, or the HTTPS origin of your cloud instance. The generated methods include `/api/` in their paths, so do not append `/api` to the base URL. Use HTTPS for remote servers.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://raw.githubusercontent.com/HumanSignal/label-studio-client-generator/master/fern/openapi/openapi.yaml` to `api/openapi.yaml`. The repository's default branch is `master`.
- oapi-codegen runs with `config.yaml`. The config maps `number` to `float64`, because numbers without a format otherwise become `float32`. Type collision resolution and prefixed enum names avoid duplicate Go declarations in the full spec. No schema patches are needed.
- `scripts/split-gen.sh` splits the generated output at declaration boundaries into `client_*.gen.go` files so Go tooling can process the large client.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-label-studio) teaches a coding agent how to use go-label-studio when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-label-studio@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-label-studio@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The release workflow publishes the skill to ClawHub:

```bash
openclaw skills install @psyb0t/go-label-studio
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from HumanSignal's published API specification. The provider's spec repository is MIT licensed; some documented endpoints require Label Studio Enterprise. The provider's copyright and license notice are preserved in [THIRD_PARTY_NOTICES](THIRD_PARTY_NOTICES). The specification itself is not redistributed here.
