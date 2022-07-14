# sample Gin API
    https://www.youtube.com/watch?v=ZI6HaPKHYsg

## run database postgre

    docker-compose up -d

## run API
    
    go  run main.go
    
## add data

    curl  POST 'localhost:8080' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "raul",
        "email": "raul.arrascue@gmail.com",
        "password": "123456"
    }'