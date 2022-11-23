package repository

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"refactoring/api/request"
	"refactoring/exceptions"
	"refactoring/models"
	"strconv"
	"time"
)

type UserJsonStruct struct {
	Increment int                    `json:"increment"`
	UserList  map[string]models.User `json:"list"`
}

type UserJsonRepository struct {
	UserJsonStruct UserJsonStruct
	StoreFileName  string
}

func NewUserJsonRepository() *UserJsonRepository {

	jsonFileName := "users.json"

	return &UserJsonRepository{
		UserJsonStruct{0, make(map[string]models.User)},
		jsonFileName,
	}
}

func (userJRepo *UserJsonRepository) CreateUser(request request.CreateUserRequest) string {

	store := userJRepo.readStorageFromFile()

	store.Increment++

	newUser := models.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := strconv.Itoa(store.Increment)
	store.UserList[id] = newUser

	userJRepo.updateFileStorage(&store)

	return id
}

func (userJRepo *UserJsonRepository) UpdateUser(request request.UpdateUserRequest) error {

	store := userJRepo.readStorageFromFile()

	_, isUserWithId := store.UserList[request.Id]
	if !isUserWithId {
		return exceptions.UserNotFound
	}

	userToUpdate := store.UserList[request.Id]
	userToUpdate.DisplayName = request.DisplayName
	store.UserList[request.Id] = userToUpdate

	userJRepo.updateFileStorage(&store)
	return nil
}

func (userJRepo *UserJsonRepository) DeleteUser(id string) error {

	store := userJRepo.readStorageFromFile()

	_, isUserWithId := store.UserList[id]
	if !isUserWithId {
		return exceptions.UserNotFound
	}

	delete(store.UserList, id)

	userJRepo.updateFileStorage(&store)

	return nil
}

func (userJRepo *UserJsonRepository) GetAllUsers() map[string]models.User {

	store := userJRepo.readStorageFromFile()
	return store.UserList
}

func (userJRepo *UserJsonRepository) readStorageFromFile() UserJsonStruct {
	file, _ := ioutil.ReadFile(userJRepo.StoreFileName)
	store := UserJsonStruct{}
	_ = json.Unmarshal(file, &store)
	return store
}

func (userJRepo *UserJsonRepository) updateFileStorage(store *UserJsonStruct) {

	byteStore, _ := json.Marshal(&store)
	_ = ioutil.WriteFile(userJRepo.StoreFileName, byteStore, fs.ModePerm)
}
