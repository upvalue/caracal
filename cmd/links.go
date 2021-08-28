package cmd

import "github.com/BurntSushi/toml"

type Link struct {
	Name string
	Url  string
}

type Variable struct {
	Key   string
	Value string
}

type Config struct {
	Link     []Link
	Variable []Variable
	Port     int

	// Derived fields
	Links     map[string]Link
	Variables map[string]string
}

func loadConfig(configPath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	cfg.Links = make(map[string]Link)
	cfg.Variables = make(map[string]string)

	for _, link := range cfg.Link {
		cfg.Links[link.Name] = link
	}

	for _, v := range cfg.Variable {
		cfg.Variables[v.Key] = v.Value
	}

	if cfg.Port == 0 {
		cfg.Port = 8080
	}

	return &cfg, nil
}
