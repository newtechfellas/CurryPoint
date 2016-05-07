package util

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"strings"
)

func SimpleJsonResponse(w http.ResponseWriter, statusCode int) {
	JsonResponse(w, nil, nil, statusCode)
}

func JsonResponse(w http.ResponseWriter, v interface{}, headers map[string]string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	//Any custom headers passed in
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(statusCode)
	if v != nil {
		b, _ := json.Marshal(v)
		fmt.Fprintf(w, "%s", string(b[:]))
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err != nil {
		b, _ := json.Marshal(map[string]interface{}{
			"message": err.Error(),
		})
		fmt.Fprintf(w, "%s", string(b[:]))
	}
}

func IsLocalHost(r *http.Request) bool {
	return strings.Contains(r.Host,"localhost:")
}

func Jsonify(obj interface{}) string {
	b, err := json.MarshalIndent(obj, " ", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func RemoveDuplicates(s []string) (t []string) {
	m := map[string]bool{}
	// walk the slice s and for each value we haven't seen so far, append it to t
	// this has the benefit of being clearer, not mutting the original underlyin array
	// as well as not hanging on to memory needlessly if the slice has few unique values
	for _, v := range s {
		if _, seen := m[v]; !seen {
			t = append(t, v)
			m[v] = true
		}
	}
	return t
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
)

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

func SecureRandomAlphaString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

func NewVerificationCode() string {
	return SecureRandomAlphaString(3) + strconv.Itoa(Random6DigitNumber())
}

func Random6DigitNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := r.Intn(999999)
	if v == 0 {
		v = Random6DigitNumber()
	}
	return v
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func LogDebug(v ...interface{}) {
	log.Println("DEBUG: ", v)
}

func LogWarning(v ...interface{}) {
	log.Println("WARNING: ", v)
}

func LogError(v ...interface{}) {
	log.Println("ERROR: ", v)
}