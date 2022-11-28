package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var PublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqYVVk0t/z8dkbzMcwDsuw7AnC
aFCGHnjPLGaTSCUbcdpq10eI+2KbJjW03/SV2WtK0zHjsp8KVSuc0LWZjrsofFd/
hrCIOslLr9UtfgwRnCYPhBKfRWHFLv/WpMOCXN4Q9lOUCkqZoG5Jd0uhJ0L3WTI3
gAt2GhlKOBbeR63SJwIDAQAB	
-----END PUBLIC KEY-----`)
var PrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQCqYVVk0t/z8dkbzMcwDsuw7AnCaFCGHnjPLGaTSCUbcdpq10eI
+2KbJjW03/SV2WtK0zHjsp8KVSuc0LWZjrsofFd/hrCIOslLr9UtfgwRnCYPhBKf
RWHFLv/WpMOCXN4Q9lOUCkqZoG5Jd0uhJ0L3WTI3gAt2GhlKOBbeR63SJwIDAQAB
AoGAGY475sGSgd6WuArBHb46YzEkTZqj6VoMuNZqFFFP74vmPPxJaIx8P9U9XQHy
h79yzmhbJVIXdM9NmVncw0mbXLpqiRTJTQQKwH4aA8ai17C+CqKHVFo7nrw1oLrZ
lMi+/i+drMjXGYDPptfaojZCGGO9t2iATsua/ARMBiilSTECQQDgNex4qxphh/df
fqYHsoSfBeIVUIeD+bykqkkkB6FkykvZ/4mAsSXlJ6o/vf7gPR75t9QOrzM80I0N
t4H4jOp1AkEAwomM0npDDTDAk8aEHKR+QmgE/96CLW6s0oPiY8Qrfqv90Nwunopx
SS277dBaiFR6NOvnoUlSBgVMUKd8sFWeqwJBAKPTAtwQuY/2FywxkAMjz3+aft1w
H/Swr8PpoNOwt567qby52LCtv7C20NEOdINIZa+1QD8SiO+wusABC/iejlkCQQCO
ABWY9cHH+RQ5SijZN4EqO/+aPfNL8oOYTsiMBn+xQR5OvMoS0/+JaSbKmtHavb9O
rfwwEPMcjfV80iUEGhttAkEAvZpnxoAPk83tL7dqm7o1SWhhwVY7muaw2ppJWgRq
ek+i0cY4XILu+lbWtVj8x3PqR7+YLJnvs6jHG32h6wPfgA==
-----END RSA PRIVATE KEY-----`)

func GetOffsetFromHeader(h http.Header) int64 {
	byteRange := h.Get("range")
	if len(byteRange) < 7 {
		return 0
	}
	if byteRange[:6] != "bytes=" {
		return 0
	}
	bytePos := strings.Split(byteRange[6:], "-")
	offset, _ := strconv.ParseInt(bytePos[0], 0, 64)
	return offset
}

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256=" {
		return ""
	}
	return digest[8:]
}

func GetSizeFromHeader(h http.Header) int64 {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
	return size
}

func CalculateHash(r io.Reader) string {
	h := sha256.New()
	io.Copy(h, r)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func RsaEncrypt(origData []byte) ([]byte, error) {

	//解密pem格式的公钥

	block, _ := pem.Decode(PublicKey)

	if block == nil {

		return nil, errors.New("public key error")

	}

	// 解析公钥

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {

		return nil, err

	}

	// 类型断言

	pub := pubInterface.(*rsa.PublicKey)

	//加密

	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

}

// 解密

func RsaDecrypt(ciphertext []byte) ([]byte, error) {

	//解密

	block, _ := pem.Decode(PrivateKey)

	if block == nil {

		return nil, errors.New("private key error!")

	}

	//解析PKCS1格式的私钥

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {

		return nil, err

	}

	// 解密

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)

}
