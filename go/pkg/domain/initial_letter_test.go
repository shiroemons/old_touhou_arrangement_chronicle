package domain

import (
	"testing"
)

func TestInitialLetter(t *testing.T) {
	t.Parallel()
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  InitialLetterType
		want1 string
	}{
		{
			name:  "記号",
			args:  args{str: "<echo>PROJECT"},
			want:  InitialLetterTypeSymbol,
			want1: "",
		},
		{
			name:  "数字",
			args:  args{str: "556ミリメートル"},
			want:  InitialLetterTypeNumber,
			want1: "",
		},
		{
			name:  "英字-C",
			args:  args{str: "COOL&CREATE"},
			want:  InitialLetterTypeAlphabet,
			want1: "C",
		},
		{
			name:  "英字-S",
			args:  args{str: "SOUND HOLIC"},
			want:  InitialLetterTypeAlphabet,
			want1: "S",
		},
		{
			name:  "ひらがな-ち",
			args:  args{str: "ちょこふぁん"},
			want:  InitialLetterTypeHiragana,
			want1: "ち",
		},
		{
			name:  "カタカナ-ふ",
			args:  args{str: "フーリンキャットマーク"},
			want:  InitialLetterTypeKatakana,
			want1: "ふ",
		},
		{
			name:  "漢字",
			args:  args{str: "森羅万象"},
			want:  InitialLetterTypeKanji,
			want1: "",
		},
		{
			name:  "その他",
			args:  args{str: "α music"},
			want:  InitialLetterTypeOther,
			want1: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := InitialLetter(tt.args.str)
			if got != tt.want {
				t.Errorf("InitialLetter() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InitialLetter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
