import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Reads data from the file in filePath and unmarshals it into target.
// target is of T type.
func LoadObject[T any](filePath string) T {
	instance := new(T)
	fileBytes := LoadRawFile(filePath)
	if err := json.Unmarshal(fileBytes, instance); err != nil {
		log.Fatalf("error parsing JSON object from file: %s, err: %s", filePath, err.Error())
	}
	return *instance
}

func LoadRawFile(filePath string) []byte {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("reading test file: %s, error:%v", filePath, err)
	}
	if err != nil {
		log.Fatalf("error loading object from file: %s. err: %s", filePath, err)
	}
	return fileBytes
}
