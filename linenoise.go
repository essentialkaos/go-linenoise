//go:build cgo || !windows
// +build cgo !windows

// Package linenoise wraps the linenoise library (https://github.com/yhirose/linenoise).
package linenoise

// ///////////////////////////////////////////////////////////////////////////////// //

// #include <stdlib.h>
// #include "linenoise.h"
// #include "utf8.h"
// #include "hooks.h"
import "C"

// ///////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"unsafe"
)

// ///////////////////////////////////////////////////////////////////////////////// //

// CompletionHandler provides possible completions for given input
type CompletionHandler func(input string) []string

// HintHandler provides hint for user input
type HintHandler func(input string) string

// ///////////////////////////////////////////////////////////////////////////////// //

// ErrKillSignal is returned returned by Line() when a user quits from prompt.
// This occurs when the user enters ctrl+C or ctrl+D.
var ErrKillSignal = errors.New("Prompt was quited with a kill signal")

// ///////////////////////////////////////////////////////////////////////////////// //

// complHandler is completion handler function
var complHandler CompletionHandler = emptyCompomplHandler

// hintHandler is hint handler function
var hintHandler HintHandler = emptyHintHandler

// hintColor contains hint color ANSI code (dark grey by default)
var hintColor uint8 = 90

// hintBold contains flag for bold hints
var hintBold bool = false

// ///////////////////////////////////////////////////////////////////////////////// //

var emptyCompomplHandler = func(input string) []string { return nil }
var emptyHintHandler = func(input string) string { return "" }

// ///////////////////////////////////////////////////////////////////////////////// //

// Line displays given string and returns line from user input
func Line(prompt string) (string, error) {
	promptCString := C.CString(prompt)
	resultCString := C.linenoise(promptCString)

	C.free(unsafe.Pointer(promptCString))

	defer C.free(unsafe.Pointer(resultCString))

	if resultCString == nil {
		return "", ErrKillSignal
	}

	result := C.GoString(resultCString)

	return result, nil
}

// AddHistory adds a line to history. Returns non-nil error on fail
func AddHistory(line string) error {
	lineCString := C.CString(line)
	res := C.linenoiseHistoryAdd(lineCString)

	C.free(unsafe.Pointer(lineCString))

	if res != 1 {
		return errors.New("Could not add line to history")
	}

	return nil
}

// SetHistoryCapacity changes the maximum length of history. Returns non-nil error on fail
func SetHistoryCapacity(capacity int) error {
	res := C.linenoiseHistorySetMaxLen(C.int(capacity))

	if res != 1 {
		return errors.New("Could not set history max len")
	}

	return nil
}

// SaveHistory saves from file with given filename. Returns non-nil error on fail
func SaveHistory(filename string) error {
	filenameCString := C.CString(filename)
	res := C.linenoiseHistorySave(filenameCString)

	C.free(unsafe.Pointer(filenameCString))

	if res != 0 {
		return errors.New("Could not save history to file")
	}

	return nil
}

// LoadHistory loads from file with given filename. Returns non-nil error on fail
func LoadHistory(filename string) error {
	filenameCString := C.CString(filename)
	res := C.linenoiseHistoryLoad(filenameCString)

	C.free(unsafe.Pointer(filenameCString))

	if res != 0 {
		return errors.New("Could not load history from file")
	}

	return nil
}

// Clear clears the screen
func Clear() {
	C.linenoiseClearScreen()
}

// SetMultiline sets linenoise to multiline or single line.
// In multiline mode the user input will be wrapped to a new line when the length exceeds the amount of available rows in the terminal.
func SetMultiline(enable bool) {
	if enable {
		C.linenoiseSetMultiLine(1)
	} else {
		C.linenoiseSetMultiLine(0)
	}
}

// SetMaskMode sets mask mode. When it is enabled, instead of the input that the user
// is typing, the terminal will just display a corresponding number of asterisks,
// like "****". This is useful for passwords and other secrets that should not
// be displayed.
func SetMaskMode(enable bool) {
	if enable {
		C.linenoiseMaskModeEnable()
	} else {
		C.linenoiseMaskModeDisable()
	}
}

// SetCompletionHandler sets the CompletionHandler to be used for completion (using Tab key)
//
// You can pass nil as the handler to remove the previously set handler
func SetCompletionHandler(h CompletionHandler) {
	if h == nil {
		complHandler = emptyCompomplHandler
		return
	}

	complHandler = h
}

// SetHintHandler sets the HintHandler to be used for input hints
//
// You can pass nil as the handler to remove the previously set handler
func SetHintHandler(h HintHandler) {
	if h == nil {
		hintHandler = emptyHintHandler
		return
	}

	hintHandler = h
}

// SetHintColor sets hint text color to color with given ANSI code
//
// Color codes: https://github.com/essentialkaos/fmtc/wiki#816-colors
func SetHintColor(color uint8, bold bool) error {
	if color < 30 || color > 97 || (color > 37 && color < 90) {
		return fmt.Errorf("\"%d\" is not valid ANSI color code", color)
	}

	hintColor, hintBold = color, bold

	return nil
}

// PrintKeyCodes puts linenoise in key codes debugging mode.
// Press keys and key combinations to see key codes. Type 'quit' at any time to exit.
// PrintKeyCodes blocks until user enters 'quit'.
func PrintKeyCodes() {
	C.linenoisePrintKeyCodes()
}

// ///////////////////////////////////////////////////////////////////////////////// //

// init inits callbacks
func init() {
	C.linenoiseSetupCompletionCallbackHook()
	C.linenoiseSetupHintCallbackHook()
}

//export linenoiseGoCompletionCallbackHook
func linenoiseGoCompletionCallbackHook(input *C.char, completions *C.linenoiseCompletions) {
	completionsSlice := complHandler(C.GoString(input))
	completionsLen := len(completionsSlice)
	completions.len = C.size_t(completionsLen)

	if completionsLen == 0 {
		return
	}

	cvec := C.malloc(C.size_t(int(unsafe.Sizeof(*(**C.char)(nil))) * completionsLen))
	cvecSlice := (*(*[999999]*C.char)(cvec))[:completionsLen]

	for i, str := range completionsSlice {
		cvecSlice[i] = C.CString(str)
	}

	completions.cvec = (**C.char)(cvec)
}

//export linenoiseGoHintCallbackHook
func linenoiseGoHintCallbackHook(input *C.char, color *C.int, bold *C.int) *C.char {
	hintText := hintHandler(C.GoString(input))

	if hintText == "" {
		return nil
	}

	*color = C.int(hintColor)

	if hintBold {
		*bold = C.int(1)
	} else {
		*bold = C.int(0)
	}

	return C.CString(hintText)
}
