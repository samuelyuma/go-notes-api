# Notes API

Simple notes api with CRUD implementation.

## ⚡ Technologies

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## 🔧 Installation

To install and run the project locally, follow these steps:

-   Clone the repository: `git clone https://github.com/samuelyuma/go-notes-api.git`
-   Navigate to the project directory: `cd go-notes-api`
-   Copy the `.env.example` for your local version (`.env`)
-   Build and start the container: `docker compose up --build`

The app will now be running at http://localhost:8080/.

## 🛣️ Available Routes

### POST

-   '/api/notes'

### GET

-   '/api/'
-   '/api/notes'
-   '/api/notes/{id:[0-9]+}'

### PATCH

-   '/api/notes/{id:[0-9]+}'

### DELETE

-   '/api/notes/{id:[0-9]+}'
