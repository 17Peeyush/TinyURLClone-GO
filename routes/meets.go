package routes

import (
	"fmt"
	"net/http"

	"github.com/17Peeyush/TinyURLClone-GO/models"
	"github.com/gin-gonic/gin"
)

func newMeeting(context *gin.Context){
	var meet models.Meet

	err := context.ShouldBindJSON(&meet)
	if err !=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the request data."})
	}
	//meet.Url = context.Param("url")
	err = meet.Save()
	if err !=nil{
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	
	context.JSON(http.StatusOK, gin.H{"message": meet.UniqueID})
}