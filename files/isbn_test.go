package isbn

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var verifyTestResults = []struct {
    argument string
    expected bool
} {
    {"0-321-14653-0", true},
    {"0 471 60695 2", true},
    {"0 471 60695 3", false},
    {"something strange", false},
    {"978-0-13-149505-0", true},
    {"978-0-13-149515-0", false},
}

func TestVefify(t *testing.T) {
    for _, r := range verifyTestResults {
        assert.Equal(t, r.expected, Verify(r.argument), r.argument)
    }
}
