package fizzbuzz

import "testing"

func TestRunBasic(t *testing.T) {
	got := Run(21, DefaultRules())
	want15 := "FizzBuzz"
	if got[14] != want15 { t.Fatalf("15 => %s want %s", got[14], want15) }
	// 21 divisible by 3 and 7 => FizzPop
	if got[20] != "FizzPop" { t.Fatalf("21 => %s want FizzPop", got[20]) }
}
