package isbn

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCleanUp(t *testing.T) {
    assert.Equal(t, "0470845252", cleanUp("0-470-84525-2"))

func TestVefify(t *testing.T) {
    assert.True(t, Verify("0-321-14653-0"))
}
