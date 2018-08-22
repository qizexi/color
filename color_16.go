package color

import (
	"fmt"
	"strings"
)

// Color value type
// 3(2^3=8) OR 4(2^4=16) bite color.
// ESC 操作的表示:
// 	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
type Color uint8

/*************************************************************
 * Basic 16 color definition
 *************************************************************/

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 99
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	OpReset         Color = iota // 0 重置所有设置
	OpBold                       // 1 加粗
	OpFuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	OpItalic                     // 3 斜体(不是所有的终端仿真器都支持)
	OpUnderscore                 // 4 下划线
	OpBlink                      // 5 闪烁
	OpFastBlink                  // 5 快速闪烁(未广泛支持)
	OpReverse                    // 7 颠倒的 交换背景色与前景色
	OpConcealed                  // 8 隐匿的
	OpStrikethrough              // 9 删除的，删除线(未广泛支持)
)

// There are basic foreground color alias
const (
	Red     = FgRed
	Cyan    = FgCyan
	Gray    = FgDarkGray
	Blue    = FgBlue
	Black   = FgBlack
	Green   = FgGreen
	White   = FgWhite
	Yellow  = FgYellow
	Magenta = FgMagenta
	Bold    = OpBold
	Normal  = FgDefault
)

/*************************************************************
 * Color render methods
 *************************************************************/

// Text render a text message
func (c Color) Text(message string) string {
	return RenderString(c.String(), message)
}

// Render messages by color setting
// usage:
// 		green := color.FgGreen.Render
// 		fmt.Println(green("message"))
func (c Color) Render(a ...interface{}) string {
	return RenderCode(c.String(), a...)
}

// Sprint render messages by color setting. is alias of the Render()
func (c Color) Sprint(a ...interface{}) string {
	return RenderCode(c.String(), a...)
}

// Sprintf format and render message.
// Usage:
// 	green := color.Green.Sprintf
//  colored := green("message")
func (c Color) Sprintf(format string, args ...interface{}) string {
	return RenderString(c.String(), fmt.Sprintf(format, args...))
}

// Print messages.
// Usage:
// 		color.Green.Print("message")
// OR:
// 		green := color.FgGreen.Print
// 		green("message")
func (c Color) Print(args ...interface{}) {
	message := fmt.Sprint(args...)
	if isLikeInCmd {
		winPrint(message, c)
	} else {
		fmt.Print(RenderString(c.String(), message))
	}
}

// Printf format and print messages.
// usage:
// 		color.Cyan.Printf("string %s", "arg0")
func (c Color) Printf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if isLikeInCmd {
		winPrint(msg, c)
	} else {
		fmt.Print(RenderString(c.String(), msg))
	}
}

// Println messages with new line
func (c Color) Println(a ...interface{}) {
	if isLikeInCmd {
		winPrintln(fmt.Sprint(a...), c)
	} else {
		fmt.Println(RenderCode(c.String(), a...))
	}
}

// String to code string. eg "35"
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// IsValid color value
func (c Color) IsValid() bool {
	return c < 107
}

/*************************************************************
 * basic color maps
 *************************************************************/

// FgColors foreground colors map
var FgColors = map[string]Color{
	"black":   FgBlack,
	"red":     FgRed,
	"green":   FgGreen,
	"yellow":  FgYellow,
	"blue":    FgBlue,
	"magenta": FgMagenta,
	"cyan":    FgCyan,
	"white":   FgWhite,
	"default": FgDefault,
}

// BgColors background colors map
var BgColors = map[string]Color{
	"black":   BgBlack,
	"red":     BgRed,
	"green":   BgGreen,
	"yellow":  BgYellow,
	"blue":    BgBlue,
	"magenta": BgMagenta,
	"cyan":    BgCyan,
	"white":   BgWhite,
	"default": BgDefault,
}

// ExFgColors extra foreground colors map
var ExFgColors = map[string]Color{
	"darkGray":     FgDarkGray,
	"lightRed":     FgLightRed,
	"lightGreen":   FgLightGreen,
	"lightYellow":  FgLightYellow,
	"lightBlue":    FgLightBlue,
	"lightMagenta": FgLightMagenta,
	"lightCyan":    FgLightCyan,
	"lightWhite":   FgLightWhite,
}

// ExBgColors extra background colors map
var ExBgColors = map[string]Color{
	"darkGray":     BgDarkGray,
	"lightRed":     BgLightRed,
	"lightGreen":   BgLightGreen,
	"lightYellow":  BgLightYellow,
	"lightBlue":    BgLightBlue,
	"lightMagenta": BgLightMagenta,
	"lightCyan":    BgLightCyan,
	"lightWhite":   BgLightWhite,
}

// Options color options map
var Options = map[string]Color{
	"reset":      OpReset,
	"bold":       OpBold,
	"fuzzy":      OpFuzzy,
	"italic":     OpItalic,
	"underscore": OpUnderscore,
	"blink":      OpBlink,
	"reverse":    OpReverse,
	"concealed":  OpConcealed,
}

/*************************************************************
 * helper methods
 *************************************************************/

// convert colors to code. return like "32;45;3"
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}
