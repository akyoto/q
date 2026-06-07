const {
	KEY_RIGHT = 262
	KEY_LEFT = 263
	KEY_DOWN = 264
	KEY_UP = 265
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
	r byte
	g byte
	b byte
	a byte
}