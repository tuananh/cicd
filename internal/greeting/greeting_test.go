package greeting_test

import (
	"testing"

	"github.com/calmonr/cicd/internal/greeting"
	"github.com/stretchr/testify/assert"
)

func TestHi(t *testing.T) {
	assert.Equal(t, "Hi, Calmon Ribeiro", greeting.Hi("Calmon Ribeiro"))
}
