package npncore

// Returns a bool indicating if a provided value is present in the provided string array
func StringArrayContains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}
