package branchio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	testable *Config
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.testable = NewConfig("foo", "bar")
}

func (suite *ConfigTestSuite) TestNewConfigByDefault() {
	assert.Equal(suite.T(), ProdAPIUrl, suite.testable.Uri)
	assert.Equal(suite.T(), "foo", suite.testable.Key)
	assert.Equal(suite.T(), "bar", suite.testable.Secret)
}

func (suite *ConfigTestSuite) TestIsValidSuccess() {
	assert.Nil(suite.T(), suite.testable.IsValid())
	assert.NoError(suite.T(), suite.testable.IsValid())
}

func (suite *ConfigTestSuite) TestIsValidEmptyUri() {
	suite.testable.Uri = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "uri" is empty`, result.Error())
}

func (suite *ConfigTestSuite) TestIsValidEmptyPublicKey() {
	suite.testable.Key = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "key" is empty`, result.Error())
}

func (suite *ConfigTestSuite) TestIsValidEmptySecretKey() {
	suite.testable.Secret = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "secret" is empty`, result.Error())
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
