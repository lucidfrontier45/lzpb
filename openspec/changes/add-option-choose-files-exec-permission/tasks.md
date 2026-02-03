## 1. CLI Argument Parsing

- [x] 1.1 Add --exec flag to CLI argument parser with default value "bootstrap,run.sh"
- [x] 1.2 Store exec files list in pack command struct
- [x] 1.3 Handle empty --exec value to disable exec permissions

## 2. File Parsing and Validation

- [x] 2.1 Implement comma-separated file list parsing function
- [x] 2.2 Add whitespace trimming for file names
- [x] 2.3 Validate non-empty file list after parsing

## 3. Permission Application

- [x] 3.1 function to set file mode to 0755
- [x] 3.2 Apply exec permissions to specified files during pack
- [x] 3.3 Handle files not found gracefully with warning

## 4. Testing

- [x] 4.1 Add unit tests for comma-separated parsing
- [x] 4.2 Add integration test for --exec with single file
- [x] 4.3 Add integration test for --exec with multiple files
- [x] 4.4 Add integration test for default --exec behavior (no --exec flag)
- [x] 4.5 Add integration test for empty --exec value

## 5. Documentation

- [x] 5.1 Update CLI help text for --exec option
- [x] 5.2 Add examples to README
