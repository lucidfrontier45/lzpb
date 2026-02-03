## Context

`lzpb` is a CLI tool that creates zip archives from directories. The tool currently:
- Uses a hardcoded `--exec` flag (default: `bootstrap,run.sh`) to set executable permissions
- Hardcodes LF normalization only for `run.sh` file (line 148 in main.go)

Users need flexibility to normalize line endings for multiple files (shell scripts, Makefiles, Python scripts, config files) without modifying source code. The current hardcoded approach is too rigid.

**Constraints:**
- Must maintain backward compatibility (default behavior unchanged)
- Must follow existing patterns (similar to `--exec` flag implementation)
- No external dependencies (stdlib only)

## Goals / Non-Goals

**Goals:**
- Add `--lf` CLI flag for configurable LF normalization
- Remove hardcoded `run.sh` special-casing
- Maintain backward compatibility with default `run.sh` behavior
- Follow existing code patterns (comma-separated parsing, similar to `--exec`)

**Non-Goals:**
- Pattern matching or glob support (exact filename matching only, like `--exec`)
- Automatic detection of text files (user must specify files explicitly)
- Line ending conversion options beyond CRLF→LF (e.g., no LF→CRLF conversion)

## Decisions

### 1. Flag name: `--lf` over alternatives
**Chosen:** `--lf=<files>`

**Alternatives considered:**
- `--newline`: More generic but could imply other conversions
- `--normalize`: More verbose, less clear about what's normalized
- `--crlf-to-lf`: Too specific, harder to type

**Rationale:** `--lf` is short, clear (indicates target line ending), and mirrors the existing `--exec` flag pattern (short, descriptive name).

### 2. Default value: `run.sh` for backward compatibility
**Chosen:** Default to `run.sh` to preserve existing behavior

**Rationale:** Current code hardcodes `run.sh` conversion. Users may rely on this behavior. Defaulting to `run.sh` ensures no breaking changes while allowing users to override or disable.

### 3. Implementation approach: Mirror `--exec` pattern
**Chosen:** Reuse the comma-separated parsing pattern from `--exec`

**Rationale:**
- Code consistency: similar flags work similarly
- Proven pattern: `parseExecFiles` already handles comma-separated values with whitespace trimming
- Maintainability: one familiar pattern for file selection flags

### 4. Data structure: Reuse existing string parsing
**Chosen:** Add new `lfFiles` field to `packOptions`, reuse `parseExecFiles` logic

**Rationale:** The parsing logic is identical (comma-separated, trim whitespace). No need to duplicate code.

## Risks / Trade-offs

**Risk:** Users might expect glob patterns (e.g., `--lf "*.sh"`)
- **Mitigation:** Explicit in help text that exact filenames are required; pattern matching is out of scope

**Risk:** Default `run.sh` behavior might not be desired by all users
- **Mitigation:** Users can disable with `--lf ""` (empty string); this is documented in usage text

**Trade-off:** Exact filename matching vs. pattern matching
- Chose exact matching for simplicity and consistency with `--exec`
- Pattern matching would add complexity (filepath.Match or glob package) without clear requirement

## Migration Plan

**Deployment:** No migration needed - backward compatible

**Rollback:** If issues arise, revert to hardcoded `run.sh` behavior by removing `--lf` flag

**Testing:**
- Verify default behavior (no flags) still converts `run.sh`
- Test multiple files via `--lf "file1,file2,file3"`
- Test empty `--lf ""` disables conversion
- Test whitespace handling in comma-separated list
