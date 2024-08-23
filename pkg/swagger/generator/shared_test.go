package generator

import (
	"log"
	"testing"
)

func TestCamelCaseToSnakeCase(t *testing.T) {
	log.Println(camelCaseToSnakeCase("GroupKindVersion"))
}
