package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"members/common"
	"members/utils"
	"os"

	"github.com/rs/zerolog/log"

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
		Create   bool   `mapstructure:"create" env:"CREATE,overwrite"`
	}

	PortJoin struct {
		Service []uint32 `mapstructure:"service" env:"SERVICE,overwrite"`
		Health  []uint32 `mapstructure:"health" env:"HEALTH,overwrite"`
	}

	Tls struct {
		ServerName string   `mapstructure:"name" env:"NAME,overwrite"`
		CertFile   string   `mapstructure:"cert" env:"CERT,overwrite"`
		KeyFile    string   `mapstructure:"key" env:"KEY,overwrite"`
		Ca         []string `mapstructure:"ca" env:"CA,overwrite"`
	}

	ClientTlsConfig struct {
		Trusted map[string][]string `mapstructure:"trusted"`
		Ca      []string            `mapstructure:"ca" env:"CA,overwrite"`
	}

	TlsConfig struct {
		Enabled    bool `env:"ENABLED"`
		Validation bool `env:"VALIDATION"`
		Registry   *Tls `env:",prefix=REGISTRY_"`
		Health     *Tls `env:",prefix=HEALTH_"`
	}

	ClientArgs struct {
		Dns       string   `mapstructure:"dns" env:"DNS,overwrite"`
		Addresses []string `mapstructure:"addresses" env:"ADDRESSES,overwrite"`

		Servers map[string][]string `mapstructure:"servers"`
	}

	Service struct {
		Svc    *PortJoin   `mapstructure:"server" env:",prefix=SERVER_"`
		Client *ClientArgs `mapstructure:"client" env:",prefix=CLIENT_"`
	}

	Members struct {
		Protocol              string           `mapstructure:"protocol" env:"PROTOCOL,overwrite"`
		Dns                   string           `mapstructure:"dns" env:"DNS,overwrite"`
		Bind                  string           `mapstructure:"bind" env:"BIND,overwrite"`
		Join                  []string         `mapstructure:"join" env:"JOIN,overwrite"`
		Member                uint32           `mapstructure:"member" env:"MEMBER,overwrite"`
		ConnectionsPerService uint32           `mapstructure:"connections" env:"CONNECTIONS,overwrite"`
		Registry              *Service         `mapstructure:"registry" env:",prefix=REGISTRY_"`
		Admin                 *Service         `mapstructure:"admin" env:",prefix=ADMIN_"`
		Client                *ClientTlsConfig `env:",prefix=CLIENT_"`
	}

	Config struct {
		Services []string               `mapstructure:"services" env:"SERVICES,overwrite"`
		Members  *Members               `mapstructure:"cluster" env:",prefix=CLUSTER_"`
		Storage  *Storage               `mapstructure:"storage" env:",prefix=STORAGE_"`
		Tls      *TlsConfig             `mapstructure:"tls" env:",prefix=TLS_"`
		List     *memberlist.Memberlist `mapstructure:"-"`
	}
)

func (t *TlsConfig) GetService(key common.Service) *Tls {
	switch key {
	case common.ServiceAdmin:
		return t.Health
	case common.ServiceRegistry:
		return t.Registry
	case common.ServiceHealth:
		return t.Health
	}
	return nil
}

func (m *Members) GetClient(key common.Service) *ClientArgs {
	switch key {
	case common.ServiceAdmin:
		return m.Admin.Client
	case common.ServiceRegistry:
		return m.Registry.Client
	}
	return nil
}

func (t *Tls) LoadCA() (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	for _, fi := range t.Ca {
		bts, err := os.ReadFile(fi)
		if err != nil {
			return nil, err
		}
		if ok := pool.AppendCertsFromPEM(bts); !ok {
			return nil, fmt.Errorf("invalid ca %s", fi)
		}
	}
	return pool, nil
}

func (t *Tls) Build() (*tls.Config, error) {
	tc, err := tls.LoadX509KeyPair(t.CertFile, t.KeyFile)
	if err != nil {
		return nil, err
	}
	pool, err := t.LoadCA()
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		ServerName:   t.ServerName,
		Certificates: []tls.Certificate{tc},
		RootCAs:      pool,
	}, nil
}

func (m *Members) GetService(key common.Service) *Service {
	switch key {
	case common.ServiceAdmin:
		return m.Admin
	case common.ServiceRegistry:
		return m.Registry
	default:
		return nil
	}
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

	filecfg := ctx.String(ConfigYaml.Key)
	if !utils.ZeroStr(filecfg) {
		v.SetConfigName(filecfg)
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := envconfig.Process(ctx.Context, &cfg); err != nil {
		return nil, err
	}
	log.Info().Interface("config", cfg).Send()
	return &cfg, nil
}
