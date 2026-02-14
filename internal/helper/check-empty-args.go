package helper

func CheckIfArrayEmpty(listArray []string) bool {
	if len(listArray) > 0 {
		return false
	}
	return true
}
