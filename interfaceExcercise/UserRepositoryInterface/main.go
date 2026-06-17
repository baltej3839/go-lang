package main

import "fmt"

type UserRepoer interface {
	GetUser(int)
	CreateUser(int)
	DeleteUser(int)
}

type PostSQL struct {
	conn string
}

type MongoDB struct {
	conn string
}

type MemoryRepo struct {
	LocalLocation string
}

func (f *PostSQL) GetUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in fileAudit GetUser", amount)
}

func (f *PostSQL) CreateUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in fileAudit GetUser", amount)
}


func (f *PostSQL) DeleteUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in fileAudit GetUser", amount)
}


func NewPostSQL() *PostSQL {
	return &PostSQL{conn: "asdsadsad"}
}


func (r *MongoDB) GetUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in database GetUser", amount)
}

func (r *MongoDB) CreateUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in database GetUser", amount)
}

func (r *MongoDB) DeleteUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in database GetUser", amount)
}


func NewMongoDB() *MongoDB {
	return &MongoDB{conn: "asdsadsad"}
}

func (r *MemoryRepo) GetUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in remote GetUser", amount)
}

func (r *MemoryRepo) CreateUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in remote GetUser", amount)
}


func (r *MemoryRepo) DeleteUser(amount int) {
	fmt.Printf("amount transferred was %s , is GetUserged in remote GetUser", amount)
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{LocalLocation: "asdsadsad"}
}

type UserService struct {
	UserRepoer
}

func NewUserService(u UserRepoer) *UserService {
	return &UserService{u}
}



func main() {

	pSvc:=NewPostSQL()

	userRepoSvc:=NewUserService(pSvc)

	userRepoSvc.CreateUser(54)
	userRepoSvc.DeleteUser(234)
	userRepoSvc.GetUser(234)
}