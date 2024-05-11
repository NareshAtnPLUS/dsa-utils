package maps

type VisitedMap[K comparable] map[K]bool

func CopyMap[K comparable, V any](src, dst *map[K]V) {
	for key, value := range *src {
		(*dst)[key] = value
	}
}

func IntersectMap[K comparable](map1, map2 map[K]bool) []K {
	var intersection []K

	// Loop over the first map and check if the node is also `true` in the second map
	for key, value1 := range map1 {
		if value1 { // Check if the value is true in the first map
			if value2, found := map2[key]; found && value2 {
				intersection = append(intersection, key)
				// Optionally, remove the item from the second map to prevent duplicates
				// delete(map2, key)
			}
		}
	}

	return intersection
}
