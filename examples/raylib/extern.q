Key const {
	Right = 262
	Left = 263
	Down = 264
	Up = 265
}

extern {
	raylib {
		BeginDrawing()
		ClearBackground(color Color)
		CloseWindow()
		DrawFPS(x int, y int)
		DrawRectangle(x int, y int, width int, height int, color Color)
		EndDrawing()
		InitWindow(width int, height int, title *byte)
		IsKeyDown(key int) -> bool
		SetTargetFPS(fps int)
		WindowShouldClose() -> bool
	}
}

Color {
	r uint8
	g uint8
	b uint8
	a uint8
}