package main

import "fmt"
import "sort"

func wordBreak(s string, wordDict []string) []string {
	sort.Strings(wordDict)
	varience := make(map[string]string)
	varience[""] = ""
	for _, elem := range s {
		char := string(elem)
		var appendix []string
		for key, value := range varience {
			varience[key] = value + char
			i := sort.SearchStrings(wordDict, varience[key])
			if i < len(wordDict) {
				if wordDict[i] == varience[key] {
					sep := ""
					if len(key) > 0 {
						sep = " "
					}
					appendix = append(appendix, key+sep+varience[key])
				} else if len(wordDict[i]) > len(varience[key]) {
					if wordDict[i][0:len(varience[key])] != varience[key] {
						delete(varience, key)
						if len(varience) == 0 {
							return []string{}
						}
					}
				}
			}

		}
		for _, value := range appendix {
			varience[value] = ""
		}
	}
	var result []string
	for key, value := range varience {
		if value == "" {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
