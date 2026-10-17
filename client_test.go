package labelstudio_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	labelstudio "github.com/psyb0t/go-label-studio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "EXAMPLE-DO-NOT-USE"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *labelstudio.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := labelstudio.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Token "+testAPIKey)

			return nil
		},
	)

	client, err := labelstudio.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestRetrieveProject(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/projects/42/", r.URL.Path)
		assert.Equal(t, "3", r.URL.Query().Get("members_limit"))
		assert.Equal(t, "Token "+testAPIKey, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"id":42,"title":"Example project"}`))
		assert.NoError(t, err)
	})

	res, err := client.ApiProjectsRetrieveWithResponse(
		context.Background(),
		42,
		&labelstudio.ApiProjectsRetrieveParams{MembersLimit: new(3)},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	assert.Equal(t, 42, *res.JSON200.Id)
	assert.Equal(t, "Example project", *res.JSON200.Title)
}

func TestCreateProject(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/projects/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Token "+testAPIKey, r.Header.Get("Authorization"))

		var body labelstudio.LseProjectCreateRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "Example project", *body.Title)

		if assert.NotNil(t, body.ControlWeights) {
			weight := (*body.ControlWeights)["label"]
			assert.InDelta(t, 0.123456789, weight.Overall, 0)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`{
			"id":7,
			"title":"Example project",
			"control_weights":{
				"label":{"overall":0.123456789,"type":"Choices"}
			}
		}`))
		assert.NoError(t, err)
	})

	res, err := client.ApiProjectsCreateWithResponse(
		context.Background(),
		labelstudio.ApiProjectsCreateJSONRequestBody{
			Title: new("Example project"),
			ControlWeights: &map[string]labelstudio.ControlTagWeightRequest{
				"label": {Overall: 0.123456789, Type: "Choices"},
			},
		},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, res.StatusCode())
	require.NotNil(t, res.JSON201)
	assert.Equal(t, 7, *res.JSON201.Id)
	assert.Equal(t, "Example project", *res.JSON201.Title)
	require.NotNil(t, res.JSON201.ControlWeights)
	weight := (*res.JSON201.ControlWeights)["label"]
	assert.InDelta(t, 0.123456789, weight.Overall, 0)
}

func TestUnauthorizedResponse(t *testing.T) {
	t.Parallel()

	const body = `{"detail":"Invalid token."}`

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/projects/42/", r.URL.Path)
		assert.Equal(t, "Token "+testAPIKey, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(body))
		assert.NoError(t, err)
	})

	res, err := client.ApiProjectsRetrieveWithResponse(
		context.Background(),
		42,
		nil,
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
	assert.JSONEq(t, body, string(res.Body))
	assert.Nil(t, res.JSON200)
}

func TestCanceledRequest(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(_ http.ResponseWriter, _ *http.Request) {
		t.Error("canceled request reached the server")
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	res, err := client.ApiProjectsRetrieveWithResponse(ctx, 42, nil)
	require.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, res)
}
