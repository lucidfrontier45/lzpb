## Why

Currently, lzpb does not provide a convenient way to set executable permissions on files during packaging. Users must manually set exec permissions after extraction or use external tools. Adding a `--exec` option allows users to specify which files should receive executable permissions directly during the build/pack process.

## What Changes

- Add `--exec` command-line option to specify files that should be given executable permissions
- Support comma-separated list of files for batch permission assignment
- Default value of `bootstrap,run.sh` for backward compatibility
- Files listed will have their mode set to include exec bits during packaging

## Capabilities

### New Capabilities
- `exec-permission`: Allows specifying files to receive executable permissions via `--exec` option with comma-separated file list; defaults to `bootstrap,run.sh` for backward compatibility

### Modified Capabilities
- (none)

## Impact

- CLI argument parsing module (new `--exec` flag)
- File mode handling in packaging logic
- User-facing documentation updated
