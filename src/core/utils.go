package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

/*
IntToHex 将一个64位int整数转换为byte字节数组
 */
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

/*
DataToHash 对byte字节数组类型的数据进行SHA256哈希计算
 */
func DataToHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}