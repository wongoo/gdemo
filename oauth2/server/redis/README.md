
Run Ouath2 Server:
```
go get -u -v gopkg.in/oauth2.v3/errors
go get -u -v github.com/satori/go.uuid
go get -u -v github.com/codegangsta/inject

go get -u -v github.com/go-fast/oauth2-storage-redis

go build server.go

./server

curl http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read

{"access_token":"FZGYOSWDMQMX23BM5BHCWQ","expires_in":7200,"scope":"read","token_type":"Bearer"}

```
