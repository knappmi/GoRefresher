package reverseunicode

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("ab") != "ba" { t.Fatal("simple reverse failed") }
	if Reverse("你好") != "好你" { t.Fatal("unicode reverse failed") }
}
