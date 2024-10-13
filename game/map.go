package game

var GRASSMAP = [][]int{
	{00, 01, 01, 02, 10, 00, 01, 02},
	{03, 04, 04, 04, 01, 07, 04, 05},
	{03, 04, 04, 04, 05, 10, 03, 05},
	{03, 04, 04, 04, 05, 10, 03, 05},
	{03, 04, 04, 04, 04, 01, 04, 05},
	{03, 04, 04, 04, 04, 04, 04, 05},
	{03, 04, 04, 04, 04, 04, 04, 05},
	{03, 04, 07, 04, 04, 04, 04, 05},
	{03, 05, 10, 03, 04, 04, 04, 05},
	{03, 05, 10, 03, 04, 04, 04, 05},
	{03, 05, 10, 03, 04, 04, 04, 05},
	{03, 05, 10, 03, 04, 04, 04, 05},
	{06, 07, 01, 04, 04, 07, 04, 05},
	{10, 10, 06, 07, 8, 10, 06, 8},
}

func getFoamPos() [][]int {
	limy := len(GRASSMAP[0])
	limx := len(GRASSMAP)

	slc := make([][]int, limx)
	for i := range slc {
		s := make([]int, limy)
		slc[i] = s
	}

	for x := 0; x < limx; x++ {
		for y := 0; y < limy; y++ {
			if GRASSMAP[x][y] == 10 {
				if x-1 > 0 && GRASSMAP[x-1][y] != 10 {
					slc[x-1][y] = 1
				}
				if x+1 < limx && GRASSMAP[x+1][y] != 10 {
					slc[x+1][y] = 1
				}
				if y-1 > 0 && GRASSMAP[x][y-1] != 10 {
					slc[x][y-1] = 1
				}
				if y+1 < limy && GRASSMAP[x][y+1] != 10 {
					slc[x][y+1] = 1
				}
			}
		}
	}
	for i, row := range slc {
		if GRASSMAP[i][0] != 10 {
			row[0] = 1
		}
		if GRASSMAP[i][limy-1] != 10 {
			row[limy-1] = 1
		}
	}
	

	return slc
}
