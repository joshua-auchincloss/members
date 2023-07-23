package config

import (
	"log"

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
	for _, v := range ctx.Uint64Slice(name) {
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

func stringSliceGetter(ctx *cli.Context, name string) interface{} {
	return ctx.StringSlice(name)
}

func (vx *default_v[T]) withOverride(ctx *cli.Context) {
	value := vx.get(ctx, vx.Key)
	if value != nil {
		if v, ok := value.(T); ok {
			vx.value = v
		}
	}
	log.Println(vx.Key, vx.value)
}

var (
	ServicesDefn = default_v[[]string]{"services", "SERVICES", stringGetter, []string{"all"}}

	MemberDefn  = default_v[uint32]{"members-member", "MEMBER_PORT", uint32Getter, 8091}
	KnownParent = default_v[string]{"members-known", "KNOWN_HOST", stringGetter, ""}
	BindIpDefn  = default_v[string]{"members-bind", "BIND_IP", stringGetter, "127.0.0.1"}

	RegistrySvcDefn    = default_v[[]uint32]{"members-registry-service", "MEMBERS_REGISTRY_SERVICE", uint32SliceGetter, []uint32{9009}}
	RegistryHealthDefn = default_v[[]uint32]{"members-registry-health", "MEMBERS_REGISTRY_HEALTH", uint32SliceGetter, []uint32{4200}}

	StoreTypeDefn = default_v[string]{"storage-type", "STORAGE_TYPE", stringGetter, "memory"}
	StoreUriDefn  = default_v[string]{"storage-uri", "STORAGE_URI", stringGetter, ""}
	StoreUserDefn = default_v[string]{"storage-username", "STORAGE_USERNAME", stringGetter, ""}
	StorePwDefn   = default_v[string]{"storage-password", "STORAGE_PASSWORD", stringGetter, ""}
	StorePortDefn = default_v[uint32]{"storage-port", "STORAGE_PORT", uint32Getter, 5432}
	StoreDbDefn   = default_v[string]{"storage-db", "STORAGE_DBNAME", stringGetter, ""}
	StoreSslDefn  = default_v[bool]{"storage-ssl", "STORAGE_SSL", stringGetter, false}
	StoreDropDefn = default_v[bool]{"storage-drop", "STORAGE_DROP", stringGetter, false}

	uint_opts = []default_v[uint32]{
		MemberDefn,
		StorePortDefn,
	}
	uint_slc_opts = []default_v[[]uint32]{
		RegistrySvcDefn,
		RegistryHealthDefn,
	}

	string_opts = []default_v[string]{
		KnownParent,
		BindIpDefn,
		StoreTypeDefn,
		StoreUriDefn,
		StoreUserDefn,
		StorePwDefn,
		StoreDbDefn,
	}

	bool_opts = []default_v[bool]{
		StoreSslDefn,
		StoreDropDefn,
	}

	slice_opts = []default_v[[]string]{
		ServicesDefn,
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
		&ServicesDefn,
	}
)

func Flags() []cli.Flag {
	flags := []cli.Flag{}
	for _, def := range uint_opts {
		flags = append(flags, &cli.Uint64Flag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
		})
	}
	for _, def := range string_opts {
		flags = append(flags, &cli.StringFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
		})
	}
	for _, def := range bool_opts {
		flags = append(flags, &cli.BoolFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
		})
	}

	for _, def := range slice_opts {
		flags = append(flags, &cli.StringFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
		})
	}

	for _, def := range uint_slc_opts {
		u6 := []uint64{}
		for _, v := range def.value {
			u6 = append(u6, uint64(v))
		}
		flags = append(flags, &cli.Uint64SliceFlag{
			Name:    def.Key,
			EnvVars: []string{def.Env},
		})
	}
	return flags
}
