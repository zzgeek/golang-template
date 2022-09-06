package config

import "time"

var JwtConf *JwtConfig

type JwtConfig struct {
	Seed         []byte
	TokenExpired time.Duration
	Issuer       string
}

func InitJwtConfig(seed, issuer string, tokenExpired time.Duration) *JwtConfig {
	JwtConf = &JwtConfig{
		Seed:         []byte(seed),
		Issuer:       issuer,
		TokenExpired: tokenExpired,
	}
	return JwtConf
}
