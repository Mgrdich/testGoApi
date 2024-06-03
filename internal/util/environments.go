package util

type EnvironmentsType int

const (
	DevEnvironment EnvironmentsType = iota
	ProdEnvironment
)

var envs = [...]string{
	DevEnvironment:  "dev",
	ProdEnvironment: "prod",
}

var envsMap map[string]EnvironmentsType

func init() {
	envsMap = make(map[string]EnvironmentsType)
	for i := 0; i < len(envs); i++ {
		envsMap[envs[i]] = EnvironmentsType(i)
	}
}

func LookUpEnv(env string) (EnvironmentsType, bool) {
	value, ok := envsMap[env]

	return value, ok
}
