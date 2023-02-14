package model

type Config struct {
	Logger     *Logger     `mapstructure:"logger"yaml:"logger"`
	Database   *Database   `mapstructure:"database" yaml:"database"`
	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
}
