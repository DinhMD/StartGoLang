package ArrayUtils

func IsContains(arr interface{}, val interface{}) bool {
	switch v := arr.(type) {
	case []int:
		for _, item := range v {
			if item == val.(int) {
				return true
			}
		}
	case []string:
		for _, item := range v {
			if item == val.(string) {
				return true
			}
		}
		// Add cases for other types as needed
	}
	return false
}
