package n_queue

/**
51. N 皇后

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var solution [][]string

func solveNQueens(n int) [][]string {
	solution = [][]string{}
	queues := make([]int, n)
	for i := 0; i < len(queues); i++ {
		queues[i] = -1
	}

	column := make(map[int]bool, n)
	dialogs1 := make(map[int]bool, n)
	dialogs2 := make(map[int]bool, n)
	backtrack(queues, 0, n, column, dialogs1, dialogs2)
	return solution
}

func backtrack(queues []int, row int, n int, column map[int]bool, dialogs1, dialogs2 map[int]bool) {
	if row == n {
		board := generateBoard(queues, n)
		solution = append(solution, board)
		return
	}

	for j := 0; j < n; j++ {
		if column[j] {
			continue
		}

		dialog1 := row - j
		if dialogs1[dialog1] {
			continue
		}

		dialog2 := row + j
		if dialogs2[dialog2] {
			continue
		}

		column[j], dialogs1[dialog1], dialogs2[dialog2] = true, true, true
		queues[row] = j
		backtrack(queues, row+1, n, column, dialogs1, dialogs2)
		column[j], dialogs1[dialog1], dialogs2[dialog2] = false, false, false
		queues[row] = -1
	}
}

func generateBoard(queues []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		t := make([]byte, n)
		for j := 0; j < n; j++ {
			t[j] = '.'
		}
		t[queues[i]] = 'Q'
		board = append(board, string(t))
	}

	return board
}
