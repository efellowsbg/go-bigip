package main

import (
	"context"
	"log"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	const wrkspcName = "exampleWorkspace"
	err := f5.CreateWorkspace(ctx, wrkspcName)
	if err != nil {
		panic(err)
	}
	result, err := f5.GetWorkspace(ctx, wrkspcName)
	if err != nil {
		panic(err)
	}
	log.Printf("Workspace: %v", result)
	opts := bigip.ExtensionConfig{
		WorkspaceName: wrkspcName,
		Name:          "exampleExt",
		Partition:     "Common",
	}
	err = f5.CreateExtension(ctx, opts)
	if err != nil {
		panic(err)
	}
	err = f5.UploadExtensionFiles(ctx, opts, "cmd/ilx_example/ilx")
	if err != nil {
		panic(err)
	}
	content, err := f5.ReadExtensionFiles(ctx, opts)
	if err != nil {
		panic(err)
	}
	spew.Dump(content)
}
