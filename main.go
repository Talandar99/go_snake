package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawGrid(head rl.Vector2) {
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if (int)(head.X) == x && (int)(head.Y) == y {

				rl.DrawRectangleRec(
					rl.NewRectangle(
						(float32)(x*30),
						(float32)(y*30),
						30, 30), rl.DarkGreen)
			} else {
				rl.DrawRectangleLinesEx(
					rl.NewRectangle(
						(float32)(x*30),
						(float32)(y*30),
						30, 30), 1, rl.DarkGreen)
			}
		}
	}
}

func main() {

	rl.InitWindow(600, 600, "go snake")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var frame, direction int
	direction = 0
	frame = 0
	var head rl.Vector2 = rl.NewVector2(0, 0)
	for !rl.WindowShouldClose() {
		frame++
		if frame > 60 {
			frame = 0
		}
		// update
		if rl.IsKeyPressed(rl.KeyW) && direction != 3 {
			direction = 1
		}
		if rl.IsKeyPressed(rl.KeyA) && direction != 4 {
			direction = 2
		}
		if rl.IsKeyPressed(rl.KeyS) && direction != 1 {
			direction = 3
		}
		if rl.IsKeyPressed(rl.KeyD) && direction != 2 {
			direction = 4
		}
		if frame%5 == 0 {
			switch direction {
			case 1:
				head.Y = head.Y - 1
				if head.Y < 0 {
					head.Y = 0
				}
			case 2:
				head.X = head.X - 1
				if head.X < 0 {
					head.X = 0
				}
			case 3:
				if direction != 1 {
					head.Y = head.Y + 1
					if head.Y >= 20 {
						head.Y = 19
					}
				}
			default:
				head.X = head.X + 1
				if head.X >= 20 {
					head.X = 19
				}
			}
		}
		//
		rl.BeginDrawing()
		// draw
		rl.ClearBackground(rl.Black)
		drawGrid(head)

		rl.EndDrawing()
	}
}
