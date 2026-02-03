## ADDED Requirements

### Requirement: CLI accepts --lf option
The CLI SHALL accept a `--lf` option that specifies one or more files to receive LF newline normalization (CRLF → LF conversion).

#### Scenario: Single file with --lf
- **WHEN** user runs `lzpb pack --lf script.sh`
- **THEN** the file `script.sh` has its CRLF line endings converted to LF in the archive

#### Scenario: Multiple files with comma-separated list
- **WHEN** user runs `lzpb pack --lf "script.sh,config.py,Makefile"`
- **THEN** all three files (`script.sh`, `config.py`, `Makefile`) have CRLF line endings converted to LF in the archive

#### Scenario: --lf option with spaces in file list
- **WHEN** user runs `lzpb pack --lf "script.sh, config.py ,Makefile"`
- **THEN** whitespace around file names is trimmed and all three files receive LF normalization

#### Scenario: Default value when --lf is not specified
- **WHEN** user runs `lzpb pack` without --lf
- **THEN** the file `run.sh` receives LF normalization by default

#### Scenario: Empty --lf disables default behavior
- **WHEN** user runs `lzpb pack --lf ""`
- **THEN** no files receive LF normalization

### Requirement: LF normalization converts CRLF to LF
The system SHALL replace all CRLF (`\r\n`) line endings with LF (`\n`) for files specified via --lf.

#### Scenario: CRLF file is converted
- **WHEN** a file specified via --lf contains CRLF line endings
- **THEN** all CRLF sequences are replaced with LF in the archive

#### Scenario: Mixed line endings are normalized
- **WHEN** a file specified via --lf contains mixed CRLF and LF line endings
- **THEN** all CRLF sequences are replaced with LF, leaving existing LF unchanged

#### Scenario: LF-only files are unchanged
- **WHEN** a file specified via --lf contains only LF line endings
- **THEN** the file content remains unchanged in the archive
