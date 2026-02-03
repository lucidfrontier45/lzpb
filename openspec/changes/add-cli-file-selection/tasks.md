## 1. CLI Flag Setup

- [ ] 1.1 Add `lfFiles` field to `packOptions` struct with default value "run.sh"
- [ ] 1.2 Add `--lf` flag in `init()` function with help text
- [ ] 1.3 Update `main()` to pass `lfFiles` to `zipDirectory` function (parse via existing `parseExecFiles` logic)

## 2. Core Implementation

- [ ] 2.1 Modify `zipDirectory()` function signature to accept `lfFiles` parameter
- [ ] 2.2 Remove hardcoded `RUN_SH` constant check (line 148)
- [ ] 2.3 Replace hardcoded logic with loop over `lfFiles` list (similar to `execFiles` loop)
- [ ] 2.4 Implement CRLF→LF conversion for files matching names in `lfFiles` list

## 3. Testing & Verification

- [ ] 3.1 Test default behavior (no flags) - verify `run.sh` still gets LF normalization
- [ ] 3.2 Test multiple files: `--lf "run.sh,bootstrap.sh,config.mk"`
- [ ] 3.3 Test empty flag: `--lf ""` disables normalization
- [ ] 3.4 Test whitespace handling: `--lf "run.sh, config.sh ,other.sh"`
- [ ] 3.5 Run `golangci-lint run` to verify code quality
- [ ] 3.6 Run `go test ./...` to verify tests pass
