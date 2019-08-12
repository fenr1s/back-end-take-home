# Guestlogix Take Home Test - Backend

### Technologies used

Go - Docker - Swagger

- to run the project locally using docker
```
    docker build -t go-docker .
    docker run -d -p 8081:8081 go-docker
```
or 

```
    go run main.go
```
### Patterns, Best Practices, Algorithm
 - This project is using SOLID concepts starting by the depency injection that on the layers
 - As each layers has a dependecy on a interface not a pure implementation i can create mocks for them, giving me the possibilty to make real unit tests
 - So solve the shortest path problem i used the BFS breadth-first-search algorithm
 - controller layer and service layer have good test coverage

```
    go test ./... -cover
    ok      github.com/fenr1s/back-end-take-home/api/controllers    1.129s  coverage: 90.0% of statements
    ok      github.com/fenr1s/back-end-take-home/domain/services    0.578s  coverage: 82.7% of statements
```

### hosting
 - this project is hosted at http://167.71.171.15:8081
 - swagger url http://167.71.171.15:8081/swagger/index.html
 - or via http get http://167.71.171.15:8081/api/routes?origin=[YYZ]&destination=[JFK]