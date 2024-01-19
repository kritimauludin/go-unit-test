/**
Menjalankan Unit Test

● Untuk menjalankan unit test kita bisa menggunakan perintah :
go test
● Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan
perintah :
go test -v
● Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah
:
'go test -v -run TestNamaFunction'

Menjalankan Semua Unit Test

● Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa gunakan perintah :
go test ./...
*/

package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
*
Menggagalkan Unit Test

● Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
● Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
● Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

t.Fail() dan t.FailNow()

● Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa
bedanya?
● Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir
ketika selesai, maka unit test tersebut dianggap gagal
● FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

t.Error(args...) dan t.Fatal(args...)

● Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
● Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia
akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
● Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai
selesai
● Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil
FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
*/
func TestHelloWorldKriti(t *testing.T) {
	result := HelloWorld("Kriti")
	if result != "Hello Kriti" {
		//error
		// t.Fail()
		t.Error("Result must be 'Hello Kriti")
	}

	fmt.Println("TestHelloWorldKriti Done")
}

func TestHelloWorldMauludin(t *testing.T) {
	result := HelloWorld("Mauludin")
	if result != "Hello Mauludin" {
		//error
		// t.FailNow()
		t.Fatal("Result must be 'Hello Kriti")
	}

	fmt.Println("TestHelloWorldMauludin Done")
}

/**
Before dan After Test

● Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit
test dieksekusi
● Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka
membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode
duplikat jadinya
● Untungnya di Go-Lang terdapat fitur yang bernama testing.M
● Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga
bisa kita gunakan untuk melakukan Before dan After di unit test

testing.M

● Untuk mengatur ekeskusi unit test, kita cukup membuat sebuah function bernama TestMain
dengan parameter testing.M
● Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi
function ini tiap kali akan menjalankan unit test di sebuah package
● Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
● Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function
unit test
*/

func TestMain(m *testing.M) {
	//before testing
	fmt.Println("Before testing")

	m.Run()

	//after testing
	fmt.Println("After testing")
}

/*
*
Assertion

● Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
● Apalagi jika result data yang harus di cek itu banyak
● Oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan
● Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh
menambahkan library untuk melakukan assertion ini
*/
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Kriti")
	if result != "Hello Kriti" {
		//error
		assert.Equal(t, "Hello Kriti", result, "Result mmust be 'Hello Kriti") //jika gagal panggol fail
	}

	fmt.Println("TestHelloWorld with assert Done")
}

/*
*
Testify

● Salah satu library assertion yang paling populer di Go-Lang adalah Testify
● Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
● https://github.com/stretchr/testify
● Kita bisa menambahkannya di Go module kita :
go get github.com/stretchr/testify
*/
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Kriti")
	if result != "Hello Kriti" {
		//error
		require.Equal(t, "Hello Kriti", result, "Result mmust be 'Hello Kriti") //jika gagal panggol failnow
	}

	fmt.Println("TestHelloWorld with require Done")
}

/**
Skip Test

● Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
● Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
● Untuk membatalkan unit test kita bisa menggunakan function Skip()
*/

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Kriti")
	require.Equal(t, "Hello Kriti", result, "Result mmust be 'Hello Kriti") //jika gagal panggol failnow
}

/*
*
Sub Test

● Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
● Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang
lainnya
● Untuk membuat sub test, kita bisa menggunakan function Run()

# Menjalankan Hanya Sub Test

● Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
go test -run TestNamaFunction
● Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
go test -run TestNamaFunction/NamaSubTest
● Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
go test -run /NamaSubTest
*/
func TestSubTest(t *testing.T) {
	t.Run("kriti", func(t *testing.T) {
		result := HelloWorld("kriti")
		require.Equal(t, "Hello kriti", result, "Result mmust be 'Hello kriti") //jika gagal panggol failnow
	})
	t.Run("mauludin", func(t *testing.T) {
		result := HelloWorld("mauludin")
		require.Equal(t, "Hello mauludin", result, "Result mmust be 'Hello mauludin") //jika gagal panggol failnow
	})
}

/*
*
Table Test

● Sebelumnya kita sudah belajar tentang sub test
● Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
● Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan
konsep table test
● Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi
hasil dari unit test
● Lalu slice tersebut kita iterasi menggunakan sub test
*/
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "kriti",
			request:  "kriti",
			expected: "Hello kriti",
		},
		{
			name:     "mauludin",
			request:  "mauludin",
			expected: "Hello mauludin",
		},
		{
			name:     "aul",
			request:  "aul",
			expected: "Hello aul",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

/**
Benchmark Function

● Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama function nya,
harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
● Selain itu, harus memiliki parameter (b *testing.B)
● Dan tidak boleh mengembalikan return value
● Untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go
*/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("kriti")
	}
}
func BenchmarkHelloWorldMauludin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mauludin")
	}
}

/*
*
Sub Benchmark

● Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function
Run()

# Menjalankan Hanya Sub Benchmark

● Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
● Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :
go test -v -bench=BenchmarkNama/NamaSub
*/
func BenchmarkSub(b *testing.B) {
	b.Run("kriti", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("kriti")
		}
	})
	b.Run("mauludin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("mauludin")
		}
	})
}

/*
*
Table Benchmark

● Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga
● Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data
berbeda-beda tanpa harus membuat banyak benchmark function
*/
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "kriti",
			Request: "kriti",
		},
		{
			Name:    "mauludin",
			Request: "mauludin",
		},
		{
			Name:    "kritimauludin",
			Request: "kriti mauludin",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}
