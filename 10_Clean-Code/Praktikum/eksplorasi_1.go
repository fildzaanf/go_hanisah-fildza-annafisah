package main

// penulisan nama struct juga perlu diperhatikan huruf awalnya menyesuakan kebutuhan jika struct ingin diakses diluar package maka perlu menggunakan huruf awal kapital yaitu User
type user struct {
	id       int
	username int // data type yang digunakan tidak sesuai dengan fungsi nama variable yang digunakan. perbaikannya dapat menggunakan data type string
	password int // data type yang digunakan tidak sesuai dengan fungsi nama variable yang digunakan. perbaikannya dapat menggunakan data type string
}

// penulisan nama struct tidak menggunakan camelCase maka perbaikannya yaitu userService
// penulisan nama struct juga perlu diperhatikan huruf awalnya menyesuakan kebutuhan jika struct ingin diakses diluar package maka perlu menggunakan huruf awal kapital yaitu UserService
type userservice struct {
	t []user // penulisan variabel t juga kurang mendeskripsikan fungsinya sehingga sulit dipahami
}

// penulisan nama struct tidak menggunakan camelCase maka perbaikannya yaitu userService dan getAllUsers
// penulisan receiver name nya kurang mendeskripsikan fungsinya sehingga sulit dipahami yaitu dengan u
// penggunaan tipe receiver nya disarankan menggunakan tipe pointer agar dapat melakukan perubahan data dalam method tersebut ; *UserService
func (u userservice) getallusers() []user {
	return u.t // penulisan u.t kurang mendeskripsikan fungsinya sehingga sulit dipahami
}

// penulisan nama struct tidak menggunakan camelCase maka perbaikannya yaitu userService dan getUserById
// penulisan receiver name nya kurang mendeskripsikan fungsinya sehingga sulit dipahami yaitu dengan u
// penggunaan tipe receiver nya disarankan menggunakan tipe pointer agar dapat melakukan perubahan dalam method tersebut ; *UserService
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t { // penulisan variabel r juga kurang mendeskripsikan fungsinya sehingga sulit dipahami
		if id == r.id {
			return r
		}
	}

	return user{}
}
