package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

type Tls struct {
	SelfsignedHostname string `json:"hostname" yaml:"hostname" env:"RC_HOSTNAME" description:"Server hostname"`

	TLSCertFile string `json:"certFile" yaml:"certFile" env:"RC_TLS_CERT_FILE" description:"Path to the TLS certificate file"`
	TLSKeyFile  string `json:"keyFile" yaml:"keyFile" env:"RC_TLS_KEY_FILE" description:"Path to the TLS key file"`
	TLSCAFile   string `json:"caFile" yaml:"caFile" env:"RC_TLS_CA_FILE" description:"Path to the TLS CA file"`

	TLSCert string `json:"tlsCert" yaml:"tlsCert" env:"RC_TLS_CERT" description:"TLS certificate"`
	TLSKey  string `json:"tlsKey" yaml:"tlsKey" env:"RC_TLS_KEY" description:"TLS key"`
	TLSCA   string `json:"tlsCa" yaml:"tlsCa" env:"RC_TLS_CA" description:"TLS CA"`
}

func (s *Tls) GetTLSConfig() (*tls.Config, error) {
	var cert tls.Certificate
	var caCert *x509.CertPool
	var err error

	// Try loading the TLS files first
	if s.TLSCertFile != "" && s.TLSKeyFile != "" {
		cert, err = tls.LoadX509KeyPair(s.TLSCertFile, s.TLSKeyFile)
		if err != nil {
			return nil, err
		}
	}

	// Append CA certificate if provided
	if s.TLSCAFile != "" {
		caCert, err := os.ReadFile(s.TLSCAFile)
		if err != nil {
			return nil, err
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to append CA certificate")
		}
	}

	// Append additional CA Certificate if provided
	if s.TLSCA != "" {
		decoded, err := base64.StdEncoding.DecodeString(s.TLSCA)
		if err != nil {
			return nil, fmt.Errorf("failed to decode CA certificate: %v", err)
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(decoded) {
			return nil, fmt.Errorf("failed to append CA certificate")
		}
	}

	// Load certificate from string if provided
	if s.TLSCert != "" && s.TLSKey != "" {
		certData := base64.StdEncoding.EncodeToString([]byte(s.TLSCert))
		keyData := base64.StdEncoding.EncodeToString([]byte(s.TLSKey))

		cert, err = tls.X509KeyPair([]byte(certData), []byte(keyData))
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS certificate: %v", err)
		}
	}

	if cert.Certificate == nil || cert.PrivateKey == nil {
		selfsignedTemplate := &x509.Certificate{
			SerialNumber:          big.NewInt(1),
			Subject:               pkix.Name{CommonName: "localhost"},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().Add(365 * 24 * time.Hour),
			KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			BasicConstraintsValid: true,
			DNSNames:              []string{"localhost"},
		}

		if s.SelfsignedHostname != "" {
			selfsignedTemplate.Subject.CommonName = s.SelfsignedHostname
		}

		privkey, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return nil, fmt.Errorf("failed to generate RSA key: %v", err)
		}

		// If no certificate is provided, create a self-signed certificate
		ncert, err := x509.CreateCertificate(rand.Reader,
			selfsignedTemplate,
			selfsignedTemplate,
			privkey.Public(),
			privkey,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to create self-signed certificate: %v", err)
		}

		x509Cert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ncert})
		x509Key := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privkey)})

		cert, err = tls.X509KeyPair(x509Cert, x509Key)
		if err != nil {
			return nil, fmt.Errorf("failed to load self-signed certificate: %v", err)
		}
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCert,
		MinVersion:   tls.VersionTLS12,
	}

	return tlsConfig, nil
}
