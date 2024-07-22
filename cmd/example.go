package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/efellowsbg/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	config := bigip.Config{
		Address:           os.Getenv("BIG_IP_HOST"),
		Username:          os.Getenv("BIG_IP_USER"),
		Password:          os.Getenv("BIG_IP_PASSWORD"),
		CertVerifyDisable: true,
	}

	f5 := bigip.NewSession(&config)
	result, err := f5.GetWorkspace("sso")
	if err != nil {
		panic(err)
	}
	spew.Dump(result)
}
