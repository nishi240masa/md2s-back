package services

import (
	"md2s/models"
	"md2s/repositorys"
)

func GetMember(id uint)(*models.Member,error){
	result,err:=repositorys.GetMemberById(1)
	if err!=nil {
		return nil,err
	}
	return result,nil
}