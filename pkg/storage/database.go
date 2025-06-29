package storage

import (
	"encoding/json"
	"gymtraker/pkg/core"
	"os"
	"path/filepath"
)

func SaveWorkout(workout core.Workout) error {
	filePath := filepath.Join("Data", workout.Date+".json")
	file, _ := json.MarshalIndent(workout, "", "")
	return os.WriteFile(filePath, file, 0644)
}

//func LoadWorkout(date string) (core.Workout, error) {
// Implementar carga desde archivo
//}
