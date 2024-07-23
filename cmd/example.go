package main

import (
	"log"
	"os"

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

	const wrkspcName = "exampleWorspace"
	err := f5.CreateWorkspace(wrkspcName)
	if err != nil {
		panic(err)
	}
	result, err := f5.GetWorkspace(wrkspcName)
	if err != nil {
		panic(err)
	}
	log.Printf("Workspace: %v", result)
	opts := bigip.ExtensionConfig{
		WorkspaceName: wrkspcName,
		Name:          "exampleExt",
		Partition:     "Common",
	}
	err = f5.CreateExtension(opts)
	if err != nil {
		panic(err)
	}
	err = f5.UploadExtensionFiles(opts, "./ilx_example/ilx")
	if err != nil {
		panic(err)
	}
}
