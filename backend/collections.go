package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// App ... the app collection
type App struct {
	AppID   primitive.ObjectID `bson:"AppID" json:"AppID"`
	Name    string             `bson:"Name" json:"Name"`
	Reports []reports          `bson:"Reports" json:"Reports"`
}

type reports struct {
	CommitHash         string    `bson:"CommitHash" json:"CommitHash"`
	CoveragePercentage float64   `bson:"CoveragePercentage" json:"CoveragePercentage"`
	CreationDate       time.Time `bson:"CreationDate" json:"CreationDate"`
}

var collection, err = getDbCollection("mycollection", "mydb")

func newApp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit endpoint: newApp")

	w.Header().Set("Content-Type", "application/json")

	var app App

	err = r.ParseMultipartForm(32 << 20)

	if err != nil {
		http.Error(w, "error parsing multipart request", http.StatusBadRequest)
	}

	app.Name = r.FormValue("Name")

	f, _, err := r.FormFile("CoverageFile")
	if err != nil {
		log.Fatalf("Parsing file: %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	_, err = io.Copy(&b, f)

	if err != nil {
		http.Error(w, "error getting file", http.StatusInternalServerError)
	}

	content := b.String()

	b.Reset()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"Name": app.Name}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatalf("Counting document: %s", err)
	}

	percentage, err := parseCoverageFile(content)
	if err != nil {
		log.Fatalf("parsing coverage file: %s", err)
	}

	cHash := r.FormValue("CommitHash")

	if count >= 1 {

		fmt.Printf("Document by the name %s already exists, updated", app.Name)

		enableCORS(&w)

		newReport := reports{
			CommitHash:         cHash,
			CoveragePercentage: percentage,
			CreationDate:       time.Now().UTC(),
		}

		op := bson.M{"$push": bson.M{"Reports": newReport}}
		collection.UpdateOne(ctx, filter, op)

		w.WriteHeader(http.StatusOK)

		var updatedApp App
		err = collection.FindOne(ctx, filter).Decode(&updatedApp)

		if err != nil {
			log.Printf("getting an app: %s", err)
		}

		json.NewEncoder(w).Encode(updatedApp)
		return
	}

	app.AppID = primitive.NewObjectID()

	newReport := reports{
		CommitHash:         cHash,
		CoveragePercentage: percentage,
		CreationDate:       time.Now().UTC(),
	}

	app.Reports = append(app.Reports, newReport)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, app)

	if err != nil {
		log.Fatalf("insert one to db: %s", err)
	}

	enableCORS(&w)
	// 201 Created
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(app)
}

func parseCoverageFile(givenString string) (float64, error) {
	if strings.Contains(givenString, "total:") {
		line := givenString[strings.LastIndex(givenString, ")")+1:]

		spacedPercentage := strings.Split(line, "%")

		trimmedPercentage := strings.TrimSpace(spacedPercentage[0])

		floatedPercentage, err := strconv.ParseFloat(trimmedPercentage, 64)
		if err != nil {
			log.Fatalf("parsing coverage file percentage: %s", err)
		}
		return floatedPercentage, nil
	}
	return 0.0, errors.New("coverage file not supported or invalid file")
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents

func getOneApp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit endpoint: getOneApp")
	w.Header().Set("Content-Type", "application/json")

	p := mux.Vars(r)
	appName := p["Name"]

	filter := bson.M{"Name": appName}

	var app App

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&app)
	if err != nil {
		log.Printf("getting an app: %s", err)
	}
	fmt.Println(app)
	json.NewEncoder(w).Encode(app)
}

func getAllApps(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit endpoint: getAllApps")
	w.Header().Set("Content-Type", "application/json")

	var apps []App

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("getting cursor: %s", err)
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &apps)

	if err != nil {
		log.Fatalf("getting all apps: %s", err)
	}
	fmt.Println(apps)

	json.NewEncoder(w).Encode(apps)
}
