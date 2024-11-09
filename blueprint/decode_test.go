package blueprint

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBook(t *testing.T) {
	resp, err := http.Get("https://raw.githubusercontent.com/raynquist/balancer/0bdafe7e7fe3137f17b1b39242c66152ad19094a/blueprints/balancer_book.txt")
	assert.NoError(t, err)

	var b Book
	err = Decode(resp.Body, &b)
	assert.NoError(t, err)

	assert.NotZero(t, len(b.Blueprints))
}
