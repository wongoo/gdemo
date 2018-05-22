
Run Ouath2 Server:
```
go get -u -v gopkg.in/oauth2.v3/errors
go get -u -v github.com/satori/go.uuid
go get -u -v github.com/codegangsta/inject

# BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support
go get -u -v github.com/tidwall/buntdb


go build server.go

./server

curl http://localhost:9096/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read

{"access_token":"FZGYOSWDMQMX23BM5BHCWQ","expires_in":7200,"scope":"read","token_type":"Bearer"}

```
