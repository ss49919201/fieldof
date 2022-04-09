package fieldof

import (
	"reflect"
	"testing"
)

func TestFieldOf(t *testing.T) {
	type testStruct struct {
		name string
		age  int
	}

	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"struct",
			args{
				testStruct{
					name: "name",
					age:  20,
				},
			},
			[]string{"name", "age"},
			false,
		},
		{
			"struct/ptr",
			args{
				&testStruct{
					name: "name",
					age:  20,
				},
			},
			[]string{"name", "age"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldOf(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
