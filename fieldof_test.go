package fieldof

import (
	"reflect"
	"testing"
)

func TestFieldOf(t *testing.T) {
	type testStruct struct {
		name    string
		age     int
		regions []string
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
					name:    "name",
					age:     20,
					regions: nil,
				},
			},
			[]string{"name", "age", "regions"},
			false,
		},
		{
			"struct/ptr",
			args{
				&testStruct{
					name:    "name",
					age:     20,
					regions: nil,
				},
			},
			[]string{"name", "age", "regions"},
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

func TestPublicFieldOf(t *testing.T) {
	type testStruct struct {
		Public  bool
		private bool
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
					Public:  true,
					private: true,
				},
			},
			[]string{"Public"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublicFieldOf(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicFieldOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicFieldOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
