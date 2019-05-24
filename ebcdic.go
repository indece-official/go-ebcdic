/**
 * Copyright 2019 indece UG (haftungsbeschränkt)
 **/

package ebcdic

import (
	"errors"
)

const mapLength = 0xFF

/**
 * Encode unicode string to EBCDIC byte array
 *
 * Supported code pages:
 *    * EBCDIC037
 *	  * EBCDIC273
 *    * EBCDIC1140
 *    * EBCDIC1141
 **/
func Encode(unicodeStr string, codePage int) ([]byte, error) {
	ebcdicMap, ok := mapsUnicodeToEBCDIC[codePage]
	if !ok {
		return nil, errors.New("Unknown code page")
	}

	runes := []rune(unicodeStr)

	var ebcdic []byte

	for _, r := range runes {
		switch {
		case ebcdicMap.HasEuroPatch && r == '€':
			/* Apply EURO-Patch (character outside of mapLength) */
			ebcdic = append(ebcdic, ebcdicMap.EuroChar)
			break
		case r <= mapLength:
			ebcdic = append(ebcdic, ebcdicMap.Map[r])
			break
		default:
			/* If outside of map, replace with 0x00 */
			ebcdic = append(ebcdic, 0)
			break
		}
	}

	return ebcdic, nil
}

/**
 * Decode EBCDIC byte array to unicode string
 *
 * Supported code pages:
 *    * EBCDIC037
 *	  * EBCDIC273
 *	  * EBCDIC1140
 *	  * EBCDIC1141
 **/
func Decode(ebcdic []byte, codePage int) (string, error) {
	ebcdicMap, ok := mapsEBCDICToUnicode[codePage]
	if !ok {
		return "", errors.New("Unknown code page")
	}

	buf := make([]rune, len(ebcdic))
	for i, b := range ebcdic {
		if ebcdicMap.HasEuroPatch && b == ebcdicMap.EuroChar {
			buf[i] = '€'
		} else {
			buf[i] = ebcdicMap.Map[b]
		}
	}

	return string(buf), nil
}
