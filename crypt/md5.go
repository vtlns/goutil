package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// MD5 生成md5摘要
func MD5(s string) string {
	data := []byte(s)
	m := md5.New()
	_, _ = m.Write(data)

	return hex.EncodeToString(m.Sum(nil))
}

// Md5Stream 流式处理
func Md5Stream(reader io.Reader) (string, error) {
	m := md5.New()
	_, err := io.Copy(m, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(m.Sum(nil)), nil
}

func Md5Sum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	return Md5Stream(file)
}
