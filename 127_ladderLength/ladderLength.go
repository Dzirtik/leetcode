package ladderLength

// Item stores information about elements
type Item struct {
	elem string
	cnt  int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) || beginWord == endWord {
		return 0
	}

	hash := make(map[string][]string)
	for _, str := range wordList {
		length := len(str)
		for i := 0; i < length; i++ {
			key := str[0:i] + "*" + str[i+1:length]
			hash[key] = append(hash[key], str)
		}
	}

	var items []Item
	item := Item{beginWord, 1}

	for {
		length := len(item.elem)
		for i := 0; i < length; i++ {
			key := item.elem[0:i] + "*" + item.elem[i+1:length]
			if val, ok := hash[key]; ok {
				delete(hash, key)
				for _, next := range val {
					if next != item.elem {
						if next == endWord {
							return item.cnt + 1
						}
						items = append(items, Item{next, item.cnt + 1})
					}
				}
			}
		}
		if len(items) > 0 {
			item, items = items[0], items[1:len(items)]
		} else {
			return 0
		}
	}
}
