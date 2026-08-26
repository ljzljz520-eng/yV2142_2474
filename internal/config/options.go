package config

func (c Config) WithPageSize(n int) Config {
	if n > 0 && n < 1000 {
		c.PageSize = n
	}
	return c
}
func (c Config) Validate() error {
	if c.DBPath == "" {
		return ErrMissingPath
	}
	if c.Session == "" {
		return ErrMissingSession
	}
	return nil
}

type configError string

func (e configError) Error() string { return string(e) }

var ErrMissingPath = configError("database path required")
var ErrMissingSession = configError("session required")
