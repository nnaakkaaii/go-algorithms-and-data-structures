package dictionary

const M = 1046527

type Dict struct {
	arr [M]string
}

func NewDict () *Dict {
	return new(Dict)
}

// getKey get key from keyString
func (dict *Dict) getKey(keyString string) int {
	chs := []rune(keyString)
	sum := 0
	p := 1

	charToInt := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}

	for _, ch := range chs {
		sum += p * charToInt[ch]
		p *= 5
	}
	return sum
}

func (dict *Dict) createHash(key int, i int) int {
	h1 := func (key int) int {return key % M}
	h2 := func (key int) int {return 1 + key % (M - 1)}
	return (h1(key) + i * h2(key)) % M
}

func (dict *Dict) Find(keyString string) bool {
	key := dict.getKey(keyString)
	for i := 0; ; i++ {
		h := dict.createHash(key, i)
		if dict.arr[h] == keyString {
			return true
		} else if len(dict.arr[h]) == 0 {
			return false
		}
	}
}

func (dict *Dict) Insert(keyString string) bool {
	key := dict.getKey(keyString)
	for i := 0; ;i++ {
		h := dict.createHash(key, i)
		if dict.arr[h] == keyString {
			return true
		} else if len(dict.arr[h]) == 0 {
			dict.arr[h] = keyString
			return false
		}
	}
}
