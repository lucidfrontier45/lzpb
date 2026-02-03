## ADDED Requirements

### Requirement: CLI accepts --exec option
The CLI SHALL accept an `--exec` option that specifies one or more files to receive executable permissions.

#### Scenario: Single file with --exec
- **WHEN** user runs `lzpb pack --exec script.sh`
- **THEN** the file `script.sh` is marked with executable permission in the archive

#### Scenario: Multiple files with comma-separated list
- **WHEN** user runs `lzpb pack --exec "script.sh,bin/app,tools/build"`
- **THEN** all three files (`script.sh`, `bin/app`, `tools/build`) are marked with executable permissions in the archive

#### Scenario: --exec option with spaces in file list
- **WHEN** user runs `lzpb pack --exec "script.sh, bin/app ,tools/build"`
- **THEN** whitespace around file names is trimmed and all three files receive executable permissions

#### Scenario: Default value when --exec is not specified
- **WHEN** user runs `lzpb pack` without --exec
- **THEN** the files `bootstrap` and `run.sh` receive executable permissions by default

#### Scenario: Empty --exec disables default behavior
- **WHEN** user runs `lzpb pack --exec ""`
- **THEN** no files receive executable permissions

### Requirement: Files receive 0755 permissions
The system SHALL set file mode to 0755 (owner read/write/execute, group/others read/execute) for files specified via --exec.

#### Scenario: Exec permission includes owner execute
- **WHEN** user specifies a file via --exec
- **THEN** the file mode includes owner execute permission (bit 0100)

#### Scenario: Exec permission includes group execute
- **WHEN** user specifies a file via --exec
- **THEN** the file mode includes group execute permission (bit 0010)

#### Scenario: Exec permission includes others execute
- **WHEN** user specifies a file via --exec
- **THEN** the file mode includes others execute permission (bit 0001)

### Requirement: Invalid --exec input is rejected
The system SHALL validate --exec input and report errors for invalid file specifications.

#### Scenario: Empty --exec value is rejected
- **WHEN** user runs `lzpb pack --exec ""`
- **THEN** the system reports an error indicating no files were specified

#### Scenario: Non-existent file shows warning
- **WHEN** user runs `lzpb pack --exec "nonexistent.sh"`
- **THEN** the system MAY warn that the file does not exist but SHALL still create the archive
