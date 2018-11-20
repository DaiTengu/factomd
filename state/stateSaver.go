// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package state

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/FactomProject/factomd/common/primitives"
)

type StateSaverStruct struct {
	FastBoot         bool
	FastBootLocation string
	TmpDBHt          uint32
	TmpState []byte
	Mutex    sync.Mutex
	Stop     bool
}

//To be increased whenever the data being saved changes from the last version
const version = 8

func (sss *StateSaverStruct) StopSaving() {
	sss.Mutex.Lock()
	defer sss.Mutex.Unlock()
	sss.Stop = true
}

func (sss *StateSaverStruct) SaveDBStateList(ss *DBStateList, networkName string) error {
	//For now, to file. Later - to DB
	if sss.Stop == true {
		return nil
	}
	sss.Mutex.Lock()
	defer sss.Mutex.Unlock()

	hsb := int(ss.GetHighestSavedBlk())
	//Save only every FastSaveRate states

	if (hsb+1)%ss.State.FastSaveRate == 0 {
		//Marshal state for future saving
		b, err := ss.MarshalBinary()
		if err != nil {
			return err
		}
		//adding an integrity check
		h := primitives.Sha(b)
		b = append(h.Bytes(), b...)
		sss.TmpState = b
		sss.TmpDBHt = ss.State.LLeaderHeight
	}

	if hsb%ss.State.FastSaveRate == 0  &&  len(sss.TmpState) > 0 {
		//Actually save data from previous cached state to prevent dealing with rollbacks
		err := SaveToFile(sss.TmpState, NetworkIDToFilename(networkName, sss.FastBootLocation))
		if err != nil {
			return err
		}
		sss.TmpState = nil
	}

	return nil
}

func (sss *StateSaverStruct) DeleteSaveState(networkName string) error {
	return DeleteFile(NetworkIDToFilename(networkName, sss.FastBootLocation))
}

func (sss *StateSaverStruct) LoadDBStateList(ss *DBStateList, networkName string) error {
	filename := NetworkIDToFilename(networkName, sss.FastBootLocation)
	fmt.Println(ss.State.FactomNodeName, "Loading from", filename)
	b, err := LoadFromFile(filename)
	if err != nil {
		return nil
	}
	if b == nil {
		return nil
	}
	h := primitives.NewZeroHash()
	b, err = h.UnmarshalBinaryData(b)
	if err != nil {
		return nil
	}
	h2 := primitives.Sha(b)
	if h.IsSameAs(h2) == false {
		fmt.Printf("LoadDBStateList - Integrity hashes do not match!")
		return nil
		//return fmt.Errorf("Integrity hashes do not match")
	}

	return ss.UnmarshalBinary(b)
}

func NetworkIDToFilename(networkName string, fileLocation string) string {
	file := fmt.Sprintf("FastBoot_%s_v%v.db", networkName, version)
	if fileLocation != "" {
		i := len(fileLocation) - 1
		if fileLocation[i] == '/' {
			fileLocation = fileLocation[:i] // trim trailing '/'
		}
		return fmt.Sprintf("%v/%v", fileLocation, file)
	}
	return file
}

func SaveToFile(b []byte, filename string) error {
	fmt.Fprintf(os.Stderr, "Saving %s\n", filename)
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return err
	}
	return nil
}

func LoadFromFile(filename string) ([]byte, error) {
	fmt.Fprintf(os.Stderr, "Load state from %s\n", filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteFile(filename string) error {
	return os.Remove(filename)
}
