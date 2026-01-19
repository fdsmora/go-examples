package backtracking

//import "fmt"

func NumOfPathsToDest(n int) int {
	if n == 1 {
		return 1
	}
	var res []string
	res = visit(0, 0, n, "", res)
	return len(res)
}

func visit(i, j, n int, path string, res []string) []string {
	if i >= n || j >= n {
		return res
	}
	if i == n-1 && j == n-1 {
		return append(res, path)
	}
	if i+1 < n {
		res = visit(i+1, j, n, path+"E", res)
	}
	if j+1 < n && j+1 <= i {
		res = visit(i, j+1, n, path+"N", res)
	}
	return res
}
