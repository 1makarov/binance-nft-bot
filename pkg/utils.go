package utils

func WW(ids *[]string, id string) bool {
	for _, i := range *ids {
		if i == id {
			return true
		}
	}
	return false
}
