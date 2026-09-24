Map {
	grid [4][4]int
}

main() {
	map := new(Map)

	loop x := 0..4 {
		loop y := 0..4 {
			map.grid[y][x] = x * y
		}
	}

	assert map.grid[0][0] == 0
	assert map.grid[0][1] == 0
	assert map.grid[0][2] == 0
	assert map.grid[0][3] == 0

	assert map.grid[1][0] == 0
	assert map.grid[1][1] == 1
	assert map.grid[1][2] == 2
	assert map.grid[1][3] == 3

	assert map.grid[2][0] == 0
	assert map.grid[2][1] == 2
	assert map.grid[2][2] == 4
	assert map.grid[2][3] == 6

	assert map.grid[3][0] == 0
	assert map.grid[3][1] == 3
	assert map.grid[3][2] == 6
	assert map.grid[3][3] == 9
}