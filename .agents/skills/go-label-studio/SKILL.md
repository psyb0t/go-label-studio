---
name: "go-label-studio"
description: "Use the unofficial Label Studio Go client for projects, tasks, annotations, predictions, imports, exports, storage, webhooks and administration. Use when Go code needs the Label Studio API or already imports go-label-studio."
homepage: https://github.com/psyb0t/go-label-studio
user-invocable: true
metadata:
  openclaw:
    emoji: "🏷️"
    requires:
      bins:
        - go
      env:
        - LABEL_STUDIO_API_KEY
        - LABEL_STUDIO_URL
permissions:
  filesystem: "Guides edits to Go source in the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "`go get` fetches the module from the Go module proxy. Runtime requests send API credentials and task data only to the user-supplied trusted Label Studio server. Use HTTPS for remote servers."
---

# go-label-studio

Unofficial, typed Go client for the Label Studio API, in package `labelstudio` at the module root. This module is not affiliated with HumanSignal.

For install, operation groups, pagination and response fields, see [references/setup.md](references/setup.md).

## Security & safety

- The spec uses `Authorization: Token <token>`. Read credentials from `LABEL_STUDIO_API_KEY`, and never hardcode, commit or log them or request headers.
- The caller supplies the base URL. Read it from trusted configuration such as `LABEL_STUDIO_URL`, verify that it belongs to the intended installation, and use HTTPS for remote servers. Annotation and task content goes to that server.
- Creating, updating or deleting projects, tasks, annotations, storage connections, webhooks or users changes server data. Perform those actions only when requested, on resources within the user's scope.

## When to use

- Go code needs to read or manage Label Studio projects, tasks, annotations or predictions.
- Go code needs Label Studio imports, exports, storage, ML backends, webhooks or administration.
- The project already imports `github.com/psyb0t/go-label-studio`.

## When NOT to use

- The code is not Go. This module ships only a Go package.
- You need a client maintained by HumanSignal. This one is unofficial.
- You want to change the generated client itself. See `AGENTS.md` in the repo.

## Usage

```go
import (
	"context"
	"net/http"
	"os"

	labelstudio "github.com/psyb0t/go-label-studio"
)

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
	return err
}

project, err := client.ApiProjectsRetrieveWithResponse(ctx, 1, nil)
if err != nil {
	return err
}

if project.JSON200 != nil {
	_ = project.JSON200
}
```

The error return covers transport and decoding failures. An HTTP error can return nil error with `JSON200` nil; check `StatusCode()`, `Status()` and `Body`. Check pointer fields before dereferencing them.

Operation IDs determine the generated method names. The spec includes Enterprise operations that a Community Edition server may not support. The server URL is its origin, without an appended `/api`, because generated paths already include that prefix.

Complete an integration by compiling its call sites and testing method, path, auth and decoded responses against a local `httptest` server. Never call a live Label Studio instance just to verify client wiring.
