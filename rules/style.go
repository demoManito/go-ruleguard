package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary reports redundant parentheses
//doc:before  f(x, (y))
//doc:after   f(x, y)
//doc:tags    style
func exprUnparen(m dsl.Matcher) {
	m.Match(`$f($*_, ($x), $*_)`).
		Report(`the parentheses around $x are superfluous`).
		Suggest(`$f($x)`)
}

//doc:summary reports empty declaration blocks
//doc:before  var ()
//doc:after   /* nothing */
//doc:tags    style
func emptyDecl(m dsl.Matcher) {
	m.Match(`var()`).Report(`empty var() block`)
	m.Match(`const()`).Report(`empty const() block`)
	m.Match(`type()`).Report(`empty type() block`)
}

//doc:summary reports empty errors creation
//doc:before  errors.New("")
//doc:after   errors.New("can't open the cache file")
//doc:tags    style
func emptyError(m dsl.Matcher) {
	m.Match(`fmt.Errorf("")`, `errors.New("")`).
		Report(`empty errors are hard to debug`)
}

//doc:summary reports empty slice declaration
//doc:before  x := []int{}
//doc:after   var x []int
//doc:tags    style
func emptySlice(m dsl.Matcher) {
	m.Match(`var $name = make([]$type, 0)`, `$name := []$type{}`, `$name := make([]$type, 0, 0)`, `$name := make([]$type, 0)`).
		Report(`zero-length slice declaring nil slice is better`).
		Suggest(`var $name []$type`)
}
