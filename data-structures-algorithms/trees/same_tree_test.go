package trees

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			name: "both nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "simple equal",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		{
			name: "different structure",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:    &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			name: "one nil",
			p:    nil,
			q:    &TreeNode{Val: 1},
			want: false,
		},
		{
			name: "same inorder different shape",
			// p:   2
			//     /
			//    1     3
			p: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			// q:   1
			//        \
			//         3
			//        /
			//       2
			q: &TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
			// inorder for both: [1,2,3]
			// they are not the same tree (different shapes)
			want: false,
		},
		{
			name: "duplicate values different shape but same inorder",
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 1}},
			q:    &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Fatalf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
