package strings

import "testing"

// TestSimplifyPath provides comprehensive test cases for the simplify path problem
func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		// Basic cases
		{
			name: "root only",
			path: "/",
			want: "/",
		},
		{
			name: "simple path",
			path: "/home",
			want: "/home",
		},
		{
			name: "multiple directories",
			path: "/home/user/documents",
			want: "/home/user/documents",
		},

		// Multiple slashes
		{
			name: "double slash",
			path: "//home",
			want: "/home",
		},
		{
			name: "multiple consecutive slashes",
			path: "////home",
			want: "/home",
		},
		{
			name: "slashes in middle",
			path: "/home///user",
			want: "/home/user",
		},
		{
			name: "trailing slash",
			path: "/home/",
			want: "/home",
		},
		{
			name: "multiple trailing slashes",
			path: "/home///",
			want: "/home",
		},

		// Single dot (current directory)
		{
			name: "single dot at end",
			path: "/home/.",
			want: "/home",
		},
		{
			name: "single dot in middle",
			path: "/home/./user",
			want: "/home/user",
		},
		{
			name: "multiple single dots",
			path: "/./home/./user/./",
			want: "/home/user",
		},
		{
			name: "only single dot",
			path: "/.",
			want: "/",
		},
		{
			name: "multiple single dots only",
			path: "/./././.",
			want: "/",
		},

		// Double dot (parent directory)
		{
			name: "double dot at end",
			path: "/home/user/..",
			want: "/home",
		},
		{
			name: "double dot in middle",
			path: "/home/../user",
			want: "/user",
		},
		{
			name: "multiple double dots",
			path: "/home/user/../../root",
			want: "/root",
		},
		{
			name: "double dot beyond root",
			path: "/../",
			want: "/",
		},
		{
			name: "multiple double dots beyond root",
			path: "/../../..",
			want: "/",
		},
		{
			name: "complex with double dots",
			path: "/a/./b/../../c/",
			want: "/c",
		},
		{
			name: "all cancelled by double dots",
			path: "/home/user/../..",
			want: "/",
		},

		// Triple or more dots (should be treated as directory names)
		{
			name: "triple dot",
			path: "/home/.../user",
			want: "/home/.../user",
		},
		{
			name: "four dots",
			path: "/home/..../user",
			want: "/home/..../user",
		},
		{
			name: "many dots",
			path: "/home/......./user",
			want: "/home/......./user",
		},

		// Mixed single and double dots
		{
			name: "single then double dot",
			path: "/home/./user/..",
			want: "/home",
		},
		{
			name: "double then single dot",
			path: "/home/../user/.",
			want: "/user",
		},
		{
			name: "complex mixed dots",
			path: "/a/b/./c/../d/..",
			want: "/a/b",
		},

		// Directory names with letters and underscores
		{
			name: "underscore in name",
			path: "/home/user_name/docs",
			want: "/home/user_name/docs",
		},
		{
			name: "multiple underscores",
			path: "/home/__test__/docs",
			want: "/home/__test__/docs",
		},
		{
			name: "mixed letters and underscores",
			path: "/home_dir/user_name_123",
			want: "/home_dir/user_name_123",
		},

		// Edge cases with dots in directory names
		{
			name: "dots as part of name (triple)",
			path: "/home/.git/objects",
			want: "/home/.git/objects",
		},
		{
			name: "name starting with dots",
			path: "/home/...hidden/file",
			want: "/home/...hidden/file",
		},

		// Complex real-world scenarios
		{
			name: "typical navigation up",
			path: "/home/user/documents/../downloads",
			want: "/home/user/downloads",
		},
		{
			name: "multiple navigation ups",
			path: "/a/b/c/d/../../e/f",
			want: "/a/b/e/f",
		},
		{
			name: "all operations combined",
			path: "/a/./b///../c/../././../d",
			want: "/d",
		},
		{
			name: "complex with trailing slash",
			path: "/a/b/c/../.././d/",
			want: "/a/d",
		},

		// Stress test cases
		{
			name: "very long path",
			path: "/home/user/documents/work/project/src/main/java/com/example/app",
			want: "/home/user/documents/work/project/src/main/java/com/example/app",
		},
		{
			name: "many double dots at start",
			path: "/../../../../../a/b/c",
			want: "/a/b/c",
		},
		{
			name: "alternating dots and dirs",
			path: "/a/../b/../c/../d",
			want: "/d",
		},
		{
			name: "only slashes and dots",
			path: "///.///.///./",
			want: "/",
		},

		// Leetcode examples
		{
			name: "leetcode example 1",
			path: "/home/",
			want: "/home",
		},
		{
			name: "leetcode example 2",
			path: "/../",
			want: "/",
		},
		{
			name: "leetcode example 3",
			path: "/home//foo/",
			want: "/home/foo",
		},
		{
			name: "leetcode example 4",
			path: "/a/./b/../../c/",
			want: "/c",
		},

		// Tricky cases
		{
			name: "single letter directories",
			path: "/a/b/c/d/e/f",
			want: "/a/b/c/d/e/f",
		},
		{
			name: "back to root and forward",
			path: "/a/../../b",
			want: "/b",
		},
		{
			name: "dots and slashes mixed",
			path: "/./././a/./b/././c/./.",
			want: "/a/b/c",
		},
		{
			name: "complex cancellation",
			path: "/a/b/c/d/../../../e",
			want: "/a/e",
		},

		// Empty or minimal
		{
			name: "root with single dot",
			path: "/.",
			want: "/",
		},
		{
			name: "root with double dot",
			path: "/..",
			want: "/",
		},

		// Real-world unix paths
		{
			name: "home to root and back",
			path: "/home/user/../..",
			want: "/",
		},
		{
			name: "relative navigation",
			path: "/var/log/../../etc/hosts",
			want: "/etc/hosts",
		},
		{
			name: "hidden files",
			path: "/home/user/.config/..",
			want: "/home/user",
		},

		// Additional edge cases
		{
			name: "underscore only directory",
			path: "/home/_/docs",
			want: "/home/_/docs",
		},
		{
			name: "letter after many slashes",
			path: "//////a",
			want: "/a",
		},
		{
			name: "mixed everything",
			path: "/a//b/./c/../d/../../e/f/..",
			want: "/a/e", // Corrected: was "/e", should be "/a/e"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath(tt.path)
			if got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

// TestSimplifyPathEdgeCases focuses on boundary conditions
func TestSimplifyPathEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "empty after simplification",
			path: "/..",
			want: "/",
		},
		{
			name: "max double dots",
			path: "/a/b/c/d/e/../../../../..",
			want: "/",
		},
		{
			name: "single char dirs with dots",
			path: "/a/./b/../c/./d/..",
			want: "/a/c",
		},
		{
			name: "only special chars",
			path: "///..//./",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := simplifyPath(tt.path)
			if got != tt.want {
				t.Errorf("simplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}

// BenchmarkSimplifyPath benchmarks the solution
func BenchmarkSimplifyPath(b *testing.B) {
	testCases := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/a/./b/../../c/",
		"/a/b/c/d/e/../../../../..",
		"///.///.///./",
		"/home/user/documents/work/project/src/main/java/com/example/app",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				simplifyPath(tc)
			}
		})
	}
}
