package backtracking

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	initialColor := image[sr][sc]

	fill(&image, sr, sc, len(image), len(image[0]), color, initialColor)

	return image
}

func fill(image *[][]int, i int, j int, rc int, cc int, color int, initialColor int) {
	if i < 0 || j < 0 || i > rc-1 || j > cc-1 || (*image)[i][j] != initialColor {
		return
	}
	(*image)[i][j] = color
	fill(image, i+1, j, rc, cc, color, initialColor)
	fill(image, i-1, j, rc, cc, color, initialColor)
	fill(image, i, j+1, rc, cc, color, initialColor)
	fill(image, i, j-1, rc, cc, color, initialColor)
}
