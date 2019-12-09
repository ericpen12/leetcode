**第一次提交**

执行用时 :276 ms, 在所有 golang 提交中击败了72.10%的用户

内存消耗 :273.2 MB, 在所有 golang 提交中击败了5.55%的用户

```go
func groupAnagrams(strs []string) [][]string {
    maps := make(map[interface{}][]string)
	values := [][]string{}


	for _, str := range strs {
		arr := make([]int, 26)

		for _,  v := range str {
			arr[v - 'a']++
		}
		key := ""

		for _, val := range arr {
			key += "#"+ strconv.Itoa(val)
		}
        if key != "" {
            maps[key] = append(maps[key], str)
        }
	}

	for _, value := range maps {
        if len(value) != 0 {
            values = append(values, value)
        }
	}

	return values
}
```

