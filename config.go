package branchio

import "fmt"

//Config structure
type Config struct {
	AppId  string
	Uri    string
	Key    string
	Secret string
}

//NewConfig Create new config from credentials
func NewConfig(key string, secret string) *Config {
	cfg := &Config{
		Uri:    ProdAPIUrl,
		Key:    key,
		Secret: secret,
	}
	return cfg
}

//IsValid check is valid config parameters
func (c *Config) IsValid() error {
	var err error
	if c.Uri == "" {
		err = fmt.Errorf(`parameter "uri" is empty`)
	} else if c.Key == "" {
		err = fmt.Errorf(`parameter "key" is empty`)
	} else if c.Secret == "" {
		err = fmt.Errorf(`parameter "secret" is empty`)
	}
	return err
}
