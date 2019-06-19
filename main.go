//初期化と各種呼び出しをここで行う
package main

import (
	"bufio"
	"fmt"
	"os"

	bl "ryuhei/Go_Dedup/cmd/bloomfilter"
	de "ryuhei/Go_Dedup/cmd/dedupe"
	st "ryuhei/Go_Dedup/cmd/store"
)

// ScanData はデータのパスを返します
func ScanData() (datapath string) {
	fmt.Println("保存するデータへのパスを入力してください")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	datapath = stdin.Text()
	return
}

func newApp() string {
	path := ScanData()
	return path
}

func main() {
	path := newApp()

	blfilter := bl.NewBloomFilter(10240000) // bloomfilterの初期化を行う
	var strideCount uint32                  // stride_count はデータのストライド数
	var strideSize uint32 = 4096            // 1024 2048 4096
	var blFlag bool                         // このフラグはBloomFilterの判定に使用
	var duplicater []uint32

	file, openerr := os.Open(path)
	if openerr != nil {
		panic(openerr)
	}
	fileinfo, infoerr := file.Stat()
	if infoerr != nil {
		panic(openerr)
	}
	checker := uint32(fileinfo.Size()) % strideSize
	if checker == 0 {
		strideCount = uint32(fileinfo.Size()) / strideSize
	} else {
		strideCount = uint32(fileinfo.Size())/strideSize + 1
	}

	// StrideHashを初期化してからストライドのハッシュを計算する
	sHashes := de.NewStrideHash(strideCount)
	sHashes.CreateStrideHash(file, strideSize)

	for i := range sHashes.Hashes {
		blFlag = blfilter.Exists(sHashes.Hashes[i])
		if blFlag == false {
			blfilter.Add(sHashes.Hashes[i])
		} else {
			duplicater = append(duplicater, uint32(i))
			fmt.Println("dedupe true")
		}
	}

	defer file.Close()
	file2, openerr2 := os.Open(path)
	if openerr2 != nil {
		panic(openerr)
	}
	st.SaveData(file2, &sHashes, strideSize)
	defer file2.Close()
}
