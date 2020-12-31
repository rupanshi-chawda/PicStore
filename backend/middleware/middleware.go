package middleware

import (
	"context"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"time"
	"bytes"
	"path"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/bson"
)

	
// SimpleServer is ...
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

// InitiateMongoClient intiates the connection to the MongoDb
func InitiateMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client
	uri := "mongodb://localhost:27017"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}
	return client
}

// UploadFile : Uploads the file to the database
func UploadFile(file, filename string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	conn := InitiateMongoClient()
	bucket, err := gridfs.NewBucket(
		conn.Database("Picstore"),
	)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	uploadStream, err :=bucket.OpenUploadStream(
		filename,
		)
		if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
}

// DownloadFile : helper		
func DownloadFile(fileName string) {
	conn := InitiateMongoClient()

	// For CRUD operations, here is an example
	db := conn.Database("Picstore")
	fsFiles := db.Collection("fs.files")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	// you can print out the results
	fmt.Println(results)
	bucket, _ := gridfs.NewBucket(
        db,
    )
	var buf bytes.Buffer
    dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
    if err != nil {
        log.Fatal(err)
	}
	fmt.Printf("File size to download: %v\n", dStream)
    ioutil.WriteFile(fileName, buf.Bytes(), 0600)	
}