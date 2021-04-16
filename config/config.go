package config

import (
	"fmt"
	"os"
	"strings"

	fsession "github.com/fasthttp/session/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/redis"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

//Config is configuration struct
type Config struct {
	*viper.Viper

	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
}

func New() *Config {
	config := &Config{
		Viper: viper.New(),
	}

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	// Automatically refresh environment variables
	config.AutomaticEnv()

	// Read configuration
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}

	config.setFiberConfig()

	return config
}

func (config *Config) setFiberConfig() {
	config.fiber = &fiber.Config{
		Prefork:                   config.GetBool("FIBER_PREFORK"),
		ServerHeader:              config.GetString("FIBER_SERVERHEADER"),
		StrictRouting:             config.GetBool("FIBER_STRICTROUTING"),
		CaseSensitive:             config.GetBool("FIBER_CASESENSITIVE"),
		Immutable:                 config.GetBool("FIBER_IMMUTABLE"),
		UnescapePath:              config.GetBool("FIBER_UNESCAPEPATH"),
		ETag:                      config.GetBool("FIBER_ETAG"),
		BodyLimit:                 config.GetInt("FIBER_BODYLIMIT"),
		Concurrency:               config.GetInt("FIBER_CONCURRENCY"),
		Views:                     config.getFiberViewsEngine(),
		ReadTimeout:               config.GetDuration("FIBER_READTIMEOUT"),
		WriteTimeout:              config.GetDuration("FIBER_WRITETIMEOUT"),
		IdleTimeout:               config.GetDuration("FIBER_IDLETIMEOUT"),
		ReadBufferSize:            config.GetInt("FIBER_READBUFFERSIZE"),
		WriteBufferSize:           config.GetInt("FIBER_WRITEBUFFERSIZE"),
		CompressedFileSuffix:      config.GetString("FIBER_COMPRESSEDFILESUFFIX"),
		ProxyHeader:               config.GetString("FIBER_PROXYHEADER"),
		GETOnly:                   config.GetBool("FIBER_GETONLY"),
		ErrorHandler:              config.errorHandler,
		DisableKeepalive:          config.GetBool("FIBER_DISABLEKEEPALIVE"),
		DisableDefaultDate:        config.GetBool("FIBER_DISABLEDEFAULTDATE"),
		DisableDefaultContentType: config.GetBool("FIBER_DISABLEDEFAULTCONTENTTYPE"),
		DisableHeaderNormalizing:  config.GetBool("FIBER_DISABLEHEADERNORMALIZING"),
		DisableStartupMessage:     config.GetBool("FIBER_DISABLESTARTUPMESSAGE"),
		ReduceMemoryUsage:         config.GetBool("FIBER_REDUCEMEMORYUSAGE"),
	}
}

//GetFiberConfig get fiber configuration
func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}

func (config *Config) getFiberViewsEngine() fiber.Views {
	var viewsEngine fiber.Views

	engine := html.New(config.GetString("FIBER_VIEWS_DIRECTORY"), config.GetString("FIBER_VIEWS_EXTENSION"))
	engine.Reload(config.GetBool("FIBER_VIEWS_RELOAD")).
		Debug(config.GetBool("FIBER_VIEWS_DEBUG")).
		Layout(config.GetString("FIBER_VIEWS_LAYOUT")).
		Delims(config.GetString("FIBER_VIEWS_DELIMS_L"), config.GetString("FIBER_VIEWS_DELIMS_R"))
	viewsEngine = engine

	return viewsEngine
}

//GetSessionConfig get session configuration
func (config *Config) GetSessionConfig() session.Config {
	var provider fsession.Provider
	switch strings.ToLower(config.GetString("SESSION_PROVIDER")) {
	case "redis":
		sessionProvider, err := redis.New(redis.Config{
			KeyPrefix: config.GetString("SESSION_KEYPREFIX"),
			Addr:      config.GetString("SESSION_HOST") + ":" + config.GetString("SESSION_PORT"),
			Password:  config.GetString("SESSION_PASSWORD"),
			DB:        config.GetInt("SESSION_DATABASE"),
		})
		if err != nil {
			fmt.Println("failed to initialized redis session provider:", err.Error())
			break
		}
		provider = sessionProvider
		break
	}

	return session.Config{
		Lookup:     config.GetString("SESSION_LOOKUP"),
		Secure:     config.GetBool("SESSION_SECURE"),
		Domain:     config.GetString("SESSION_DOMAIN"),
		SameSite:   config.GetString("SESSION_SAMESITE"),
		Expiration: config.GetDuration("SESSION_EXPIRATION"),
		Provider:   provider,
		GCInterval: config.GetDuration("SESSION_GCINTERVAL"),
	}
}
