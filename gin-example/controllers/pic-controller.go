package controllers

import (
	"gin-example/models"
	"github.com/gin-gonic/gin"
	"sync"
)

type PicController interface {
	GETALL(context *gin.Context)

	UPDATE(context *gin.Context)

	CREATE(context *gin.Context)

	DELETE(context *gin.Context)
}



type controller struct{
	pics []models.Pics
}

type generator struct {
	counter int
	mtx sync.Mutex
}

func (g *generator)getNextId()int  {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.counter++
	return g.counter
}

var g *generator = &generator{}

func (c *controller) GETALL(context *gin.Context) {
	context.JSON(200,c.pics)
}

func (c *controller) UPDATE(context *gin.Context) {
	var picToUpdate models.Pics
	if err := context.ShouldBindUri(&picToUpdate); err!=nil{
		context.String(400, "bad request", err)
	}
	if err := context.BindJSON(&picToUpdate); err!=nil{
		context.String(400, "bad request", err)
		return
	}
	for idx, pic := range c.pics {
		if pic.Id == picToUpdate.Id{
			c.pics[idx] = picToUpdate
			context.String(200, "success", picToUpdate.Id)
			return
		}
	}
	context.String(400, "bad request")
	return
}

func (c *controller) CREATE(context *gin.Context) {
	picToCreate := models.Pics{Id: g.getNextId()}
	if err := context.BindJSON(&picToCreate); err!=nil{
		context.String(400, "bad request", err)
	}
	c.pics = append(c.pics, picToCreate)
	context.String(200,"success")
}

func (c *controller) DELETE(context *gin.Context) {
	var picToDelete models.Pics
	if err := context.ShouldBindUri(&picToDelete); err!=nil{
		context.String(400, "bad request", err)
		return
	}
	for idx, pic := range c.pics {
		if pic.Id == picToDelete.Id{
			c.pics = append(c.pics[0:idx], c.pics[idx+1:len(c.pics)]...)
			context.String(200,"success delete")
			return
		}
	}
	context.String(400, "bad request")
	return
}

func NewPicController() PicController  {
	return &controller{pics: make([]models.Pics,0)}
}


