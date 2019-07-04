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