package dotmatrix

import "tinygo.org/x/tinyfont"

var Bold = tinyfont.Font{
	Glyphs: []tinyfont.Glyph{

		/*   */ {Rune: 32, Width: 0, Height: 8, XAdvance: 2, XOffset: 0, YOffset: -7, Bitmaps: []uint8{}},

		/* ! */ {Rune: 33, Width: 2, Height: 8, XAdvance: 3, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xff, 0xCC}},

		/* " */ {Rune: 34, Width: 5, Height: 8, XAdvance: 6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xDE, 0xC0}},

		/* # */ {Rune: 35, Width: 7, Height: 8, XAdvance: 8, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x6D, 0xFD, 0xB7, 0xF6, 0xC0}},

		/* $ */ {Rune: 36, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x33, 0xF0, 0xE1, 0xF9, 0x80}},

		/* % */ {Rune: 37, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xCF, 0x31, 0x8C, 0x63, 0x3C, 0xC0}},

		/* & */ {Rune: 38, Width: 7, Height: 8, XAdvance: 8, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x73, 0x6D, 0x9C, 0xDF, 0x66, 0xC0}},

		/* ' */ {Rune: 39, Width: 2, Height: 8, XAdvance: 3, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xF0}},

		/* ( */ {Rune: 40, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x36, 0xCC, 0xC6, 0x30}},

		/* ) */ {Rune: 41, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x33, 0x36, 0xC0}},

		/* * */ {Rune: 42, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x30, 0xCF, 0xDE, 0xCC}},

		/* + */ {Rune: 43, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x30, 0xCF, 0xCC, 0x30}},

		/* , */ {Rune: 44, Width: 2, Height: 8, XAdvance: 3, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xF6}},

		/* - */ {Rune: 45, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xF0}},

		/* . */ {Rune: 46, Width: 2, Height: 8, XAdvance: 3, XOffset: 0, YOffset: -2, Bitmaps: []uint8{0xF0}},

		/* / */ {Rune: 47, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x0C, 0x31, 0x8C, 0x63, 0x0C}},

		/* 0 */ {Rune: 48, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x26, 0xF7, 0xBD, 0xEC, 0x80}},

		/* 1 */ {Rune: 49, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x6E, 0x66, 0x66, 0xF0}},

		/* 2 */ {Rune: 50, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xC6, 0x66, 0x63, 0xE0}},

		/* 3 */ {Rune: 51, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xC6, 0x61, 0xED, 0xC0}},

		/* 4 */ {Rune: 52, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3B, 0xF7, 0xF1, 0x8C, 0x60}},

		/* 5 */ {Rune: 53, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFE, 0x3C, 0x31, 0xED, 0xC0}},

		/* 6 */ {Rune: 54, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0x3D, 0xBD, 0xED, 0xC0}},

		/* 7 */ {Rune: 55, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF8, 0xC6, 0x66, 0x63}},

		/* 8 */ {Rune: 56, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF6, 0xED, 0xED, 0xC0}},

		/* 9 */ {Rune: 57, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF7, 0xB7, 0x8D, 0xC0}},

		/* : */ {Rune: 58, Width: 2, Height: 8, XAdvance: 3, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xCC}},

		/* ; */ {Rune: 59, Width: 3, Height: 8, XAdvance: 4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x61, 0xE0}},

		/* < */ {Rune: 60, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x36, 0xC6, 0x30}},

		/* = */ {Rune: 61, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xF0, 0xF0}},

		/* > */ {Rune: 62, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xC6, 0x36, 0xC0}},

		/* ? */ {Rune: 63, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xC6, 0xE6, 0x01, 0x80}},

		/* @ */ {Rune: 64, Width: 8, Height: 8, XAdvance: 9, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x7F, 0xC1, 0xDD, 0xD5, 0xDF, 0xC0, 0x7F}},

		/* A */ {Rune: 65, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF7, 0xFD, 0xEF, 0x60}},

		/* B */ {Rune: 66, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF6, 0xF7, 0xED, 0xEF, 0xC0}},

		/* C */ {Rune: 67, Width: 5, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF1, 0x8C, 0x6D, 0xC0}},

		/* D */ {Rune: 68, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF6, 0xF7, 0xBD, 0xEF, 0xC0}},

		/* E */ {Rune: 69, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFE, 0x31, 0xEC, 0x63, 0xE0}},

		/* F */ {Rune: 70, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFE, 0x31, 0xEC, 0x63}},

		/* G */ {Rune: 71, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF1, 0xBD, 0xED, 0xE0}},

		/* H */ {Rune: 72, Width: 5, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xDE, 0xF7, 0xFD, 0xEF, 0x60}},

		/* I */ {Rune: 73, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF6, 0x66, 0x66, 0xF0}},

		/* J */ {Rune: 74, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3C, 0x61, 0x86, 0x18, 0x6D, 0x9C}},

		/* K */ {Rune: 75, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xDE, 0xFD, 0xCF, 0x6F, 0x60}},

		/* L */ {Rune: 76, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x31, 0x8C, 0x63, 0xE0}},

		/* M */ {Rune: 77, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0xFF, 0xF3, 0xCF, 0x3C, 0xC0}},

		/* N */ {Rune: 78, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0x3E, 0xFB, 0xDF, 0x7C, 0xC0}},

		/* O */ {Rune: 79, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF7, 0xBD, 0xED, 0xC0}},

		/* P */ {Rune: 80, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF6, 0xF7, 0xEC, 0x63}},

		/* Q */ {Rune: 81, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x76, 0xF7, 0xBD, 0xED, 0xC3}},

		/* R */ {Rune: 82, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF6, 0xF7, 0xED, 0xEF, 0x60}},

		/* S */ {Rune: 83, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x7E, 0x30, 0xE1, 0x8F, 0xC0}},

		/* T */ {Rune: 84, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFC, 0xC3, 0x0C, 0x30, 0xC3}},

		/* U */ {Rune: 85, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xDE, 0xF7, 0xBD, 0xED, 0xC0}},

		/* V */ {Rune: 86, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0x3C, 0xF3, 0x79, 0xE3}},

		/* W */ {Rune: 87, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0x3C, 0xF3, 0xFF, 0xFC, 0xC0}},

		/* X */ {Rune: 88, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0x37, 0x8C, 0x7B, 0x3C, 0xC0}},

		/* Y */ {Rune: 89, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xCF, 0x37, 0x8C, 0x63, 0x0C}},

		/* Z */ {Rune: 90, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xF8, 0xC6, 0x66, 0x63, 0xE0}},

		/* [ */ {Rune: 91, Width: 3, Height: 8, XAdvance: 4, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFB, 0x6D, 0xB8}},

		/* \ */ {Rune: 92, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC3, 0x06, 0x0C, 0x18, 0x30, 0xC0}},

		/* ] */ {Rune: 93, Width: 3, Height: 8, XAdvance: 4, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xED, 0xB6, 0xF8}},

		/* ^ */ {Rune: 94, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x6F}},

		/* _ */ {Rune: 95, Width: 4, Height: 8, XAdvance: 5, XOffset: 1, YOffset: -1, Bitmaps: []uint8{0xF0}},

		/* a */ {Rune: 97, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x70, 0xDF, 0xB7, 0x80}},

		/* b */ {Rune: 98, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x3D, 0xBD, 0xEF, 0xC0}},

		/* c */ {Rune: 99, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x76, 0xF1, 0xB7}},

		/* d */ {Rune: 100, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x18, 0xDF, 0xBD, 0xED, 0xE0}},

		/* e */ {Rune: 101, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x76, 0xFF, 0x87}},

		/* f */ {Rune: 102, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x36, 0x6F, 0x66, 0x60}},

		/* g */ {Rune: 103, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x7E, 0xF7, 0xB7, 0x8F, 0xC0}},

		/* h */ {Rune: 104, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x3D, 0xBD, 0xEF, 0x60}},

		/* i */ {Rune: 105, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x60, 0xE6, 0x66, 0xF0}},

		/* j */ {Rune: 106, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x18, 0x0E, 0x31, 0x8F, 0x6E}},

		/* k */ {Rune: 107, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x37, 0xEE, 0x7B, 0x60}},

		/* l */ {Rune: 108, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xE6, 0x66, 0x66, 0xF0}},

		/* m */ {Rune: 109, Width: 7, Height: 8, XAdvance: 8, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xED, 0xAF, 0x5E, 0xBD, 0x60}},

		/* n */ {Rune: 110, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xF6, 0xF7, 0xBD, 0x80}},

		/* o */ {Rune: 111, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x76, 0xF7, 0xB7}},

		/* p */ {Rune: 112, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xF6, 0xF7, 0xBF, 0x63}},

		/* q */ {Rune: 113, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x7E, 0xF7, 0xB7, 0x8C, 0x60}},

		/* r */ {Rune: 114, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xFF, 0x31, 0x8C}},

		/* s */ {Rune: 115, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x7E, 0x1C, 0x3F}},

		/* t */ {Rune: 116, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x66, 0x6F, 0x66, 0x30}},

		/* u */ {Rune: 117, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xDE, 0xF7, 0xB7, 0x80}},

		/* v */ {Rune: 118, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xCF, 0x3C, 0xDE, 0x30}},

		/* w */ {Rune: 119, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xCF, 0x3F, 0xFF, 0xCC}},

		/* x */ {Rune: 120, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xCD, 0xE3, 0x1E, 0xCC}},

		/* y */ {Rune: 121, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xDE, 0xF7, 0xB7, 0x8F, 0xC0}},

		/* z */ {Rune: 122, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xF3, 0x6C, 0xF0}},

		/* { */ {Rune: 123, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x36, 0x6C, 0x66, 0x30}},

		/* | */ {Rune: 124, Width: 2, Height: 8, XAdvance: 3, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xFF, 0xFC}},

		/* } */ {Rune: 125, Width: 4, Height: 8, XAdvance: 5, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0xC6, 0x63, 0x66, 0xC0}},

		/* ~ */ {Rune: 126, Width: 5, Height: 8, XAdvance: 6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x6E, 0x80}},

		/* £ */ {Rune: 163, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x39, 0xB6, 0x3C, 0x61, 0x87, 0xC0}},

		/* © */ {Rune: 169, Width: 8, Height: 8, XAdvance: 9, XOffset: 0, YOffset: -8, Bitmaps: []uint8{0x3C, 0x66, 0xDB, 0xF3, 0xF3, 0xDB, 0x66, 0x3C}},

		/* « */ {Rune: 171, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x6F, 0x66, 0xC0}},

		/* » */ {Rune: 187, Width: 6, Height: 8, XAdvance: 7, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xD9, 0xBD, 0x80}},

		///* Å */ {Rune: 197, Width: 5, Height: 8, XAdvance: 6, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x76,0xff,0xbd } },

		///* Æ */ {Rune: 198, Width: 6, Height: 8, XAdvance: 7, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0xf3,0xcf,0xbc } },

		///* Ø */ {Rune: 216, Width: 4, Height: 8, XAdvance: 5, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0xff,0xff } },

		///* å */ {Rune: 229, Width: 4, Height: 8, XAdvance: 5, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x06,0x7f } },

		///* æ */ {Rune: 230, Width: 6, Height: 8, XAdvance: 7, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x01,0xf7,0xff } },

		///* ò */ {Rune: 242, Width: 4, Height: 8, XAdvance: 5, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x60,0xff } },

		///* ó */ {Rune: 243, Width: 4, Height: 8, XAdvance: 5, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x60,0xff } },

		///* ø */ {Rune: 248, Width: 5, Height: 8, XAdvance: 6, XOffset: 1, YOffset: 0, Bitmaps: []uint8{ 0x03,0xbf } },

		/* ‐ */ {Rune: 8208, Width: 3, Height: 8, XAdvance: 4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xE0}},
	},
	YAdvance: 8,
}
