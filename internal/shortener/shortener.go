package shortener

import (
	"hash/fnv"
	"url_shortner/internal/utils"
)

type URL struct {
	Original string
	Short    string
}

var urls map[string]*URL = make(map[string]*URL)

func generateShortURL(long string) (string, error) {
	number := fnv.New32a()
	_, err := number.Write([]byte(long))
	if err != nil {
		return "", err
	}

	short := utils.Encode62(uint64(number.Sum32()))

	return short, nil
}

// ------------
//
//	EXPORTED
//
// ------------
func Find(short string) (*URL, bool) {
	u, exists := urls[short]
	return u, exists
}

// Constructor
func NewURL(original string) (*URL, error) {
	short, err := generateShortURL(original)
	if err != nil {
		return nil, err
	}

	url := &URL{
		Original: original,
		Short:    short,
	}

	urls[short] = url

	return url, nil
}

// Receiver fn
func (url *URL) Info() string {
	return "Short: " + url.Short + " points to: " + url.Original
}
