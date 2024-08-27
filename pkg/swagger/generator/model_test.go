package generator

import (
	"log"
	"testing"
)

func TestMakeGenDefinition(t *testing.T) {

	opts := &GenOpts{}
	opts.Spec = "testdata/unit/model/swagger.json"
	doc, err := opts.loadSpec()
	if err != nil {
		panic(err)
	}
	model, err := makeGenDefinition("io.k8s.api.core.v1.Namespace", "models", *doc.Schema(), doc, opts)

	if err != nil {
		panic(err)
	}
	log.Println("model:", model)
}
