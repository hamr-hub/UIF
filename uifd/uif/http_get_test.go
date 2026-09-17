package uif

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSecureSubscribeURL(t *testing.T) {
	assert.True(t, isSecureSubscribeURL("https://fnsuburl.kilxs.cn/api/v1/client/secureSubscribe?token=abc"))
	assert.True(t, isSecureSubscribeURL("https://example.com/API/V1/CLIENT/SECURESUBSCRIBE?token=abc"))
	assert.False(t, isSecureSubscribeURL("https://example.com/api/v1/client/subscribe?token=abc"))
	assert.False(t, isSecureSubscribeURL("://bad-url"))
}
