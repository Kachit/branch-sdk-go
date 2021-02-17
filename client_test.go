package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil)
	assert.NotEmpty(t, client)
}

func Test_Client_GetExportResource(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Export()
	assert.NotEmpty(t, result)
}
