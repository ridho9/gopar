package gopar

import "testing"

func TestString(t *testing.T) {
	type args struct {
		pattern string
		input   string
	}
	tests := []struct {
		name          string
		args          args
		wantNextInput string
		wantResult    any
		wantErr       bool
	}{
		{
			name:    "err if input empty",
			args:    args{input: ""},
			wantErr: true,
		},
		{
			name:          "err if pattern longer",
			args:          args{input: "abc", pattern: "abcde"},
			wantErr:       true,
			wantNextInput: "abc",
		},
		{
			name:          "err no match",
			args:          args{input: "abc123", pattern: "xyz"},
			wantNextInput: "abc123",
			wantErr:       true,
		},
		{
			name:          "success",
			args:          args{input: "abc123", pattern: "abc"},
			wantNextInput: "123",
			wantResult:    "abc",
		},
		{
			name:          "success utf8",
			args:          args{input: "日本語123", pattern: "日本語"},
			wantNextInput: "123",
			wantResult:    "日本語",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := String(tt.args.pattern)
			res := parser.Run(tt.args.input)
			// gotNextInput, gotResult, err := parser.Run(tt.args.input)
			gotNextInput := res.input.peekString()
			gotResult := res.result
			err := res.err

			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNextInput != tt.wantNextInput {
				t.Errorf("String() gotNextInput = %v, want %v", gotNextInput, tt.wantNextInput)
			}
			if gotResult != tt.wantResult {
				t.Errorf("String() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTakeWhile(t *testing.T) {
	pred := func(r rune) bool {
		return ('a' <= r) && (r <= 'z')
	}
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		wantNextInput string
		wantResult    any
		wantErr       bool
	}{
		{
			name:          "success > 1",
			args:          args{input: "abc123"},
			wantNextInput: "123",
			wantResult:    "abc",
		},
		{
			name:          "success empty",
			args:          args{input: "123"},
			wantNextInput: "123",
			wantResult:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := TakeWhile(pred)
			res := parser.Run(tt.args.input)
			// gotNextInput, gotResult, err := parser.Run(tt.args.input)
			gotNextInput := res.input.peekString()
			gotResult := res.result
			err := res.err

			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNextInput != tt.wantNextInput {
				t.Errorf("String() gotNextInput = %v, want %v", gotNextInput, tt.wantNextInput)
			}
			if gotResult != tt.wantResult {
				t.Errorf("String() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
