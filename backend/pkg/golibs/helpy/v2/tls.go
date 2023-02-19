package helpy

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

func TLSInsecure() credentials.TransportCredentials {
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = true
	return credentials.NewTLS(&tlsConfig)
}
