package generator

import (
	"log"
	"testing"
)

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
