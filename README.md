# go-ebcdic
Go-Package for transcoding EBCDIC <-> Unicode / UTF8 / ISO 8859-1

Features:
* Supports multiple code pages (including Euro-Patches)
* Characters that are not included in EBCDIC code page are replaced with `0x00`

## Usage
Install:
```
go get github.com/indece/go-ebcdic
```

Use it in your code:
```
import (
    "github.com/indece-official/go-ebcdic"
)

...

func myEncode() {
    in := "This string special chars like § and Ö"

    encoded, err := ebcdic.Encode(in, ebcdic.EBCDIC037)
    if err != nil {
        return
    }
}

func myDecode() {
    in := []byte{0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2}
		
    decoded, err := ebcdic.Decode(in, ebcdic.EBCDIC037)
    if err != nil {
        return
    }
}
```

## Supported code pages
Don't known which code page you have to use? The most common one used is EBCDIC037 ;-)

| Code Page | Countries | Source | Comment |
| --- | --- | --- | --- |
| **EBCDIC037** | AUS, CAN, NZL, PRT, ZAF, USA | See https://en.wikipedia.org/wiki/EBCDIC_037 |
| EBCDIC273 | DEU, AUT | See https://en.wikipedia.org/wiki/EBCDIC_273 |
| EBCDIC1140 | AUS, CAN, NZL, PRT, ZAF, USA | See https://en.wikipedia.org/wiki/EBCDIC_1140 | EBCDIC037 with Euro-Patch |
| EBCDIC1141 | DEU, AUT | See https://en.wikipedia.org/wiki/EBCDIC_1141 | EBCDIC273 with Euro-Patch |
