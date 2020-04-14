package queue_stack

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    visited := make(map[int]bool)
    imageRender(image, sr, sc, image[sr][sc], newColor, visited)
    return image
}

func imageRender(image [][]int, sr int, sc int, srcColor int, newColor int, visited map[int]bool)  {
    if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) {
        return
    }
    _, ok := visited[getMapKey(sr, sc)]
    if ok {
        return
    }
    if image[sr][sc] != srcColor {
        return
    }
    image[sr][sc] = newColor
    visited[getMapKey(sr, sc)] = true
    
    imageRender(image, sr - 1, sc, srcColor, newColor, visited)
    imageRender(image, sr + 1, sc, srcColor, newColor, visited)
    imageRender(image, sr, sc - 1, srcColor, newColor, visited)
    imageRender(image, sr, sc + 1, srcColor, newColor, visited)
}

func getMapKey(sr int, sc int) int {
    return sr * 1000 + sc
}
