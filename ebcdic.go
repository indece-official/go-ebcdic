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
 **/
func Encode(unicodeStr string, codePage int) ([]byte, error) {
	ebcdicMap, ok := mapsUnicodeToEBDIC[codePage]
	if ! ok {
		return nil, errors.New("Unknown code page")
	}

	runes := []rune(unicodeStr)

	var ebcdic []byte

	for _, r := range runes {
		/* If outside of map, replace with 0x00 */
		if r > mapLength {
			ebcdic = append(ebcdic, 0)
		} else {
			ebcdic = append(ebcdic, ebcdicMap[r])
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
 **/
func Decode(ebcdic []byte, codePage int) (string, error) {
	ebcdicMap, ok := mapsEBDICToUnicode[codePage]
	if ! ok {
		return "", errors.New("Unknown code page")
	}

	buf := make([]rune, len(ebcdic))
    for i, b := range ebcdic {
        buf[i] = rune(ebcdicMap[b])
	}
	
    return string(buf), nil
}
