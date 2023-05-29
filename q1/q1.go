package q1

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
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
