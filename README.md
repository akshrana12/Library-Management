# Library Management API

This is a library management API system build in GoLang Language. There are three main tables.


## DataBase Info

1. Users - It stores the information about the user. 
Schema of users = {
  StudentId,
  StudentName,
  EmailId
}

2. Books - It stores the information about the books.
Schema of books = {
  BookId,
  BookName,
  BookAuther,
  BookGenre
}

3. BookIssue - This table contains which user has issued which book.
Schema of BookIssue = {
  UserId,
  BookId
}


Below are the few endpoints:

## Book Related API'S

1. POST - /api/books/addBook

It adds the new book to the Books table.
#### JSON it accepts
{
  BookId INTEGER,
  BookName STRING,
  BookAuther STRING,
  BookGenre STRING
}

#### On Success It returns
status 200 
{
  BookId INTEGER,
  BookName STRING,
  BookAuther STRING,
  BookGenre STRING
}

2. GET - /api/books/getAllBooks

#### On Success It returns
status 200  with list of all the books that are available.

3. GET - /api/books/getBook/:id

#### On Success It returns
status 200  with book with given id.

4. GET - /getBook/genre/:genre

#### On Success It returns
status 200  with all books of given genre.


## User Related API'S

1. POST - /api/user/addUser

It adds the new user to the Users table.
#### JSON it accepts
{
  StudentId INTEGER,
  StudentName STRING,
  EmailId STRING
}

#### On Success It returns
status 200 
{
  StudentId INTEGER,
  StudentName STRING,
  EmailId STRING
}

2. GET - /api/user/getAllUsers

#### On Success It returns
status 200  with list of all the users that are available.

3. GET - /api/user/getUser/:id

#### On Success It returns
status 200  with user with given id.


## BookIssue Related API'S

1. POST - /api/action/issueBook?userId={valid user Id}&bookId={valid book Id}

#### On Success It returns
status 200 
{
  UserId INTEGER,
  BookId INTEGER,
}

2. DELETE - /api/action/returnBook?userId={valid user Id}&bookId={valid book Id}

#### On Success It returns
status 200

3. GET - /api/action/getUserDetail?userId={valid user Id}

#### On Success It returns
status 200  with all the books issued by this user

