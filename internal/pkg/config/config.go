package config

import (
	"github.com/Aoi-hosizora/ahlib-mx/xvalidator"
	"github.com/Aoi-hosizora/ahlib/xdefault"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Config struct {
	Meta    *MetaConfig    `yaml:"meta"    validate:"required"`
	Message *MessageConfig `yaml:"message" validate:"required"`
}

type MetaConfig struct {
	Port    uint16 `yaml:"port"     validate:"required"`
	Host    string `yaml:"host"     default:"0.0.0.0"`
	RunMode string `yaml:"run-mode" default:"debug"`
	LogName string `yaml:"log-name" default:"./logs/console"`
	Pprof   bool   `yaml:"pprof"    default:"false"`
	Swagger bool   `yaml:"swagger"  default:"false"`
	DocHost string `yaml:"doc-host"`

	BucketPrd int64 `yaml:"bucket-prd" default:"60"  validate:"gt=0"`
	BucketCap int64 `yaml:"bucket-cap" default:"200" validate:"gt=0"`
	BucketQua int64 `yaml:"bucket-qua" default:"50"  validate:"gt=0"`
	DefLimit  int32 `yaml:"def-limit"  default:"20"  validate:"gt=0"`
	MaxLimit  int32 `yaml:"max-limit"  default:"50"  validate:"gt=0"`
}

type MessageConfig struct {
	GitHubToken string `yaml:"github-token" validate:"required"`
}

var _debugMode = true

func IsDebugMode() bool {
	return _debugMode
}

func Load(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	f = xstring.FastStob(os.ExpandEnv(xstring.FastBtos(f)))
	if err = yaml.Unmarshal(f, cfg); err != nil {
		return nil, err
	}
	if _, err = xdefault.FillDefaultFields(cfg); err != nil {
		return nil, err
	}
	if err = validateConfig(cfg); err != nil {
		return nil, err
	}

	_debugMode = strings.ToLower(cfg.Meta.RunMode) != "release"
	return cfg, nil
}

func validateConfig(cfg *Config) error {
	val := xvalidator.NewMessagedValidator()
	val.SetValidateTagName("validate")
	val.SetMessageTagName("message")
	val.UseTagAsFieldName("yaml", "json")

	err := val.ValidateStruct(cfg)
	if err != nil {
		ut, _ := xvalidator.ApplyEnglishTranslator(val.ValidateEngine())
		translated := err.(*xvalidator.MultiFieldsError).Translate(ut, false)
		return xvalidator.MergeMapToError(translated)
	}
	return nil
}
