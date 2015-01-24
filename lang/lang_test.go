package lang

import (
	"testing"
)

func TestLang(t *testing.T) {
	s := `import map "github.com/awalterschulze/picolang/funcs.Map"
		import inc "github.com/awalterschulze/picolang/funcs.Inc"

	let a = inc(2)

	let c = map(inc, {2.3, 4})

	println(c)
	`
	ss, err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", ss[0].Generate())
}
