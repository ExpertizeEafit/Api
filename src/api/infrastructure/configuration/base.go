package configuration

type baseConfiguration struct{}

func (*baseConfiguration) BasePath() string {
	return basePath
}

func (*baseConfiguration) DBDriverName() string {
	return dbDriverName
}
