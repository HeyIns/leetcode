func topKFrequent(words []string, k int) []string {
    heap := []string{}
    words_map := map[string]int{}
    for _, n := range(words) {
        words_map[n]++
    }
    
    for key, value := range(words_map) {
        if len(heap) < k {
            //setup
            heap = append(heap, key)
            i := len(heap) - 1
            j := (i + 1) / 2 - 1
            for i > 0 && (words_map[heap[i]] < words_map[heap[j]] || words_map[heap[i]] == words_map[heap[j]] && heap[i] > heap[j]) {
                heap[i], heap[j] = heap[j], heap[i]
                i = j
                j = (i + 1) / 2 - 1
            }
        } else {
            if value > words_map[heap[0]] || value == words_map[heap[0]] && key < heap[0] {
                heap[0] = key
                //adjust
                adjustTop(heap, words_map)
            }
        }
    }
    // fmt.Println(heap)
    res := []string{}
    res_equal := []string{}
    for len(heap) > 0 {
        if len(res_equal) == 0 || words_map[heap[0]] == words_map[res_equal[0]] {
            res_equal = append(res_equal, heap[0])
        } else {
            // sort.Strings(res_equal)
            if len(res_equal) > 1 {
                sort.Strings(res_equal)
                reverse(res_equal)
            }
            res = append(res, res_equal...)
            res_equal = append(res_equal[0:0], heap[0])
        }
        heap[0] = heap[len(heap) - 1]
        heap = heap[:len(heap) - 1]
        //adjust
        adjustTop(heap, words_map)
    }
    if len(res_equal) > 1 {
        sort.Strings(res_equal)
        reverse(res_equal)
    }
    res = append(res, res_equal...)
    reverse(res)
    return res
}

func adjustTop(heap []string, words_map map[string]int) {
    i := 0
    for i < len(heap) {
        child := i
        if 2 * i + 2 < len(heap) {
            if words_map[heap[2 * i + 1]] < words_map[heap[2 * i + 2]] {
                child = 2 * i + 1
            } else if words_map[heap[2 * i + 1]] > words_map[heap[2 * i + 2]] {
                child = 2 * i + 2
            } else {
                if heap[2 * i + 1] < heap[2 * i + 2] {
                    child = 2 * i + 2
                } else {
                    child = 2 * i + 1
                }
            }
        } else if 2 * i + 1 < len(heap) {
            child = 2 * i + 1
        }
        if words_map[heap[i]] > words_map[heap[child]] || words_map[heap[i]] == words_map[heap[child]] && heap[i] < heap[child] {
            heap[i], heap[child] = heap[child], heap[i]
            i = child
        } else {
            break
        }
    }
}
func reverse(s []string) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}


//用golang自带sort工具包，重写less方法
// type node struct {
//     word string
//     count int
// }

// type list []node

// func (l list) Len() int {
//     return len(l)
// }
// func (l list) Swap(i, j int) {
//     l[i], l[j] = l[j], l[i]
// }
// func (l list) Less(i, j int) bool {
//     if l[i].count == l[j].count {
//         return l[i].word < l[j].word
//     }

//     return l[i].count > l[j].count
// }

// func topKFrequent(words []string, k int) []string {
//     dict := map[string]int{}

//     for _, word := range words {
//         dict[word]++
//     }

//     list := make(list, len(dict))
//     for word, count := range dict {
//         list = append(list, node{word, count, })
//     }

//     sort.Sort(list)

//     ret := make([]string, k)

//     for i := 0; i < k; i++ {
//         ret[i] = list[i].word
//     }

//     return ret
// }