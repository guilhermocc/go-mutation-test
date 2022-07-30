package validator

import "testing"

func Test_validateIdentifier(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid identifier medium",
			args{id: "abc12"},
			true,
		},
		{"valid identifier small size",
			args{id: "a"},
			true,
		},
		{"valid identifier large size",
			args{id: "abc123"},
			true,
		},
		{"invalid small identifier",
			args{id: ""},
			false,
		},
		{"invalid large identifier",
			args{id: "abcde12"},
			false,
		},
		{"invalid identifier starting with a number",
			args{id: "2bcd2"},
			false,
		},
		{"invalid identifier starting with a special character",
			args{id: "&abc1"},
			false,
		},
		{"invalid identifier with correct size but composed by invalid characters",
			args{id: "abc1&"},
			false,
		},
		{"invalid identifier with nil id",
			args{id: "nil"},
			false,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			id := &testCase.args.id
			if *id == "nil" {
				id = nil
			}
			if got := validateIdentifier(id); got != testCase.want {
				t.Errorf("validateIdentifier() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func FuzzValidateIdentifier(f *testing.F) {
	f.Add("abc12")
	f.Fuzz(func(t *testing.T, id string) {
		print(id)
		result := validateIdentifier(&id)
		if result == false {
			//t.Errorf("validateIdentifier(%q) = %v, want %v", id, result, true)
		}
	})

}
