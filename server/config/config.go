package config

import "github.com/spf13/viper"

var (
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string

	LDAP_BASE               string
	LDAP_HOST               string
	LDAP_PORT               int
	LDAP_SSL                bool
	LDAP_BIND_DN            string
	LDAP_BIND_PASSWORD      string
	LDAP_USER_FILTER        string
	LDAP_ATTRIBUTES         []string
	LDAP_USERNAME_ATTRIBUTE string
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	DB_HOST = viper.GetString("DB_HOST")

	viper.SetDefault("DB_PORT", "3306")
	DB_PORT = viper.GetString("DB_PORT")

	viper.SetDefault("DB_USERNAME", "dev")
	DB_USERNAME = viper.GetString("DB_USERNAME")

	viper.SetDefault("DB_PASSWORD", "dev")
	DB_PASSWORD = viper.GetString("DB_PASSWORD")

	viper.SetDefault("DB_NAME", "dev")
	DB_NAME = viper.GetString("DB_NAME")

	viper.SetDefault("LDAP_BASE", "ou=people,dc=test,dc=test")
	LDAP_BASE = viper.GetString("LDAP_BASE")

	viper.SetDefault("LDAP_HOST", "localhost")
	LDAP_HOST = viper.GetString("LDAP_HOST")

	viper.SetDefault("LDAP_PORT", 389)
	LDAP_PORT = viper.GetInt("LDAP_PORT")

	viper.SetDefault("LDAP_SSL", false)
	LDAP_SSL = viper.GetBool("LDAP_SSL")

	viper.SetDefault("LDAP_BIND_DN", "cn=admin,dc=dev,dc=test")
	LDAP_BIND_DN = viper.GetString("LDAP_BIND_DN")

	viper.SetDefault("LDAP_BIND_PASSWORD", "s3cr3t")
	LDAP_BIND_PASSWORD = viper.GetString("LDAP_BIND_PASSWORD")

	viper.SetDefault("LDAP_USER_FILTER", "(uid=%s)")
	LDAP_USER_FILTER = viper.GetString("LDAP_USER_FILTER")

	viper.SetDefault("LDAP_ATTRIBUTES", []string{"uid", "mail"})
	LDAP_ATTRIBUTES = viper.GetStringSlice("LDAP_ATTRIBUTES")

	viper.SetDefault("LDAP_USERNAME_ATTRIBUTE", "uid")
	LDAP_USERNAME_ATTRIBUTE = viper.GetString("LDAP_USERNAME_ATTRIBUTE")
}
