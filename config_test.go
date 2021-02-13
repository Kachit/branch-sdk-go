package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	result := NewConfig("foo", "bar")
	assert.Equal(t, BranchIOProductionUri, result.Uri)
	assert.Equal(t, "foo", result.Key)
	assert.Equal(t, "bar", result.Secret)
}
