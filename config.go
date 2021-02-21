package branchio

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
