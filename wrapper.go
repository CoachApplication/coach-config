package config

type Wrapper interface {
	Get(key string) (ScopedConfig, error)
	List() ([]string, error)
}
