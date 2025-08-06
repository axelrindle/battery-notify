package config

import (
	"log"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func init() {
	config.AddDriver(yaml.Driver)

	config.WithOptions(config.ParseEnv, config.ParseTime, config.ParseDefault)

	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "key"
	})
}

func (c *Config) Load(file string) {
	if err := config.LoadFiles(file); err != nil {
		log.Fatal(err)
	}

	userDir := path.Join(path.Dir(file), "conf.d")
	if err := config.LoadFromDir(userDir, "yaml"); err != nil {
		log.Fatal(err)
	}

	if err := config.Decode(c); err != nil {
		log.Fatal(err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(c); err != nil {
		log.Fatal(err)
	}
}
