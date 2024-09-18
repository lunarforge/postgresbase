package apis_test

import (
	"net/http"
	"testing"

	"github.com/lunarforge/xbase/tests"
)

func TestHealthAPI(t *testing.T) {
	t.Parallel()

	scenarios := []tests.ApiScenario{
		{
			Name:           "health status returns 200",
			Method:         http.MethodGet,
			Url:            "/api/health",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"code":200`,
				`"data":{`,
				`"canBackup":true`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
