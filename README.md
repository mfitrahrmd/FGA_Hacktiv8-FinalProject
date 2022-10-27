# Final Project DTS FGA & Hacktiv8
api for managing user's photos and social medias

### Prerequisites

Prequisites package:
* Go (Go Programming Language)

Optional package:
* Docker (Application Containerization)

### Configuraion

**.env file is required for this project to run, you can create new one or rename the existing one (see .env.example)**

### Installing

Run following command to add required package to your host
```
go mod tidy
```

### Containerization

To run this code with docker container you can run following command:
```
docker compose up
```

### API Access

You can access any endpoint under base path or you can read the full documentations on 'http://127.0.0.1:8001/swagger/index.html' after you run this code.