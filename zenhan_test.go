package zenhan

import (
	"testing"
)

const original = "ﾟabcＤＥﾞＦ123４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"

func TestH2zAsciiOnly(t *testing.T) {
	actual := H2z(original, ASCII)
	expected := "ﾟａｂｃＤＥﾞＦ123４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zDigitOnly(t *testing.T) {
	actual := H2z(original, DIGIT)
	expected := "ﾟabcＤＥﾞＦ１２３４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zKanaOnly(t *testing.T) {
	actual := H2z(original, KANA)
	expected := "゜abcＤＥ゛Ｆ123４５６アガサダナバビプペ゜"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zAsciiAndDigit(t *testing.T) {
	actual := H2z(original, ASCII|DIGIT)
	expected := "ﾟａｂｃＤＥﾞＦ１２３４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zAsciiAndKana(t *testing.T) {
	actual := H2z(original, ASCII|KANA)
	expected := "゜ａｂｃＤＥ゛Ｆ123４５６アガサダナバビプペ゜"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zDigitAndKana(t *testing.T) {
	actual := H2z(original, DIGIT|KANA)
	expected := "゜abcＤＥ゛Ｆ１２３４５６アガサダナバビプペ゜"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zIgnore(t *testing.T) {
	actual := H2z(original, ALL, "a", "1", "2")
	expected := "゜aｂｃＤＥ゛Ｆ12３４５６アガサダナバビプペ゜"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zInvalid(t *testing.T) {
	text := string([]byte{0xff, 0xfe, 0xfd})
	actual := H2z(text, ALL)
	expected := text
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestH2zEmpty(t *testing.T) {
	text := ""
	actual := H2z(text, ALL)
	expected := text
	if expected != actual {
		t.Errorf("Expected empty, but %s found", actual)
	}
}

func TestH2zAll(t *testing.T) {
	actual1 := H2z(original, ALL)
	expected := "゜ａｂｃＤＥ゛Ｆ１２３４５６アガサダナバビプペ゜"
	if expected != actual1 {
		t.Errorf("Expected %s, but %s found", expected, actual1)
	}
	actual2 := H2z(original, ASCII|DIGIT|KANA)
	if expected != actual2 {
		t.Errorf("Expected %s, but %s found", expected, actual2)
	}
}

func TestZ2hAsciiOnly(t *testing.T) {
	actual := Z2h(original, ASCII)
	expected := "ﾟabcDEﾞF123４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hDigitOnly(t *testing.T) {
	actual := Z2h(original, DIGIT)
	expected := "ﾟabcＤＥﾞＦ123456ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hKanaOnly(t *testing.T) {
	actual := Z2h(original, KANA)
	expected := "ﾟabcＤＥﾞＦ123４５６ｱｶﾞｻﾀﾞﾅﾊﾞﾋﾞﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hAsciiAndDigit(t *testing.T) {
	actual := Z2h(original, ASCII|DIGIT)
	expected := "ﾟabcDEﾞF123456ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hAsciiAndKana(t *testing.T) {
	actual := Z2h(original, ASCII|KANA)
	expected := "ﾟabcDEﾞF123４５６ｱｶﾞｻﾀﾞﾅﾊﾞﾋﾞﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hDigitAndKana(t *testing.T) {
	actual := Z2h(original, DIGIT|KANA)
	expected := "ﾟabcＤＥﾞＦ123456ｱｶﾞｻﾀﾞﾅﾊﾞﾋﾞﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hIgnore(t *testing.T) {
	actual := Z2h(original, ALL, "Ｄ", "４", "５")
	expected := "ﾟabcＤEﾞF123４５6ｱｶﾞｻﾀﾞﾅﾊﾞﾋﾞﾌﾟﾍﾟﾟ"
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hInvalid(t *testing.T) {
	text := string([]byte{0xff, 0xfe, 0xfd})
	actual := Z2h(text, ALL)
	expected := text
	if expected != actual {
		t.Errorf("Expected %s, but %s found", expected, actual)
	}
}

func TestZ2hEmpty(t *testing.T) {
	text := ""
	actual := Z2h(text, ALL)
	expected := text
	if expected != actual {
		t.Errorf("Expected empty, but %s found", actual)
	}
}

func TestZ2hAll(t *testing.T) {
	actual1 := Z2h(original, ALL)
	expected := "ﾟabcDEﾞF123456ｱｶﾞｻﾀﾞﾅﾊﾞﾋﾞﾌﾟﾍﾟﾟ"
	if expected != actual1 {
		t.Errorf("Expected %s, but %s found", expected, actual1)
	}
	actual2 := Z2h(original, ASCII|DIGIT|KANA)
	if expected != actual2 {
		t.Errorf("Expected %s, but %s found", expected, actual2)
	}
}

func getAllHankakuAlpha() []string {
	allHankakuAlpha := make([]string, 0, 52)
	// a to z
	for i := 97; i < 123; i++ {
		allHankakuAlpha = append(allHankakuAlpha, string(i))
	}
	// A to Z
	for i := 65; i < 91; i++ {
		allHankakuAlpha = append(allHankakuAlpha, string(i))
	}
	return allHankakuAlpha
}

func getAllZenkakuAlpha() []string {
	allZenkakuAlpha := make([]string, 0, 52)
	// ａ to ｚ
	for i := 0xff41; i < 0xff5b; i++ {
		allZenkakuAlpha = append(allZenkakuAlpha, string(i))
	}
	// Ａ to Ｚ
	for i := 0xff21; i < 0xff3b; i++ {
		allZenkakuAlpha = append(allZenkakuAlpha, string(i))
	}
	return allZenkakuAlpha
}

func TestAllAlpha(t *testing.T) {
	allHankakuAlpha := getAllHankakuAlpha()
	allZenkakuAlpha := getAllZenkakuAlpha()
	lh := len(allHankakuAlpha)
	lz := len(allZenkakuAlpha)
	if lh != lz {
		t.Fatalf("allHankakuAlpha length %s not equal allZenkakuAlpha length %s", lh, lz)
	}
	m := make(map[string]string)
	for i, s := range allHankakuAlpha {
		m[s] = allZenkakuAlpha[i]
	}
	for h, z := range m {
		if H2z(h, ALL) != z {
			t.Errorf("Expected %s, but %s found", z, h)
		}
		if Z2h(z, ALL) != h {
			t.Errorf("Expected %s, but %s found", h, z)
		}
	}
}
