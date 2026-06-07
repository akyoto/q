// This example is currently Windows only and it
// needs the raylib.dll file next to the executable.
// It creates a window showing the number of frames
// per second and a rectangle that can be moved with
// the arrow keys.
//
// Linux users can use WINE after adding raylib.dll:
//
//     q build examples/raylib -os windows
//     wine examples/raylib/raylib.exe
//

const {
	canvasWidth = 1280
	canvasHeight = 720
}

main() {
	raylib.InitWindow(canvasWidth, canvasHeight, "raylib example\0".ptr)
	raylib.SetTargetFPS(1000)

	player := new(Player) {
		x: canvasWidth / 2,
		y: canvasHeight / 2,
		width: 10,
		height: 10
	}

	loop {
		if raylib.WindowShouldClose() {
			raylib.CloseWindow()
			delete(player)
			return
		}

		input(player)
		update(player)
		draw(player)
	}
}

input(player *Player) {
	if raylib.IsKeyDown(KEY_RIGHT) {
		player.x += 1
	}

	if raylib.IsKeyDown(KEY_LEFT) {
		player.x -= 1
	}

	if raylib.IsKeyDown(KEY_DOWN) {
		player.y += 1
	}

	if raylib.IsKeyDown(KEY_UP) {
		player.y -= 1
	}
}

update(player *Player) {
	minX := player.width / 2

	if player.x < minX {
		player.x = minX
	}

	maxX := canvasWidth - minX

	if player.x > maxX {
		player.x = maxX
	}

	minY := player.height / 2

	if player.y < minY {
		player.y = minY
	}

	maxY := canvasHeight - minY

	if player.y > maxY {
		player.y = maxY
	}
}

draw(player *Player) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Color{r: 0, g: 0, b: 0, a: 255})
	raylib.DrawFPS(10, 10)
	raylib.DrawRectangle(player.x-player.width/2, player.y-player.height/2, player.width, player.height, Color{r: 0, g: 255, b: 0, a: 255})
	raylib.EndDrawing()
}

Player {
	x int
	y int
	width int
	height int
}