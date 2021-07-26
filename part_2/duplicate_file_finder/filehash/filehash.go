package filehash

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
	"sync"
)

type FileHash struct {
	Path string
	Size int64
	Hash [32]byte
}

func HashFiles(fileHashingWG *sync.WaitGroup, fileHashingChannel chan string, updateDuplicatesChannel chan FileHash) {
	defer fileHashingWG.Done()
	for path := range fileHashingChannel {
		fileHash, err := generateFileHash(path)
		if err != nil {
			log.Println("Error hashing file:", err)
			continue
		}
		updateDuplicatesChannel <- fileHash
	}
}
func generateFileHash(path string) (FileHash, error) {
	file, err := os.Open(path)
	if err != nil {
		return FileHash{}, err
	}
	defer file.Close()

	hasher := sha256.New()
	bytesWritten, err := io.Copy(hasher, file)
	if err != nil {
		return FileHash{}, err
	}

	var hash [32]byte
	copy(hash[:], hasher.Sum(nil)[:32])
	return FileHash{Path: path, Size: bytesWritten, Hash: hash}, nil
}
