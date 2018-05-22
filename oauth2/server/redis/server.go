// authors: wangoo
// created: 2018-05-21
// ouath2 server demo based on redis storage

package main

import (
	"log"
	"net/http"
	"github.com/go-fast/oauth2-storage-redis"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/models"
)

// redis token store
func configTokenStore(manager *manage.Manager, redisConfig *redis.Config) {
	manager.MustTokenStorage(redis.NewTokenStore(redisConfig))
}

// redis client  store
func configClientStore(manager *manage.Manager, redisConfig *redis.Config) (err error) {
	clientStore, err := redis.NewClientStore(redisConfig)
	if err != nil {
		return
	}
	err = clientStore.Add(&models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	if err != nil {
		return
	}

	manager.MapClientStorage(clientStore)
	return
}

func main() {
	manager := manage.NewDefaultManager()
	// use redis token store
	config := &redis.Config{
		Addr: "127.0.0.1:6379",
	}
	configTokenStore(manager, config)
	err := configClientStore(manager, config)
	if err != nil {
		log.Fatal(err)
		return
	}

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
