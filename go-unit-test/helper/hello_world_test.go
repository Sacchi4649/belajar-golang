package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{{
		name:    "BenchAdi",
		request: "Adi",
	},
		{
			name:    "BenchJoko",
			request: "Joko",
		}}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// go test -v -run=  -bench=. atau nama
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; 1 < b.N; i++ {
		HelloWorld("Sany")
	}
}

// go test -v -run=TestTidakAda -bench=BenchmarkSub/Adika
func BenchmarkSub(b *testing.B) {
	b.Run("Sany", func(b *testing.B) {
		for i := 0; 1 < b.N; i++ {
			HelloWorld("Sany")
		}
	})
	b.Run("Adika", func(b *testing.B) {
		for i := 0; 1 < b.N; i++ {
			HelloWorld("Adika")
		}
	})
}
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(adi)",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "HelloWorld(joko)",
			request:  "joko",
			expected: "Hello Joko",
		},
		{
			name:     "HelloWorld(sany)",
			request:  "Sany",
			expected: "Hello Sany",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Sany")
	assert.Equal(t, "Hello Sany", result, "Result must be 'Hello Sany'")
	// fmt.Println("test hello world assert passed")
}

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "darwin" {
		t.Skip("Cant run on Windows OS")
	}

	result := HelloWorld("Sany")
	assert.Equal(t, "Hello Sany", result, "Result must be 'Hello Sany'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adika")
	require.Equal(t, "Hello Adika", result, "Result must be 'Hello Adika'")
	// fmt.Println("test hello world require passed")
}

func TestHelloWorldSany(t *testing.T) {
	result := HelloWorld("Sany")

	if result != "Hello Sany" {
		t.Fatal("Result must be 'Hello Sany';")
	}

	t.Log("test hello sany passed")
}

func TestHelloWorldAdika(t *testing.T) {
	result := HelloWorld("Adika")

	if result != "Hello Adika" {
		t.Fatal("Result must be 'Hello Adika';")
	}

	t.Log("test hello adika passed")
}

// test main itu akan dijalankan sebelum dan sesudah unit test
// hanya dijalankan dalam 1 package sekali
func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run()
	fmt.Println("After unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Sub Test 1", func(t *testing.T) {
		result := HelloWorld("adi")
		assert.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	})

	t.Run("Sub Test 2", func(t *testing.T) {
		result := HelloWorld("Joko")
		assert.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")
	})
}
