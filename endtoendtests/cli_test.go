package endtoendtests

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	buildCommand := exec.Command("mosmlc", "-P", "full", "-o", "cccat", "cccat.sml")
	buildCommand.Dir = "./.."
	err := buildCommand.Run()
	if err != nil {
		panic(err)
	}

	result := m.Run()

	removeCommand := exec.Command("rm", "cccat")
	removeCommand.Dir = "./.."
	err = removeCommand.Run()
	if err != nil {
		panic(err)
	}
	os.Exit(result)
}

func TestCccat(t *testing.T) {
	t.Run("Print file contents 1", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "hi.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "hi\nhello", out.String())
	})

	t.Run("Print file contents 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "quotes.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t,
			`"Your heart is the size of an ocean. Go find yourself in its hidden depths."
"The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."
"Thinking is the capital, Enterprise is the way, Hard Work is the solution."
"If You Can'T Make It Good, At Least Make It Look Good."
"Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
"It is bad for a young man to sin; but it is worse for an old man to sin."
"If You Are Out To Describe The Truth, Leave Elegance To The Tailor."
"O man you are busy working for the world, and the world is busy trying to turn you out."
"While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
"These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
`,
			out.String())
	})
}
