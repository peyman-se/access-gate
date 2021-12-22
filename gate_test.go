package gate

import "testing"

func TestDefiningARuleWillAddThatToGate(t *testing.T) {
	g := NewGate()
	g.Define("foo-bar", func(args ...interface{}) bool {
		a := args[0].(string)
		b := args[1].(string)
		return a == b
	})

	if len(g.Rules) != 1 {
		t.Error("Rule should be added to gate")
	}
}

func TestGatePasses(t *testing.T) {
	g := NewGate()
	g.Define("foo-bar", func(args ...interface{}) bool {
		a := args[0].(string)
		b := args[1].(string)
		return a == b
	})

	a := "randomstring"
	b := "randomstring"
	if !g.Can("foo-bar", a, b) {
		t.Error("Access denied but it should have passed")
	}
}
