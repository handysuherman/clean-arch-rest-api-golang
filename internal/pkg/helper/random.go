package helper

import (
	"crypto/ed25519"
	crypRand "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")
var unescaper = strings.NewReplacer("99", "9", "90", "-", "91", "_")

const (
	Environment = "local"
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	numeric     = "123456789"
)

func init() {
	seed, err := secureRandomInt64()
	if err != nil {
		// Handle the error, e.g., by falling back to a less secure method
		seed = time.Now().UnixNano()
	}

	rand.New(rand.NewSource(seed))
	// rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
}

func secureRandomInt64() (int64, error) {
	maxInt64 := big.NewInt(1<<63 - 1)
	seed, err := crypRand.Int(crypRand.Reader, maxInt64)
	if err != nil {
		return 0, err
	}
	return seed.Int64(), nil
}

// RandomInt generate a rando minteger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min) + 1
}

// RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomStringSelection(item ...string) string {
	n := len(item)
	return item[rand.Intn(n)]
}

func RandomStringInt(n int) string {
	var sb strings.Builder
	k := len(numeric)

	for i := 0; i < n; i++ {
		c := numeric[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPhoneNumber() string {
	var sb strings.Builder
	k := len(numeric)

	for i := 0; i < 10; i++ {
		c := numeric[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return "628" + sb.String()
}

// RandomName generate a random name
func RandomName() string {
	return RandomString(12)
}

// RandomNumber generate a random amount of number
func RandomNumber() int64 {
	return RandomInt(100000000000, 999999999999)
}

// RandomCurrencies generate a random currencies
func RandomCurrencies() string {
	currencies := []string{"IDR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomBool() bool {
	bools := []bool{true, false}
	n := len(bools)
	return bools[rand.Intn(n)]
}

func RandomUrl() string {
	return fmt.Sprintf("https://%s.com", RandomString(6))
}

// RandomEmail generate a random email string
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomUUID generate a random uuid
func RandomUUID() uuid.UUID {
	id, _ := uuid.NewRandom()
	return id
}

func RandomKeyPair() (privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) {

	generatedPublicKey, generatedPrivateKey, _ := ed25519.GenerateKey(crypRand.Reader)

	return generatedPrivateKey, generatedPublicKey
}

func Escaper() {
	id, _ := uuid.New().MarshalBinary()
	ori, _ := uuid.FromBytes(id)
	rawEnc := escaper.Replace(base64.RawURLEncoding.EncodeToString(id))

	dec, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(rawEnc))
	if err != nil {
		fmt.Print(err)
	}
	decOri, _ := uuid.FromBytes(dec)

	fmt.Printf("original val = %v\n", ori)
	fmt.Printf("encrypted val = %v\n", rawEnc)
	fmt.Printf("decoded val = %v\n", decOri)
}

func RandomUUIDEncoding() string {
	id, _ := uuid.New().MarshalBinary()

	rawUrl := escaper.Replace(base64.RawURLEncoding.EncodeToString(id))

	return rawUrl
}

func UUIDEncode(uid uuid.UUID) string {
	id, _ := uid.MarshalBinary()

	raw := escaper.Replace(base64.RawURLEncoding.EncodeToString(id))

	return raw
}

func UUIDDecode(raw string) (uuid.UUID, bool) {
	dec, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(raw))
	if err != nil {
		return uuid.Nil, false
	}
	id, err := uuid.FromBytes(dec)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}

func UUIDChecker(uid string) (uuid.UUID, bool) {
	dec, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(uid))
	if err != nil {
		return uuid.Nil, false
	}
	id, err := uuid.FromBytes(dec)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}

func GenerateULID() (ulid.ULID, error) {
	seed, err := secureRandomInt64()
	if err != nil {
		seed = time.Now().UnixNano()
	}
	entropy := rand.New(rand.NewSource(seed))
	ms := ulid.Timestamp(time.Now())

	s, err := ulid.New(ms, entropy)
	if err != nil {
		return ulid.ULID{}, err
	}

	return s, nil
}
