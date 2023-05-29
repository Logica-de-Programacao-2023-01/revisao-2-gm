package q1

import (
	"fmt"
)

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func mergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	// Copia as informações do studentData1 para o mapa mergedData
	for key, student := range studentData1 {
		mergedData[key] = student
	}

	// Atualiza as informações existentes ou adiciona novos alunos do studentData2 para o mapa mergedData
	for key, student := range studentData2 {
		existingStudent, ok := mergedData[key]
		if ok {
			// Atualiza as matérias e notas do aluno existente com as informações mais recentes
			for subject, grade := range student.Subjects {
				existingStudent.Subjects[subject] = grade
			}
		} else {
			// Adiciona um novo aluno ao mapa mergedData
			mergedData[key] = student
		}
	}

	return mergedData
}

func main() {
	studentData1 := map[string]Student{
		"John": {
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Math":   85.0,
				"Science": 78.0,
			},
		},
		"Alice": {
			Name: "Alice",
			Age:  19,
			Subjects: map[string]float64{
				"English": 90.0,
				"Physics": 82.0,
			},
		},
	}

	studentData2 := map[string]Student{
		"John": {
			Name: "John",
			Age:  20,
			Subjects: map[string]float64{
				"Math":   90.0,
				"History": 88.0,
			},
		},
		"Bob": {
			Name: "Bob",
			Age:  21,
			Subjects: map[string]float64{
				"Chemistry": 79.0,
				"Physics":   85.0,
			},
		},
	}

	mergedData := mergeStudentData(studentData1, studentData2)

	// Exibe as informações combinadas dos alunos
	for _, student := range mergedData {
		fmt.Printf("Name: %s, Age: %d\n", student.Name, student.Age)
		fmt.Println("Subjects:")
		for subject, grade := range student.Subjects {
			fmt.Printf("- %s: %.2f\n", subject, grade)
		}
		fmt.Println()
	}
}
