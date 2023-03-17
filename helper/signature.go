package helper

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrParsePEM   = errors.New("failed to parse PEM block containing the key")
	ErrKeyTypeRSA = errors.New("key type is not RSA")
)

func CreateSignatureHeader(signingSecret, body []byte) (string, error) {
	h := hmac.New(sha256.New, signingSecret)
	_, err := h.Write(body)
	if err != nil {
		return "", err
	}

	hash := hex.EncodeToString(h.Sum(nil))
	return hash, nil
}

func RSAVerifyPSS(data []byte, signature, pubPEM string) error {
	sig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	pk, err := ParseRSAPublicKey(pubPEM)
	if err != nil {
		return err
	}

	digest := sha256.Sum256(data)

	return rsa.VerifyPSS(pk, crypto.SHA256, digest[:], sig, &rsa.PSSOptions{})
}

func ParseRSAPublicKey(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, ErrParsePEM
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}

	return nil, ErrKeyTypeRSA
}

func ParseRSAPrivateKey(privatePEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privatePEM))
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return private, nil
}

func RSASignPSS(data []byte, privPEM string) (string, error) {
	privKey, err := ParseRSAPrivateKey(privPEM)
	if err != nil {
		return "", err
	}

	rng := rand.Reader
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPSS(rng, privKey, crypto.SHA256, hashed[:], &rsa.PSSOptions{})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func GenJWTToken(secretKey []byte, payload interface{}) (string, error) {
	payloadData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	var claims jwt.MapClaims
	err = json.Unmarshal(payloadData, &claims)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

// VerifyJWTTokenWithProcessClaims with processClaimFunc will process claims and return the secret key
func VerifyJWTTokenWithProcessClaims(tokenStr string, processClaimsFunc func(map[string]interface{}) ([]byte, error)) error {
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, nil
		}

		secretKey, err := processClaimsFunc(claims)
		if err != nil {
			return nil, err
		}

		return secretKey, nil
	})

	return err
}

func GenerateRSAKeys(k KeySize) (privatePem, publicPem string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, int(k))
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Save the private key as a PEM block
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	var privateKeyBuf bytes.Buffer
	err = pem.Encode(&privateKeyBuf, privateKeyPEM)
	if err != nil {
		fmt.Println("Error encoding private key:", err)
		return
	}
	privatePem = privateKeyBuf.String()

	// Extract the public key from the private key
	publicKey := privateKey.PublicKey

	// Save the public key as a PEM block
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("Error marshaling public key:", err)
		return
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	var publicKeyBuf bytes.Buffer
	err = pem.Encode(&publicKeyBuf, publicKeyPEM)
	if err != nil {
		fmt.Println("Error encoding public key:", err)
		return
	}
	publicPem = publicKeyBuf.String()

	return
}
