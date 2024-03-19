package handlers

import (
	"os"
	"testing"
)

func TestGetMostActiveCookieMultiple(t *testing.T) {
	// Prepare a CSV file for testing
	csvData := []byte(`cookie,timestamp
AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00
`)
	WrongCsvData := []byte(`cookiess,timestamp
AtY0laUfhglK3lC72018-12-09T14:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
AtY0laUfhglK3lC7,2018-12-09T06:19:00+00:00
SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00
`)
	WrongTimeCsvData := []byte(`cookiess,timestamp
AtY0laUfhglK3lC7,Jan 2018
SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00

`)

	csvFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary CSV file: %v", err)
	}
	WrongCsvFile, err := os.CreateTemp("", "test_wrong.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary CSV file: %v", err)
	}
	WrongTimeCsvFile, err := os.CreateTemp("", "test_wrong_time.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary CSV file: %v", err)
	}
	defer os.Remove(csvFile.Name())
	defer csvFile.Close()

	if _, err := csvFile.Write(csvData); err != nil {
		t.Fatalf("Failed to write data to temporary CSV file: %v", err)
	}
	if _, err := WrongCsvFile.Write(WrongCsvData); err != nil {
		t.Fatalf("Failed to write data to temporary CSV file: %v", err)
	}
	if _, err := WrongTimeCsvFile.Write(WrongTimeCsvData); err != nil {
		t.Fatalf("Failed to write data to temporary CSV file: %v", err)
	}

	// Test cases
	tests := []struct {
		name           string
		logFile        string
		specifiedDate  string
		expectedResult []string
		expectedError  bool
	}{
		{
			name:           "Valid and incorrect date",
			logFile:        csvFile.Name(),
			specifiedDate:  "2019-12-09",
			expectedResult: []string{},
			expectedError:  false,
		},
		{
			name:           "Valid Date and Single output",
			logFile:        csvFile.Name(),
			specifiedDate:  "2018-12-09",
			expectedResult: []string{"AtY0laUfhglK3lC7"},
			expectedError:  false,
		},
		{
			name:           "Valid and correct date",
			logFile:        csvFile.Name(),
			specifiedDate:  "2018-12-08",
			expectedResult: []string{"SAZuXPGUrfbcn5UA", "4sMM2LxV07bPJzwf", "fbcn5UAVanZf6UtG"},
			expectedError:  false,
		},
		{
			name:           "Invalid Date",
			logFile:        csvFile.Name(),
			specifiedDate:  "Jan 2018",
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			name:           "No File Provided",
			logFile:        "",
			specifiedDate:  "2018-12-08",
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			name:           "Improper CSV Data",
			logFile:        WrongCsvFile.Name(),
			specifiedDate:  "2018-12-08",
			expectedResult: []string{},
			expectedError:  true,
		},
		{
			name:           "Improper Timestamp Data",
			logFile:        WrongTimeCsvFile.Name(),
			specifiedDate:  "2018-12-08",
			expectedResult: []string{},
			expectedError:  true,
		},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetMostActiveCookie(tc.logFile, tc.specifiedDate)

			if (err != nil) != tc.expectedError {
				t.Fatalf("Expected error: %v, got: %v", tc.expectedError, err != nil)
			}
			if len(result) != len(tc.expectedResult) {
				t.Errorf("Arrays have different lengths. Expected: %d, Actual: %d", len(tc.expectedResult), len(result))
				return
			}

		})
	}
}
