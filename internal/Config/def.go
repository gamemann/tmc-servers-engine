package Config

type Config struct {
	RetrieveURL string `json:"retrieveurl"`
	UpdateURL   string `json:"updateurl"`
	BasicAuth   bool   `json:"basicauth"`
	Token       string `json:"token"`

	// For basic auth. Name of parameter we're using with the token.
	KeyParam string `json:"keyparam"`

	// For verifying a server.
	ClaimKeyField string `json:"claimkeyfield"`

	// Server retrieving.
	MaxServers       uint `json:"maxservers"`
	MaxServersPerReq uint `json:"maxserverspr"`

	// Query settings.
	WaitInterval  uint `json:"waitinterval"`
	FetchInterval uint `json:"fetchinterval"`
}

func (cfg *Config) SetDefaults() {
	cfg.KeyParam = "key"
	cfg.ClaimKeyField = "claimkey"

	cfg.RetrieveURL = "https://mydomain.example/servers?sort=laststatupdate"
	cfg.UpdateURL = "https://mydomain.example/servers/{id}"

	cfg.MaxServers = 40
}
