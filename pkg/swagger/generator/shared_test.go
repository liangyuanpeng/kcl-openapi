package generator

import (
	"log"
	"strings"
	"testing"
)

func TestH(t *testing.T) {
	log.Println(strings.Replace("_a_b", "_", "", 1))
}

func TestCamelCaseToSnakeCase(t *testing.T) {
	cases := []struct {
		value  string
		expect string
	}{
		{
			value:  "k8s.api.admissionregistration.v1.ValidatingAdmissionPolicyList",
			expect: "k8s.api.admissionregistration.v1.validating_admission_policy_list",
		},
		{
			value:  "ValidatingAdmissionPolicyList",
			expect: "validating_admission_policy_list",
		},
		{
			value:  "k8s.api.core.v1.PersistentVolumeClaim",
			expect: "k8s.api.core.v1.persistent_volume_claim",
		},
		{
			value:  "PersistentVolumeClaim",
			expect: "persistent_volume_claim",
		},
	}
	for _, testcase := range cases {
		t.Run(testcase.value, func(t *testing.T) {
			got := camelCaseToSnakeCase(testcase.value)
			if got != testcase.expect {
				t.Fatalf("unexpected output, expect:\n%s\ngot:%s\n", testcase.expect, got)
			}
		})
	}

	log.Println(camelCaseToSnakeCase("GroupKindVersion"))
}
