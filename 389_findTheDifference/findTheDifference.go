package findTheDifference

func findTheDifference(s string, t string) byte {
	var i byte
	for _, v := range []byte(s + t) {
		i = i ^ v
	}
	return i
}
