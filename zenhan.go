package zenhan

type flag uint

const (
	ALL flag = 1 << iota
	ASCII
	DIGIT
	KANA
)

func H2z(text string, mode flag) string {
	return "ﾟａｂｃＤＥﾞＦ123４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
}

func Z2h(text string, mode flag) string {
	return "ﾟａｂｃＤＥﾞＦ123４５６ｱｶﾞｻダナバビﾌﾟﾍﾟﾟ"
}