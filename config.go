package branchio

//Common config
type Config struct {
	AppId  string
	Uri    string
	Key    string
	Secret string
}

//Create new config from credentials
func NewConfig(key string, secret string) *Config {
	cfg := &Config{
		Uri:    ProdAPIUrl,
		Key:    key,
		Secret: secret,
	}
	return cfg
}
