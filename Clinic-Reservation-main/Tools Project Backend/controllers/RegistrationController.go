package controllers

import (
	"toolsProject/models"
)

func SignUp(user models.User) models.User {
	if user.Type == "Doctor" {
		doctors := GetAllDoctors()
		var check bool = true
		for _, d := range doctors {
			if user.Name == d.Name || user.Mail == d.Mail {
				check = false
				break
			}
		}
		if !check {
			user.ID = 0
			return user
		}
		typeUser := models.Doctor{Name: user.Name, Mail: user.Mail, Type: user.Type, Password: user.Password}
		typeUser = AddDoctor(typeUser)
		user.ID = typeUser.ID
	} else {
		patients := GetAllPatients()
		var check bool = true
		for _, p := range patients {
			if user.Name == p.Name || user.Mail == p.Mail {
				check = false
				break
			}
		}
		if !check {
			user.ID = 0
			return user
		}
		typeUser := models.Patient{Name: user.Name, Mail: user.Mail, Type: user.Type, Password: user.Password}
		typeUser = AddPatient(typeUser)
		user.ID = typeUser.ID
	}
	user.UUID = GenerateUuid(user.ID, user.Type)
	return user
}

func SignIn(user models.User) models.User {
	if user.Type == "Doctor" {
		doctors := GetAllDoctors()
		var check bool = false
		for _, d := range doctors {
			if user.Name == d.Name && user.Password == d.Password {
				user.ID = d.ID
				user.Mail = d.Mail
				check = true
				break
			}
		}
		if !check {
			user.ID = 0
			return user
		}
	} else {
		patients := GetAllPatients()
		var check bool = false
		for _, p := range patients {
			if user.Name == p.Name && user.Password == p.Password {
				user.ID = p.ID
				user.Mail = p.Mail
				check = true
				break
			}
		}
		if !check {
			user.ID = 0
			return user
		}
	}
	user.UUID = GenerateUuid(user.ID, user.Type)
	return user
}
