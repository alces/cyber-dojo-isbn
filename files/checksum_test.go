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
    {[]byte{0, 4, 7, 1, 6, 0, 6, 9, 5}, 2},
}

func TestISBN10Checksum(t *testing.T) {
    for _, r := range isbn10ChecksumTestResults {
        assert.Equal(t, r.expected, isbn10Checksum(r.argument), r.expected)
    }
}

var isbn13ChecksumTestResults = []struct {
    argument []byte
    expected byte
} {
    {[]byte{9, 7, 8, 0, 4, 7, 0, 0, 5, 9, 0, 2}, 9},
    {[]byte{9, 7, 8, 0, 5, 9, 6, 8, 0, 9, 4, 8}, 5},
    {[]byte{9, 7, 8, 0, 1, 3, 1, 4, 9, 5, 0, 5}, 0},
}

func TestISBN13Checksum(t *testing.T) {
    for _, r := range isbn13ChecksumTestResults {
        assert.Equal(t, r.expected, isbn13Checksum(r.argument), r.expected)
    }
}

var verifyChecksumTestResults = []struct {
    argument string
    expected bool
} {

    {"0471958697", true},
    {"0471958698", false},
    {"9780470059029", true},
    {"9780470059026", false},
    {"980470059029", false},
}

func TestVerifyChecksum(t *testing.T) {
    for _, r := range verifyChecksumTestResults {
        assert.Equal(t, r.expected, verifyChecksum(r.argument), r.argument)
    }
}