package ebcdic

import (
	"errors"
)

// ToUnicode converts ebcdic data to unicode data
func ToUnicode(ebcdicData []byte, codePage CodePage) ([]rune, error) {
	conv, ok := codeTables[codePage]
	if !ok {
		return nil, errors.New("code page not supported")
	}

	unicodeData := make([]rune, len(ebcdicData))

	for i, ebcdicDatum := range ebcdicData {
		switch {
		case conv.EuroChar != 0 && ebcdicDatum == conv.EuroChar:
			unicodeData[i] = 0x20AC
		default:
			unicodeData[i] = conv.ToUnicode[ebcdicDatum]
		}
	}

	return unicodeData, nil
}

// FromUnicode converts unicode data to ebcdic data
func FromUnicode(unicodeData []rune, codePage CodePage) ([]byte, error) {
	conv, ok := codeTables[codePage]
	if !ok {
		return nil, errors.New("code page not supported")
	}

	ebcdicData := make([]byte, len(unicodeData))

	for i, unicodeDatum := range unicodeData {
		switch {
		case conv.EuroChar != 0 && unicodeDatum == 0x20AC:
			ebcdicData[i] = conv.EuroChar
		case unicodeDatum <= 0xFF:
			ebcdicData[i] = conv.FromUnicode[unicodeDatum]
		default:
			ebcdicData[i] = 0x00
		}
	}

	return ebcdicData, nil
}
