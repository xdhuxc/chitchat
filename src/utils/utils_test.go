package utils

import "testing"

func TestStringInSlice(t *testing.T) {
	type args struct {
		s     string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				s:     "a",
				slice: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "failed",
			args: args{
				s:     "d",
				slice: []string{"a", "b", "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.s, tt.args.slice); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
