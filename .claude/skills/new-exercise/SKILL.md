---
name: new-exercise
description: Prepares a new exercise folder of this repository. Fixes the statement (leetcode<number>.md), creates the Go skeleton and the tests if they do not exist yet, and verifies the tests with a reference solution. Use when the user starts a new exercise or asks to prepare a folder.
argument-hint: "[folder number or name, e.g. 19]"
---

# New exercise

Prepare an exercise folder following the rules of `CLAUDE.md` (statement format, skeleton and tests). Do not repeat them here: read `CLAUDE.md` first and follow it.

## Steps

1. **Find the folder.** The argument is the folder number (`19`) or part of its name. Match it with the folders named `NN_LC<number>_<name>`.
   - If there is no argument, look for folders that have a `leetcode<number>.md` but are missing the solution `.go` or the `_test.go`, and work on those.
   - If more than one folder matches or none does, tell the user and ask.
2. **Read the statement** (`leetcode<number>.md`) and any image next to it.
3. **Fix the statement** only if it is still the raw text copied from LeetCode.
4. **Skeleton.** Create `<package>.go` only if the folder has no solution `.go` file. Never overwrite.
5. **Tests.** Create `<package>_test.go` only if the folder has no `_test.go` file. Never overwrite.
   - If the only thing missing is one of them, create just that one.
   - If both already exist, say so and stop.
6. **Verify the tests** in a temporary folder under `/tmp` with a reference solution, as `CLAUDE.md` says. Delete the temporary folder afterwards. The reference solution must not stay in the repository.
7. **Check the repository**: `gofmt -l` and `go vet` on the folder. The skeleton compiles, so `go vet` must pass.
8. **Do not solve the exercise** and do not add it to the README, unless the user asks.

## Final answer

Reply in the language of the user's message. If there is no clear signal (for example, only the command), use English. Keep it short and mention:
- which files were created and which already existed,
- the function signature the user has to implement and the package name,
- what the tests cover (a few bullet points) and which cases fail with the skeleton,
- how to run them, with the command alone in its own fenced code block (```` ```bash ````), so the user can copy it directly (for example, `go test -v ./18_LC567_permutationinstring/`, with the real folder name), and
- the best time and space complexities achievable for the problem.

Give hints on how to solve the problem only if the user asks for them.
