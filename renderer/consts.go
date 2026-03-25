package renderer

import _ "embed"

//go:embed header.html
var HTML_SVG_HEADER string

//go:embed footer.html
var HTML_SVG_FOOTER string

//go:embed stats.html
var HTML_STATS_BOX string

const (
	BACKGROUND_COLOR    string = "#000000"
	CIRCLE_FILL_COLOR   string = "#578cff"
	CIRCLE_RADIUS       int    = 45
	STROKE_WIDTH        int    = 1
	STROKE_COLOR        string = "#cecece"
	TEXT_COLOR          string = "#FFFFFF"
	NODE_TEXT_FONT_SIZE int    = 12
)
