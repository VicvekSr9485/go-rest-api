package controller

import (
	"diary-api/model"
	"diary-api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddEntry(context *gin.Context) {
	var input model.entry
	if err := context.ShouldBindJSON(&input), err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	user err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Entriess": user.Entries()})
}
