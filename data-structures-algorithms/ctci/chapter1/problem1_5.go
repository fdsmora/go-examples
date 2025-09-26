package chapter1

func CheckEdits(a, b string) bool {
	// Quick check: if length difference > 1, impossible
	if abs(len(a)-len(b)) > 1 {
		return false
	}

	// Identify longer and shorter
	var longer, shorter string
	if len(a) >= len(b) {
		longer, shorter = a, b
	} else {
		longer, shorter = b, a
	}

	diff := false
	i, j := 0, 0
	for i < len(longer) && j < len(shorter) {
		if longer[i] != shorter[j] {
			if diff {
				return false
			}
			diff = true
			// If same length, move both (replacement case)
			if len(longer) == len(shorter) {
				j++
			}
		} else {
			j++
		}
		i++
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
