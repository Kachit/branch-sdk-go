package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Export_1(t *testing.T) {
	cfg := BuildStubConfig()
	transport := NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	assert.NotEmpty(t, resource)
}
