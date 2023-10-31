package config

import (
	"os"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"

	

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Token string
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error("Ошибка reading config file", "err", err)
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		logger.Error("Ошибка parsing yaml", "err", err)
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) Token() string {
	return s.config.Token
}

func (s *Service) GetConfig() Config {
	return s.config
}
