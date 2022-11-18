# go-ebcdic

Go-Package for transcoding EBCDIC ↔ Unicode / UTF8

[![GoDoc](https://godoc.org/github.com/indece-official/go-ebcdic?status.svg)](https://godoc.org/github.com/indece-official/go-ebcdic)

Features:

- Supports multiple code pages (including Euro-Patches)
- Optimized performance

## Usage

Install:

```
go get github.com/indece-official/go-ebcdic
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

Don't known which code page you have to use? The one most commonly used ist EBCDIC037 ;-)

| Code Page     | Countries                    | Source                                             | Comment                   |
| ------------- | ---------------------------- | -------------------------------------------------- | ------------------------- |
| **EBCDIC037** | AUS, CAN, NZL, PRT, ZAF, USA | See https://en.wikipedia.org/wiki/EBCDIC_037       |
| EBCDIC273     | DEU, AUT                     | See https://en.wikipedia.org/wiki/EBCDIC_273       |
| EBCDIC500     | International                | See https://en.wikipedia.org/wiki/Code_page_37#500 |
| EBCDIC1140    | AUS, CAN, NZL, PRT, ZAF, USA | See https://en.wikipedia.org/wiki/EBCDIC_1140      | EBCDIC037 with Euro-Patch |
| EBCDIC1141    | DEU, AUT                     | See https://en.wikipedia.org/wiki/EBCDIC_1141      | EBCDIC273 with Euro-Patch |
| EBCDIC1148    | International                | See https://en.wikipedia.org/wiki/Code_page_37#500 | EBCDIC500 with Euro-Patch |
