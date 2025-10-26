package fizzbuzz

import "fmt"

// Rule defines a divisibility mapping. TODO: extend with precedence if desired.
type Rule struct { Divisor int; Word string }

// Run returns a slice of outputs for numbers 1..n based on rules.
func Run(n int, rules []Rule) []string {
	if n <= 0 { return []string{} }
	valid := make([]Rule, 0, len(rules))
	for _, r := range rules { if r.Divisor > 0 && r.Word != "" { valid = append(valid, r) } }
	out := make([]string, n)
	for i:=1; i<=n; i++ {
		var s string
		for _, r := range valid { if i%r.Divisor==0 { s += r.Word } }
		if s=="" { s = fmt.Sprintf("%d", i) }
		out[i-1] = s
	}
	return out
}

// DefaultRules returns classic FizzBuzz+Pop rules.
func DefaultRules() []Rule { return []Rule{{3,"Fizz"},{5,"Buzz"},{7,"Pop"}} }
