// This example is currently Windows only and it
// needs the raylib.dll file next to the executable.
// It creates a window showing the number of frames
// per second.
//
// Linux users can use WINE after adding raylib.dll:
//
//     q build examples/raylib --os windows
//     wine examples/raylib/raylib.exe
//
// NOTE:
// - the `float` data type is not implemented yet

main() {
	raylib.InitWindow(1280, 720, "raylib example\0".ptr)

	loop {
		if raylib.WindowShouldClose() {
			return
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(Color{r: 0, g: 0, b: 0, a: 255})
		raylib.DrawFPS(10, 10)
		raylib.EndDrawing()
	}
}

extern {
	raylib {
		BeginDrawing()
		ClearBackground(color Color)
		DrawFPS(x int, y int)
		EndDrawing()
		InitWindow(width int, height int, title *byte)
		WindowShouldClose() -> bool
	}
}