## 1. CLI Argument Parsing

- [ ] 1.1 Add --exec flag to CLI argument parser with default value "bootstrap,run.sh"
- [ ] 1.2 Store exec files list in pack command struct
- [ ] 1.3 Handle empty --exec value to disable exec permissions

## 2. File Parsing and Validation

- [ ] 2.1 Implement comma-separated file list parsing function
- [ ] 2.2 Add whitespace trimming for file names
- [ ] 2.3 Validate non-empty file list after parsing

## 3. Permission Application

- [ ] 3.1 function to set file mode Create to 0755
- [ ] 3.2 Apply exec permissions to specified files during pack
- [ ] 3.3 Handle files not found gracefully with warning

## 4. Testing

- [ ] 4.1 Add unit tests for comma-separated parsing
- [ ] 4.2 Add integration test for --exec with single file
- [ ] 4.3 Add integration test for --exec with multiple files
- [ ] 4.4 Add integration test for default --exec behavior (no --exec flag)
- [ ] 4.5 Add integration test for empty --exec value

## 5. Documentation

- [ ] 5.1 Update CLI help text for --exec option
- [ ] 5.2 Add examples to README
