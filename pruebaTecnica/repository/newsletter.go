package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/perajim/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Email struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"Email"`
}

type Newsletter struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"Name"`
	SentEmails int32              `bson:"SentEmails"`
	Recipients []Email            `bson:"Recipients"`
}

type NewsletterRepository struct {
	db *mongo.Collection
}

var db *mongo.Database

func CreateNewsletterRepository(collection string) *NewsletterRepository {
	db = initDB()
	return &NewsletterRepository{
		db: db.Collection(collection),
	}
}

func (r NewsletterRepository) CreateBookmark(createNewsletter models.CreateNewsletter, ctx context.Context) (models.Newsletter, error) {
	var newsletterModel models.Newsletter
	newsletterModel.ID = primitive.NewObjectID()
	newsletterModel.Name = createNewsletter.Name
	newsletterModel.Recipients = []string{}
	newsletterModel.SentEmails = 0

	res, err := r.db.InsertOne(ctx, newsletterModel)
	if err != nil {
		return newsletterModel, err
	}

	newsletterModel.ID = res.InsertedID.(primitive.ObjectID)
	return newsletterModel, nil
}

func (r NewsletterRepository) SaveFile(storeFile models.StoreFile, ctx context.Context) (models.StoreFile, error) {
	storeFile.ID = primitive.NewObjectID()

	res, err := r.db.InsertOne(ctx, storeFile)
	if err != nil {
		return storeFile, err
	}

	storeFile.ID = res.InsertedID.(primitive.ObjectID)
	return storeFile, nil
}

func (r NewsletterRepository) AddRecipient(updateNewsletter models.UpdateNewsletter, ctx context.Context) error {
	var email Email

	objID, err := primitive.ObjectIDFromHex(updateNewsletter.IDNewsletter)
	if err != nil {
		panic(err)
	}

	var newsletter models.Newsletter
	error := r.db.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&newsletter)

	if error != nil {
		fmt.Println("No se pudo encontrar el newsletter indicado: " + error.Error())
	}
	email.ID = primitive.NewObjectID()
	email.Email = updateNewsletter.EmailRecipient
	_, err = r.db.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.D{
			{"$push", bson.M{"Recipients": email}},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		fmt.Println("EXISTIO UN ERROR EN DECODE del usuario: " + err.Error())
	}

	return err
}

func (r NewsletterRepository) RemoveEmail(email string, newsletterId string, ctx context.Context) error {
	newsId, _ := primitive.ObjectIDFromHex(newsletterId)
	emailId, _ := primitive.ObjectIDFromHex(email)
	//filter := bson.M{"_id": newsId, "Recipients._id": docID}
	//filter := bson.D{primitive.E{Key: "_id", Value: newsId}}
	//res, err := r.db.DeleteOne(ctx, filter)
	_, err := r.db.UpdateOne(
		context.Background(),
		bson.M{"_id": newsId},
		bson.M{
			"$pull": bson.M{"Recipients": bson.M{"_id": emailId}}},
	)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func (r NewsletterRepository) GetEmailList(id string, ctx context.Context) ([]Email, error) {
	uid, _ := primitive.ObjectIDFromHex(id)
	var newsletter Newsletter
	err := r.db.FindOne(ctx, bson.M{"_id": uid}).Decode(&newsletter)

	if err != nil {
		return nil, err
	}

	var list []Email
	for _, email := range newsletter.Recipients {
		list = append(list, email) // note the = instead of :=
	}

	return list, nil
}

func (r NewsletterRepository) GetNewsletters(ctx context.Context) ([]*Newsletter, error) {
	var newsletters []*Newsletter
	cur, err := r.db.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var newsletter Newsletter
		err := cur.Decode(&newsletter)
		if err != nil {
			return newsletters, err
		}

		newsletters = append(newsletters, &newsletter)
	}

	return newsletters, nil
}

func (r NewsletterRepository) GetFile(id string, ctx context.Context) (models.StoreFile, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	var storeFile models.StoreFile
	error := r.db.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&storeFile)

	if error != nil {
		return storeFile, error
	}
	return storeFile, error
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db:27017/"))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("usersMongo")
}
