package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

const (
	sectionTwitter           = "twitter"
	keyTwitterConsumerKey    = "consumer_key"
	keyTwitterConsumerSecret = "consumer_secret"
	keyTwitterAccessToken    = "access_token"
	keyTwitterAccessSecret   = "access_secret"

	sectionOpenAi   = "openai"
	keyOpenAiApiKey = "api_key"
)

type Config struct {
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string
	OpenAiApiKey          string
}

func ReadConfig(iniFilePath string) (Config, error) {
	if !fileExists(iniFilePath) {
		return Config{}, fmt.Errorf("%s does not exist or is a directory", iniFilePath)
	}
	cfg, err := ini.Load(iniFilePath)
	if err != nil {
		return Config{}, err
	}

	if err := isIniFileValid(cfg); err != nil {
		return Config{}, err
	}

	twitterSection := cfg.Section(sectionTwitter)
	openAiSection := cfg.Section(sectionOpenAi)
	config := Config{
		TwitterConsumerKey:    twitterSection.Key(keyTwitterConsumerKey).String(),
		TwitterConsumerSecret: twitterSection.Key(keyTwitterConsumerSecret).String(),
		TwitterAccessToken:    twitterSection.Key(keyTwitterAccessToken).String(),
		TwitterAccessSecret:   twitterSection.Key(keyTwitterAccessSecret).String(),
		OpenAiApiKey:          openAiSection.Key(keyOpenAiApiKey).String(),
	}

	return config, nil
}

func isIniFileValid(cfg *ini.File) error {
	twitterSection, err := cfg.GetSection(sectionTwitter)
	if err != nil {
		return fmt.Errorf("missing section: %s", sectionTwitter)
	}

	requiredKeys := []string{keyTwitterConsumerKey, keyTwitterConsumerSecret, keyTwitterAccessToken, keyTwitterAccessSecret}
	for _, key := range requiredKeys {
		if !twitterSection.HasKey(key) {
			return fmt.Errorf("missing key '%s' in section '%s'", key, sectionTwitter)
		}
	}

	openAiSection, err := cfg.GetSection(sectionOpenAi)
	if err != nil {
		return fmt.Errorf("missing section: %s", sectionOpenAi)
	}

	if !openAiSection.HasKey(keyOpenAiApiKey) {
		return fmt.Errorf("missing key '%s' in section '%s'", keyOpenAiApiKey, sectionOpenAi)
	}

	return nil
}

func fileExists(iniFilePath string) bool {
	info, err := os.Stat(iniFilePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
