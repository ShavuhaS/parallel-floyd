package floyd

import (
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ShavuhaS/parallel-floyd/internal/utils"
)

func TestSequentialSP(t *testing.T) {
	for _, testCase := range floydTestCases {
		name, input, expected := testCase.name, testCase.input, testCase.expectedDist
		t.Run(name, func(t *testing.T) {
			got := SequentialSP(input)
			utils.AssertMatricesEqual(t, got, expected, INF)
		})
	}
}

func TestFloydWarshallWithPath(t *testing.T) {
	// Шлях до стандартної директорії з тестовими даними
	testDataDir := "testdata"

	// Читаємо всі файли в директорії
	files, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Fatalf("Не вдалося прочитати директорію testdata: %v", err)
	}

	for _, file := range files {
		// Шукаємо тільки файли, які закінчуються на "_input.txt"
		if file.IsDir() || !strings.HasSuffix(file.Name(), "_input.txt") {
			continue
		}

		// Витягуємо базове ім'я тесту (наприклад, "graph1_input.txt" -> "graph1")
		baseName := strings.TrimSuffix(file.Name(), "_input.txt")

		// Запускаємо окремий підтест для кожного набору файлів
		t.Run(baseName, func(t *testing.T) {
			// Формуємо шляхи до всіх трьох файлів
			inputPath := filepath.Join(testDataDir, baseName+"_input.txt")
			distPath := filepath.Join(testDataDir, baseName+"_dist.txt")
			prevPath := filepath.Join(testDataDir, baseName+"_prev.txt")

			// 1. Зчитуємо вхідну матрицю
			inputMat, err := utils.InputFromFile(inputPath)
			if err != nil {
				t.Fatalf("Помилка читання input файлу: %v", err)
			}

			// 2. Зчитуємо очікувані матриці результатів
			expectedDist, err := utils.DistFromFile(distPath)
			if err != nil {
				t.Fatalf("Помилка читання dist файлу: %v", err)
			}

			expectedPrev, err := utils.DistFromFile(prevPath)
			if err != nil {
				t.Fatalf("Помилка читання prev файлу: %v", err)
			}

			actualDist, actualPrev := SequentialSPWithPath(inputMat)

			// 4. Порівнюємо результати за допомогою твоєї функції
			// Примітка: я припускаю, що твоя функція приймає *testing.T для виведення помилок.
			// Якщо вона просто повертає bool/error, оброби це відповідним чином.
			utils.AssertMatricesEqual(t, actualDist, expectedDist, INF)
			utils.AssertMatricesEqual(t, actualPrev, expectedPrev, math.MaxInt)
		})
	}
}
