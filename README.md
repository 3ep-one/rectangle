# Rectangle Overlap Detector

This package get rectangle dimenstion and check whether they overlap or not.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Require GO 1.6+
Require Redis 5.0.6+



### Installing

installed with:
```
go get -v github.com/3ep-one/rectangle
```

## Running the tests

You can run test using HTTPie:
```
  http POST http://127.0.0.1:8080 <inp.json
```
and then get response using :
```
  http Get http://127.0.0.1:8080
```


### And coding style tests

All code linted by flake8.

## Deployment
You should first set configs.
You can start program by:
```
  go run $GOPATH/src/github.com/3ep-one/rectangle/main.go 
```

After running the program you can send POST request JSON(exp. inp.json) and recive JSON answers by GET request.
All accepted values are stored in Redis database.

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
