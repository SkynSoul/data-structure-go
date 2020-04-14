package queue_stack

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	defRooms(rooms, 0, visited)
	return len(visited) == len(rooms)
}

func defRooms(rooms [][]int, roomNum int, visited map[int]bool) {
	if _, ok := visited[roomNum]; ok {
		return
	}
	visited[roomNum] = true
	for i := 0; i < len(rooms[roomNum]); i++ {
		defRooms(rooms, rooms[roomNum][i], visited)
	}
}
