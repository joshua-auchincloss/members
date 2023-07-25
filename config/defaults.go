package config

import (
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type (
	getter = func(ctx *cli.Context, name string) interface{}

	option interface {
		applyViper(v *viper.Viper)
		withOverride(ctx *cli.Context)
	}

	default_v[T any] struct {
		Key string
		Env string

		get   getter
		value T
	}
)

var (
	_ option = ((*default_v[any])(nil))
)

func (v default_v[T]) Get(vi *viper.Viper) T {
	return vi.Get(v.Key).(T)
}

func (vx *default_v[T]) applyViper(vi *viper.Viper) {
	vi.BindEnv(
		vx.Key,
		vx.Env,
	)
	vi.SetDefault(vx.Key, vx.value)
}

func uint32Getter(ctx *cli.Context, name string) interface{} {
	i := uint32(ctx.Uint64(name))
	if i == 0 {
		return nil
	}
	return i
}

func uint32SliceGetter(ctx *cli.Context, name string) interface{} {
	u3 := []uint32{}
	u6 := ctx.Uint64Slice(name)
	if len(u6) == 0 {
		return nil
	}
	for _, v := range u6 {
		u3 = append(u3, uint32(v))
	}
	return u3
}

func stringGetter(ctx *cli.Context, name string) interface{} {
	v := ctx.String(name)
	if v == "" {
		return nil
	}
	return v
}

func boolGetter(ctx *cli.Context, name string) interface{} {
	v := ctx.Bool(name)
	return v
}

func stringSliceGetter(ctx *cli.Context, name string) interface{} {
	slc := ctx.StringSlice(name)
	if len(slc) == 0 {
		return nil
	}
	return slc
}

func (vx *default_v[T]) withOverride(ctx *cli.Context) {
	value := vx.get(ctx, vx.Key)
	if value != nil {
		if v, ok := value.(T); ok {
			vx.value = v
		}
	}
	log.Debug().Interface(vx.Key, vx.value).Msg("overriden")
}

var (
	ConfigYaml = default_v[string]{"config", "CONFIG", stringGetter, ""}

	ServicesDefn = default_v[[]string]{"services", "SERVICES", stringSliceGetter, []string{"all"}}

	MemberDefn            = default_v[uint32]{"cluster-member", "CLUSTER_MEMBER", uint32Getter, 8091}
	MemperConnsPerService = default_v[uint32]{"cluster-connections", "CLUSTER_CONNECTIONS", uint32Getter, uint32(0)}
	KnownParent           = default_v[string]{"cluster-known", "KNOWN_HOST", stringGetter, ""}
	BindIpDefn            = default_v[string]{"cluster-bind", "BIND_IP", stringGetter, "127.0.0.1"}

	RemoteTlsDefn         = default_v[bool]{"tls", "TLS", boolGetter, true}
	TlsAdminServerCert    = default_v[string]{"tls-admin-server-cert", "TLS_ADMIN_SERVER_CERT", stringGetter, ""}
	TlsAdminServerKey     = default_v[string]{"tls-admin-server-key", "TLS_ADMIN_SERVER_KEY", stringGetter, ""}
	TlsRegistryServerCert = default_v[string]{"tls-registry-server-cert", "TLS_REGISTRY_SERVER_CERT", stringGetter, ""}
	TlsRegistryServerKey  = default_v[string]{"tls-registry-server-key", "TLS_REGISTRY_SERVER_KEY", stringGetter, ""}
	TlsHealthServerCert   = default_v[string]{"tls-health-server-cert", "TLS_HEALTH_SERVER_CERT", stringGetter, ""}
	TlsHealthServerKey    = default_v[string]{"tls-health-server-key", "TLS_HEALTH_SERVER_KEY", stringGetter, ""}

	RemoteDebugDefn = default_v[bool]{"debug", "DEBUG", boolGetter, false}

	RegistrySvcDefn = default_v[[]uint32]{"cluster-registry-server-service",
		"CLUSTER_REGISTRY_SERVER_SERVICE",
		uint32SliceGetter,
		[]uint32{9009},
	}
	RegistryHealthDefn = default_v[[]uint32]{"cluster-registry-server-health", "CLUSTER_REGISTRY_SERVER_HEALTH", uint32SliceGetter, []uint32{4200}}

	RegistryCliDnsDefn = default_v[string]{"cluster-registry-client-dns",
		"CLUSTER_REGISTRY_CLIENT_DNS",
		stringGetter,
		"localhost",
	}
	RegistryCliAddrDefn = default_v[[]string]{"cluster-registry-client-addresses",
		"CLUSTER_REGISTRY_CLIENT_ADDRESSES",
		stringSliceGetter,
		[]string{},
	}

	AdminSvcDefn = default_v[[]uint32]{"cluster-admin-server-service",
		"CLUSTER_ADMIN_SERVER_SERVICE",
		uint32SliceGetter,
		[]uint32{9010},
	}
	AdminHealthDefn = default_v[[]uint32]{"cluster-admin-server-health",
		"CLUSTER_ADMIN_SERVER_HEALTH",
		uint32SliceGetter,
		[]uint32{4201},
	}

	AdminCliDnsDefn = default_v[string]{"cluster-admin-client-dns",
		"CLUSTER_ADMIN_CLIENT_DNS",
		stringGetter,
		"localhost",
	}
	AdminCliAddrDefn = default_v[[]string]{"cluster-admin-client-addresses",
		"CLUSTER_ADMIN_CLIENT_ADDRESSES",
		stringSliceGetter,
		[]string{},
	}

	StoreTypeDefn   = default_v[string]{"storage-type", "STORAGE_TYPE", stringGetter, "memory"}
	StoreUriDefn    = default_v[string]{"storage-uri", "STORAGE_URI", stringGetter, ""}
	StoreUserDefn   = default_v[string]{"storage-username", "STORAGE_USERNAME", stringGetter, ""}
	StorePwDefn     = default_v[string]{"storage-password", "STORAGE_PASSWORD", stringGetter, ""}
	StorePortDefn   = default_v[uint32]{"storage-port", "STORAGE_PORT", uint32Getter, 5432}
	StoreDbDefn     = default_v[string]{"storage-db", "STORAGE_DBNAME", stringGetter, ""}
	StoreSslDefn    = default_v[bool]{"storage-ssl", "STORAGE_SSL", boolGetter, false}
	StoreDropDefn   = default_v[bool]{"storage-drop", "STORAGE_DROP", boolGetter, false}
	StoreDebugDefn  = default_v[bool]{"storage-debug", "STORAGE_DEBUG", boolGetter, false}
	StoreCreateDefn = default_v[bool]{"storage-create", "STORAGE_CREATE", boolGetter, false}

	uint_opts = []default_v[uint32]{
		MemberDefn,
		StorePortDefn,
		MemperConnsPerService,
	}
	uint_slc_opts = []default_v[[]uint32]{
		RegistrySvcDefn,
		RegistryHealthDefn,
		AdminHealthDefn,
		AdminSvcDefn,
	}

	string_opts = []default_v[string]{
		TlsAdminServerCert,
		TlsAdminServerKey,
		TlsHealthServerCert,
		TlsHealthServerKey,
		TlsRegistryServerCert,
		TlsRegistryServerKey,
		ConfigYaml,
		KnownParent,
		BindIpDefn,
		StoreTypeDefn,
		StoreUriDefn,
		StoreUserDefn,
		StorePwDefn,
		StoreDbDefn,
		RegistryCliDnsDefn,
		AdminCliDnsDefn,
	}

	bool_opts = []default_v[bool]{
		StoreSslDefn,
		StoreDropDefn,
		StoreCreateDefn,
		StoreDebugDefn,
		RemoteDebugDefn,
		RemoteTlsDefn,
	}

	slice_opts = []default_v[[]string]{
		ServicesDefn,
		AdminCliAddrDefn,
		RegistryCliAddrDefn,
	}

	options = []option{
		&MemberDefn,
		&RegistryHealthDefn,
		&RegistrySvcDefn,
		&KnownParent,
		&BindIpDefn,
		&StoreTypeDefn,
		&StoreUriDefn,
		&StoreUserDefn,
		&StorePwDefn,
		&StorePortDefn,
		&StoreDbDefn,
		&StoreSslDefn,
		&StoreDropDefn,
		&StoreCreateDefn,
		&StoreDebugDefn,
		&ServicesDefn,
	}
)

func ClusterFlags() []cli.Flag {
	return Flags(
		uint_opts,
		string_opts,
		bool_opts,
		slice_opts,
		uint_slc_opts,
	)
}

func RemoteFlags() []cli.Flag {
	return Flags(
		[]default_v[uint32]{
			MemperConnsPerService,
		},
		[]default_v[string]{
			ConfigYaml,
			RegistryCliDnsDefn,
			AdminCliDnsDefn,
		},
		[]default_v[bool]{
			RemoteDebugDefn,
			RemoteTlsDefn,
		},
		[]default_v[[]string]{
			AdminCliAddrDefn,
			RegistryCliAddrDefn,
		},
		[]default_v[[]uint32]{},
	)
}
func Flags(
	uint_opts []default_v[uint32],
	string_opts []default_v[string],
	bool_opts []default_v[bool],
	slice_opts []default_v[[]string],
	uint_slc_opts []default_v[[]uint32],
) []cli.Flag {
	flags := []cli.Flag{}
	for _, def := range uint_opts {
		flags = append(flags, &cli.Uint64Flag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
			Value:   uint64(def.value),
		})
	}
	for _, def := range string_opts {
		flags = append(flags, &cli.StringFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
			Value:   def.value,
		})
	}
	for _, def := range bool_opts {
		flags = append(flags, &cli.BoolFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
			Value:   def.value,
		})
	}

	for _, def := range slice_opts {
		slc := cli.StringSlice{}
		for _, v := range def.value {
			slc.Set(v)
		}
		flags = append(flags, &cli.StringSliceFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
			Value:   &slc,
		})
	}

	for _, def := range uint_slc_opts {
		slc := cli.Uint64Slice{}
		for _, v := range def.value {
			slc.Set(strconv.Itoa(int(v)))
		}
		flags = append(flags, &cli.Uint64SliceFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
			Value:   &slc,
		})
	}
	return flags
}
