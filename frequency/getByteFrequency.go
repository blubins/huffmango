package frequency

// returns a map of byte : frequency from an array of bytes
func GetByteFrequency(data *[]byte) map[byte]uint64 {
	fqtbl := make(map[byte]uint64)

	for _, b := range *data {
		fqtbl[b] = fqtbl[b] + 1
	}

	return fqtbl
}
