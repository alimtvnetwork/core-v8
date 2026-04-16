---
name: macos runner root cause analysis
description: Root cause and fix for macOS run.ps1 terminal escape output and misleading runtime failure reports.
type: report
---
Root cause analysis for `run.ps1 -tc` on macOS:

1. `scripts/DashboardTheme.psm1` used an OSC 11 terminal background-color probe on Linux/macOS by writing `ESC ] 11 ; ? ESC \` directly to the console.
2. In non-interactive or partially compatible PowerShell hosts on macOS, that probe leaked into the shell as literal text, producing errors like:
   - `]11;rgb:1919/1a1a/1b1b\ : The term ... is not recognized`
   - stray `]` command errors after the script completed.
3. The real problem was not Go test execution itself; it was dashboard theme auto-detection running too aggressively for the host.
4. A second diagnosis issue existed in coverage reporting: compile-blocked packages were also inserted into `runtime-failures.txt`, which mixed compile failures with true runtime crashes and obscured root cause analysis.

Implemented fix:

- Guard OSC 11 probing behind interactive console checks:
  - `Environment.UserInteractive`
  - `Console.In/Out` availability
  - input/output not redirected
- Always restore `stty` state in a `finally` block.
- Keep fallback theme detection so the dashboard still works when probing is skipped.
- Stop adding compile-check blocked packages to runtime failure maps.
- Filter runtime failure reporting to exclude blocked packages and keep that report focused on true runtime crashes / missing profiles.

How to diagnose this class of issue in future:

- If terminal output shows raw `]11;rgb...` or similar OSC fragments, inspect `scripts/DashboardTheme.psm1` first.
- If `runtime-failures.txt` lists packages that also appear in blocked compile packages, inspect `scripts/CoverageCompileCheck.psm1` and `scripts/CoverageReportJson.psm1` for classification leakage.
- Root cause reports must separate:
  - terminal-host issues,
  - compile/setup failures,
  - runtime crashes.

Expected outcome after fix:

- `run.ps1 -tc` should no longer emit stray OSC command text on macOS.
- `runtime-failures.txt` should only show actual runtime/package execution failures, not compile-blocked packages.
- Build root cause should be read from `blocked-packages.txt` / `build-errors.txt`, while runtime crashes remain in `runtime-failures.txt`.
