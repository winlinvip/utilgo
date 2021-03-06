package urlutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upyun/utilgo/urlutil"
)

func TestParseRawStreamURL(t *testing.T) {
	{
		vhost, app, stream := urlutil.ParseRawStreamURL("aaa.com/aa/bb")
		assert.Equal(t, "aaa.com", vhost, "vhost should be parsed OK")
		assert.Equal(t, "aa", app, "app should be parsed OK")
		assert.Equal(t, "bb", stream, "stream should be parsed OK")
	}

	{
		vhost, app, stream := urlutil.ParseRawStreamURL("aaa.com/aa/bb/cc")
		assert.Equal(t, "aaa.com", vhost, "vhost should be parsed OK")
		assert.Equal(t, "aa", app, "app should be parsed OK")
		assert.Equal(t, "bb/cc", stream, "stream should be parsed OK")
	}
}
