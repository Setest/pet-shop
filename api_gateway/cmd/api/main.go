package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/setest/pet-shop/api-gateway/internal/config"
	"github.com/setest/pet-shop/api-gateway/internal/log"
	"github.com/setest/pet-shop/api-gateway/internal/resources"
	//"github.com/PetStores/go-simple/internal/diagnostics"
	//categoryc "github.com/PetStores/go-simple/internal/petstore/category"
	//categorydp "github.com/PetStores/go-simple/internal/petstore/category/withdb"
	//petc "github.com/PetStores/go-simple/internal/petstore/pet"
	//petdp "github.com/PetStores/go-simple/internal/petstore/pet/withdb"
	//"github.com/PetStores/go-simple/internal/resources"
	//"github.com/PetStores/go-simple/internal/restapi"
	//
	//"go.uber.org/zap"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	log.DefaultLogger = log.NewLoggerWithConfig(c.Logger)

	log.Info("Starting the application...")
	log.Info("Reading configuration and initializing resources...")

	rsc, err := resources.New(c.Resources)
	if err != nil {
		log.Panicln("Can't initialize resources:", err)
	}
	defer func() {
		err = rsc.Release()
		if err != nil {
			log.Errorln("Got an error during resources release.", "err", err)
		}
	}()

	//log.Info("Configuring the application units...")
	//catdb := categorydp.New(rsc.DB)
	//cc := categoryc.NewController(catdb)
	//
	//petdb := petdp.New(rsc.DB)
	//pc := petc.NewController(petdb)
	//
	//slogger.Info("Starting the servers...")
	//rapi := restapi.New(slogger, rsc.Config.RESTAPIPort, cc, pc)
	//rapi.Start()
	//
	//diag := diagnostics.New(slogger, rsc.Config.DiagPort, rsc.Healthz)
	//diag.Start()
	//slogger.Info("The application is ready to serve requests.")
	//

	go func() {
		time.Sleep(time.Second * 5)
		log.Info("im wake up")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		log.Infoln("Received a signal", x.String())
		//case err := <-diag.Notify():
		//	log.Error("Received an error from the diagnostics server.", "err", err)
		//case err := <-rapi.Notify():
		//	log.Error("Received an error from the business logic server.", "err", err)
	}

	log.Info("Stopping the servers...")
	//err = rapi.Stop()
	//if err != nil {
	//	slogger.Error("Got an error while stopping the business logic server.", "err", err)
	//}
	//
	//err = diag.Stop()
	//if err != nil {
	//	slogger.Error("Got an error while stopping the diag logic server.", "err", err)
	//}
	//
	log.Info("The app is calling the last defers and will be stopped.")
}
