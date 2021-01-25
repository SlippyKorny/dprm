package main

// isInMap checks whether a value is in a map[string][]string under some key
func isInMap(m map[string][]string, v string) bool {
	for k, arr := range m {
		if k == v {
			return true
		}

		for i := 0; i < len(arr); i++ {
			if arr[i] == v {
				return true
			}
		}
	}

	return false
}
