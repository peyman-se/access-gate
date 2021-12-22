package gate

// Gate is a gate. It can be used to check if a user has access to a resource.
type Gate struct {
	Rules map[string]func(...interface{}) bool
}

// NewGate creates a new gate.
func NewGate() *Gate {
	return &Gate{
		Rules: make(map[string]func(args ...interface{}) bool),
	}
}

//Define defines a new rule for the gate.
func (g *Gate) Define(name string, rule func(arg ...interface{}) bool) {
	g.Rules[name] = rule
}

// Can checks if the rule is defined and returns the result of the rule.
// If the rule is not defined, it returns false.
func (g *Gate) Can(name string, args ...interface{}) bool {
	if rule, ok := g.Rules[name]; ok {
		return rule(args...)
	}
	return false
}
