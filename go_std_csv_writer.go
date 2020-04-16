package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{
	{"Feature", "Hardware Required", "Vendor", "Introduced Release", "Asset", "Current Release"},
	{"100,000 simultaneous RPM probes from RPM clients for offload RPM", "MS-MICs", "Juniper", "14.1R3", "5a3e8cfa-cc54-4657-a9aa-b1ecbafe21e9", "12.2R3"},
	{"100,000 simultaneous RPM probes from RPM clients for offload RPM", "MS-MICs", "Juniper", "14.1R3", "7da2e4d1-941d-45e5-b2c0-e602511dc7ee", "12.2R3"},
	{"3GPP Fixed wireless access subscribers on BNGs", "", "Juniper", "17.1R3", "4has54d1-941d-45e5-b2c0-e602511dc7ee", "12.2R3"},
	{"3GPP Fixed wireless access subscribers on BNGs", "", "Juniper", "17.1R3", "7dgsbwe1-941d-45e5-b2c0-e602511dc7ee", "12.2R3"},
	{"Additional DS-Lite features", "MS-MICs", "Juniper", "14.1R3", "941d-45e5-941d-45e5-b2c0-e602511dc7ee", "12.2R3"},
	{"Enhancing the current Inline JFlow scale limits", "MS-MICs", "Juniper", "14.1R3", "941d-45e5-4has54d1-b2c0-e602511dc7ee", "12.2R3"},
	{"ANSI TIA-1057 LLDP - MED Support", "", "Cisco", "15.0SE", "b6161de2-7e50-11ea-bc55-0242ac130003", "14.0SE"},
	{"Auto QoS enhancements for video endpoints", "", "Cisco", "15.0SE", "b6161de2-7e50-11ea-bc55-0242ac130003", "14.0SE"},
}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	//writes to a buffer, when finished you want to write to the file, so you must call csv.Writer.Flush
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
