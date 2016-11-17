// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Sqlbeat SqlbeatConfig
}

type SqlbeatConfig struct {
	Period            string  `yaml:"period"`
	DBType            string  `yaml:"dbtype"`
	Hostname          string  `yaml:"hostname"`
	Port              string  `yaml:"port"`
	Username          string  `yaml:"username"`
	Password          string  `yaml:"password"`
	EncryptedPassword string  `yaml:"encryptedpassword"`
	Database          string  `yaml:"database"`
	PostgresSSLMode   string  `yaml:"postgressslmode"`
	Queries           []Query `yaml:"queries"`
	DeltaWildcard     string  `yaml:"deltawildcard"`
}

type Query struct {
	Sql         string `yaml:"sql"`
	Type        string `yaml:"type"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Interval    string `yaml:"interval"`
}
