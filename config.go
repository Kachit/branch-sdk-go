package branchio

const BranchIOProductionUri = "https://api2.branch.io"

//Common config
type Config struct {
	Uri    string
	Key    string
	Secret string
}

//Create new config from credentials
func NewConfig(key string, secret string) *Config {
	cfg := &Config{
		Uri:    BranchIOProductionUri,
		Key:    key,
		Secret: secret,
	}
	return cfg
}
