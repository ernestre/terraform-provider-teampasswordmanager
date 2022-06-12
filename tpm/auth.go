package tpm

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

const (
	AuthHeaderPublicKey        = "X-Public-Key"
	AuthHeaderRequestHash      = "X-Request-Hash"
	AuthHeaderRequestTimestamp = "X-Request-Timestamp"
)

// https://teampasswordmanager.com/docs/api-authentication/
func generateAuthHash(endpoint string, timestamp int64, requestBody []byte, privateKey string) string {
	data := fmt.Sprintf("%s%d%s", endpoint, timestamp, string(requestBody))

	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

func generateAuthHeaders(
	publicKey string,
	hash string,
	timestamp int64,
) map[string]string {
	return map[string]string{
		AuthHeaderPublicKey:        publicKey,
		AuthHeaderRequestHash:      hash,
		AuthHeaderRequestTimestamp: strconv.Itoa(int(timestamp)),
	}
}
