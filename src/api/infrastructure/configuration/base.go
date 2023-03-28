package configuration

type baseConfiguration struct{}

func (*baseConfiguration) BasePath() string {
	return basePath
}
