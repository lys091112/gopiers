package mystrings

import (
	"testing"
)

func BenchmarkSumString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddString()
	}
	b.StopTimer()
}

func BenchmarkSprintfString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SprintfString()
	}
	b.StopTimer()
}

func BenchmarkBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuilderString()
	}
	b.StopTimer()
}

func BenchmarkBytesBuffString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BufferString()
	}
	b.StopTimer()
}

func BenchmarkJoinstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinString()
	}
	b.StopTimer()
}

func BenchmarkByteSliceString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ByteSliceString()
	}
	b.StopTimer()
}

func TestMakeTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "t01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeTest()
		})
	}
}
