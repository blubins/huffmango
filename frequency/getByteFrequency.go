package frequency

func GetByteFrequency(data *[]byte) map[byte]uint64 {
	fqtbl := make(map[byte]uint64)

	for _, b := range *data {
		fqtbl[b] = fqtbl[b] + 1
	}

	return fqtbl
}
