package isbn

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var isbn10ChecksumTestResults = []struct {
    argument []byte
    expected byte
} {
    {[]byte{0, 4, 7, 1, 9, 5, 8, 6, 9}, 7},
}

func TestISBN10Checksum(t *testing.T) {
    for _, r := range isbn10ChecksumTestResults {
        assert.Equal(t, r.expected, isbn10Checksum(r.argument), r.expected)
    }
}    