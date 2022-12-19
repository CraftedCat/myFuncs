package myFuncs

import (
	"errors"
	"os"
)

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func GetExist() []string {
	var disksClear []string
	//disks := []string{"C:\\", "D:\\", "E:\\", "F:\\", "G:\\", "H:\\", "I:\\", "J:\\", "K:\\", "L:\\", "M:\\", "N:\\",
	//	"O:\\", "P:\\", "Q:\\", "R:\\", "S:\\", "T:\\", "U:\\", "V:\\", "W:\\", "X:\\", "Y:\\", "Z:\\"}

	disks := []string{"C:\\Users\\gdvvl\\go"}

	for i := 0; i < len(disks); i++ {
		disk := disks[i]
		result, err := Exists(disk)
		if err == nil {
			if result {
				disksClear = append(disksClear, disk)
			}
		}
	}
	//fmt.Println(disksClear)
	return disksClear
}
