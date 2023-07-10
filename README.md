# Task Manager with Golang
*** ***
## To run the code
* Build
```
go  build
```
* Run
```
./main  
```

** **
## ! It's experimental project

* ## Routes
* ### Get All Tasks (GET Request)
```
 http://localhost:3000/all
```
**  **

* ### Add new Task (POST request)
```
 http://localhost:3000/add
```
 * request body has to have 
 ```
 {
  "Title":"Reading",
  "Status":"doing"
  }
```
**  **
* ### Get by id (GET Request)
```
 http://localhost:3000?id=1
```
** **
* ### Update by id (PUT Request)
```
 http://localhost:3000/edit?id=1
```
* request body can have Status and / or Title
 ```
 {
  "Title":"Reading",
  "Status":"doing"
  }
```
** **
* ### Delete task by id (DELETE Request)
```
 http://localhost:3000/delete?id=1
```

** **

# Thanks


