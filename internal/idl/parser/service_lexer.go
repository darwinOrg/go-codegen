// Code generated from Service.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ServiceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ServiceLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func servicelexerLexerInit() {
	staticData := &ServiceLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'with_client'", "'='", "'namespace'", "'go_import'", "'ALL_HINT'",
		"'SVC_HINT'", "'STRUCT_HINT'", "'true'", "'false'", "'struct'", "'{'",
		"'}'", "'extends'", "'POINT'", "'service'", "'required'", "'optional'",
		"'POST'", "'('", "')'", "'GET'", "'URL'", "'PAGEABLE'", "'list'", "'<'",
		"'>'", "'void'", "'not_login'", "'map'", "'['", "']'", "':'", "';'",
		"", "", "", "'bool'", "'byte'", "'i16'", "'i32'", "'i64'", "'double'",
		"'string'", "", "", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL", "TYPE_BYTE", "TYPE_I16",
		"TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING", "LITERAL", "IDENTIFIER",
		"COMMA", "WS", "SL_COMMENT", "ML_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL", "TYPE_BYTE", "TYPE_I16",
		"TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING", "LITERAL", "ESC_SEQ",
		"IDENTIFIER", "COMMA", "LETTER", "DIGIT", "HEX_DIGIT", "WS", "SL_COMMENT",
		"ML_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 49, 465, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 33, 3, 33, 297, 8, 33, 1, 33, 4, 33, 300, 8, 33, 11, 33, 12,
		33, 301, 1, 34, 3, 34, 305, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 4, 34, 311,
		8, 34, 11, 34, 12, 34, 312, 1, 35, 3, 35, 316, 8, 35, 1, 35, 4, 35, 319,
		8, 35, 11, 35, 12, 35, 320, 1, 35, 1, 35, 4, 35, 325, 8, 35, 11, 35, 12,
		35, 326, 3, 35, 329, 8, 35, 1, 35, 1, 35, 4, 35, 333, 8, 35, 11, 35, 12,
		35, 334, 3, 35, 337, 8, 35, 1, 35, 1, 35, 3, 35, 341, 8, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 5, 43, 382, 8, 43, 10, 43, 12,
		43, 385, 9, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 391, 8, 43, 10, 43,
		12, 43, 394, 9, 43, 1, 43, 3, 43, 397, 8, 43, 1, 44, 1, 44, 1, 44, 1, 45,
		1, 45, 3, 45, 404, 8, 45, 1, 45, 1, 45, 1, 45, 5, 45, 409, 8, 45, 10, 45,
		12, 45, 412, 9, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1,
		49, 3, 49, 422, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 4, 50, 428, 8, 50, 11,
		50, 12, 50, 429, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 3, 51, 437, 8, 51,
		1, 51, 5, 51, 440, 8, 51, 10, 51, 12, 51, 443, 9, 51, 1, 51, 3, 51, 446,
		8, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 456,
		8, 52, 10, 52, 12, 52, 459, 9, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		457, 0, 53, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 0, 91,
		45, 93, 46, 95, 0, 97, 0, 99, 0, 101, 47, 103, 48, 105, 49, 1, 0, 10, 2,
		0, 43, 43, 45, 45, 2, 0, 69, 69, 101, 101, 2, 0, 34, 34, 92, 92, 2, 0,
		39, 39, 92, 92, 6, 0, 34, 34, 39, 39, 92, 92, 110, 110, 114, 114, 116,
		116, 2, 0, 46, 46, 95, 95, 2, 0, 65, 90, 97, 122, 2, 0, 65, 70, 97, 102,
		2, 0, 9, 9, 32, 32, 1, 0, 10, 10, 488, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0,
		0, 0, 105, 1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 119, 1, 0, 0, 0, 5, 121,
		1, 0, 0, 0, 7, 131, 1, 0, 0, 0, 9, 141, 1, 0, 0, 0, 11, 150, 1, 0, 0, 0,
		13, 159, 1, 0, 0, 0, 15, 171, 1, 0, 0, 0, 17, 176, 1, 0, 0, 0, 19, 182,
		1, 0, 0, 0, 21, 189, 1, 0, 0, 0, 23, 191, 1, 0, 0, 0, 25, 193, 1, 0, 0,
		0, 27, 201, 1, 0, 0, 0, 29, 207, 1, 0, 0, 0, 31, 215, 1, 0, 0, 0, 33, 224,
		1, 0, 0, 0, 35, 233, 1, 0, 0, 0, 37, 238, 1, 0, 0, 0, 39, 240, 1, 0, 0,
		0, 41, 242, 1, 0, 0, 0, 43, 246, 1, 0, 0, 0, 45, 250, 1, 0, 0, 0, 47, 259,
		1, 0, 0, 0, 49, 264, 1, 0, 0, 0, 51, 266, 1, 0, 0, 0, 53, 268, 1, 0, 0,
		0, 55, 273, 1, 0, 0, 0, 57, 283, 1, 0, 0, 0, 59, 287, 1, 0, 0, 0, 61, 289,
		1, 0, 0, 0, 63, 291, 1, 0, 0, 0, 65, 293, 1, 0, 0, 0, 67, 296, 1, 0, 0,
		0, 69, 304, 1, 0, 0, 0, 71, 315, 1, 0, 0, 0, 73, 342, 1, 0, 0, 0, 75, 347,
		1, 0, 0, 0, 77, 352, 1, 0, 0, 0, 79, 356, 1, 0, 0, 0, 81, 360, 1, 0, 0,
		0, 83, 364, 1, 0, 0, 0, 85, 371, 1, 0, 0, 0, 87, 396, 1, 0, 0, 0, 89, 398,
		1, 0, 0, 0, 91, 403, 1, 0, 0, 0, 93, 413, 1, 0, 0, 0, 95, 415, 1, 0, 0,
		0, 97, 417, 1, 0, 0, 0, 99, 421, 1, 0, 0, 0, 101, 427, 1, 0, 0, 0, 103,
		436, 1, 0, 0, 0, 105, 451, 1, 0, 0, 0, 107, 108, 5, 119, 0, 0, 108, 109,
		5, 105, 0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 104, 0, 0, 111, 112,
		5, 95, 0, 0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 108, 0, 0, 114, 115, 5,
		105, 0, 0, 115, 116, 5, 101, 0, 0, 116, 117, 5, 110, 0, 0, 117, 118, 5,
		116, 0, 0, 118, 2, 1, 0, 0, 0, 119, 120, 5, 61, 0, 0, 120, 4, 1, 0, 0,
		0, 121, 122, 5, 110, 0, 0, 122, 123, 5, 97, 0, 0, 123, 124, 5, 109, 0,
		0, 124, 125, 5, 101, 0, 0, 125, 126, 5, 115, 0, 0, 126, 127, 5, 112, 0,
		0, 127, 128, 5, 97, 0, 0, 128, 129, 5, 99, 0, 0, 129, 130, 5, 101, 0, 0,
		130, 6, 1, 0, 0, 0, 131, 132, 5, 103, 0, 0, 132, 133, 5, 111, 0, 0, 133,
		134, 5, 95, 0, 0, 134, 135, 5, 105, 0, 0, 135, 136, 5, 109, 0, 0, 136,
		137, 5, 112, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139, 5, 114, 0, 0, 139,
		140, 5, 116, 0, 0, 140, 8, 1, 0, 0, 0, 141, 142, 5, 65, 0, 0, 142, 143,
		5, 76, 0, 0, 143, 144, 5, 76, 0, 0, 144, 145, 5, 95, 0, 0, 145, 146, 5,
		72, 0, 0, 146, 147, 5, 73, 0, 0, 147, 148, 5, 78, 0, 0, 148, 149, 5, 84,
		0, 0, 149, 10, 1, 0, 0, 0, 150, 151, 5, 83, 0, 0, 151, 152, 5, 86, 0, 0,
		152, 153, 5, 67, 0, 0, 153, 154, 5, 95, 0, 0, 154, 155, 5, 72, 0, 0, 155,
		156, 5, 73, 0, 0, 156, 157, 5, 78, 0, 0, 157, 158, 5, 84, 0, 0, 158, 12,
		1, 0, 0, 0, 159, 160, 5, 83, 0, 0, 160, 161, 5, 84, 0, 0, 161, 162, 5,
		82, 0, 0, 162, 163, 5, 85, 0, 0, 163, 164, 5, 67, 0, 0, 164, 165, 5, 84,
		0, 0, 165, 166, 5, 95, 0, 0, 166, 167, 5, 72, 0, 0, 167, 168, 5, 73, 0,
		0, 168, 169, 5, 78, 0, 0, 169, 170, 5, 84, 0, 0, 170, 14, 1, 0, 0, 0, 171,
		172, 5, 116, 0, 0, 172, 173, 5, 114, 0, 0, 173, 174, 5, 117, 0, 0, 174,
		175, 5, 101, 0, 0, 175, 16, 1, 0, 0, 0, 176, 177, 5, 102, 0, 0, 177, 178,
		5, 97, 0, 0, 178, 179, 5, 108, 0, 0, 179, 180, 5, 115, 0, 0, 180, 181,
		5, 101, 0, 0, 181, 18, 1, 0, 0, 0, 182, 183, 5, 115, 0, 0, 183, 184, 5,
		116, 0, 0, 184, 185, 5, 114, 0, 0, 185, 186, 5, 117, 0, 0, 186, 187, 5,
		99, 0, 0, 187, 188, 5, 116, 0, 0, 188, 20, 1, 0, 0, 0, 189, 190, 5, 123,
		0, 0, 190, 22, 1, 0, 0, 0, 191, 192, 5, 125, 0, 0, 192, 24, 1, 0, 0, 0,
		193, 194, 5, 101, 0, 0, 194, 195, 5, 120, 0, 0, 195, 196, 5, 116, 0, 0,
		196, 197, 5, 101, 0, 0, 197, 198, 5, 110, 0, 0, 198, 199, 5, 100, 0, 0,
		199, 200, 5, 115, 0, 0, 200, 26, 1, 0, 0, 0, 201, 202, 5, 80, 0, 0, 202,
		203, 5, 79, 0, 0, 203, 204, 5, 73, 0, 0, 204, 205, 5, 78, 0, 0, 205, 206,
		5, 84, 0, 0, 206, 28, 1, 0, 0, 0, 207, 208, 5, 115, 0, 0, 208, 209, 5,
		101, 0, 0, 209, 210, 5, 114, 0, 0, 210, 211, 5, 118, 0, 0, 211, 212, 5,
		105, 0, 0, 212, 213, 5, 99, 0, 0, 213, 214, 5, 101, 0, 0, 214, 30, 1, 0,
		0, 0, 215, 216, 5, 114, 0, 0, 216, 217, 5, 101, 0, 0, 217, 218, 5, 113,
		0, 0, 218, 219, 5, 117, 0, 0, 219, 220, 5, 105, 0, 0, 220, 221, 5, 114,
		0, 0, 221, 222, 5, 101, 0, 0, 222, 223, 5, 100, 0, 0, 223, 32, 1, 0, 0,
		0, 224, 225, 5, 111, 0, 0, 225, 226, 5, 112, 0, 0, 226, 227, 5, 116, 0,
		0, 227, 228, 5, 105, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230, 5, 110, 0,
		0, 230, 231, 5, 97, 0, 0, 231, 232, 5, 108, 0, 0, 232, 34, 1, 0, 0, 0,
		233, 234, 5, 80, 0, 0, 234, 235, 5, 79, 0, 0, 235, 236, 5, 83, 0, 0, 236,
		237, 5, 84, 0, 0, 237, 36, 1, 0, 0, 0, 238, 239, 5, 40, 0, 0, 239, 38,
		1, 0, 0, 0, 240, 241, 5, 41, 0, 0, 241, 40, 1, 0, 0, 0, 242, 243, 5, 71,
		0, 0, 243, 244, 5, 69, 0, 0, 244, 245, 5, 84, 0, 0, 245, 42, 1, 0, 0, 0,
		246, 247, 5, 85, 0, 0, 247, 248, 5, 82, 0, 0, 248, 249, 5, 76, 0, 0, 249,
		44, 1, 0, 0, 0, 250, 251, 5, 80, 0, 0, 251, 252, 5, 65, 0, 0, 252, 253,
		5, 71, 0, 0, 253, 254, 5, 69, 0, 0, 254, 255, 5, 65, 0, 0, 255, 256, 5,
		66, 0, 0, 256, 257, 5, 76, 0, 0, 257, 258, 5, 69, 0, 0, 258, 46, 1, 0,
		0, 0, 259, 260, 5, 108, 0, 0, 260, 261, 5, 105, 0, 0, 261, 262, 5, 115,
		0, 0, 262, 263, 5, 116, 0, 0, 263, 48, 1, 0, 0, 0, 264, 265, 5, 60, 0,
		0, 265, 50, 1, 0, 0, 0, 266, 267, 5, 62, 0, 0, 267, 52, 1, 0, 0, 0, 268,
		269, 5, 118, 0, 0, 269, 270, 5, 111, 0, 0, 270, 271, 5, 105, 0, 0, 271,
		272, 5, 100, 0, 0, 272, 54, 1, 0, 0, 0, 273, 274, 5, 110, 0, 0, 274, 275,
		5, 111, 0, 0, 275, 276, 5, 116, 0, 0, 276, 277, 5, 95, 0, 0, 277, 278,
		5, 108, 0, 0, 278, 279, 5, 111, 0, 0, 279, 280, 5, 103, 0, 0, 280, 281,
		5, 105, 0, 0, 281, 282, 5, 110, 0, 0, 282, 56, 1, 0, 0, 0, 283, 284, 5,
		109, 0, 0, 284, 285, 5, 97, 0, 0, 285, 286, 5, 112, 0, 0, 286, 58, 1, 0,
		0, 0, 287, 288, 5, 91, 0, 0, 288, 60, 1, 0, 0, 0, 289, 290, 5, 93, 0, 0,
		290, 62, 1, 0, 0, 0, 291, 292, 5, 58, 0, 0, 292, 64, 1, 0, 0, 0, 293, 294,
		5, 59, 0, 0, 294, 66, 1, 0, 0, 0, 295, 297, 7, 0, 0, 0, 296, 295, 1, 0,
		0, 0, 296, 297, 1, 0, 0, 0, 297, 299, 1, 0, 0, 0, 298, 300, 3, 97, 48,
		0, 299, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301,
		302, 1, 0, 0, 0, 302, 68, 1, 0, 0, 0, 303, 305, 5, 45, 0, 0, 304, 303,
		1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 5, 48,
		0, 0, 307, 308, 5, 120, 0, 0, 308, 310, 1, 0, 0, 0, 309, 311, 3, 99, 49,
		0, 310, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312,
		313, 1, 0, 0, 0, 313, 70, 1, 0, 0, 0, 314, 316, 7, 0, 0, 0, 315, 314, 1,
		0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 336, 1, 0, 0, 0, 317, 319, 3, 97, 48,
		0, 318, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320,
		321, 1, 0, 0, 0, 321, 328, 1, 0, 0, 0, 322, 324, 5, 46, 0, 0, 323, 325,
		3, 97, 48, 0, 324, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 324, 1,
		0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 322, 1, 0, 0,
		0, 328, 329, 1, 0, 0, 0, 329, 337, 1, 0, 0, 0, 330, 332, 5, 46, 0, 0, 331,
		333, 3, 97, 48, 0, 332, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 332,
		1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 318, 1, 0,
		0, 0, 336, 330, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 339, 7, 1, 0, 0,
		339, 341, 3, 67, 33, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341,
		72, 1, 0, 0, 0, 342, 343, 5, 98, 0, 0, 343, 344, 5, 111, 0, 0, 344, 345,
		5, 111, 0, 0, 345, 346, 5, 108, 0, 0, 346, 74, 1, 0, 0, 0, 347, 348, 5,
		98, 0, 0, 348, 349, 5, 121, 0, 0, 349, 350, 5, 116, 0, 0, 350, 351, 5,
		101, 0, 0, 351, 76, 1, 0, 0, 0, 352, 353, 5, 105, 0, 0, 353, 354, 5, 49,
		0, 0, 354, 355, 5, 54, 0, 0, 355, 78, 1, 0, 0, 0, 356, 357, 5, 105, 0,
		0, 357, 358, 5, 51, 0, 0, 358, 359, 5, 50, 0, 0, 359, 80, 1, 0, 0, 0, 360,
		361, 5, 105, 0, 0, 361, 362, 5, 54, 0, 0, 362, 363, 5, 52, 0, 0, 363, 82,
		1, 0, 0, 0, 364, 365, 5, 100, 0, 0, 365, 366, 5, 111, 0, 0, 366, 367, 5,
		117, 0, 0, 367, 368, 5, 98, 0, 0, 368, 369, 5, 108, 0, 0, 369, 370, 5,
		101, 0, 0, 370, 84, 1, 0, 0, 0, 371, 372, 5, 115, 0, 0, 372, 373, 5, 116,
		0, 0, 373, 374, 5, 114, 0, 0, 374, 375, 5, 105, 0, 0, 375, 376, 5, 110,
		0, 0, 376, 377, 5, 103, 0, 0, 377, 86, 1, 0, 0, 0, 378, 383, 5, 34, 0,
		0, 379, 382, 3, 89, 44, 0, 380, 382, 8, 2, 0, 0, 381, 379, 1, 0, 0, 0,
		381, 380, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383,
		384, 1, 0, 0, 0, 384, 386, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 386, 397,
		5, 34, 0, 0, 387, 392, 5, 39, 0, 0, 388, 391, 3, 89, 44, 0, 389, 391, 8,
		3, 0, 0, 390, 388, 1, 0, 0, 0, 390, 389, 1, 0, 0, 0, 391, 394, 1, 0, 0,
		0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394,
		392, 1, 0, 0, 0, 395, 397, 5, 39, 0, 0, 396, 378, 1, 0, 0, 0, 396, 387,
		1, 0, 0, 0, 397, 88, 1, 0, 0, 0, 398, 399, 5, 92, 0, 0, 399, 400, 7, 4,
		0, 0, 400, 90, 1, 0, 0, 0, 401, 404, 3, 95, 47, 0, 402, 404, 5, 95, 0,
		0, 403, 401, 1, 0, 0, 0, 403, 402, 1, 0, 0, 0, 404, 410, 1, 0, 0, 0, 405,
		409, 3, 95, 47, 0, 406, 409, 3, 97, 48, 0, 407, 409, 7, 5, 0, 0, 408, 405,
		1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 408, 407, 1, 0, 0, 0, 409, 412, 1, 0,
		0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 92, 1, 0, 0, 0,
		412, 410, 1, 0, 0, 0, 413, 414, 5, 44, 0, 0, 414, 94, 1, 0, 0, 0, 415,
		416, 7, 6, 0, 0, 416, 96, 1, 0, 0, 0, 417, 418, 2, 48, 57, 0, 418, 98,
		1, 0, 0, 0, 419, 422, 3, 97, 48, 0, 420, 422, 7, 7, 0, 0, 421, 419, 1,
		0, 0, 0, 421, 420, 1, 0, 0, 0, 422, 100, 1, 0, 0, 0, 423, 428, 7, 8, 0,
		0, 424, 425, 5, 13, 0, 0, 425, 428, 5, 10, 0, 0, 426, 428, 5, 10, 0, 0,
		427, 423, 1, 0, 0, 0, 427, 424, 1, 0, 0, 0, 427, 426, 1, 0, 0, 0, 428,
		429, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 431,
		1, 0, 0, 0, 431, 432, 6, 50, 0, 0, 432, 102, 1, 0, 0, 0, 433, 434, 5, 47,
		0, 0, 434, 437, 5, 47, 0, 0, 435, 437, 5, 35, 0, 0, 436, 433, 1, 0, 0,
		0, 436, 435, 1, 0, 0, 0, 437, 441, 1, 0, 0, 0, 438, 440, 8, 9, 0, 0, 439,
		438, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 441, 442,
		1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 444, 446, 5, 13,
		0, 0, 445, 444, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0,
		447, 448, 5, 10, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 6, 51, 0, 0, 450,
		104, 1, 0, 0, 0, 451, 452, 5, 47, 0, 0, 452, 453, 5, 42, 0, 0, 453, 457,
		1, 0, 0, 0, 454, 456, 9, 0, 0, 0, 455, 454, 1, 0, 0, 0, 456, 459, 1, 0,
		0, 0, 457, 458, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 460, 1, 0, 0, 0,
		459, 457, 1, 0, 0, 0, 460, 461, 5, 42, 0, 0, 461, 462, 5, 47, 0, 0, 462,
		463, 1, 0, 0, 0, 463, 464, 6, 52, 0, 0, 464, 106, 1, 0, 0, 0, 27, 0, 296,
		301, 304, 312, 315, 320, 326, 328, 334, 336, 340, 381, 383, 390, 392, 396,
		403, 408, 410, 421, 427, 429, 436, 441, 445, 457, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ServiceLexerInit initializes any static state used to implement ServiceLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewServiceLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ServiceLexerInit() {
	staticData := &ServiceLexerLexerStaticData
	staticData.once.Do(servicelexerLexerInit)
}

// NewServiceLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewServiceLexer(input antlr.CharStream) *ServiceLexer {
	ServiceLexerInit()
	l := new(ServiceLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ServiceLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Service.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ServiceLexer tokens.
const (
	ServiceLexerT__0        = 1
	ServiceLexerT__1        = 2
	ServiceLexerT__2        = 3
	ServiceLexerT__3        = 4
	ServiceLexerT__4        = 5
	ServiceLexerT__5        = 6
	ServiceLexerT__6        = 7
	ServiceLexerT__7        = 8
	ServiceLexerT__8        = 9
	ServiceLexerT__9        = 10
	ServiceLexerT__10       = 11
	ServiceLexerT__11       = 12
	ServiceLexerT__12       = 13
	ServiceLexerT__13       = 14
	ServiceLexerT__14       = 15
	ServiceLexerT__15       = 16
	ServiceLexerT__16       = 17
	ServiceLexerT__17       = 18
	ServiceLexerT__18       = 19
	ServiceLexerT__19       = 20
	ServiceLexerT__20       = 21
	ServiceLexerT__21       = 22
	ServiceLexerT__22       = 23
	ServiceLexerT__23       = 24
	ServiceLexerT__24       = 25
	ServiceLexerT__25       = 26
	ServiceLexerT__26       = 27
	ServiceLexerT__27       = 28
	ServiceLexerT__28       = 29
	ServiceLexerT__29       = 30
	ServiceLexerT__30       = 31
	ServiceLexerT__31       = 32
	ServiceLexerT__32       = 33
	ServiceLexerINTEGER     = 34
	ServiceLexerHEX_INTEGER = 35
	ServiceLexerDOUBLE      = 36
	ServiceLexerTYPE_BOOL   = 37
	ServiceLexerTYPE_BYTE   = 38
	ServiceLexerTYPE_I16    = 39
	ServiceLexerTYPE_I32    = 40
	ServiceLexerTYPE_I64    = 41
	ServiceLexerTYPE_DOUBLE = 42
	ServiceLexerTYPE_STRING = 43
	ServiceLexerLITERAL     = 44
	ServiceLexerIDENTIFIER  = 45
	ServiceLexerCOMMA       = 46
	ServiceLexerWS          = 47
	ServiceLexerSL_COMMENT  = 48
	ServiceLexerML_COMMENT  = 49
)
