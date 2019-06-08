package main

import (
    "github.com/emikohmann/go-structuring/domain-driven-full/repository/db"
    "github.com/emikohmann/go-structuring/domain-driven/domain/ban"
    "github.com/emikohmann/go-structuring/domain-driven/http"
    "github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
)

func main() {
    //repo := ds.New()
    repo := db.New()
    serv := ban.NewService(repo)

    handler := http.NewHandler(serv)
    router := mlhandlers.DefaultMeliRouter()

    router.POST("/ban", handler.ApplyBAN)
    router.DELETE("/ban", handler.RemoveBAN)

    router.Run()
}
