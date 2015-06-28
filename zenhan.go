// Package zenhan provides convertion between Zenkaku(fullwidth Japanese) and Hankaku(halfwidth Japanese)
package zenhan

import (
	"strings"
	"unicode/utf8"
)

type flag uint

const (
	// ASCII converts only ascii
	ASCII flag = 1 << iota
	// DIGIT converts only digit
	DIGIT
	// KANA converts only kana
	KANA
	// ALL converts ascii, digit, and kana
	ALL = ASCII | DIGIT | KANA
)

const (
	dakuten    = "ﾞ"
	handakuten = "ﾟ"
)

var (
	zhASCII map[string]string
	hzASCII map[string]string
	zhDIGIT map[string]string
	hzDIGIT map[string]string
	zhKANA  map[string]string
	hzKANA  map[string]string
)

func zip(a []string, b []string) map[string]string {
	m := make(map[string]string)
	for i, s := range a {
		m[s] = b[i]
	}
	return m
}

func join(m1 map[string]string, m2 map[string]string) map[string]string {
	m := make(map[string]string)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}

func init() {
	zASCII := []string{"ａ", "ｂ", "ｃ", "ｄ", "ｅ", "ｆ", "ｇ", "ｈ", "ｉ",
		"ｊ", "ｋ", "ｌ", "ｍ", "ｎ", "ｏ", "ｐ", "ｑ", "ｒ",
		"ｓ", "ｔ", "ｕ", "ｖ", "ｗ", "ｘ", "ｙ", "ｚ",
		"Ａ", "Ｂ", "Ｃ", "Ｄ", "Ｅ", "Ｆ", "Ｇ", "Ｈ", "Ｉ",
		"Ｊ", "Ｋ", "Ｌ", "Ｍ", "Ｎ", "Ｏ", "Ｐ", "Ｑ", "Ｒ",
		"Ｓ", "Ｔ", "Ｕ", "Ｖ", "Ｗ", "Ｘ", "Ｙ", "Ｚ",
		"！", "”", "＃", "＄", "％", "＆", "’", "（", "）",
		"＊", "＋", "，", "−", "．", "／", "：", "；", "＜",
		"＝", "＞", "？", "＠", "［", "￥", "］", "＾", "＿",
		"‘", "｛", "｜", "｝", "〜", "　"}
	zDigit := []string{"０", "１", "２", "３", "４",
		"５", "６", "７", "８", "９"}
	zKana := []string{"ア", "イ", "ウ", "エ", "オ",
		"カ", "キ", "ク", "ケ", "コ",
		"サ", "シ", "ス", "セ", "ソ",
		"タ", "チ", "ツ", "テ", "ト",
		"ナ", "ニ", "ヌ", "ネ", "ノ",
		"ハ", "ヒ", "フ", "ヘ", "ホ",
		"マ", "ミ", "ム", "メ", "モ",
		"ヤ", "ユ", "ヨ",
		"ラ", "リ", "ル", "レ", "ロ",
		"ワ", "ヲ", "ン",
		"ァ", "ィ", "ゥ", "ェ", "ォ",
		"ッ", "ャ", "ュ", "ョ", "ヴ",
		"ガ", "ギ", "グ", "ゲ", "ゴ",
		"ザ", "ジ", "ズ", "ゼ", "ゾ",
		"ダ", "ヂ", "ヅ", "デ", "ド",
		"バ", "ビ", "ブ", "ベ", "ボ",
		"パ", "ピ", "プ", "ペ", "ポ",
		"。", "、", "・", "゛", "゜", "「", "」", "ー"}
	hASCII := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
		"!", `"`, "#", "$", "%", "&", "'", "(", ")",
		"*", "+", ",", "-", ".", "/", ":", ";", "<",
		"=", ">", "?", "@", "[", "\\", "]", "^", "_",
		"`", "{", "|", "}", "~", " "}
	hDigit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	hKana := []string{"ｱ", "ｲ", "ｳ", "ｴ", "ｵ",
		"ｶ", "ｷ", "ｸ", "ｹ", "ｺ",
		"ｻ", "ｼ", "ｽ", "ｾ", "ｿ",
		"ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ",
		"ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ",
		"ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ",
		"ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ",
		"ﾔ", "ﾕ", "ﾖ",
		"ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ",
		"ﾜ", "ｦ", "ﾝ",
		"ｧ", "ｨ", "ｩ", "ｪ", "ｫ",
		"ｯ", "ｬ", "ｭ", "ｮ", "ｳﾞ",
		"ｶﾞ", "ｷﾞ", "ｸﾞ", "ｹﾞ", "ｺﾞ",
		"ｻﾞ", "ｼﾞ", "ｽﾞ", "ｾﾞ", "ｿﾞ",
		"ﾀﾞ", "ﾁﾞ", "ﾂﾞ", "ﾃﾞ", "ﾄﾞ",
		"ﾊﾞ", "ﾋﾞ", "ﾌﾞ", "ﾍﾞ", "ﾎﾞ",
		"ﾊﾟ", "ﾋﾟ", "ﾌﾟ", "ﾍﾟ", "ﾎﾟ",
		"｡", "､", "･", "ﾞ", "ﾟ", "｢", "｣", "ｰ"}

	zhASCII = zip(zASCII, hASCII)
	hzASCII = zip(hASCII, zASCII)

	zhDIGIT = zip(zDigit, hDigit)
	hzDIGIT = zip(hDigit, zDigit)

	zhKANA = zip(zKana, hKana)
	hzKANA = zip(hKana, zKana)
}

func makeHan2zenDict(mode flag) map[string]string {
	m := make(map[string]string)
	if mode&ASCII == ASCII {
		m = join(m, hzASCII)
	}
	if mode&DIGIT == DIGIT {
		m = join(m, hzDIGIT)
	}
	if mode&KANA == KANA {
		m = join(m, hzKANA)
	}
	return m
}

func makeZen2hanDict(mode flag) map[string]string {
	m := make(map[string]string)
	if mode&ASCII == ASCII {
		m = join(m, zhASCII)
	}
	if mode&DIGIT == DIGIT {
		m = join(m, zhDIGIT)
	}
	if mode&KANA == KANA {
		m = join(m, zhKANA)
	}
	return m
}

func getValueFromMap(m map[string]string, key string, defaultValue string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

func any(array []string, value string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}

// H2z converts string from hankaku to zenkaku
func H2z(text string, mode flag, ignore ...string) string {
	if !utf8.ValidString(text) {
		return text
	}
	m := makeHan2zenDict(mode)
	t := []rune(text)
	converted := make([]string, 0, len(t))
	for i, v := range t {
		curr := string(v)
		if any(ignore, curr) {
			converted = append(converted, curr)
		} else if curr == dakuten || curr == handakuten && i > 0 {
			prev := string(t[i-1])
			if z, ok := m[prev+curr]; ok {
				converted = converted[:len(converted)-1]
				converted = append(converted, z)
			} else {
				converted = append(converted, getValueFromMap(m, curr, curr))
			}
		} else {
			converted = append(converted, getValueFromMap(m, curr, curr))
		}
	}
	return strings.Join(converted, "")
}

// Z2h converts string from zenkaku to hankaku
func Z2h(text string, mode flag, ignore ...string) string {
	if !utf8.ValidString(text) {
		return text
	}
	m := makeZen2hanDict(mode)
	t := []rune(text)
	converted := make([]string, 0, len(t))
	for _, v := range t {
		curr := string(v)
		if any(ignore, curr) {
			converted = append(converted, curr)
		} else {
			converted = append(converted, getValueFromMap(m, curr, curr))
		}
	}
	return strings.Join(converted, "")
}
