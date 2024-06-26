package client_test

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/diginfra/client-go/pkg/api/outputs"
	"github.com/diginfra/client-go/pkg/api/version"
	"github.com/diginfra/client-go/pkg/client"
)

// The simplest use of a Client, just create and Close it.
func ExampleClient() {
	//Set up a connection to the server.
	c, err := client.NewForConfig(context.Background(), &client.Config{
		Hostname:   "localhost",
		Port:       5060,
		CertFile:   "/etc/diginfra/certs/client.crt",
		KeyFile:    "/etc/diginfra/certs/client.key",
		CARootFile: "/etc/diginfra/certs/ca.crt",
	})
	if err != nil {
		log.Fatalf("unable to create a Diginfra client: %v", err)
	}
	defer c.Close()
}

// A client that is created and then used to get the Diginfra output events.
func ExampleClient_outputsGet() {
	// Set up a connection to the server.
	c, err := client.NewForConfig(context.Background(), &client.Config{
		Hostname:   "localhost",
		Port:       5060,
		CertFile:   "/etc/diginfra/certs/client.crt",
		KeyFile:    "/etc/diginfra/certs/client.key",
		CARootFile: "/etc/diginfra/certs/ca.crt",
	})
	if err != nil {
		log.Fatalf("unable to create a Diginfra client: %v", err)
	}
	defer c.Close()
	outputsClient, err := c.Outputs()
	if err != nil {
		log.Fatalf("unable to obtain an output client: %v", err)
	}

	ctx := context.Background()
	fcs, err := outputsClient.Get(ctx, &outputs.Request{})
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		res, err := fcs.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error closing stream after EOF: %v", err)
		}
		fmt.Printf("rule: %s\n", res.Rule)
	}
}

func ExampleClient_version() {
	// Set up a connection to the server.
	c, err := client.NewForConfig(context.Background(), &client.Config{
		Hostname:   "localhost",
		Port:       5060,
		CertFile:   "/etc/diginfra/certs/client.crt",
		KeyFile:    "/etc/diginfra/certs/client.key",
		CARootFile: "/etc/diginfra/certs/ca.crt",
	})
	if err != nil {
		log.Fatalf("unable to create a Diginfra client: %v", err)
	}
	defer c.Close()
	versionClient, err := c.Version()
	if err != nil {
		log.Fatalf("unable to obtain a version client: %v", err)
	}

	ctx := context.Background()
	res, err := versionClient.Version(ctx, &version.Request{})
	if err != nil {
		log.Fatalf("error obtaining the Diginfra version: %v", err)
	}
	fmt.Printf("%v\n", res)
}
