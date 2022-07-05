package common

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type ConfigObject map[string]string

const (
	configFilename = ".env"
	envPrefix      = "SRVEXEC_"
)

var (
	Config ConfigObject
)

// init Config
func init() {
	var err error
	Config, err = godotenv.Read(configFilename)

	if err != nil {
		LogWarn("error loading environment: " + err.Error())
	}

	// Remove prefix from keys
	for k := range Config {
		Config[strings.TrimPrefix(k, envPrefix)] = Config[k]
		delete(Config, k)
	}

	// for logs
	dotenvKeys := Config.Keys()

	// Env var take precedence over .env file
	for _, e := range os.Environ() {
		// split key and value
		pair := strings.Split(e, "=")
		// check if key starts with SRVEXEC_
		if strings.HasPrefix(pair[0], envPrefix) {
			// remove prefix
			key := strings.TrimPrefix(pair[0], envPrefix)
			// add key and value to config
			Config[key] = pair[1]
		}
	}

	// for logs
	envKeys := Config.Keys()

	// Remove values from dotenvKeys present in envKeys
	for _, d := range dotenvKeys {
		for i, e := range envKeys {
			if d == e {
				envKeys = append(envKeys[:i], envKeys[i+1:]...)
			}
		}
	}

	// Set log level
	LogLevel = Config.GetDefault("LOG_LEVEL", "info")
	LogInfo("setting log level to " + LogLevel)

	LogDebug(".env loaded: " + strings.Join(dotenvKeys, ", "))
	LogDebug("env vars loaded: " + strings.Join(envKeys, ", "))
	LogInfo("env vars loaded")
}

// GetSafe returns the value for the given key.
// If the key is not found, an empty string is returned along with a false value (false = don't exists).
func (c *ConfigObject) GetSafe(key string) (string, bool) {
	val, exist := (*c)[key]
	if !exist {
		// accessed but not set
		LogWarn("config key '" + key + "' accessed but not set")
		return "", false
	}
	return val, true
}

// GetDefault returns the value for the given key.
// If the key is not found in .env nor env, the default value is returned.
func (c *ConfigObject) GetDefault(key, defaultValue string) string {
	val, exist := c.GetSafe(key)
	if !exist {
		LogWarn("using default value '" + defaultValue + "' for config key '" + key + "'")
		return defaultValue
	}
	return val
}

func (c *ConfigObject) Get(key string) string {
	val, _ := c.GetSafe(key)
	return val
}

// Get keys
func (c *ConfigObject) Keys() []string {
	keys := make([]string, len(*c))
	i := 0
	for k := range *c {
		keys[i] = k // a bit quicker than append
		i++
	}
	return keys
}
