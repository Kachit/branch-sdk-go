package branchio

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Resources_Resource_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)

	body, _ := LoadStubResponseData("stubs/data/exports/export.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.Get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_Resources_Resource_Post(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)

	body, _ := LoadStubResponseData("stubs/data/exports/export.success.json")

	httpmock.RegisterResponder("POST", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.Post("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_Resources_NewResourceAbstract(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	assert.NotEmpty(t, resource)
}
