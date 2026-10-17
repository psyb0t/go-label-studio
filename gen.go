// Package labelstudio is a Go client for the Label Studio API,
// generated with oapi-codegen from the OpenAPI document HumanSignal publishes.
package labelstudio

//go:generate bash scripts/fetch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.yaml
//go:generate bash scripts/split-gen.sh
