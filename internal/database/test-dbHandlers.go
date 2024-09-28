package database

import "github.com/andreadebortoli2/GO-Live-Chat/internal/models"

func (m *testDB) Login(email, password string) (models.User, error) {
	var result models.User
	return result, nil
}

func (m *testDB) AddUser(userName, email, password string) error {
	return nil
}

func (m *testDB) GetAllUsers() ([]models.User, error) {
	var users []models.User
	return users, nil
}

func (m *testDB) SetModerator(accLvl string, id string) error {
	return nil
}

func (m *testDB) DeleteUserByID(id string) error {
	return nil
}

func (m *testDB) GetLastMessages() ([]models.Message, error) {
	var messages []models.Message
	return messages, nil
}

func (m *testDB) GetOlderMessages() ([]models.Message, error) {
	var messages []models.Message
	return messages, nil
}

func (m *testDB) PostNewMessage(id int, msg string) error {
	return nil
}
