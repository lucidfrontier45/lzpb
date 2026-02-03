## MODIFIED Requirements

### Requirement: Hardcoded run.sh newline conversion
The system previously converted line endings to LF only for the hardcoded filename `run.sh`. This requirement is removed in favor of configurable file selection via the --lf option.

**Original Behavior:**
- Only `run.sh` received LF normalization
- No way to specify additional files
- No way to disable the behavior

**New Behavior:**
- LF normalization is configurable via --lf flag
- Defaults to `run.sh` for backward compatibility
- Can be extended to multiple files
- Can be disabled with empty string

#### Scenario: Backward compatibility maintained
- **WHEN** user runs `lzpb pack` without any flags
- **THEN** `run.sh` still receives LF normalization (default behavior preserved)

#### Scenario: Multiple files can be normalized
- **WHEN** user runs `lzpb pack --lf "run.sh,bootstrap.sh,config.mk"`
- **THEN** all three files receive LF normalization

#### Scenario: Normalization can be disabled
- **WHEN** user runs `lzpb pack --lf ""`
- **THEN** no files receive LF normalization
