package duplicateobj

import (
	"fmt"

	"gopkg.in/oleiade/reflections.v1"
)

func getObjFieldString(fieldsToExtract []string, object interface{}) string {
	output := ""
	for _, fieldName := range fieldsToExtract {
		value, _ := reflections.GetField(object, fieldName)
		output = output + ">>" + fmt.Sprintf("%v", value)
	}
	return output
}

// FindDuplicateObj - reads each object and detects objects with the same fields
func FindDuplicateObj(fieldsToExtract []string, objects []interface{}) []int {
	unique := make(map[string]bool)
	repeatIndexes := []int{}
	for index, object := range objects {
		name := getObjFieldString(fieldsToExtract, object)
		_, alreadyExists := unique[name]
		if alreadyExists {
			repeatIndexes = append(repeatIndexes, index)
		} else {
			unique[name] = true
		}
	}
	return repeatIndexes
}
