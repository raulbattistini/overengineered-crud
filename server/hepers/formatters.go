package hepers

import (
	"math/rand"
	"server/globals"
	"server/types"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/internal/uuid"
	"golang.org/x/text/unicode/norm"
)

func GenRandomNumber(min, max int) int {
	return 0
}

func GenRandomPostTitle(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}

func GenRandomPostContent(length int) string {
	if length < 100 {
		length = 100
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}

func GenUuidStr() string {
	return uuid.New().String()
}

func CleanAllInput(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\n", "")
	input = norm.NFD.String(input)
	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, input)
}

func CleanInputWithLineBreaks(input string) string {
	input = strings.TrimSpace(input)
	input = norm.NFD.String(input)
	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, input)
}

func CleanPost(pst types.Post) types.Post {
	title := pst.Title.(string)
	title = CleanAllInput(pst.Title.(string))
	pst.Title = title
	pst.Content = CleanInputWithLineBreaks(pst.Content)
	return pst
}

func CleanPostId(id string) (*int, error) {
	id = CleanAllInput(id)
	if id == "" {
		return &globals.EmptyId, nil
	}
	fmtId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return &fmtId, nil
}
