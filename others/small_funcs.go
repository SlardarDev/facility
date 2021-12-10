package others

func Extend(dst, src map[string]interface{}) {

	for k, v := range src {
		if _, ok := dst[k]; !ok {
			dst[k] = v
		}
	}
}

func ExtendAndOverwrite(dst, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}
