package mapReduce

import (
	"fmt"
	"notebook/common"
	"strings"
	"sync"
	"unicode"
)

// Goroutine池大小
const poolSize = 5

// KeyValue 结构体
type KeyValue struct {
	Key   string
	Value int
}

type MapReduce struct {
}

// Mapper 函数
func (m MapReduce) Mapper(document string) []KeyValue {
	var results []KeyValue
	words := strings.FieldsFunc(document, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	for _, word := range words {
		results = append(results, KeyValue{word, 1})
	}
	return results
}

// Reducer 函数
func (m MapReduce) Reducer(kvPairs []KeyValue) map[string]int {
	result := make(map[string]int)
	for _, kv := range kvPairs {
		result[kv.Key] += kv.Value
	}
	return result
}

// 并行MapReduce优化版本
func (m MapReduce) mapReduceOptimized(documents []string) map[string]int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var intermediate []KeyValue
	kvChannel := make(chan []KeyValue)
	semaphore := make(chan struct{}, poolSize)

	// 启动多个goroutine并行处理
	for _, doc := range documents {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(document string) {
			defer wg.Done()
			kvChannel <- m.Mapper(document)
			<-semaphore
		}(doc)
	}

	// 收集所有中间结果
	go func() {
		wg.Wait()
		close(kvChannel)
	}()

	// 汇总中间结果
	for kvPairs := range kvChannel {
		mu.Lock() // 避免不同的goroutine可能会同时修改intermediate，导致数据不一致或程序崩溃
		intermediate = append(intermediate, kvPairs...)
		mu.Unlock()
	}

	// 调用Reducer函数
	return m.Reducer(intermediate)
}

func (m MapReduce) Handle() {
	documents := []string{
		"hello world",
		"world of golang",
		"hello golang",
	}
	result := m.mapReduceOptimized(documents)
	m.Desc(documents, result)
}

func (m MapReduce) Desc(param any, result any) {
	fmt.Println("【--知识点--】", common.GetPackageName())
	common.PrintDescription("MapReduce 是一种编程模型，用于处理和生成大规模数据集，主要包含两个步骤：\nMap：将输入数据分割成小块，并对每一小块数据进行处理，产生一组中间结果。\nReduce：将中间结果汇总，生成最终结果")
	common.PrintParams(param)
	common.PrintResult(result)
}
