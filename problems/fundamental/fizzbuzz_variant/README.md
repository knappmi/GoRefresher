# FizzBuzz Variant

Goal: Implement an extensible FizzBuzz style engine.

Requirements:
- Function Run(n int, rules []Rule) []string
- Rule: Divisor int, Word string, Precedence (lower number = higher priority) optional stretch.
- If number matches multiple divisibility rules, concatenate words in rule slice order (respect precedence if implemented).
- If no rule matches, emit the number.
- Support dynamic addition/removal of rules.
- Provide example rules (3->"Fizz", 5->"Buzz", 7->"Pop").

Edge Cases:
- n <= 0: return empty slice.
- Divisor == 0: ignore rule (avoid panic) and return error from validation.

Testing:
- Table tests for standard set up to 21.
- Test custom ordering & overlapping divisors (e.g., 15 => FizzBuzz).
- Benchmark default rules for n=10_000.

Stretch Ideas:
- Allow predicate based rules (func(int) bool) via interface.
- Add concurrent version producing a channel.
- CLI: accept rules via flags (e.g., -rule 3:Fizz -rule 5:Buzz).
