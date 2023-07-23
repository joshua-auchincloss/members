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
	ServicesDefn = default_v[[]string]{"services", "SERVICES", stringSliceGetter, []string{"all"}}

	MemberDefn  = default_v[uint32]{"members-member", "MEMBER_PORT", uint32Getter, 8091}
	KnownParent = default_v[string]{"members-known", "KNOWN_HOST", stringGetter, ""}
	BindIpDefn  = default_v[string]{"members-bind", "BIND_IP", stringGetter, "127.0.0.1"}

	RegistrySvcDefn    = default_v[[]uint32]{"members-registry-service", "MEMBERS_REGISTRY_SERVICE", uint32SliceGetter, []uint32{9009}}
	RegistryHealthDefn = default_v[[]uint32]{"members-registry-health", "MEMBERS_REGISTRY_HEALTH", uint32SliceGetter, []uint32{4200}}

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
		StoreCreateDefn,
		StoreDebugDefn,
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
		&StoreCreateDefn,
		&StoreDebugDefn,
		&ServicesDefn,
	}
)

func Flags() []cli.Flag {
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
