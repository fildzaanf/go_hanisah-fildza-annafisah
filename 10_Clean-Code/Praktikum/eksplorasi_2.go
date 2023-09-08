package main

// untuk penulisan nama struct sesuai kebutuhan jika ingin dapat diakses diluar package dapat ditulis dengan awalan huruf besar atau PascalCase
type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

// untuk penulisan nama struct sesuai kebutuhan jika ingin dapat diakses diluar package dapat ditulis dengan awalan huruf besar atau PascalCase
type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}

	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Mobil{}

	mobilLamban.berjalan()
}
