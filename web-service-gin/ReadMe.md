https://medium.com/@cavdy/creating-restful-api-using-golang-and-postgres-part-1-58fe83c6f1ee
https://github.com/cavdy-play/go_db


https://go.dev/doc/tutorial/web-service-gin

mkdir web-service-gin
cd web-service-gin
go mod init syadav214/web-service-gin
touch main.go

go get .

go run .

go build
then run the executable/binary



GET
curl http://localhost:8080/albums

GET by id
curl http://localhost:8080/albums/2

POST
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
