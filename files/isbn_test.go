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
        assert.Equal(t, r.expected, cleanUp(r.argument), r.argument)
    }
}

var isISBN10TestResults = []struct {
    argument string
    expected bool
} {
    {"0470845252", true},
    {"470845252", false},
    {"47084525211", false},
    {"847084525X", true},
    {"X470845258", false},
    {"47084525X", false},
    {"4700840525X", false},
}

func TestIsISBN10(t *testing.T) {
    for _, r := range isISBN10TestResults {
        assert.Equal(t, r.expected, isISBN10(r.argument), r.argument)
    }
}

var isISBN13TestResults = []struct {
    argument string
    expected bool
} {
    {"9780470059029", true},
}

func TestIsISBN10(t *testing.T) {
    for _, r := range isISBN13TestResults {
        assert.Equal(t, r.expected, isISBN13(r.argument), r.argument)
    }
}

func TestVefify(t *testing.T) {
    assert.True(t, Verify("0-321-14653-0"))
}
