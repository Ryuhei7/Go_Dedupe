package bloomfilter

import (
	"strconv"

	util "ryuhei/Go_Dedup/cmd/util"
)

var (
	k = 3
)

// BloomFilter is a slise of boolen.
type BloomFilter []bool

// NewBloomFilter はコンストラクタ
func NewBloomFilter(size int) BloomFilter {
	return make([]bool, size)
}

// Add はフィルターに要素を追加する
func (bf BloomFilter) Add(element string) {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		bf[util.DoubleHashing(i64HashA, i64HashB, i, len(bf))] = true
	}
}

// Exists は対象がフィルターに引っかかるかチェックする
func (bf BloomFilter) Exists(element string) bool {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		if !bf[util.DoubleHashing(i64HashA, i64HashB, i, len(bf))] {
			return false
		}
	}
	return true
}
