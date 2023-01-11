package domain

import (
	"unicode"

	"github.com/goark/kkconv"
)

type InitialLetterType string

const (
	InitialLetterTypeSymbol   InitialLetterType = "symbol"
	InitialLetterTypeNumber   InitialLetterType = "number"
	InitialLetterTypeHiragana InitialLetterType = "hiragana"
	InitialLetterTypeKatakana InitialLetterType = "katakana"
	InitialLetterTypeKanji    InitialLetterType = "kanji"
	InitialLetterTypeAlphabet InitialLetterType = "alphabet"
	InitialLetterTypeOther    InitialLetterType = "other"
)

var (
	AlphabetUpperCase = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0041, 0x005A, 1},
		},
		LatinOffset: 1,
	}
	AlphabetLowerCase = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0061, 0x007A, 1},
		},
		LatinOffset: 1,
	}
	Hiragana = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x3041, 0x3096, 1}, // ぁ 〜 ゖ
		},
	}
	Katakana = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x30A1, 0x30FA, 1}, // ァ 〜 ヺ
			{0x31F0, 0x31FF, 1}, // ㇰ 〜 ㇿ
		},
	}
	HalfWidthKatakana = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0xFF61, 0xFF9F, 1},
		},
	}
	Kanji = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x4E00, 0x9FFF, 1},
		},
	}
)

func runeCheck(r rune) (InitialLetterType, string) {
	if unicode.IsPunct(r) || unicode.IsSymbol(r) {
		return InitialLetterTypeSymbol, ""
	}
	if unicode.IsDigit(r) && unicode.IsNumber(r) {
		return InitialLetterTypeNumber, ""
	}
	if unicode.Is(Hiragana, r) {
		return InitialLetterTypeHiragana, kkconv.Chokuon(string(r), true)
	}
	if unicode.Is(Katakana, r) || unicode.Is(HalfWidthKatakana, r) {
		return InitialLetterTypeKatakana, kkconv.Hiragana(kkconv.Chokuon(string(r), true), true)
	}
	if unicode.Is(Kanji, r) {
		return InitialLetterTypeKanji, ""
	}
	if unicode.In(r, AlphabetUpperCase, AlphabetLowerCase) {
		return InitialLetterTypeAlphabet, string(unicode.ToUpper(r))
	}
	return InitialLetterTypeOther, ""
}

// InitialLetter 文字列の頭文字の文字種別を判定する。
// ひらがなやカタカナ、アルファベットの場合は、2つ目の戻り値に詳細情報が返されます。
func InitialLetter(str string) (InitialLetterType, string) {
	return runeCheck([]rune(str)[0])
}
