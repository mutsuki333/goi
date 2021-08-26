package value

import (
	"github.com/mutsuki333/goi/modules/log"
	"github.com/spf13/viper"
)

var store *viper.Viper

func init() {
	store = viper.New()
	store.SetConfigType("json")
}

func SetDefault(key string, value interface{}) {
	store.SetDefault(key, value)
}

func Set(key string, value interface{}) {
	store.Set(key, value)
}

func GetString(key string) string {
	return store.GetString(key)
}

func GetBool(key string) bool {
	return store.GetBool(key)
}

func GetInt(key string) int {
	return store.GetInt(key)
}

func GetFloat64(key string) float64 {
	return store.GetFloat64(key)
}

// Read read in data
func Read() {
	store.SetConfigName("data")
	// store.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}
}

// Save write current data to file
func Save() {
	store.WriteConfigAs("data.json")
}
