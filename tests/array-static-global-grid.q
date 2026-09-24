global {
	grid [4][3]int
}

main() {
	grid[2][1] = 42

	assert grid[0][0] == 0
	assert grid[0][1] == 0
	assert grid[0][2] == 0

	assert grid[1][0] == 0
	assert grid[1][1] == 0
	assert grid[1][2] == 0

	assert grid[2][0] == 0
	assert grid[2][1] == 42
	assert grid[2][2] == 0

	assert grid[3][0] == 0
	assert grid[3][1] == 0
	assert grid[3][2] == 0
}