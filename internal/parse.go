package internal

import (
	"github.com/FerdinaKusumah/rumpi/model"
	"github.com/FerdinaKusumah/rumpi/utils"
	"flag"
	"os"
)

type Option struct {
	Source  string
	ListApi []*model.ConfigFile
}

func ParseOption() *Option {
	var opt = new(Option)

	flag.StringVar(&opt.Source, "source", "./config.json", "Please Input JSON File Source")
	flag.StringVar(&opt.Source, "s", "./config.json", "Please Input JSON File Source")
	flag.Parse()

	// if source is empty
	if opt.Source == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	opt.ListApi = utils.LoadConfig(opt.Source)
	return opt
}
