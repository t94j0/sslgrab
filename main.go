package sslgrab

import (
	"crypto/tls"
	"crypto/x509"
)

// Grab grabs the specified SSL certificate.
//
// Requires port number in host
func Grab(host string) (certs []*x509.Certificate, err error) {
	conn, err := tls.Dial("tcp", host, &tls.Config{
		InsecureSkipVerify: true,
	})

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if err := conn.Handshake(); err != nil {
		return nil, err
	}

	return conn.ConnectionState().PeerCertificates, nil
}
