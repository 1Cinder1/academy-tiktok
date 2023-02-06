package model

type Logger struct {
	LogLevel   string `mapstructure:"logLevel" yaml:"logLevel"`
	SavePath   string `mapstructure:"savePath" yaml:"savePath"`
	MaxSize    int    `mapstructure:"maxSize" yaml:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge" yaml:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups" yaml:"maxBackups"`
	IsCompres  bool   `mapstructure:"isCompres" yaml:"isCompres"`
}
