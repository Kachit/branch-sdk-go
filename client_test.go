package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	client, err := NewClientFromConfig(BuildStubConfig(), nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
}

func Test_Client_NewClientFromConfigInvalid(t *testing.T) {
	cfg := BuildStubConfig()
	cfg.Uri = ""
	client, err := NewClientFromConfig(cfg, nil)
	assert.Error(t, err)
	assert.Empty(t, client)
}

func Test_Client_GetExportResource(t *testing.T) {
	client, _ := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Export()
	assert.NotEmpty(t, result)
}
