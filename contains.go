package slice

func Contains(sl []interface{}, v interface{}) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return -1, false
}

func ContainsInt(sl []int, v int) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return -1, false
}

func ContainsInt64(sl []int64, v int64) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return -1, false
}

func ContainsString(sl []string, v string) (int, bool) {
	for i, vv := range sl {
		if vv == v {
			return i, true
		}
	}
	return -1, false
}
