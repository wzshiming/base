package base

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"
)

func RsaGenerateKey(bits int) (*PrivateKey, *PublicKey, error) {
	var priv, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return &PrivateKey{
			PrivateKey: *priv,
		}, &PublicKey{
			PublicKey: priv.PublicKey,
		}, nil
}

type PublicKey struct {
	rsa.PublicKey
}

func ParsePublicKey(code []byte) *PublicKey {
	var block, _ = pem.Decode(code)

	var pub, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	var p, ok = pub.(*rsa.PublicKey)
	if !ok {
		return nil
	}
	return &PublicKey{
		PublicKey: *p,
	}
}

func (pk *PublicKey) Encrypt(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &pk.PublicKey, data)
}

func (pk *PublicKey) Verify(data []byte, sig []byte) error {
	var h = md5.New()
	h.Write(data)
	var hashed = h.Sum(nil)
	var err = rsa.VerifyPKCS1v15(&pk.PublicKey, crypto.MD5, hashed, sig)
	if err != nil {
		return err
	}
	return nil
}

func (pk *PublicKey) Marshal(ss ...string) ([]byte, error) {
	var k, err = x509.MarshalPKIXPublicKey(&pk.PublicKey)
	if err != nil {
		return nil, err
	}
	var block pem.Block
	block.Bytes = k
	block.Type = strings.Join(ss, " ")
	return pem.EncodeToMemory(&block), nil
}

type PrivateKey struct {
	rsa.PrivateKey
}

func ParsePrivateKey(code []byte) *PrivateKey {
	var block, _ = pem.Decode(code)
	if block == nil {
		return nil
	}
	var p, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return &PrivateKey{
		PrivateKey: *p,
	}
}

func (pk *PrivateKey) Decrypt(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, &pk.PrivateKey, data)
}

func (pk *PrivateKey) Sign(data []byte) ([]byte, error) {
	var h = md5.New()
	h.Write(data)
	var hashed = h.Sum(nil)
	var res, err = rsa.SignPKCS1v15(rand.Reader, &pk.PrivateKey, crypto.MD5, hashed)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pk *PrivateKey) Marshal(ss ...string) ([]byte, error) {
	var k = x509.MarshalPKCS1PrivateKey(&pk.PrivateKey)
	var block pem.Block
	block.Bytes = k
	block.Type = strings.Join(ss, " ")
	return pem.EncodeToMemory(&block), nil
}
