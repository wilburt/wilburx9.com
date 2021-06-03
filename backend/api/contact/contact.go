package contact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wilburt/wilburx9.dev/backend/api/internal"
	"github.com/wilburt/wilburx9.dev/backend/configs"
	"gopkg.in/gomail.v2"
	"net/http"
)

type requestData struct {
	SenderEmail string `json:"sender_email"`
	SenderName  string `json:"sender_name"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}

// Handler validates request body and reCAPTCHA and possibly sends an email
func Handler(c *gin.Context) {
	var data requestData
	err := c.ShouldBindJSON(&data)
	message := validate(data)
	if err != nil || message != "" {
		c.JSON(http.StatusBadRequest, internal.MakeErrorResponse(message))
		return
	}

	// TODO: Validate reCAPTCHA

	err = sendEmail(data)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Warning("Couldn't send email")
		c.JSON(http.StatusBadGateway, internal.MakeErrorResponse("Email not sent"))
		return
	}
	c.JSON(http.StatusOK, internal.MakeSuccessResponse("Email sent successfully"))
}

func sendEmail(data requestData) error {
	config := configs.Config
	sender := fmt.Sprintf("%v <%v>", data.SenderName, data.SenderEmail)
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", config.ContactEmail)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/plain", data.Message)

	d := gomail.NewDialer(config.SmtpHost, config.SmtpPort, config.SmtpUsername, config.SmtpPassword)
	err := d.DialAndSend(m)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Sent")
	}
	return err
}

func validateRecaptcha() error {
	// TODO implement reCAPTCHA verification
	return nil
}
