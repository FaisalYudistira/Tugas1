/*package main

import (
	"fmt"
	"math"
)

func main() {
	//menhitung luas persegi
	var Sisi = 4
	fmt.Printf("1. Luas persegi dengan panjang sisi %d adalah %.f\n", Sisi, LuasPersegi(float64(Sisi)))

	//menghitung luas segitiga
	var Alas, Tinggi = 4, 5
	fmt.Printf("2. Luas segitiga dengan panjang alas %d dan tinggi %d adalah %.f\n", Alas, Tinggi, LuasSegitiga(float64(Alas), float64(Tinggi)))

	//menghitung luas lingkaran
	var Radius = 10
	fmt.Printf("3. Luas lingkaran dengan jari-jari %d adalah %.4f\n", Radius, LuasLingkaran(float64(Radius)))

	//menghitung energi potensial dan energi kinetik
	var Massa, Ketinggian, Kecepatan = 5, 7, 3
	var Ep, Ek = Energi(float64(Massa), float64(Ketinggian), float64(Kecepatan))
	fmt.Printf("4.a. Energi potensial suatu benda dengan massa %d dan ketinggian %d adalah %.f joule\n", Massa, Ketinggian, Ep)
	fmt.Printf("4.b. energi kinetik suatu benda dengan massa %d dan kecepatan %d adalah %.2f joule\n", Massa, Kecepatan, Ek)

	//menghitung frekuensi
	var Getaran, Waktu = 150, 20
	fmt.Printf("5. Frekuensi suatu benda dengan banyak getaran %d dan waktu %d detik adalah %.2f Hz\n", Getaran, Waktu, Frekuensi(float64(Getaran), float64(Waktu)))

	//menghitung konversi suhu
	var C, F, K, R = 60, 90, 350, 40
	KonversiSuhu(float64(C), float64(F), float64(K), float64(R))
}

// fungsi menghitung luas persegi
func LuasPersegi(s float64) (area float64) {
	area = math.Pow(s, 2)
	return
}

// fungsi menghitung luas segitiga
func LuasSegitiga(a float64, t float64) (area float64) {
	area = 0.5 * a * t
	return
}

// fungsi menghitung luas lingkaran
func LuasLingkaran(r float64) (area float64) {
	area = math.Pi * math.Pow(r, 2)
	return
}

// fungsi menghitung besar energi potensial dan energi kinetik
func Energi(m float64, h float64, v float64) (EnergiPotensial float64, EnergiKinetik float64) {
	const g float64 = 10
	EnergiPotensial = m * g * h
	EnergiKinetik = 0.5 * m * math.Pow(v, 2)
	return
}

// fungsi menghitung frekuensi
func Frekuensi(n float64, t float64) (f float64) {
	f = n / t
	return
}

// fungsi menghitung konversi suhu
func KonversiSuhu(C float64, F float64, K float64, R float64) {
	var Celcius, Fahrenheit, Kelvin, Reamur float64
	fmt.Println("6. Konversi satuan suhu")
	fmt.Println("Konversi satuan Celcius")
	//konversi celcius menjadi fahrenheit
	Fahrenheit = 1.8*C + 32
	fmt.Printf("%.f°C = %.2f°F\n", C, Fahrenheit)
	//konversi celcius menjadi kelvin
	Kelvin = C + 273.15
	fmt.Printf("%.f°C = %.2f°K\n", C, Kelvin)
	//konversi celcius menjadi reamur
	Reamur = 0.8 * C
	fmt.Printf("%.f°C = %.2f°R\n", C, Reamur)

	fmt.Println("Konversi satuan Fahrenheit")
	//konversi fahrenheit menjadi celcius
	Celcius = (F - 32) * 5 / 9
	fmt.Printf("%.f°F = %.2f°C\n", F, Celcius)
	//konversi fahrenheit menjadi kelvin
	Kelvin = (F + 459.67) * 5 / 9
	fmt.Printf("%.f°F = %.2f°K\n", F, Kelvin)
	//konversi fahrenheit menjadi reamur
	Reamur = (F - 32) * 4 / 9
	fmt.Printf("%.f°F = %.2f°R\n", F, Reamur)

	fmt.Println("Konversi satuan Kelvin")
	//konversi kelvin menjadi celcius
	Celcius = K - 273.15
	fmt.Printf("%.f°K = %.2f°C\n", K, Celcius)
	//konversi kelvin menjadi fahrenheit
	Fahrenheit = (K * 9 / 5) - 459.67
	fmt.Printf("%.f°K = %.2f°F\n", K, Fahrenheit)
	//konversi kelvin menjadi reamur
	Reamur = 0.8 * (K - 273.15)
	fmt.Printf("%.f°K = %.2f°R\n", K, Reamur)

	fmt.Println("Konversi satuan Reamur")
	//konversi reamur menjadi celcius
	Celcius = R / 0.8
	fmt.Printf("%.f°R = %.2f°C\n", R, Celcius)
	//konversi reamur menjadi fahrenheit
	Fahrenheit = R*2.25 + 32
	fmt.Printf("%.f°R = %.2f°F\n", R, Fahrenheit)
	//konversi reamur menjadi kelvin
	Kelvin = R/0.8 + 273.15
	fmt.Printf("%.f°R = %.2f°K\n", R, Kelvin)
}