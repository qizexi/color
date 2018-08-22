/*
Package color is command line color library, written using golang

Source code and other details for the project are available at GitHub:

	https://github.com/gookit/color

More usage please see README and tests.
*/
package color

import (
	"fmt"
	"regexp"
)

// console color mode
const (
	ModeNormal    = iota
	Mode256        // 8 bite
	ModeRGB        // 24 bite
	ModeGrayscale
)

// color render templates
const (
	SettingTpl   = "\x1b[%sm"
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
)

// ResetCode value
const ResetCode = "0"

// ResetSet 重置/正常 关闭所有属性。
const ResetSet = "\x1b[0m"

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

// Enable switch color display
var Enable = true

var (
	// mark current env, It's like in cmd.exe
	isLikeInCmd bool
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
	// mark current env is support color.
	// Always: isLikeInCmd != isSupportColor
	isSupportColor = IsSupportColor()
)

/*************************************************************
 * global settings
 *************************************************************/

// Set set console color attributes
func Set(colors ...Color) (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	// on windows cmd.exe
	if isLikeInCmd {
		return winSet(colors...)
	}

	return fmt.Printf(SettingTpl, colors2code(colors...))
}

// Reset reset console color attributes
func Reset() (int, error) {
	if !Enable { // not enable
		return 0, nil
	}

	// on windows cmd.exe
	if isLikeInCmd {
		return winReset()
	}

	return fmt.Print(ResetSet)
}

// Disable disable color output
func Disable() {
	Enable = false
}

/*************************************************************
 * render color code
 *************************************************************/

// RenderCode render message by color code.
// Usage:
// 	msg := RenderCode("3;32;45", "some", "message")
func RenderCode(code string, args ...interface{}) string {
	message := fmt.Sprint(args...)
	if len(code) == 0 {
		return message
	}

	// disabled OR not support color
	if !Enable || !isSupportColor {
		return ClearCode(message)
	}

	return fmt.Sprintf(FullColorTpl, code, message)
}

// RenderString render a string with color code.
// Usage:
// 	msg := RenderString("3;32;45", "a message")
func RenderString(code string, str string) string {
	// some check
	if len(code) == 0 || str == "" {
		return str
	}

	// disabled OR not support color
	if !Enable || !isSupportColor {
		return ClearCode(str)
	}

	return fmt.Sprintf(FullColorTpl, code, str)
}

// ClearCode clear color codes.
// eg:
// 		"\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	return codeRegex.ReplaceAllString(str, "")
}
