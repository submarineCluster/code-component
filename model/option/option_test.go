package option

import "testing"

func TestCanonicalKey(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{
			"businessCode",
			"business_code",
		},
		{
			"title",
			"title",
		},
		{
			"Title",
			"title",
		},
		{
			"exp_group_key",
			"exp_group_key",
		},
		{
			"expGroupKey",
			"exp_group_key",
		},
	}

	for _, tt := range tests {
		result := canonicalKey(tt.args)
		if result != tt.want {
			t.Fatalf("[tt=%+v]invalid result=%v", tt, result)
		}
	}
}
