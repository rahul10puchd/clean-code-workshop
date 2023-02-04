package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// Approach-1 with go testing library
func TestToReadableSize(t *testing.T) {

	got := toReadableSize(4000)
	want := "4 KB"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Approach-2 with go testing library with multiple multiple tests scenerios(Table Driven Test)
type toReadableSizeTest struct {
	nbytes   int
	expected string
}

var toReadableSizeTests = []toReadableSizeTest{
	{1000, "1000 B"},
	{1000 * 1000, "1000 KB"},
	{1000 * 1000 * 1000, "1000 MB"},
	{1000 * 1000 * 1000 * 1000, "1000 GB"},
}

func TestToReadableSizeMultiple(t *testing.T) {

	for _, test := range toReadableSizeTests {
		if output := toReadableSize(int64(test.nbytes)); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// Approach-3 with go testing library with multiple multiple tests scenerios and name for each scenerio(Table Driven Test)
func TestToReadableSizeMultipleApproachTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    int64
		expected string
	}{
		{"byte_return", 125, "125 B"},
		{"kilobyte_return", 1010, "1 KB"},
		{"megabyte_return", 1988909, "1 MB"},
		{"gigabyte_return", 29121988909, "29 GB"},
		{"gigabyte_return", 890929121988909, "890 TB"},
		{"byte_return", 1000, "1000 B"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output := toReadableSize(tc.input)
			if output != tc.expected {
				t.Errorf("input %d, unexpected output: %s", tc.input, output)
			}
		})
	}
}

// Approach-4 with ginko ang gomega
func TestToReadableSizeWithGinkoAndGomega(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "To readable size")
}

var _ = Describe("Main", func() {
	Context("To readable size", func() {
		It("should return in bytes if nbytes is less than 1000", func() {
			result := toReadableSize(125)
			Expect(result).To(Equal("125 B"))
		})

		It("should return in kilo bytes if nbytes is less than MB", func() {
			result := toReadableSize(10000)
			Expect(result).To(Equal("10 KB"))
		})
	})
})

// Approach-5 with ginko ang gomega(Table Driven Test)
var _ = Describe("Main", func() {
	Context("To readable size", func() {

		DescribeTable("To readable size", func(nbytes int64, expectedResult string) {
			actualResult := toReadableSize(nbytes)
			Expect(actualResult).To(Equal(expectedResult))
		},
			Entry("should return in bytes if nbytes is less than 1000", int64(125), "125 B"),
			Entry("should return in kilo bytes if nbytes is less than MB", int64(10000), "10 KB"),
		)
	})

})

func Test_traverseDir(t *testing.T) {
	type args struct {
		hashes     map[string]string
		duplicates map[string]string
		dupeSize   *int64
		entries    []os.FileInfo
		directory  string
	}
	type output struct {
		lengthOfhashes     int
		lengthOfduplicates int
		expectedpanic      bool
	}
	tests := []struct {
		name           string
		args           args
		expectedResult output
	}{
		{
			name: "traverseDir-Test-1-Pass",
			args: args{
				hashes:     map[string]string{},
				duplicates: map[string]string{},
				dupeSize:   new(int64),
				entries:    []os.FileInfo{},
				directory:  "duplicates_files_directory",
			},
			expectedResult: output{
				lengthOfhashes:     2,
				lengthOfduplicates: 2,
			},
		},
		{
			name: "traverseDir-Test-2-Pass",
			args: args{
				hashes:     map[string]string{},
				duplicates: map[string]string{},
				dupeSize:   new(int64),
				entries:    []os.FileInfo{},
				directory:  "/Users/rahul/thoughtworks/clean-code-workshop",
			},
			expectedResult: output{
				lengthOfhashes:     33,
				lengthOfduplicates: 3,
			},
		},
		{
			name: "traverseDir-Test-3-Fail",
			args: args{
				hashes:     map[string]string{},
				duplicates: map[string]string{},
				dupeSize:   new(int64),
				entries:    []os.FileInfo{},
				directory:  "",
			},
			expectedResult: output{
				lengthOfhashes:     0,
				lengthOfduplicates: 0,
				expectedpanic:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedResult.expectedpanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
				traverseDir(tt.args.hashes, tt.args.duplicates, tt.args.dupeSize, tt.args.entries, tt.args.directory)
			}
			//var err error
			tt.args.entries, _ = ioutil.ReadDir(tt.args.directory)
			// if err != nil {
			// 	panic(err)
			// }
			traverseDir(tt.args.hashes, tt.args.duplicates, tt.args.dupeSize, tt.args.entries, tt.args.directory)
			lengthOfHashed := len(tt.args.hashes)
			lengthOfDuplicates := len(tt.args.duplicates)
			log.Println("tt.args.dupeSize", tt.args.dupeSize)
			if lengthOfHashed != tt.expectedResult.lengthOfhashes || lengthOfDuplicates != tt.expectedResult.lengthOfduplicates {
				t.Errorf("input %v, unexpected output:\nlengthOfHashed: %d\nlengthOfDuplicates: %d", tt.args, lengthOfHashed, lengthOfDuplicates)
			}
		})
	}
}
