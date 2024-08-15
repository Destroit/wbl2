package main
import "testing"

func testUnpack(input []rune, want []rune, t *testing.T) {
    got, _ := unpack(input)
    if string(got) != string(want) {
	t.Errorf("Result was incorrect, got: %q, want: %q", got, want)
    }
}


func TestNothing(t *testing.T) {
    input := []rune("abcd^&.")
    want := []rune("abcd^&.")
    testUnpack(input, want, t)
}

func TestBasic(t *testing.T) {
    input := []rune("a0b1c2d3")
    want := []rune("bccddd")
    testUnpack(input, want, t)
}

func TestEscape(t *testing.T) {
    input := []rune("a//b/1")
    want := []rune("a/b1")
    testUnpack(input, want, t)
}

func TestUnicode(t *testing.T) {
    input := []rune("Ʃ2Ъ1Ő3ы")
    want := []rune("ƩƩЪŐŐŐы")
    testUnpack(input, want, t)
}

func TestEmpty(t *testing.T) {
    input := []rune("")
    want := []rune("")
    testUnpack(input, want, t)
}

func TestInvalid(t *testing.T) {
    _, err := unpack([]rune("1a23"))
    if err == nil {
	t.Errorf("Should be an error(starts with number)")
    }
}
