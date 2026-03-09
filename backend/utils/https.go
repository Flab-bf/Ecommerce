package utils

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/http2"
	"io/ioutil"
)

func ConnectHttps() *tls.Config {
	cert, err := tls.LoadX509KeyPair("D:/GoProjects/src/ecommerce/ca_server/server.crt",
		"D:/GoProjects/src/ecommerce/ca_server/server.key")
	if err != nil {
		return nil
	}
	certBytes, err := ioutil.ReadFile("D:/GoProjects/src/ecommerce/ca_server/ca.crt")
	if err != nil {
		return nil
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("Failed to parse root certificate.")
	}
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MaxVersion:   tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		NextProtos: []string{http2.NextProtoTLS},
	}
	return cfg
}
