# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.3.1, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.3.0, 2026-10-23

- Breaking: `ApiOrganizationsMembershipsSearchCreate` and `ApiWorkspacesMembershipsBulkCreate` now take a `*ApiOrganizationsMembershipsSearchCreateParams` or `*ApiWorkspacesMembershipsBulkCreateParams` argument after the id. This applies to the `WithBody` and `WithFormdataBody` variants, the `ClientWithResponses` `...WithResponse` variants, the matching `ClientInterface` and `ClientWithResponsesInterface` methods, and the `NewApiOrganizationsMembershipsSearchCreateRequest*` and `NewApiWorkspacesMembershipsBulkCreateRequest*` builders. Migration: pass `nil` for the new argument to keep the old behavior.
- Added `ApiOrganizationsMembershipsSearchCreateParams` (field `IncludeSkills`) and `ApiWorkspacesMembershipsBulkCreateParams` (fields `Filters`, `Role`, `Search`, `Tags`, `UserLastActivityGte`, `UserLastActivityLte` and others the spec lists).
- Added `MemberSkill`, `MemberSkillStatusEnum` and its constants `MemberSkillStatusEnumFailed`, `MemberSkillStatusEnumSelfReported` and `MemberSkillStatusEnumVerified`.
- Added `LseOrganizationMemberList.Skills` and `PaginatedProjectMember.Skills` (`*[]MemberSkill`, optional).
- Added `IncludeSkills` to `ApiOrganizationsMembershipsSearchListParams` and `ApiProjectsMembersPaginatedListParams`.
- Added `Filters`, `UserLastActivityGte` and `UserLastActivityLte` to `ApiProjectsMembersBulkCreateParams`, and `Ids` to `ApiProjectsMembersBulkDestroyParams`.
- Added `ConfigurablePermissionOption.ExcludedBy` (`*string`, optional), the broader permission that supersedes this one.

## v0.2.1, 2026-10-22

- Added `AllRolesProjectList.IsOnboardingProject` (`*bool`, JSON `is_onboarding_project`, optional). The upstream spec describes it as an onboarding project that contributors must pass before identity verification.
- `ApiProjectsPartialUpdateMultipartRequestBody` moved to another generated file. Its name and type (`PatchedLseProjectUpdateRequest`) are unchanged.
- The module now requires Go 1.26.10 or newer. The previous Go 1.26.9 has a `crypto/x509` vulnerability (GO-2026-6809) that 1.26.10 fixes.
- The development image uses Go 1.26.10 and a newer `ca-certificates` package, because the old package version is gone from the Alpine index.

## v0.2.0, 2026-10-20

- `Label.CreatedBy` and `Label.Organization` changed from `int` to `*int` and are now optional in JSON (`omitempty`).
- Removed `CreatedBy` and `Organization` from `PatchedLabelRequest`, `ProjectTemplateRequest` and `PatchedProjectTemplateRequest`. The upstream spec no longer accepts them in requests.
- The `ApiProjectsPartialUpdate*RequestBody` type aliases moved between generated files. Their names and types are unchanged.

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
