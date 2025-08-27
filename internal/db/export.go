package db

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/wallarm/gotestwaf/internal/payload/encoder"
)

func (db *DB) ExportPayloads(payloadsExportFile string, includeRequestDetails bool) error {
	csvFile, err := os.Create(payloadsExportFile)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write headers
	headers := []string{
		"Payload",
		"Check Status",
		"Response Code",
		"Placeholder",
		"Encoder",
		"Set",
		"Case",
		"Test Result",
	}
	
	if includeRequestDetails {
		headers = append(headers, []string{
			"HTTP Method",
			"Request URL",
			"Request Headers",
			"Request Body",
		}...)
	}
	
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	for _, blockedTest := range db.blockedTests {
		p := blockedTest.Payload
		e := blockedTest.Encoder
		testResult := "passed"

		ep, err := encoder.Apply(e, p)
		if err != nil {
			return err
		}

		if isFalsePositiveTest(blockedTest.Set) {
			testResult = "failed"
		}

		// Prepare data row
		dataRow := []string{
			ep,
			"blocked",
			strconv.Itoa(blockedTest.ResponseStatusCode),
			blockedTest.Placeholder,
			blockedTest.Encoder,
			blockedTest.Set,
			blockedTest.Case,
			testResult,
		}
		
		if includeRequestDetails {
			dataRow = append(dataRow, []string{
				blockedTest.HTTPMethod,
				blockedTest.RequestURL,
				blockedTest.RequestHeaders,
				blockedTest.RequestBody,
			}...)
		}
		
		err = csvWriter.Write(dataRow)
		if err != nil {
			return err
		}
	}

	for _, passedTest := range db.passedTests {
		p := passedTest.Payload
		e := passedTest.Encoder
		testResult := "failed"

		ep, err := encoder.Apply(e, p)
		if err != nil {
			return err
		}

		if isFalsePositiveTest(passedTest.Set) {
			testResult = "passed"
		}

		// Prepare data row
		dataRow := []string{
			ep,
			"passed",
			strconv.Itoa(passedTest.ResponseStatusCode),
			passedTest.Placeholder,
			passedTest.Encoder,
			passedTest.Set,
			passedTest.Case,
			testResult,
		}
		
		if includeRequestDetails {
			dataRow = append(dataRow, []string{
				passedTest.HTTPMethod,
				passedTest.RequestURL,
				passedTest.RequestHeaders,
				passedTest.RequestBody,
			}...)
		}
		
		err = csvWriter.Write(dataRow)
		if err != nil {
			return err
		}
	}

	for _, naTest := range db.naTests {
		p := naTest.Payload
		e := naTest.Encoder

		ep, err := encoder.Apply(e, p)
		if err != nil {
			return err
		}

		// Prepare data row
		dataRow := []string{
			ep,
			"unresolved",
			strconv.Itoa(naTest.ResponseStatusCode),
			naTest.Placeholder,
			naTest.Encoder,
			naTest.Set,
			naTest.Case,
			"unknown",
		}
		
		if includeRequestDetails {
			dataRow = append(dataRow, []string{
				naTest.HTTPMethod,
				naTest.RequestURL,
				naTest.RequestHeaders,
				naTest.RequestBody,
			}...)
		}
		
		err = csvWriter.Write(dataRow)
		if err != nil {
			return err
		}
	}

	return nil
}
