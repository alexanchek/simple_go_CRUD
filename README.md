# simple-CRUD-go

This is a simple example of CRUD-application (Create, Read, Update, Delete) using Golang. There is a representation of a database with books, including such fields such as Title, Year, Author.

# Setup

Clone this repository with 

git clone git@github.com:alexanchek/simple-CRUD-go.git

Set POSTGRESQL_URL in a file .env for a connection to your POSTGRES database.

Possible variants: 

postgresql://localhost
postgresql://localhost:5432
postgresql://localhost/mydb
postgresql://user@localhost
postgresql://user:secret@localhost
postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
postgresql://localhost/mydb?user=other&password=secret
