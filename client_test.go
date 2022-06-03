package branchio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) TestNewClientFromConfig() {
	client, err := NewClientFromConfig(BuildStubConfig(), nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), client)
}

func (suite *ClientTestSuite) TestNewClientFromConfigInvalid() {
	cfg := BuildStubConfig()
	cfg.Uri = ""
	client, err := NewClientFromConfig(cfg, nil)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), client)
}

func (suite *ClientTestSuite) TestGetExportResource() {
	client, _ := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Export()
	assert.NotEmpty(suite.T(), result)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
