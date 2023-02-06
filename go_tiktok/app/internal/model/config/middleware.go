package model

type Middleware struct {
	Jwt *Jwt `mapstructure:"jwt" yaml:"jwt"`
}

type Jwt struct {
	SecretKey   string `mapstructure:"secertKey" yaml:"secretKey"`
	ExpiresTime int64  `mapstructure:"expiresTime" yaml:"expiresTime"`
	BufferTime  int64  `mapstructure:"bufferTime" yaml:"bufferTime"`
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`
}
