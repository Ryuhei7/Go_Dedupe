package store

import (
	"fmt"
	"io"
	"os"

	de "ryuhei/Go_Dedup/cmd/dedupe"
	ts "ryuhei/Go_Dedup/cmd/testtool"
)

// SaveData はデータをbucketの中に保存する関数
// データを4kづつに読み込んでMD5のハッシュ値をファイル名にする
// 基本的には全てバケット１に保存する
func SaveData(file *os.File, sh *de.StrideHash, strideSize uint32) {

	buf := make([]byte, int(strideSize))
	me := ts.NewMeasure()

	me.StartAll()
	for i := range sh.Hashes {
		_, readerr := io.ReadFull(file, buf)

		if readerr != nil {
			// panic(readerr)
		}
		// binary := hex.EncodeToString(buf[:readSize])
		namehash := sh.Hashes[i]

		filename := "bucket/1/" + namehash
		func() {
			newfile, err := os.Create(filename)
			if err != nil {
				// Openエラー処理
				panic(err)
			}
			defer newfile.Close()

			//fmt.Println("bufsize = ", hex.EncodeToString(buf[:readSize]))
			newfile.Write(buf)
		}()
	}
	me.EndAll()
	me.CalcAll()
	fmt.Println("Store")
	fmt.Println("Memory is ", me.Mem)
	fmt.Println("Cpu Time is ", me.Cputime)
	fmt.Println("Process Time is ", me.Time)
	fmt.Println(" ")
}
