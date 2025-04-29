package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeaders(t *testing.T) {
	for scenario, fn := range map[string]func(
		tt *testing.T,
	){
		"Default Headers Contain Content-type": testDefaultHeadersContainContentType,
		"Adding Headers Includes Content-type": testAddingHeadersIncludesContentType,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testDefaultHeadersContainContentType(t *testing.T) {
	headers := StandardResponseHeaders(NoHeaders)

	_, ok := headers[HeaderContentType]
	assert.True(t, ok)
}

func testAddingHeadersIncludesContentType(t *testing.T) {
	var ok bool
	custom := make(Headers)
	custom[HeaderAccept] = "text/html"
	headers := StandardResponseHeaders(custom)

	_, ok = headers[HeaderContentType]
	assert.True(t, ok)
	_, ok = headers[HeaderAccept]
	assert.True(t, ok)
}
