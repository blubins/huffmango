package renderer

import (
	_ "embed"
	"fmt"
	huffmantree "huffmango/huffman"
	"huffmango/node"
	"os"
	"strconv"
	"strings"
)

func handleOverlap(n *node.Node, depth int32, counter *int32, xGap, yGap, startY int32, positions map[*node.Node]vector2) {
	if n == nil {
		return
	}
	// left node taken
	handleOverlap(n.Left, depth+1, counter, xGap, yGap, startY, positions)
	// save pos of curr node
	positions[n] = vector2{
		x: *counter * xGap / 2,
		y: startY + depth*yGap,
	}
	*counter++
	// right node taken
	handleOverlap(n.Right, depth+1, counter, xGap, yGap, startY, positions)
}

func inOrderSVGAppend(n *node.Node, positions map[*node.Node]vector2, builder *strings.Builder) {
	if n == nil {
		return
	}
	// grab the pos of the current node
	pos := positions[n]
	// inorder travere the node with drawing strokes

	// left child
	if n.Left != nil {
		circlePos := positions[n.Left]
		fmt.Fprintf(
			builder,
			`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="%d"/>%s`,
			pos.x, pos.y+20, circlePos.x, circlePos.y-20, STROKE_COLOR, STROKE_WIDTH, "\n",
		)

		inOrderSVGAppend(n.Left, positions, builder)
	}

	// right child
	if n.Right != nil {
		circlePos := positions[n.Right]
		fmt.Fprintf(
			builder,
			`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="%d"/>%s`,
			pos.x, pos.y+20, circlePos.x, circlePos.y-20, STROKE_COLOR, STROKE_WIDTH, "\n",
		)

		inOrderSVGAppend(n.Right, positions, builder)
	}

	// finally draw the circle
	fmt.Fprintf(
		builder,
		`<circle cx="%d" cy="%d" r="25" fill="%s"/>%s`,
		pos.x,
		pos.y,
		CIRCLE_FILL_COLOR,
		"\n",
	)

	// draw the text underneath
	// byte value
	fmt.Fprintf(
		builder,
		`<text x="%d" y="%d" class="nodeText" fill=%s>0x%x</text>%s`,
		pos.x-int32(CIRCLE_RADIUS)+20,
		pos.y+int32(CIRCLE_RADIUS)-5,
		TEXT_COLOR,
		n.Data.(huffmantree.Data).B,
		"\n",
	)

	// draw frequency text underneath the byte value
	fmt.Fprintf(
		builder,
		`<text x="%d" y="%d" class="nodeText" fill=%s>F: %d</text>%s`,
		pos.x-int32(CIRCLE_RADIUS)+20,
		pos.y+int32(CIRCLE_RADIUS)+int32(NODE_TEXT_FONT_SIZE)-5,
		TEXT_COLOR,
		n.Data.(huffmantree.Data).Freq,
		"\n",
	)
}

// made svg html view creator function
func CreateHTMLView(root *node.Node, stats *huffmantree.Stat, outputPath string) error {
	// initialize a map of node pointers to x,y positions in a plane
	nodes := make(map[*node.Node]vector2)
	counter := int32(1)
	// modified inorder traversal that builds the svg viewer
	// and handles x,y coordinate overlaps
	handleOverlap(root, 0, &counter, 60, 120, 40, nodes)

	svgW := counter * 60
	svgH := int32(1440)
	// initialize io and svg header
	var svgBuilder strings.Builder
	fmt.Fprintf(
		&svgBuilder,
		`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d">%s`,
		svgW,
		svgH,
		"\n",
	)

	fmt.Fprintf(
		&svgBuilder,
		`<rect width="100%%" height="100%%" fill="%s"/>%s`,
		BACKGROUND_COLOR,
		"\n",
	)
	inOrderSVGAppend(root, nodes, &svgBuilder)
	svgBuilder.WriteString("</svg>")
	// huffman tree stats display
	statsHTML := HTML_STATS_BOX
	statsHTML = strings.Replace(statsHTML, "REPLACE_NUM_BITS", fmt.Sprintf("%d", stats.NumBits), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_NUM_ENCODED_BITS", fmt.Sprintf("%d", stats.NumEncodedBits), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_TABLE_SIZE", fmt.Sprintf("%d", stats.TableSize), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_NUM_TOT_BYTES_WRITTEN", fmt.Sprintf("%d", stats.NumTotBytesWritten), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_PERCENT_IMPROVEMENT", fmt.Sprintf("%.3f", stats.PercentImprovement*100), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_PERCENT_ACTUAL_IMPROVEMENT", fmt.Sprintf("%.3f", stats.PercentActualImprovement*100), 1)
	statsHTML = strings.Replace(statsHTML, "REPLACE_AVG_SYMBOL_SIZE", fmt.Sprintf("%.4f", stats.AvgSymbolSize), 1)

	// background color, font size
	var builder strings.Builder
	builder.WriteString(strings.Replace(HTML_SVG_HEADER, "REPLACE_BACKGROUND_COLOR", BACKGROUND_COLOR, 1))
	builder.WriteString(strings.Replace(HTML_SVG_HEADER, "REPLACE_NODE_TEXT_FONT_SIZE", strconv.Itoa(NODE_TEXT_FONT_SIZE), 1))

	// build the string to insert into the file into the builder
	builder.WriteString(statsHTML)
	builder.WriteString(svgBuilder.String())
	builder.WriteString(HTML_SVG_FOOTER)

	// create the file at the outputpath
	f, err := os.Create(fmt.Sprintf("%s", outputPath))
	if err != nil {
		return err
	}
	defer f.Close()

	// finally write the data
	_, err = f.WriteString(builder.String())
	if err != nil {
		return err
	}

	return nil
}
