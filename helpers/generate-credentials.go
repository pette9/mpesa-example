package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//GenerateSecurityCredentials ...
func GenerateSecurityCredentials() string {
	pemData, err := ioutil.ReadFile(viper.GetString("mpesa_cert_filename"))
	if err != nil {
		logrus.Fatalf("read key file: %s", err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Fatalf("bad key data: %s", "not PEM-encoded")
	}
	if got, want := block.Type, "CERTIFICATE"; got != want {
		log.Fatalf("unknown key type %q, want %q", got, want)
	}
	key, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("bad private key: %s", err)
	}
	pkey := key.PublicKey.(*rsa.PublicKey)
	out, err := rsa.EncryptPKCS1v15(rand.Reader, pkey, []byte(viper.GetString("mpesa_initiator_password")))
	if err != nil {
		log.Fatalf("error encrypting initiator pass: %s", err)
	}
	return b64.StdEncoding.EncodeToString(out)
}
