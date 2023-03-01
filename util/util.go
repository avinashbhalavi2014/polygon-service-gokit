package util

import (
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid"
	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	HostIp             string `mapstructure:"HOST_IP"`
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBConnectionString string `mapstructure:"DB_CONNECTION_STRING"`
	MigrationURL       string `mapstructure:"MIGRATION_URL"`
	ServiceName        string `mapstructure:"SERVICE_NAME"`
}

// LoadConfig will load the configuration from file...
func LoadConfig(path string) (config Config, err error) {
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName("app")
	vp.SetConfigType("yaml")

	vp.AutomaticEnv()

	err = vp.ReadInConfig()
	if err != nil {
		return
	}

	err = vp.Unmarshal(&config)
	return
}

// GenerateID will generate new id and return id in string
func GenerateID() (id string, err error) {

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	ulidID, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}

	return ulidID.String(), err
}
