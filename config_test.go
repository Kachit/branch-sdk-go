package branchio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	result := NewConfig("foo", "bar")
	assert.Equal(t, ProdAPIUrl, result.Uri)
	assert.Equal(t, "foo", result.Key)
	assert.Equal(t, "bar", result.Secret)
}

func Test_Config_IsValidSuccess(t *testing.T) {
	config := Config{Uri: ProdAPIUrl, Key: "foo", Secret: "bar"}
	assert.Nil(t, config.IsValid())
	assert.NoError(t, config.IsValid())
}

func Test_Config_IsValidEmptyUri(t *testing.T) {
	filter := Config{Key: "foo", Secret: "bar"}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "uri" is empty`, result.Error())
}

func Test_Config_IsValidEmptyPublicKey(t *testing.T) {
	filter := Config{Uri: ProdAPIUrl, Key: "", Secret: "bar"}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "key" is empty`, result.Error())
}

func Test_Config_IsValidEmptySecretKey(t *testing.T) {
	filter := Config{Uri: ProdAPIUrl, Key: "foo", Secret: ""}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "secret" is empty`, result.Error())
}
