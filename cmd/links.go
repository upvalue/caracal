package cmd

import "github.com/BurntSushi/toml"

type Link struct {
	Name string
	Url  string
}

type Config struct {
	Link  []Link
	Links map[string]Link
}

func loadConfig(configPath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}
	cfg.Links = make(map[string]Link)
	for _, link := range cfg.Link {
		cfg.Links[link.Name] = link
	}
	return &cfg, nil
}
