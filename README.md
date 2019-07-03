# Redeam Coding Exercise  
## Setup

## Operation

### Deploying a Postgres DB
To deploy a Postgres container to Docker, I simply get the Postgres Docker Hub image and ran this command:
`docker run --rm --name postgres-docker -e POSTGRES_PASSWORD=dopesauce -d -p 15432:5432 postgres`

The postgres instance can be accessed through port 15432, which the Go app is configured to look for when accessing the database.

### Running the Go app
Note: Make sure to start up the Postgres container before running the REST app.  The app will verify/create all database tables as part of its startup.

From all my tests and attempts, I haven't been able to successfully deploy a container with a Go application.
To run the app, I simply:
 * run `go test && go build`
 * execute the generated executable file

Once running on the local machine, the Books app can be accessed through four REST endpoints:

#### Create a Book
URL: `localhost:8080/book`
Request Method: `POST`
Headers: `'Content-Type': 'application/json'`
Request Body:
```
{
	"Title": "Harry Potter",
	"Author": "J K Rowling",
	"Publisher": "Bloomsbury",
	"Publish_Date": "12-31-1997",
	"Rating": 2.2334,
	"Status": 2
}
```
 * Title, Author, and Publisher must be provided and non-empty
 * Publish_Date must be in MM-DD-YYYY format
 * Rating must be provided, 1.0 <= Rating <= 3.0
 * Status must be value 1-3.  Values correlate with entries from the `book_status` table.

Response Body:
```
{
	"id": 1
}
```
 * ID value corresponds with ID entry in the `book` DB table.

#### Read a Book
URL: `localhost:8080/book/{id}`
Request Method: `GET`
Headers: `Accept: application/json`
Expected Response Body (if book available):
```
{
	"Title": "Harry Potter",
	"Author": "J K Rowling",
	"Publisher": "Bloomsbury",
	"Publish_Date": "12-31-1997",
	"Rating": 2,
	"Status": 2
}
```

#### Update a Book
URL: `localhost:8080/book/{id}`
Request Method: `PUT`
Headers: `'Content-Type': 'application/json'`
Request Body: (Note: All fields are optional and may be left out to prevent updates to those fields)

#### Delete a Book
URL: `localhost:8080/book/{id}`
Request Method: `DELETE`

## Development Process

### Learning Go
The Go programming language, although very different from Java in build and operation, comes across as a very solid platform.  That being said, it took me some time to adjust to a few of the differences between the two platforms.

The first thing I noticed is that Go is a component-oriented language by design, while Java requires other libraries (Dropwizard, Spring Boot, etc.) to become one.  This threw me off initially, but it starts to make sense after spending some time trying to understand how to use its structures.

Error handling seems rudimentary at first, but is surprisingly intuitive.  Returning multiple values from a method has the added benefit of nudging developers to handle errors as they write and use methods, instead of deferring all errors to the end of a method.

One of the larger pain points was learning how to mock objects for clean testing in Go.  Without objects to assign mocks to, it took a little while to understand how a project needs to be structured in order to be unit testable (and mockable).

### Database
For the RDBMS, I chose to use Docker Hub's postgres image.  While I would never do this for a production build, for the purposes of this exercise, I hard-coded the database login credentials and table creation into the db_storage.go file and use the provided Docker commands to initialize the postgres Docker container.

### Entry Validation
Normally in a Java REST application, I would set up a custom ExceptionHandler and map a list of exceptions to the HTTP status codes they should return.  In retrospect, it's entirely possible that a Go library exists which can connect to Gorilla Mux and map outbound errors to a provided list of status codes, but ultimately I performed my error handling around what I found was possible to check in the different layers.

To accommodate the requirements for updates and creates, I added validation functions to the Book struct.  This solution came after pondering how to tailor the UPDATE SQL statement to fields provided in its request body, since the system could accept the default values for Book and write them to the database.

### Deploymenet
For some reason or another, I've had a difficult time standing up a Go image on Docker.  After some investigation, it turned out that my file structure was causing issues with `go get` because the subdirectories didn't have host names.  I flattened my application to one directory level as a last-minute effort to deploy a container locally.