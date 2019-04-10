package model

import (
	"database/sql"
)

type User struct {
	ID       int    `json:id`
	Name     string `json:name`
	Password string `json:pass`
	Email    string `json:email`
	//File     string `json:file`
}
type Book struct {
	Title          string `json:title`
	Author         string `json:author`
	Price          int    `json:price`
	ImgUrl         string `json:img_url`
	RakutenPageUrl string `json:rakuten_url`
	User_id        int    `json:id`
}

func (u *User) Insert(db *sql.DB) (*User, error) {
	_, err := db.Exec("insert into users (username,email,password) values (?,?,?)", u.Name, u.Email, u.Password)
	//fmt.Println("RES", res)
	//fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		//File:u.File
	}, nil
}
func (b *Book) InsertBook(db *sql.DB) (*Book, error) {
	_, err := db.Exec("insert into books (title,author,price,img_url,rakuten_page_url,user_id) values (?,?,?,?,?,?)", b.Title, b.Author, b.Price, b.ImgUrl, b.RakutenPageUrl, b.User_id)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (u *User) Check(db *sql.DB) (*User, error) {
	//fmt.Println("JKJK", u)
	//fmt.Println(u.Password)

	usr := &User{}
	if err := db.QueryRow(`select id,username,email,password from users where email=? AND password=?`, u.Email, u.Password).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password); err != nil {
		return nil, err
	}
	//fmt.Println("JJJ", usr)
	return usr, nil
}
func (u *User) SelectById(db *sql.DB) (*User, error) {
	usr := &User{
		ID: u.ID,
	}
	if err := db.QueryRow(`select username,email from users where id=?`, u.ID).Scan(&usr.Name, &usr.Email); err != nil {
		return nil, err
	}
	return usr, nil
}
func (u *User) SelectIdByNameANDPass(db *sql.DB) (*User, error) {
	usr := &User{}
	if err := db.QueryRow(`select id from users where username=? AND password=?`, u.Name, u.Password).Scan(&usr.ID); err != nil {
		return nil, err
	}
	return usr, nil
}
func (u *User) SelectPersonalBooks(db *sql.DB) (*Book, error) {
	bo := &Book{}
	//for i, _ := range bo {
	//	err := db.QueryRow(`select title,author,price,img_url,rakuten_page_url from books where user_id=?`, u.ID).Scan(&bo[i].Title, &bo[i].Author, &bo[i].Price, &bo[i].ImgUrl, &bo[i].RakutenPageUrl)
	//}
	if err := db.QueryRow(`select title,author,price,img_url,rakuten_page_url from books where user_id=?`, u.ID).Scan(&bo.Title, &bo.Author, &bo.Price, &bo.ImgUrl, &bo.RakutenPageUrl); err != nil {
		return nil, err
	}
	return bo, nil
}
