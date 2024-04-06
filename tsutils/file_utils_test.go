package tsutils

import "testing"

func TestGetFileExtension(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
		expected FileType
		err      error
	}{
		// SUCCESSES
		{name: "a scala file",
			filePath: "dir/file.scala",
			expected: FTScala},
		{name: "a go file",
			filePath: "dir/file.go",
			expected: FTGo},
		{name: "a go file in a hidden directory",
			filePath: ".dir/file.go",
			expected: FTGo},

		// FAILURES
		{name: "an incorrect file path",
			filePath: "filewithoutextension",
			expected: FTNone,
			err:      FilePathIncorrectErr},
		{name: "an incorrect file path in dir",
			filePath: "dir/filewithoutextension",
			expected: FTNone,
			err:      FilePathIncorrectErr},
		{name: "a file with several dots",
			filePath: "dir/file.go.temp",
			expected: FTNone,
			err:      FilePathIncorrectErr},
		{name: "a dot file",
			filePath: ".file",
			expected: FTNone,
			err:      FilePathIncorrectErr},
		{name: "a file in a hidden directory",
			filePath: ".dir/file",
			expected: FTNone,
			err:      FilePathIncorrectErr},
		{name: "a go file without a name",
			filePath: ".go",
			expected: FTNone,
			err:      FileNameIncorrectErr},
		{name: "a go file without a name in a hidden dir",
			filePath: "dir1/.dir2/.go",
			expected: FTNone,
			err:      FileNameIncorrectErr},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFileExtension(tt.filePath)
			if got != tt.expected {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
			if err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", err, tt.err)
			}
		})
	}

}
