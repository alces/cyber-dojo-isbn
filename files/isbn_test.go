package isbn

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var cleanUpTestResults = []struct{
    argument string
    expected string
} {
    {"0470845252", "0470845252"},
    {"0-470-84525-2", "0470845252"},
    {"0 470 84525 2", "0470845252"},
}

func TestCleanUp(t *testing.T) {
    for _, r := range cleanUpTestResults {
        assert.Equal(t, r.expected, cleanUp(r.argument))
    }
}

func TestIsISBN10(t *testing.T) {
    assert.True(t, isISBN10("0470845252"))  
}

func TestVefify(t *testing.T) {
    assert.True(t, Verify("0-321-14653-0"))
}
