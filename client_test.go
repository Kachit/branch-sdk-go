package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil)
	assert.NotEmpty(t, client)
}

func Test_Client_GetExportsResource(t *testing.T) {
	client := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Exports()
	assert.NotEmpty(t, result)
}
