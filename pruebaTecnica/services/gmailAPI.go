package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/perajim/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// GmailService : Gmail client for sending email
var GmailService *gmail.Service

func OAuthGmailService() {
	config := oauth2.Config{
		ClientID:     "1056273389086-c875rje1vijvuivfd9dvmqbjivcf7pf5.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  "ya29.A0AVA9y1vqdDmrma6bxDmk9qH2ZFr_dHHok1jDgwddAES1WdyPJ9Fsoy4gtLE1ZMUKkzv4iNlw5sqyphA1U3o2r6vpqLlAi1E7RuhG8xKfmNbe5et2T6_pvuhq6mQvg39QrWszmWX0dIdirXtyoaHXpKxIDps0YUNnWUtBVEFTQVRBU0ZRRTY1ZHI4aThtczgyQkVxMl9iNU9WellaQ2JLdw0163",
		RefreshToken: "1//0459pHp1OiiejCgYIARAAGAQSNwF-L9IrQpLUSbUj0N3WBkjE3LWF3piEyk-KeLSBXUJk_qEvKKcL-aD9ar63-Rc_IkwbIhlzI3I",
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is initialized \n")
	}
}

func SendEmail(dst string, id string, emails []repository.Email, content string, subject string) (bool, error) {
	boundary := randStr(32, "alphanum")

	fileBytes, err := ioutil.ReadFile(dst)
	if err != nil {
		log.Fatalf("Unable to read file for attachment: %v", err)
	}
	fileMIMEType := http.DetectContentType(fileBytes)
	fileData := base64.StdEncoding.EncodeToString(fileBytes)
	url := strings.Split(dst, "/")
	fileName := url[len(url)-1]
	AttachedContent :=
		"--" + boundary + "\n" +
			"Content-Type: " + fileMIMEType + "; name=" + string('"') + fileName + string('"') + " \n" +
			"MIME-Version: 1.0\n" +
			"Content-Transfer-Encoding: base64\n" +
			"Content-Disposition: attachment; filename=" + string('"') + fileName + string('"') + " \n\n" +
			fileData

	var message gmail.Message

	from := "apis.cdmx@gmail.com"

	for _, email := range emails {
		fmt.Println(email.Email)
		unsuscribe := "\n\n\nPara dejar de recibir este newsletter da click en el siguiente enlace:\n" + "http://localhost:8080/suscription/" + id + "/" + email.ID.Hex()
		content := content + unsuscribe
		msgEncoded := ("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
			"MIME-Version: 1.0\n" +
			"to: " + email.Email + "\n" +
			"from: " + from + "\n" +
			"subject: " + subject + "\n\n" +

			"--" + boundary + "\n" +
			"Content-Type: text/plain; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
			"MIME-Version: 1.0\n" +
			"Content-Transfer-Encoding: 7bit\n\n" +
			content + "\n\n" +
			"--" + boundary + "\n" +

			AttachedContent +
			"--" + boundary + "--")
		msg := []byte(msgEncoded)

		message.Raw = base64.URLEncoding.EncodeToString(msg)
		_, err := GmailService.Users.Messages.Send("me", &message).Do()
		//Send the message
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
