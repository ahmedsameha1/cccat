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
		ccomand := exec.Command("./cccat", "hi_jp.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "こんにちは\n幸運を", out.String())
	})

	t.Run("Print file contents 3", func(t *testing.T) {
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

	t.Run("Print from the standard in 1", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 hi.txt | ./cccat -")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "hi\n", out.String())
	})

	t.Run("Print from the standard in 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 hi_jp.txt | ./cccat -")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "こんにちは\n", out.String())
	})

	t.Run("Print from the standard in 3", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat -")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "\"Your heart is the size of an ocean. Go find yourself in its hidden depths.\"\n", out.String())
	})

	t.Run("Print from the standard in 4", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "\"Your heart is the size of an ocean. Go find yourself in its hidden depths.\"\n", out.String())
	})

	t.Run("Print files contents concatenated 1", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "hi.txt", "hello.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "hi\nhellohello\nhi", out.String())
	})

	t.Run("Print files contents concatenated 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "hi_jp.txt", "hello.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "こんにちは\n幸運をhello\nhi", out.String())
	})

	t.Run("Print files contents concatenated 3", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "quotes.txt", "quotes2.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, `"Your heart is the size of an ocean. Go find yourself in its hidden depths."
"The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."
"Thinking is the capital, Enterprise is the way, Hard Work is the solution."
"If You Can'T Make It Good, At Least Make It Look Good."
"Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
"It is bad for a young man to sin; but it is worse for an old man to sin."
"If You Are Out To Describe The Truth, Leave Elegance To The Tailor."
"O man you are busy working for the world, and the world is busy trying to turn you out."
"While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
"These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
"I Don'T Believe In Failure. It Is Not Failure If You Enjoyed The Process."
"Do not get elated at any victory, for all such victory is subject to the will of God."
"Wear gratitude like a cloak and it will feed every corner of your life."
"If you even dream of beating me you'd better wake up and apologize."
"I Will Praise Any Man That Will Praise Me."
"One Of The Greatest Diseases Is To Be Nobody To Anybody."
"I'm so fast that last night I turned off the light switch in my hotel room and was in bed before the room was dark."
"People Must Learn To Hate And If They Can Learn To Hate, They Can Be Taught To Love."
"Everyone has been made for some particular work, and the desire for that work has been put in every heart."
"The less of the World, the freer you live."
`, out.String())
	})

	t.Run("Print files contents concatenated 4", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat hi.txt - hello.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "hi\nhello\"Your heart is the size of an ocean. Go find yourself in its hidden depths.\"\nhello\nhi",
			out.String())
	})
	
	t.Run("Print files contents concatenated 5", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat hi.txt - hi_jp.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "hi\nhello\"Your heart is the size of an ocean. Go find yourself in its hidden depths.\"\nこんにちは\n幸運を",
			out.String())
	})

	t.Run("Print file contents with numbers for lines", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n3 quotes.txt | ./cccat -n -")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, `1 "Your heart is the size of an ocean. Go find yourself in its hidden depths."
2 "The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."
3 "Thinking is the capital, Enterprise is the way, Hard Work is the solution."
`,
			out.String())
	})

	t.Run("Print file contents with numbers for lines 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 test.txt | ./cccat -n")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, "", out.String())
	})

	t.Run("Print file contents with numbers for non-empty lines 1", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "-b", "quotes_emptylines.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t,
			`1 "Your heart is the size of an ocean. Go find yourself in its hidden depths."

2 "The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."

3 "Thinking is the capital, Enterprise is the way, Hard Work is the solution."

4 "If You Can'T Make It Good, At Least Make It Look Good."
5 "Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
6 "It is bad for a young man to sin; but it is worse for an old man to sin."
7 "If You Are Out To Describe The Truth, Leave Elegance To The Tailor."


8 "O man you are busy working for the world, and the world is busy trying to turn you out."
9 "While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
10 "These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
`,
			out.String())
	})

	t.Run("Print file contents with numbers for non-empty lines 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "-nb", "quotes_emptylines.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t,
			`1 "Your heart is the size of an ocean. Go find yourself in its hidden depths."

2 "The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."

3 "Thinking is the capital, Enterprise is the way, Hard Work is the solution."

4 "If You Can'T Make It Good, At Least Make It Look Good."
5 "Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
6 "It is bad for a young man to sin; but it is worse for an old man to sin."
7 "If You Are Out To Describe The Truth, Leave Elegance To The Tailor."


8 "O man you are busy working for the world, and the world is busy trying to turn you out."
9 "While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
10 "These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
`,
			out.String())
	})

	t.Run("Print file contents with numbers for non-empty lines 3", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("./cccat", "-bn", "quotes_emptylines.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t,
			`1 "Your heart is the size of an ocean. Go find yourself in its hidden depths."

2 "The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."

3 "Thinking is the capital, Enterprise is the way, Hard Work is the solution."

4 "If You Can'T Make It Good, At Least Make It Look Good."
5 "Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
6 "It is bad for a young man to sin; but it is worse for an old man to sin."
7 "If You Are Out To Describe The Truth, Leave Elegance To The Tailor."


8 "O man you are busy working for the world, and the world is busy trying to turn you out."
9 "While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
10 "These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
`,
			out.String())
	})

	t.Run("Print files contents concatenated with numbers for lines 1", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat -n hi.txt - hello.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, `1 hi
2 hello"Your heart is the size of an ocean. Go find yourself in its hidden depths."
3 hello
4 hi`,
			out.String())
	})

	t.Run("Print files contents concatenated with numbers for lines 2", func(t *testing.T) {
		var out strings.Builder
		var errOut strings.Builder
		ccomand := exec.Command("bash", "-c", "head -n1 quotes.txt | ./cccat -b quotes_emptylines.txt - hello.txt")
		ccomand.Dir = "./.."
		ccomand.Stderr = &errOut
		ccomand.Stdout = &out
		err := ccomand.Run()
		assert.NoError(t, err)
		assert.Equal(t, `1 "Your heart is the size of an ocean. Go find yourself in its hidden depths."

2 "The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."

3 "Thinking is the capital, Enterprise is the way, Hard Work is the solution."

4 "If You Can'T Make It Good, At Least Make It Look Good."
5 "Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
6 "It is bad for a young man to sin; but it is worse for an old man to sin."
7 "If You Are Out To Describe The Truth, Leave Elegance To The Tailor."


8 "O man you are busy working for the world, and the world is busy trying to turn you out."
9 "While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
10 "These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
11 "Your heart is the size of an ocean. Go find yourself in its hidden depths."
12 hello
13 hi`,
out.String())
	})
}