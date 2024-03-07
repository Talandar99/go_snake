package main

<<<<<<< HEAD
import rl "github.com/gen2brain/raylib-go/raylib"

func updateKeys(direction *int) {
	if rl.IsKeyPressed(rl.KeyW) && *direction != 3 {
		*direction = 1
	}
	if rl.IsKeyPressed(rl.KeyA) && *direction != 4 {
		*direction = 2
	}
	if rl.IsKeyPressed(rl.KeyS) && *direction != 1 {
		*direction = 3
	}
	if rl.IsKeyPressed(rl.KeyD) && *direction != 2 {
		*direction = 4
	}
}

func updateSnake(snake *[400]rl.Vector2, direction *int, snack *rl.Vector2) {
	var newHead = rl.Vector2{X: (*snake)[0].X, Y: (*snake)[0].Y}
	switch *direction {
	case 1:
		newHead.Y = (*snake)[0].Y - 1
		if newHead.Y < 0 {
			newHead.Y = 0
		}
	case 2:
		newHead.X = (*snake)[0].X - 1
		if newHead.X < 0 {
			newHead.X = 0
		}
	case 3:
		if *direction != 1 {
			newHead.Y = (*snake)[0].Y + 1
			if newHead.Y >= 20 {
				newHead.Y = 19
			}
		}
	default:
		newHead.X = (*snake)[0].X + 1
		if newHead.X >= 20 {
			newHead.X = 19
		}
	}
	if newHead.X == snack.X && newHead.Y == snack.Y {
		snack.X = (float32)(rl.GetRandomValue(0, 19))
		snack.Y = (float32)(rl.GetRandomValue(0, 19))
		for i := len((*snake)) - 1; i > 0; i-- {
			if (*snake)[i].X != -1 && (*snake)[i].Y != -1 {
				(*snake)[i+1].X = (*snake)[i].X
				(*snake)[i+1].Y = (*snake)[i].Y
			}
		}
	}
	for i := len((*snake)) - 1; i >= 0; i-- {
		if i > 0 {
			if (*snake)[i].X != -1 && (*snake)[i].Y != -1 {
				(*snake)[i].X = (*snake)[i-1].X
				(*snake)[i].Y = (*snake)[i-1].Y
			}
		} else {
			(*snake)[0].X = newHead.X
			(*snake)[0].Y = newHead.Y
		}
	}
}

func update(frame *int, direction *int, snake *[400]rl.Vector2, snack *rl.Vector2) {
	*frame++
	updateKeys(direction)
	if *frame >= 10 {
		*frame = 0
		updateSnake(snake, direction, snack)
	}
}

func drawGrid(snake [400]rl.Vector2, snack rl.Vector2) {
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			for i := 0; i < len(snake); i++ {
				if (int)(snake[i].X) == x && (int)(snake[i].Y) == y {
					rl.DrawRectangleRec(
						rl.NewRectangle((float32)(x*30), (float32)(y*30), 30, 30), rl.Lime)
					rl.DrawRectangleLinesEx(
						rl.NewRectangle((float32)(x*30), (float32)(y*30), 30, 30), 5, rl.DarkGreen)
				}
				if (int)(snack.X) == x && (int)(snack.Y) == y {
					rl.DrawRectangleRec(
						rl.NewRectangle((float32)(x*30), (float32)(y*30), 30, 30), rl.Red)
					rl.DrawRectangleLinesEx(
						rl.NewRectangle((float32)(x*30), (float32)(y*30), 30, 30), 5, rl.Maroon)
				}
			}
		}
	}
}

func draw(snake [400]rl.Vector2, snack rl.Vector2) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	drawGrid(snake, snack)
	rl.EndDrawing()
}
func main() {

	rl.InitWindow(600, 600, "go snake")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	direction := 0
	frame := 0
	var snake [400]rl.Vector2
	var snack = rl.Vector2{X: 10, Y: 10}
	for i := 0; i < len(snake); i++ {
		snake[i] = rl.Vector2{X: -1, Y: -1}
	}
	snake[2] = rl.Vector2{X: 0, Y: 0}
	snake[1] = rl.Vector2{X: 1, Y: 0}
	snake[0] = rl.Vector2{X: 2, Y: 0}

	for !rl.WindowShouldClose() {
		update(&frame, &direction, &snake, &snack)
		draw(snake, snack)
	}
=======
func main() {
	println("Hello from Go")
>>>>>>> add1977 (init)
}
