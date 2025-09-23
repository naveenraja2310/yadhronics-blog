package settings

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config Configuration

type S3Config struct {
	BucketName       string
	Region           string
	AccessKey        string
	SecretKey        string
	DistributionId   string
	CloudfrontDomain string
}

type Configuration struct {
	DBURI          string `mapstructure:"DB_URI"`
	DB_NAME        string `mapstructure:"DB_NAME"`
	DB_TIME        int    `mapstructure:"DB_TIME"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	AppPort        string `mapstructure:"APP_PORT"`
	AllowedDomains string `mapstructure:"ALLOWED_DOMAINS"`
	AWS            S3Config
}

func InitConfig() (Configuration, error) {
	var configDir, envDir string

	currentDir, err := os.Getwd()
	if err != nil {
		return Config, err
	}

	configDir = filepath.Join(currentDir, "")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		configDir = "./"
	}

	envDir = filepath.Join(currentDir, "")
	if _, err := os.Stat(envDir); os.IsNotExist(err) {
		envDir = "./"
	}

	// Load `.env`
	viper.AddConfigPath(envDir)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.MergeInConfig(); err != nil {
		return Config, err
	}

	// Bind sensitive env variables
	envVars := []string{
		"DB_URI", "DB_TIME", "APP_PORT", "DB_NAME", "JWT_SECRET", "ALLOWED_DOMAINS",
	}

	for _, envVar := range envVars {
		if err := viper.BindEnv(envVar); err != nil {
			return Config, err
		}
	}

	viper.SetDefault("AWS", map[string]interface{}{
		"BucketName":      "",
		"Region":          "",
		"AccessKey":       "",
		"SecretKey":       "",
		"DistributionKey": "",
	})

	// Unmarshal the combined configuration
	if err := viper.Unmarshal(&Config); err != nil {
		return Config, err
	}

	Config.AWS = S3Config{
		BucketName:       viper.GetString("AWS_BUCKET_NAME"),
		Region:           viper.GetString("AWS_REGION"),
		AccessKey:        viper.GetString("AWS_ACCESS_KEY"),
		SecretKey:        viper.GetString("AWS_SECRET_KEY"),
		DistributionId:   viper.GetString("AWS_DISTRIBUTION_ID"),
		CloudfrontDomain: viper.GetString("AWS_CLOUDFRONT_DOMAIN"),
	}

	fmt.Println("Configuration loaded successfully")
	return Config, nil
}
