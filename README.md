# codecoverage-thesis

* To get the DEV environment started, please run ```docker run --name mongodb -p 27017:27017 -d mongo:latest``` to fire up the database
* Then run inside ```backend``` directory the command ```go run main.go collections.go db.go``` to start the backend
* Finally inside ```frontend``` directory run ```npm start```, this will start the frontend redirect you to its localhost url
