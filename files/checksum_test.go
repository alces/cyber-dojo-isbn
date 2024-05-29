import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestISBN10Checksum(t *testing.T) {
    assert.Equals(t, 7, []byte{0, 4, 7, 1, 9, 5, 8, 6, 9}) 
}    