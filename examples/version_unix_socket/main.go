package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diginfra/client-go/pkg/api/version"
	"github.com/diginfra/client-go/pkg/client"
	"github.com/gogo/protobuf/jsonpb"
)

func main() {
	// Set up a connection to the server.
	c, err := client.NewForConfig(context.Background(), &client.Config{
		UnixSocketPath: "unix:///run/diginfra/diginfra.sock",
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
	out, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
