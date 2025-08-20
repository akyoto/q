// This example is currently Windows only and it
// needs the raylib.dll file next to the executable.
// It creates an empty window.
//
// NOTE:
// - the `float` data type is not implemented yet
// - passing structs by value to extern calls is not implemented yet

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
		raylib.EndDrawing()
	}
}

extern {
	raylib {
		BeginDrawing()
		EndDrawing()
		GetCurrentMonitor() -> int
		GetMonitorRefreshRate(monitor int) -> int
		InitWindow(width int, height int, title *byte)
		SetTargetFPS(fps int)
		WindowShouldClose() -> bool
	}
}