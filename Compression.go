package hgo

import (
	"bytes"
	"compress/flate"
	"io/ioutil"
)

const DefaultCompression = flate.DefaultCompression

// Can use default compressionLevel: flate.DefaultCompression
func CompressBytes(a []byte, compressionLevel int) (result []byte) {
	var buffer bytes.Buffer
	var writer, writerResult = flate.NewWriter(&buffer, compressionLevel)
	AssertResult(writerResult)
	var _, writeResult = writer.Write(a)
	AssertResult(writeResult)
	var flushResult = writer.Flush()
	AssertResult(flushResult)
	var closeResult = writer.Close()
	AssertResult(closeResult)
	result = buffer.Bytes()
	return
}

func DecompressBytes(a []byte) []byte {
	var reader = bytes.NewReader(a)
	var deflater = flate.NewReader(reader)
	var data, readResult = ioutil.ReadAll(deflater)
	AssertResult(readResult)
	return data
}
