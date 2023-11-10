package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"

	pkgerrors "github.com/pkg/errors"
)

// InitSetMap initialises the set mapping for the decklist parser
func (i impl) InitSetMap() error {

	sets, err := i.dbApi.GetSets()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	n := len(sets)
	if n == 0 {
		return fmt.Errorf("no sets found")
	}
	log.Printf("%d sets found", n)

	setData := make([][]string, n)
	for index, set := range sets {
		setData[index] = make([]string, 2)
		setData[index][0] = set.PtcgoCode
		setData[index][1] = set.ID
	}
	log.Printf("csv data recorded")

	err = ensureBaseDir("./setmap/setmap.csv")
	if err != nil {
		return err
	}

	file, err := os.Create(setMapPath)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	log.Printf("file created")

	defer file.Close()

	wr := csv.NewWriter(file)
	err = wr.WriteAll(setData)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	log.Printf("data written to file")

	return nil
}

// shamelessly taken from stackoverflow
func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
