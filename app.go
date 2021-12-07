package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) Initialize(){
	a.Router
	
}