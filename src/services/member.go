package services

import (
	"md2s/models"
	"md2s/repositorys"
)

func GetMember(id uint)(*models.User,error){
	result,err:=repositorys.GetMemberById(id)
	if err!=nil {
		return nil,err
	}
	return result,nil
}