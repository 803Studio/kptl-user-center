package config

type MariaTablesConfig struct {
    Users string `yaml:"users" toml:"users" json:"users"`
}

type MariaConfig struct {
    Username string            `yaml:"username" toml:"username" json:"username"`
    Password string            `yaml:"password" toml:"password" json:"password"`
    Database string            `yaml:"database" toml:"database" json:"database"`
    Addr     string            `yaml:"addr" toml:"addr" json:"addr"`
    Conn     int               `yaml:"conn" toml:"conn" json:"conn"`
    Tables   MariaTablesConfig `yaml:"tables" toml:"tables" json:"tables"`
}

type KeysConfig struct {
    Private string `yaml:"private" toml:"private" json:"private"`
    Public  string `yaml:"public" toml:"public" json:"public"`
}

type AppConfigType struct {
    Port   int         `yaml:"port" toml:"port" json:"port"`
    Secret string      `yaml:"secret" toml:"secret" json:"secret"`
    Maria  MariaConfig `yaml:"maria" toml:"maria" json:"maria"`
    Keys   KeysConfig  `yaml:"keys" toml:"keys" json:"keys"`
}
