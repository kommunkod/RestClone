package config

type BasicAuth struct {
	Enabled     bool         `json:"enabled" yaml:"enabled" env:"RC_AUTH_ENABLED" env-default:"false" description:"Enable basic authentication"`
	Credentials []Credential `json:"credentials" yaml:"credentials" description:"List of credentials for basic authentication"`
	Realm       string       `json:"realm" yaml:"realm" env:"RC_AUTH_REALM" env-default:"RestClone" description:"Basic authentication realm"`
}

type Credential struct {
	Username string `json:"username" yaml:"username" env:"RC_BASIC_USER" description:"Username for basic authentication"`
	Password string `json:"password" yaml:"password" env:"RC_BASIC_PASS" description:"Password for basic authentication"`
}

type Auth struct {
	Enabled   bool      `json:"enabled" yaml:"enabled" env:"RC_AUTH_ENABLED" env-default:"false" description:"Enable authentication"`
	BasicAuth BasicAuth `json:"basicAuth" yaml:"basicAuth" description:"Basic authentication configuration"`
}
