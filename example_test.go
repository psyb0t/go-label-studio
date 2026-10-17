package labelstudio_test

import (
	"context"
	"net/http"
	"os"

	labelstudio "github.com/psyb0t/go-label-studio"
)

//nolint:testableexamples // needs server credentials, compile check only
func Example() {
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
