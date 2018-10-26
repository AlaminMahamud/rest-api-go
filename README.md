### Simple Rest Api in GO

``` shell

export PROJECT_NAME=restApi

# Git Clone 
https://github.com/AlaminMahamud/rest-api-go $PROJECT_NAME

cd $PROJECT_NAME

# Installing Router
go get github.com/gorilla/mux

go build
./$PROJECT_NAME

```

Now go to your API Client [e.g. browser / postman / advanced rest client]. and fire the endpoints.
1. GET  `localhost:8000/people`
2. GET  `localhost:8000/people/1`
3. POST `localhost:8000/people/4`

``` json
{
  "firstname": "Alamin",
  "lastname": "Mahamud",
  "address": {
    "city": "City X",
    "state": "State X"
  }
}
```
and so on...
