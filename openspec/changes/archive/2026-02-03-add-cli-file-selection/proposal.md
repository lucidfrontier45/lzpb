## Why

Currently, `lzpb` hardcodes newline conversion for only `run.sh` (line 148). This is too rigid - users need to normalize line endings for other shell scripts, text files, or any file where CRLF would cause issues (e.g., Makefiles, Python scripts, config files). Adding a CLI option provides flexibility without code changes.

## What Changes

- Add new `--lf` CLI flag accepting comma-separated list of filenames (similar to existing `--exec` flag)
- Modify zip packing logic to convert CRLF to LF for any file matching names in `--lf` list
- Remove hardcoded `run.sh` special-casing (now handled via `--lf` flag, with `run.sh` included in default)
- Update usage/help text to document new flag

## Capabilities

### New Capabilities
- `cli-file-selection`: Add configurable file selection for newline normalization via CLI flags

### Modified Capabilities
- `zip-packing`: REQUIREMENTS changed - newline conversion is now configurable instead of hardcoded to `run.sh` only

## Impact

**Affected Code**:
- `main.go`: Add new flag, modify `zipDirectory` function, update help text

**API Changes**:
- New CLI flag: `--lf=<files>` (comma-separated, defaults to `run.sh`)
- Backward compatible: existing behavior preserved via default value

**Dependencies**: None (uses only stdlib)
