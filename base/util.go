package base

import (
	"os"
	"restapi-go/config"
	"restapi-go/logging"
	"strconv"
)

func SavePID(cf *config.AppConf) {
	log := logging.Log()
	pid := os.Getpid()
	pidFilePath := cf.Log.Dir + "/mypid"
	file, err := os.Create(pidFilePath)
	if err != nil {
		log.Info("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Error("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}
}
