package winnow

func generateFingerprint(hashList []int) map[int]int {
	fingerprint := make(map[int]int)

	for i := range hashList {
		if i+WindowSize > len(hashList) {
			break
		}

		tmpList := hashList[i : i+WindowSize]
		minHash := tmpList[WindowSize-1]
		minPos := WindowSize + i - 1

		for j := 0; j < WindowSize; j++ {
			if tmpList[j] < minHash {
				minHash = tmpList[j]
				minPos = i + j
			}
		}

		if _, ok := fingerprint[minPos]; !ok {
			fingerprint[minPos] = minHash
		}
	}

	return fingerprint
}
