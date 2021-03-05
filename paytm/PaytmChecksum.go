/**
 * Paytm uses checksum signature to ensure that API requests and responses shared between your
 * application and Paytm over network have not been tampered with. We use SHA256 hashing and
 * AES128 encryption algorithm to ensure the safety of transaction data.
 *
 * @author     Lalit Kumar
 * @version    2.0
 * @link       https://developer.paytm.com/docs/checksum/#go
 */

package PaytmChecksum

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var IV = "@@@@&&&&####$$$$"

func Encrypt(input string, key string) (string, error) {

	_key := []byte(key)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}

	b := []byte(input)
	b = pkcs5Pad(b, aes.BlockSize, len(input))
	ciphertext := make([]byte, len(b))

	_iv := []byte(IV)
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(ciphertext, b)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted string, key string) (string, error) {

	_key := []byte(key)
	text, _ := base64.StdEncoding.DecodeString(encrypted)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	_iv := []byte(IV)
	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(decrypted, text)

	return string(pkcs5Unpad(decrypted)), nil
}

func GenerateSignature(params map[string]string, key string) (string, error) {
	sorted_string := getStringByParams(params)
	return GenerateSignatureByString(sorted_string, key)
}
func VerifySignature(params map[string]string, key string, checksum string) (bool, error) {
	if _, ok := params["CHECKSUMHASH"]; ok {
		delete(params, "CHECKSUMHASH")
	}
	sorted_string := getStringByParams(params)
	return VerifySignatureByString(sorted_string, key, checksum)
}

func GenerateSignatureByString(params string, key string) (string, error) {
	salt := generateRandomString(4)
	calChckSum, err := calculateChecksum(params, key, salt)
	if err != nil {
		return "", err
	}
	return calChckSum, nil
}

func VerifySignatureByString(params string, key string, checksum string) (bool, error) {
	paytm_hash, err := Decrypt(checksum, key)
	if err != nil {
		return false, err
	}
	salt := paytm_hash[len(paytm_hash)-4:]
	calChckSum, err := calculateChecksum(params, key, salt)
	if err != nil {
		return false, err
	}

	return (checksum == calChckSum), nil
}

func generateRandomString(length int) string {
	data := []byte("9876543210ZYXWVUTSRQPONMLKJIHGFEDCBAabcdefghijklmnopqrstuvwxyz!@#$&_")
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = data[rand.Intn(len(data))]
	}
	return string(b)
}

func getStringByParams(params map[string]string) string {
	sorted_keys := make([]string, 0, len(params))
	for k := range params {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	sorted_values := make([]string, 0, len(params))
	for _, element := range sorted_keys {
		param := params[element]
		if strings.ToLower(params[element]) == "null" {
			param = ""
		}
		sorted_values = append(sorted_values, param)
	}
	return strings.Join(sorted_values, "|")
}

func calculateHash(params string, salt string) string {
	finalString := params + "|" + salt
	hash := sha256.New()
	hash.Write([]byte(finalString))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString + salt
}

func calculateChecksum(params string, key string, salt string) (string, error) {
	hashString := calculateHash(params, salt)
	checksum, err := Encrypt(hashString, key)
	if err != nil {
		return "", err
	}
	return checksum, nil
}

func pkcs5Pad(text []byte, blocksize int, after int) []byte {
	padding := (blocksize - len(text)%blocksize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func pkcs5Unpad(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}
