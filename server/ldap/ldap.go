package ldap

import (
	"github.com/kacpersaw/go-ldap-client"
	"gitlab.com/kacpersaw/intra-auctions/config"
)

var (
	LDAP *ldap.LDAPClient
)

func InitLDAP() *ldap.LDAPClient {
	client := &ldap.LDAPClient{
		Base:         config.LDAP_BASE,
		Host:         config.LDAP_HOST,
		Port:         config.LDAP_PORT,
		UseSSL:       config.LDAP_SSL,
		BindDN:       config.LDAP_BIND_DN,
		BindPassword: config.LDAP_BIND_PASSWORD,
		UserFilter:   config.LDAP_USER_FILTER,
		Attributes:   config.LDAP_ATTRIBUTES,
	}
	defer client.Close()

	return client
}
