package configTemplate

//import (
//	"fmt"
//	"github.com/spf13/viper"
//)
//
//type Config struct {
//	ListenAddr  string `mapstructure:"LISTEN_ADDR"`
//	ListenPort  string `mapstructure:"LISTEN_PORT"`
//	Env         string `mapstructure:"ENV"`
//	DatabaseUrl string `mapstructure:"DATABASE_URL"`
//}
//
//func (c Config) ListenAddrAndPort() string {
//	return fmt.Sprintf("%s:%s", c.ListenAddr, c.ListenPort)
//}
//
//func FromEnv() (*Config, error) {
//	v := viper.New()
//	v.SetDefault("LISTEN_ADDR", "0.0.0.0аа")
//	v.SetDefault("LISTEN_PORT", "8000")
//	v.SetDefault("ENV", "local")
//	v.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/magbat_user?sslmode=disable")
//	v.SetConfigName("env")
//	v.AutomaticEnv()
//
//	cfg := Config{}
//	err := v.Unmarshal(&cfg)
//	if err != nil {
//		return nil, err
//	}
//
//	return &cfg, nil
//}
