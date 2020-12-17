package main

import (
	"os"

	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	strategy "github.com/koltyakov/gosip/auth/addin" // see more https://go.spflow.com/auth/overview
)

/**

Depending on the SharePoint environment and use case, auth strategy (https://go.spflow.com/auth/strategies)
can be different. For a production installation Azure Certificate Auth
(https://go.spflow.com/auth/custom-auth/azure-certificate-auth) might be preferred.

*/

// Binds SharePoint API client
func getSP() *api.SP {
	auth := &strategy.AuthCnfg{
		SiteURL:      os.Getenv("SPAUTH_SITEURL"),
		ClientID:     os.Getenv("SPAUTH_CLIENTID"),
		ClientSecret: os.Getenv("SPAUTH_CLIENTSECRET"),
	}
	client := &gosip.SPClient{AuthCnfg: auth}
	sp := api.NewSP(client)
	return sp
}
