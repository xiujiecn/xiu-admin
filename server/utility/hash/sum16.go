package hash

import "os"

func Sum16ChecksumData(bs []byte) uint16 {
	sum := uint16(0)
	for _, b := range bs {
		sum += uint16(b)
	}
	return sum
}

func Sum16ChecksumFile(filePath string) uint16 {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()
	// 分块读取
	blockSize := 1024 * 1024
	buf := make([]byte, blockSize)
	sum := uint16(0)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		sum += Sum16ChecksumData(buf[:n])
	}
	return sum
}
