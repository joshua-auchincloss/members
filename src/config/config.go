package config

import (
	"fmt"
	"log"
	"members/common"
	"os"

	"github.com/hashicorp/memberlist"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"github.com/sethvargo/go-envconfig"
)

type (
	Storage struct {
		Kind     string `mapstructure:"kind" env:"KIND,overwrite"`
		URI      string `mapstructure:"uri" env:"URI,overwrite"`
		Username string `mapstructure:"username" env:"USERNAME,overwrite"`
		Password string `mapstructure:"password" env:"PASSWORD,overwrite"`
		Port     uint32 `mapstructure:"port" env:"PORT,overwrite"`
		DB       string `mapstructure:"db" env:"DB,overwrite"`
		SSL      bool   `mapstructure:"ssl" env:"SSL,overwrite"`
		Drop     bool   `mapstructure:"drop" env:"DROP,overwrite"`
		Debug    bool   `mapstructure:"debug" env:"DEBUG,overwrite"`
	}

	PortJoin struct {
		Service []uint32 `mapstructure:"service" env:"SERVICE,overwrite"`
		Health  []uint32 `mapstructure:"health" env:"HEALTH,overwrite"`
	}

	Members struct {
		Bind     string    `mapstructure:"bind" env:"BIND,overwrite"`
		Join     []string  `mapstructure:"join" env:"JOIN,overwrite"`
		Member   uint32    `mapstructure:"member" env:"MEMBER,overwrite"`
		Registry *PortJoin `mapstructure:"registry" env:",prefix=REGISTRY_"`
	}

	Config struct {
		Services []string               `mapstructure:"services" env:"SERVICES,overwrite"`
		Members  *Members               `mapstructure:"members" env:",prefix=CLUSTER_"`
		Storage  *Storage               `mapstructure:"storage" env:",prefix=STORAGE_"`
		List     *memberlist.Memberlist `mapstructure:"-"`
	}
)

func (m *Members) GetService(key common.Service) *PortJoin {
	switch key {
	case common.ServiceRegistry:
		return m.Registry
	}

	return nil
}

func getConfig(ctx *cli.Context) (*Config, error) {
	v := viper.NewWithOptions(
		viper.KeyDelimiter("-"),
	)
	log.Print("storage kind: ", os.Getenv("STORAGE_KIND"))
	for _, opt := range options {
		opt.withOverride(ctx)
		opt.applyViper(v)
	}
	viper.AutomaticEnv()
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := envconfig.Process(ctx.Context, &cfg); err != nil {
		return nil, err
	}
	log.Printf("config: %s", cfg.String())
	return &cfg, nil
}

func (cfg *Config) String() string {
	return fmt.Sprintf(`{
	Members: %s,
	Services: %+v,
	Storage: %+v
}`, cfg.Members.String(),
		cfg.Services,
		*cfg.Storage)
}

func (h *PortJoin) String() string {
	return fmt.Sprintf(`{
		Health: %+v,
		Service: %+v,
}`, h.Health, h.Service)
}

func (m *Members) String() string {
	return fmt.Sprintf(`{
	Bind: %+v,
	Member: %+v,
	Join: %+v
	Registry: %+v
}`, m.Bind,
		m.Member,
		m.Join,
		m.Registry.String(),
	)
}
