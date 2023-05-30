package ebcdic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Test encoding unicode to EBCDIC 037
 **/
func TestSplitEncodeEBCDIC037(t *testing.T) {
	in := "This string è contains $%? different special _chars_, z.B. § and Ö\u0085\n\x0A"
	want := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x15, 0x25, 0x25,
	}

	got, err := Encode(in, EBCDIC037)
	assert.NoError(t, err)
	assert.Equal(t, want, got, "Encode(\"%s\", EBCDIC037) = %v; want %v", in, got, want)
}

/**
 * Test decoding EBCDIC 037 to unicode
 **/
func TestSplitDecodeEBCDIC037(t *testing.T) {
	in := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x15, 0x25,
	}
	want := "This string è contains $%? different special _chars_, z.B. § and Ö\u0085\n"

	got, err := Decode(in, EBCDIC037)
	assert.NoError(t, err)
	assert.Equal(t, want, got, "Decode(%v, EBCDIC037) = \"%s\"; want \"%s\"", in, got, want)
}

/**
 * Test encoding unicode to EBCDIC 273
 **/
func TestSplitEncodeEBCDIC273(t *testing.T) {
	in := "Dieser String enthält $%? verschiedene _Sonderzeichen_, z.B. ß"
	want := []byte{
		0xC4, 0x89, 0x85, 0xA2, 0x85, 0x99, 0x40, 0xE2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x85, 0x95,
		0xA3, 0x88, 0xC0, 0x93, 0xA3, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0xA5, 0x85, 0x99, 0xA2, 0x83, 0x88,
		0x89, 0x85, 0x84, 0x85, 0x95, 0x85, 0x40, 0x6D, 0xE2, 0x96, 0x95, 0x84, 0x85, 0x99, 0xA9, 0x85,
		0x89, 0x83, 0x88, 0x85, 0x95, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xA1,
	}

	got, err := Encode(in, EBCDIC273)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode(\"%s\", EBCDIC273) = %v; want %v", in, got, want)
	}
}

/**
 * Test decoding EBCDIC 273 to unicode
 **/
func TestSplitDecodeEBCDIC273(t *testing.T) {
	in := []byte{
		0xC4, 0x89, 0x85, 0xA2, 0x85, 0x99, 0x40, 0xE2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x85, 0x95,
		0xA3, 0x88, 0xC0, 0x93, 0xA3, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0xA5, 0x85, 0x99, 0xA2, 0x83, 0x88,
		0x89, 0x85, 0x84, 0x85, 0x95, 0x85, 0x40, 0x6D, 0xE2, 0x96, 0x95, 0x84, 0x85, 0x99, 0xA9, 0x85,
		0x89, 0x83, 0x88, 0x85, 0x95, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xA1,
	}
	want := "Dieser String enthält $%? verschiedene _Sonderzeichen_, z.B. ß"

	got, err := Decode(in, EBCDIC273)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Decode(%v, EBCDIC273) = \"%s\"; want \"%s\"", in, got, want)
	}
}

/**
 * Test encoding unicode to EBCDIC 500
 **/
func TestSplitEncodeEBCDIC500(t *testing.T) {
	in := "This string è contains $%? different special _chars_, z.B. § and Ö"
	want := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC,
	}

	got, err := Encode(in, EBCDIC500)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode(\"%s\", EBCDIC500) = %v; want %v", in, got, want)
	}
}

/**
 * Test decoding EBCDIC 500 to unicode
 **/
func TestSplitDecodeEBCDIC500(t *testing.T) {
	in := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xC1, 0xC2, 0xC3,
		0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xD1, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xE2, 0xE3,
		0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91,
		0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9,
	}
	want := "This string è contains $%? different special _chars_, z.B. § and Ö 0123456789ABCDEFGHIJLMNOPQRSTUVWXYZabcdefghijlmnopqrstuvwxyz"

	got, err := Decode(in, EBCDIC500)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Decode(%v, EBCDIC500) = \"%s\"; want \"%s\"", in, got, want)
	}
}

/**
 * Test encoding unicode to EBCDIC 1140
 **/
func TestSplitEncodeEBCDIC1140(t *testing.T) {
	in := "This string è contains $%? different special _chars_, z.B. § and Ö 0123456789ABCDEFGHIJLMNOPQRSTUVWXYZabcdefghijlmnopqrstuvwxyz"
	want := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xC1, 0xC2, 0xC3,
		0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xD1, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xE2, 0xE3,
		0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91,
		0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9,
	}

	got, err := Encode(in, EBCDIC1140)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode(\"%s\", EBCDIC1140) = %v; want %v", in, got, want)
	}
}

/**
 * Test decoding EBCDIC 1140 to unicode
 **/
func TestSplitDecodeEBCDIC1140(t *testing.T) {
	in := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0x9F,
	}
	want := "This string è contains $%? different special _chars_, z.B. § and €"

	got, err := Decode(in, EBCDIC1140)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Decode(%v, EBCDIC1140) = \"%s\"; want \"%s\"", in, got, want)
	}
}

/**
 * Test encoding unicode to EBCDIC 1141
 **/
func TestSplitEncodeEBCDIC1141(t *testing.T) {
	in := "Dieser String enthält $%? verschiedene _Sonderzeichen_, z.B. €"
	want := []byte{
		0xC4, 0x89, 0x85, 0xA2, 0x85, 0x99, 0x40, 0xE2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x85, 0x95,
		0xA3, 0x88, 0xC0, 0x93, 0xA3, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0xA5, 0x85, 0x99, 0xA2, 0x83, 0x88,
		0x89, 0x85, 0x84, 0x85, 0x95, 0x85, 0x40, 0x6D, 0xE2, 0x96, 0x95, 0x84, 0x85, 0x99, 0xA9, 0x85,
		0x89, 0x83, 0x88, 0x85, 0x95, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0x9F,
	}

	got, err := Encode(in, EBCDIC1141)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode(\"%s\", EBCDIC1141) = %v; want %v", in, got, want)
	}
}

/**
 * Test decoding EBCDIC 1141 to unicode
 **/
func TestSplitDecodeEBCDIC1141(t *testing.T) {
	in := []byte{
		0xC4, 0x89, 0x85, 0xA2, 0x85, 0x99, 0x40, 0xE2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x85, 0x95,
		0xA3, 0x88, 0xC0, 0x93, 0xA3, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0xA5, 0x85, 0x99, 0xA2, 0x83, 0x88,
		0x89, 0x85, 0x84, 0x85, 0x95, 0x85, 0x40, 0x6D, 0xE2, 0x96, 0x95, 0x84, 0x85, 0x99, 0xA9, 0x85,
		0x89, 0x83, 0x88, 0x85, 0x95, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0x9F,
	}
	want := "Dieser String enthält $%? verschiedene _Sonderzeichen_, z.B. €"

	got, err := Decode(in, EBCDIC1141)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Decode(%v, EBCDIC1141) = \"%s\"; want \"%s\"", in, got, want)
	}
}

/**
 * Test encoding unicode to EBCDIC 1148
 **/
func TestSplitEncodeEBCDIC1148(t *testing.T) {
	in := "This string è contains $%? different special _chars_, z.B. § and Ö 0123456789ABCDEFGHIJLMNOPQRSTUVWXYZabcdefghijlmnopqrstuvwxyz"
	want := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xC1, 0xC2, 0xC3,
		0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xD1, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xE2, 0xE3,
		0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91,
		0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9,
	}

	got, err := Encode(in, EBCDIC1148)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode(\"%s\", EBCDIC1148) = %v; want %v", in, got, want)
	}
}

/**
 * Test decoding EBCDIC 1148 to unicode
 **/
func TestSplitDecodeEBCDIC1148(t *testing.T) {
	in := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC, 0x40, 0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xC1, 0xC2, 0xC3,
		0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xD1, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xE2, 0xE3,
		0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91,
		0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9,
	}
	want := "This string è contains $%? different special _chars_, z.B. § and Ö 0123456789ABCDEFGHIJLMNOPQRSTUVWXYZabcdefghijlmnopqrstuvwxyz"

	got, err := Decode(in, EBCDIC1148)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("Decode(%v, EBCDIC1148) = \"%s\"; want \"%s\"", in, got, want)
	}
}

func BenchmarkEncode(b *testing.B) {
	in := "This string è contains $%? different special _chars_, z.B. § and Ö"

	for n := 0; n < b.N; n++ {
		Encode(in, EBCDIC037)
	}
}

func BenchmarkDecode(b *testing.B) {
	in := []byte{
		0xE3, 0x88, 0x89, 0xA2, 0x40, 0xA2, 0xA3, 0x99, 0x89, 0x95, 0x87, 0x40, 0x54, 0x40, 0x83, 0x96,
		0x95, 0xA3, 0x81, 0x89, 0x95, 0xA2, 0x40, 0x5B, 0x6C, 0x6F, 0x40, 0x84, 0x89, 0x86, 0x86, 0x85,
		0x99, 0x85, 0x95, 0xA3, 0x40, 0xA2, 0x97, 0x85, 0x83, 0x89, 0x81, 0x93, 0x40, 0x6D, 0x83, 0x88,
		0x81, 0x99, 0xA2, 0x6D, 0x6B, 0x40, 0xA9, 0x4B, 0xC2, 0x4B, 0x40, 0xB5, 0x40, 0x81, 0x95, 0x84,
		0x40, 0xEC,
	}

	for n := 0; n < b.N; n++ {
		Decode(in, EBCDIC037)
	}
}
