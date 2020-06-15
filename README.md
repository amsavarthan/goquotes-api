# Go Quotes API
GO Quotes is a free, open source quotations API. It was built when I was getting started with GO lang. The sample database includes 500+ quotes.
[Live API Documentation](https://goquotes.docs.apiary.io)

- [Server](#server)
- [API Methods](#api-methods)
  - [Get a random quote(s)](#get-a-random-quotes)
  - [Get random quote(s) filtered by given type](#get-random-quotes-filtered-by-given-type)
  - [Get all quotes](#get-all-quotes)
  - [Get all authors](#get-all-authors)
  - [Get all tags](#get-all-tags)
  - [Get all data quotes by given type](#get-all-data-quotes-by-given-type)


## Server

| Name         |   URL                                        | Description    
|--------------|----------------------------------------------|-------------------------------------------
| Production   |  https://goquotes-api.herokuapp.com/api/v1/          | Synced with the master branch of this repo

## API Methods

### Get a random quote(s)
>Returns number of random quotes

**Path**
```HTTP
GET /random
```

**Query Parameters**

Param   | Type   | Description
--------|--------|-------------
count   |`Int`   | Number of quotes to fetch

* * *

### Get random quote(s) filtered by given type
>Returns number of random quotes filtered by type

**Path**
```HTTP
GET /random/:count
```

**Query Parameters**

Param   | Type     | Description                        |
--------|----------|------------------------------------|
count   |`Int`     | Number of quotes to fetch          |
type    | `String` | Filter quotes by type (author/tag) |
val     | `String` | Filter quotes by value of type     |

* * *

### Get all quotes
>Returns every quotes available in the server

**Path**
```HTTP
GET /all/quotes
```

* * *

### Get all authors
>Returns every authors available in the server

**Path**
```HTTP
GET /all/authors
```

* * *

### Get all tags
>Returns every tags available in the server

**Path**
```HTTP
GET /all/tags
```

* * *

### Get all data quotes by given type
>Returns all quotes filtered by given type and value

**Path**
```HTTP
GET /all
```

**Query Parameters**

Param   | Type     | Description                        |
--------|----------|------------------------------------|
type    | `String` | Filter quotes by type (author/tag) |
val     | `String` | Filter quotes by value of type     |
