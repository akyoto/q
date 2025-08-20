// This example is currently Windows only and it
// needs the raylib.dll file next to the executable.
// It creates an empty window.
//
// Linux users can use WINE after adding raylib.dll:
//
//     q build examples/raylib --os windows
//     wine examples/raylib/raylib.exe
//
// NOTE:
// - the `float` data type is not implemented yet
// - passing structs by value to extern calls
//   currently only works up to a size of 8 bytes

main() {
	raylib.InitWindow(1280, 720, "raylib example\0".ptr)
	monitor := raylib.GetCurrentMonitor()
	fps := raylib.GetMonitorRefreshRate(monitor)
	raylib.SetTargetFPS(fps)

	loop {
		if raylib.WindowShouldClose() {
			return
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(Color{r: 32, g: 32, b: 32, a: 255})
		raylib.EndDrawing()
	}
}

Color {
	r byte
	g byte
	b byte
	a byte
}

extern {
	raylib {
		BeginDrawing()
		ClearBackground(color Color)
		EndDrawing()
		GetCurrentMonitor() -> int
		GetMonitorRefreshRate(monitor int) -> int
		InitWindow(width int, height int, title *byte)
		SetTargetFPS(fps int)
		WindowShouldClose() -> bool
	}
}