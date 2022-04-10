package robot_simulation

/**
874. 模拟行走机器人

机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：

-2 ：向左转 90 度
-1 ：向右转 90 度
1 <= x <= 9 ：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。
机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）

输入：commands = [4,-1,3], obstacles = []
输出：25
解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 3 个单位，到达 (3, 4)
距离原点最远的是 (3, 4) ，距离为 32 + 42 = 25

*/

/**
 */
func robotSim(commands []int, obstacles [][]int) int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	di := 0
	x, y := 0, 0

	obstaclesMap := make(map[[2]int]struct{}, len(obstacles))
	for _, v := range obstacles {
		obstaclesMap[[2]int{v[0], v[1]}] = struct{}{}
	}
	max := 0

	for i := 0; i < len(commands); i++ {
		//右转指令，顺时针
		if commands[i] == -1 {
			di = (di + 1) % 4
		} else if commands[i] == -2 {
			//左转指令，逆时针
			di = (di + 3) % 4
		} else {
			for j := 0; j < commands[i]; j++ {
				nextX := x + dx[di]
				nextY := y + dy[di]
				if _, ok := obstaclesMap[[2]int{nextX, nextY}]; ok {
					break
				}

				x, y = nextX, nextY
				tmp := x*x + y*y
				if tmp > max {
					max = tmp
				}
			}
		}
	}

	return max
}
