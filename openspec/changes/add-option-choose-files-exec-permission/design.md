## Context

lzpb is a command-line tool for creating and extracting lzip-based archives. Currently, the tool lacks a native way to set executable permissions on files during packaging. Users who package scripts or binaries need to manually add exec permissions after extraction or use external tools like `chmod`. This creates friction in build/deploy workflows.

## Goals / Non-Goals

**Goals:**
- Add `--exec` option to specify files that should receive executable permissions
- Support comma-separated file list for batch specification
- Apply exec permissions during the packaging phase

**Non-Goals:**
- Recursive directory handling (files must be specified explicitly)
- Unix-only permissions (no Windows equivalent)
- Wildcard/glob pattern matching
- Changing ownership (only exec bit, not owner/group)

## Decisions

1. **Flag format: `--exec`** - Chosen over alternatives like `--executable` or `-x` for clarity and consistency with common CLI patterns.

2. **Comma-separated values** - Simple string parsing allows multiple files without needing to repeat the flag. Alternative of repeatable `--exec file1 --exec file2` was rejected due to added complexity.

3. **String splitting on comma** - Use Go's `strings.Split` to parse the comma-separated list. This handles spaces around commas gracefully by trimming whitespace from each element.

4. **Validation at parse time** - Validate file existence and format errors during CLI argument parsing to fail fast before packaging begins.

5. **Default value: `bootstrap,run.sh`** - For backward compatibility, the --exec option defaults to `bootstrap,run.sh`. This ensures existing workflows continue to work without requiring users to explicitly specify the option. Users can override by providing a custom value or set `--exec=""` to disable.

6. **Permission mode: 0755** - Use standard Unix permission (owner read/write/execute, group/others read/execute). This is the most common use case for packaged executables.

## Risks / Trade-offs

- **Risk**: Users may expect wildcard expansion. → Mitigation: Document that glob patterns are not supported; files must be specified explicitly.
- **Risk**: Relative paths may cause issues if archive is extracted to different directory. → Mitigation: Store paths as-specified; users are responsible for correct path references.
- **Trade-off**: No recursive directory support. → Simpler implementation; users can specify files individually or script if needed.
