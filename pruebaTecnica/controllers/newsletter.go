package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/perajim/models"
	"github.com/perajim/repository"
	"github.com/perajim/services"
)

//UserController ...
type NewsletterController struct{}

var newsletter = repository.CreateNewsletterRepository("newsletter")
var fileRepository = repository.CreateNewsletterRepository("files")

//Create ...
func (ctrl NewsletterController) Create(c *gin.Context) {
	var createNewsletter models.CreateNewsletter

	if c.ShouldBindJSON(&createNewsletter) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}

	newsletter, err := newsletter.CreateBookmark(createNewsletter, c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Hubo un error al crear el newsletter"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "newsletter creado con exito", "newsletter": gin.H{"id": newsletter.ID.Hex(), "email": newsletter.Name}})
}

//SaveImage ...
func (ctrl NewsletterController) SaveFile(c *gin.Context) {
	var storeFile models.StoreFile

	dst := "./.tmp/"

	file, error := c.FormFile("file")
	if error != nil {
		fmt.Print("\n\n\n Se presento el siguiente error: " + error.Error() + "\n\n\n\n")
	}
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst+file.Filename)

	storeFile.Path = dst + file.Filename
	storeFile.FileName = file.Filename

	fileSaved, err := fileRepository.SaveFile(storeFile, c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Hubo un error al crear el newsletter"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Archivo guardado con exito", "file": gin.H{"id": fileSaved.ID.Hex(), "email": fileSaved.FileName}})
}

//SendNewsletter ...
func (ctrl NewsletterController) SendNewsletter(c *gin.Context) {
	id := c.Param("id")
	idFile := c.Param("idFile")
	var sendNewsletter models.SendNewsletter

	if c.ShouldBindJSON(&sendNewsletter) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}

	file, error := fileRepository.GetFile(idFile, c)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "No se pudieron encontrar los archivos para enviar"})
	}

	emails, err := newsletter.GetEmailList(id, c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "No se pudieron encontrar emails para enviar"})
	}
	services.OAuthGmailService()
	status, err := services.SendEmail(file.Path, id, emails, sendNewsletter.Content, sendNewsletter.Subject)
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using OAUTH")
	}

	c.JSON(http.StatusOK, gin.H{"message": newsletter})

}

func (ctrl NewsletterController) AddRecipient(c *gin.Context) {
	var updateNewsletter models.UpdateNewsletter

	if c.ShouldBindJSON(&updateNewsletter) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}

	err := newsletter.AddRecipient(updateNewsletter, c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Hubo un error al agregar un correo al newsletter"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "correo agregado al newsletter con exito"})
}

func (ctrl NewsletterController) RemoveEmail(c *gin.Context) {
	newsletterId := c.Param("newsletter")
	email := c.Param("email")

	//getID, err := strconv.ParseInt(id, 10, 64)
	if len(email) < 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	err := newsletter.RemoveEmail(email, newsletterId, c)
	if err != nil {
		fmt.Print(err)

		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "No se pudo eliminar el Email del newsletter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email eliminado correctamente"})
}

func (ctrl NewsletterController) GetRecipients(c *gin.Context) {
	id := c.Param("newsletter")
	emails, err := newsletter.GetEmailList(id, c)
	for _, email := range emails {
		fmt.Println(email)
	}
	if err != nil {
		fmt.Println("HUBO UYN GRAN ERRROR: " + err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "data"})
}

func (ctrl NewsletterController) GetNewsletters(c *gin.Context) {
	newsletters, err := newsletter.GetNewsletters(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newsletters})
}
