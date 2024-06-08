package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/diginfra/client-go/pkg/api/outputs"
	"github.com/diginfra/client-go/pkg/client"
	"github.com/gogo/protobuf/jsonpb"
)

func printOutput(res *outputs.Response) error {
	out, err := (&jsonpb.Marshaler{}).MarshalToString(res)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}

func main() {
	c, err := client.NewForConfig(context.Background(), &client.Config{
		Hostname:   "localhost",
		Port:       5060,
		CertFile:   "/etc/diginfra/certs/client.crt",
		KeyFile:    "/etc/diginfra/certs/client.key",
		CARootFile: "/etc/diginfra/certs/ca.crt",
	})
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer c.Close()
	ctx := context.Background()

	err = c.OutputsWatch(ctx, printOutput, time.Second*1)
	if err != nil {
		log.Fatalf("outputs watch error: %v", err)
	}
}
