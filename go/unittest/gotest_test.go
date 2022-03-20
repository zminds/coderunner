package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		Println("go test run")
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := myFunc(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("myFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
