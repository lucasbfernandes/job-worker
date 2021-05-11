package security

import (
	"github.com/dgrijalva/jwt-go"

	"crypto/rsa"
	"io/ioutil"
	"os"
)

const (
	// This will create files inside pwd/cert
	defaultTLSCertFilePath = "cert/server.crt"

	// This will create files inside pwd/cert
	defaultTLSKeyFilePath = "cert/server.key"

	// This will create files inside pwd/cert
	defaultJWTCertFilePath = "cert/jwt.crt"
)

type SecService struct {
	jwtPubKey *rsa.PublicKey
}

func NewSecService() (*SecService, error) {
	jwtCertBytes, err := ioutil.ReadFile(GetJWTCertFilePath())
	if err != nil {
		return nil, err
	}

	jwtPubKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtCertBytes)
	if err != nil {
		return nil, err
	}

	return &SecService{
		jwtPubKey: jwtPubKey,
	}, nil
}

func (s *SecService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return s.jwtPubKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetTLSCertFilePath() string {
	certFilePath, envExists := os.LookupEnv("TLS_CERT_FILE_PATH")
	if envExists && certFilePath != "" {
		return certFilePath
	}
	return defaultTLSCertFilePath
}

func GetTLSKeyFilePath() string {
	keyFilePath, envExists := os.LookupEnv("TLS_KEY_FILE_PATH")
	if envExists && keyFilePath != "" {
		return keyFilePath
	}
	return defaultTLSKeyFilePath
}

func GetJWTCertFilePath() string {
	certFilePath, envExists := os.LookupEnv("JWT_CERT_FILE_PATH")
	if envExists && certFilePath != "" {
		return certFilePath
	}
	return defaultJWTCertFilePath
}
