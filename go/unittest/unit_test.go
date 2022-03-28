package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
		{name: "t2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_myFunc(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{param: "test"},
			wantErr: false,
		},
		{
			name:    "5678",
			args:    args{param: "5678"},
			wantErr: true,
		},
		{
			name:    "123",
			args:    args{param: "123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := myFunc(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("myFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
