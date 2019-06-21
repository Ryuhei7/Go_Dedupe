//ここでは読み込んだファイルを4kごとのバイナリに分解して、それをスライスで返す
//データのパスを引数としてもらってそれを分解する
//4kで分解してそれをスライスにしてポインタを返す

package dedupe

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	ts "ryuhei/Go_Dedup/cmd/testtool"
	"ryuhei/Go_Dedup/cmd/util"
)

// StrideHash の中にあるスライスに対象データのストライドの
// hash値を入れる
type StrideHash struct {
	Hashes      []string
	StrideCount uint32
}

// NewStrideHash はコンストラクタ
func NewStrideHash(strideCount uint32) StrideHash {

	var sh StrideHash
	sh.Hashes = make([]string, strideCount)
	sh.StrideCount = strideCount

	return sh
}

// CreateStrideHash はバイナリをストライドごとにハッシュ化する
func (sh *StrideHash) CreateStrideHash(file *os.File, strideSize uint32) {

	var MD5hash string
	buf := make([]byte, int(strideSize))
	var i uint32

	me := ts.NewMeasure()
	me.StartAll()
	for i = 0; i < sh.StrideCount; i++ {
		readSize, readerr := io.ReadFull(file, buf)

		if readerr != nil {
			//panic(readerr)
		}
		//fmt.Println(hex.EncodeToString(buf))
		binary := hex.EncodeToString(buf[:readSize])
		MD5hash = util.CalcMD5Hash(binary)
		sh.Hashes[i] = MD5hash
	}
	me.EndAll()
	me.CalcAll()
	fmt.Println("MD5 Hashing")
	fmt.Println("Memory is ", me.Mem)
	fmt.Println("Cpu Time is ", me.Cputime)
	fmt.Println("Process Time is ", me.Time)
	fmt.Println(" ")

}
