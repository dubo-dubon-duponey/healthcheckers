package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
  uri := fmt.Sprintf("%s", os.Getenv("HEALTHCHECK_URL"))
  if uri == "" {
    os.Exit(1)
  }

	contype := fmt.Sprintf("%s", os.Getenv("HEALTHCHECK_TYPE"))
	if contype != "tcp" {
		contype = "udp"
	}
  question := fmt.Sprintf("%s", os.Getenv("HEALTHCHECK_QUESTION"))

	tcpResolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dialer := net.Dialer{}
			return dialer.DialContext(ctx, contype, fmt.Sprintf("%s", uri))
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := tcpResolver.LookupCNAME(ctx, question)

	if err != nil || len(res) == 0 {
		os.Exit(1)
	}
	fmt.Println(res)
}
