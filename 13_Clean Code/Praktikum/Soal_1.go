package main

type user struct {
	id       int
	username int
	password int
	// 1. penamaan variabel kurang deskriptif
	// id -> ID
	// username -> Username
	// password -> Password
	// alasan : Penamaan variabel yang kurang deskriptif,
	// seperti "id", "username", dan "password"
	// yang seharusnya dinamai "ID", "Username",
	// dan "Password" agar lebih jelas dan mudah dimengerti.

	// 2. tipe data variabel kurang deskriptif
	// username -> Username string
	// password -> Password string
	// alasan : Tipe data variabel yang kurang deskriptif, seperti "username" dan "password"
	// yang seharusnya didefinisikan sebagai string agar lebih jelas.

	// 3. PENAMAAN STRUCT
	// user -> User
	// alasan : Penamaan struct yang kurang deskriptif,
	// seperti "user" yang seharusnya didefinisikan sebagai "User"

}

type userservice struct {
	t []user

	// 1. penamaan Struct dan variabel t kurang deskriptif
	// userservice -> UserService
	// t -> Users []user

	// alasan :Penamaan struct "user" seharusnya dinamai "User" agar lebih jelas dan mudah dimengerti.
	// Penamaan struct "userservice" dan variabel "t" kurang deskriptif.
	//  Seharusnya dinamai "UserService"
	// dan "Users" agar lebih jelas dan mudah dimengerti.
}

func (u userservice) getallusers() []user {
	// 1. penamaan variabel receiver kurang deskriptif
	// u -> us
	// alasan : Penamaan variabel receiver "u" kurang deskriptif tidak memperjelas Struct.

	// 2. penamaan variabel return kurang deskriptif
	// u.t -> us.Users
	// alasan : Penamaan variabel return "u.t" kurang deskriptif dan tidak jelas.

	// 3. penamaan method kurang deskriptif
	// getallusers -> GetAllUsers()
	// alasan : Penamaan method "getallusers" kurang deskriptif agar lebih jelas dan mudah dimengerti.
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	// 1. penamaan variabel receiver kurang deskriptif
	// u -> us
	// alasan : Penamaan variabel receiver "u" kurang deskriptif tidak memperjelas Struct.

	// 2. penamaan variabel return kurang deskriptif
	// u.t -> us.Users
	// alasan : Penamaan variabel return "u.t" kurang deskriptif dan tidak jelas.

	// 3. penamaan method kurang deskriptif tidak jelas
	// getuserbyid -> GetUserByID()

	//4. Pada fungsi getuserbyid(), jika tidak ditemukan user dengan
	// id tertentu, maka akan mengembalikan user kosong (zero value),
	// padahal bisa jadi ada user dengan id 0.
	// Seharusnya mengembalikan error atau menggunakan tipe data pointer user.
	// contoh:
	// func GetUserByID(id int) (*user, error) {
	// for _, u := range UserService.Users {
	// 	if id == u.ID {
	// 		return &u, nil
	// 	}
	// }
	// return nil, fmt.Errorf("user with id %d not found", id)
	// }

	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
