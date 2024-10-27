package environment

type LoadConfigI interface {
	LoadEnv(config any) error
}

func NewLoadConfig() LoadConfigI {
	return &Environment{}
}
