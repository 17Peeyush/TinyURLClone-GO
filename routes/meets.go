package routes

import (
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
	err = meet.Save()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	context.JSON(http.StatusOK, gin.H{"uniqueID": meet.UniqueID})
}

func getMeeting(context *gin.Context){
	var meet models.Meet
	paramValue := context.Param("url")
	if paramValue == ""{
		context.JSON(http.StatusBadRequest, gin.H{"error": "Path parameter 'param' is missing or empty."})
		return
	}
	meet.UniqueID = paramValue
	err := meet.GetLink()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//redirection logic will come here till then returning link
	context.Redirect(http.StatusMovedPermanently, meet.Url)
	//context.JSON(http.StatusOK, gin.H{"url":meet.Url})
}

