package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	result := []string{}
	var hs func(int, string)

	hs = func(i int, path string) {
		if i >= len(digits) {
			result = append(result, path)
			return
		}
		for _, v := range mp[string(digits[i])] {
			hs(i+1, path+string(v))
		}
	}

	hs(0, "")
	return result
}
