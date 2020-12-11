package main

import (
	"os"

	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	strategy "github.com/koltyakov/gosip/auth/addin"
)

var sp *api.SP

func init() {
	auth := &strategy.AuthCnfg{
		SiteURL:      os.Getenv("SPAUTH_SITEURL"),
		ClientID:     os.Getenv("SPAUTH_CLIENTID"),
		ClientSecret: os.Getenv("SPAUTH_CLIENTSECRET"),
	}
	client := &gosip.SPClient{AuthCnfg: auth}
	sp = api.NewSP(client)
}
