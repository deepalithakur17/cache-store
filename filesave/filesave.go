package filesave

import (
	"encoding/json"
	"gorm-gogo/storage"
	"log"
	"os"
)

func SaveDataToFile() {
	bytes, err := json.MarshalIndent(storage.Chache, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal map: %v", err)
	}

	// Write to file
	err = os.WriteFile("dump.json", bytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
