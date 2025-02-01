package output

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/yourpwnguy/refine/pkg/common"
	"github.com/yourpwnguy/refine/pkg/types"
)

// Absolutely shit !!
func BeautifyPrint(params types.Params) {
	// Format the time duration as "1mXs", "Xs", or "Xms" depending on the value
	timeTaken := params.TimeTaken
	var formattedTime string

	if timeTaken.Seconds() >= 60 {
		minutes := int(timeTaken.Minutes())
		seconds := int(timeTaken.Seconds()) % 60
		formattedTime = fmt.Sprintf("%dm%ds", minutes, seconds)
	} else {
		// Regex to match pattern with number and unit (s, ms, μs)
		re := regexp.MustCompile(`(\d+)\.\d*\s*(s|ms|µs)`)
		match := re.FindStringSubmatch(timeTaken.String())
		if len(match) > 0 {
			formattedTime = fmt.Sprintf("%s%s", match[1], match[2])
		} else {
			formattedTime = timeTaken.String()
		}
	}

	// Print the formatted output (From not stdin)
	// TODO: IS THERE ANY FORMAT TO PRINT THIS ?
	if params.OutputFile != "" {
		fmt.Fprintln(os.Stderr,
			"("+common.G.Bold(common.G.BrBlue("*"))+")",
			common.G.Bold(common.G.BrRed("OUT")),
			common.G.White("->"),
			common.G.Bold(common.G.BrYellow(params.OutputFile)),
			common.G.White("|"),
			common.G.Bold(common.G.BrBlue("TOT")),
			common.G.White("->"),
			common.G.Bold(common.G.BrYellow(strconv.Itoa(params.TotalLinesCount))),
			common.G.White("|"),
			common.G.Bold(common.G.BrBlue("UNQ")),
			common.G.White("->"),
			common.G.Bold(common.G.BrYellow(strconv.Itoa(params.UniqueLinesCount))),
			common.G.White("|"),
			common.G.Bold(common.G.BrCyan("DUR")),
			common.G.White("->"),
			common.G.Bold(common.G.BrGreen(formattedTime)),
		)
	} else {
		fmt.Fprintln(os.Stderr,
			"\n("+common.G.Bold(common.G.BrBlue("*"))+")",
			common.G.Bold(common.G.BrBlue("TOT")),
			common.G.White("->"),
			common.G.Bold(common.G.BrYellow(strconv.Itoa(params.TotalLinesCount))),
			common.G.White("|"),
			common.G.Bold(common.G.BrBlue("UNQ")),
			common.G.White("->"),
			common.G.Bold(common.G.BrYellow(strconv.Itoa(params.UniqueLinesCount))),
			common.G.White("|"),
			common.G.Bold(common.G.BrCyan("DUR")),
			common.G.White("->"),
			common.G.Bold(common.G.BrGreen(formattedTime)),
		)
	}
}
