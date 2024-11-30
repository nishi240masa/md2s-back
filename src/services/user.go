package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)

func GetUser(id uint)(*models.User,error){
	result,err:=repositorys.GetUserById(id)
	if err!=nil {
		return nil,err
	}
	return result,nil
}

func CreateUser(user dto.CreateUserData)(*models.User,error){

	newUser:=models.UserCreate{
		Name:user.Name,
		IconURL:user.IconURL,
		GoogleId:user.GoogleId,
	}


	return repositorys.CreateUser(newUser)
}

func DeleteUser(id models.UUID)error{

	if err:=repositorys.DeleteUser(id);err!=nil{
		return err
	}
	return nil

}