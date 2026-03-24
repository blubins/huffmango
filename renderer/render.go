package renderer

import (
	"fmt"
	huffmantree "huffmango/huffman"
	"huffmango/node"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func inOrderTraversalDraw(n *node.Node, x, y, xOffset int32, weight int32) {
	if n == nil {
		return
	}

	yGap := int32(80)
	if weight < 50 {
		weight = 50
	}
	weight = weight * 2 / 5
	if n.Left != nil {
		childX := x - xOffset - weight
		childY := y + yGap
		rl.DrawLineEx(rl.Vector2{X: float32(x), Y: float32(y + 20)}, rl.Vector2{X: float32(childX), Y: float32(childY)}, 2, rl.White)
		inOrderTraversalDraw(n.Left, childX, childY, xOffset/3, weight)
	}

	if n.Right != nil {
		childX := x + xOffset + weight
		childY := y + yGap
		rl.DrawLineEx(rl.Vector2{X: float32(x), Y: float32(y + 20)}, rl.Vector2{X: float32(childX), Y: float32(childY)}, 2, rl.White)
		inOrderTraversalDraw(n.Right, childX, childY, xOffset/3, weight)
	}

	rl.DrawCircle(x, y, 25, rl.SkyBlue)
	data := n.Data.(huffmantree.Data)
	rl.DrawText(fmt.Sprintf("%d\n%d", data.B, data.Freq), x-15, y+30, 15, rl.LightGray)

}

func Render(root *node.Node) {
	rl.InitWindow(2560, 1440, "Huffman Tree Viewer")
	rl.ToggleFullscreen()
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		inOrderTraversalDraw(root, 1280, 40, 720, 200)
		rl.EndDrawing()
	}
}
