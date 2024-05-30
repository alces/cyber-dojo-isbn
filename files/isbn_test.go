package isbn

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestVefify(t *testing.T) {
    assert.True(t, Verify("0-321-14653-0"))
}
