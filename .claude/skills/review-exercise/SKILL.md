---
name: review-exercise
description: Runs the tests of an exercise of this repository against the user's solution. If all the tests pass, it says so; if any fails, it explains why and gives hints (never the solution) based on the existing code. Use when the user asks to review an exercise, to check why the tests fail, or wants to know if their solution works.
argument-hint: "[folder number or name, e.g. 19]"
---

# Review exercise

Check the user's solution of an exercise and help them to fix it. Follow the section "Helping with a solution" of `CLAUDE.md`: give **hints and examples only, never the complete solution**, and do not modify any file.

## Steps

1. **Find the folder.** The argument is the folder number (`19`) or part of its name. Match it with the folders named `NN_LC<number>_<name>`.
   - If there is no argument, use the folder whose solution `.go` file was modified most recently (not the `_test.go`).
   - If more than one folder matches or none does, tell the user and ask.
2. **Read** the statement (`leetcode<number>.md`), the solution `.go` and the `_test.go`.
   - If the solution is still the skeleton (it has `// TODO: implement the solution.`), say so and stop: there is nothing to review yet.
3. **Run** `gofmt -l`, `go vet` and `go test -count=1 -timeout 60s ./<folder>/` (use `-v` to see which tests fail).
4. **If everything passes**, say so in one or two lines. Do not add improvements, documentation or README changes unless the user asks.
5. **If something fails**, find out why before answering:
   - Group the failures by root cause. Many failing tests are often one single bug.
   - For each cause, trace by hand a **concrete failing input** on the user's code (step by step, with the values of the variables), so the user sees where it goes wrong. Running a copy of their code with prints under `/tmp` is fine; delete it afterwards.
   - Compile errors, `go vet` warnings, panics and timeouts are failures too: explain what they mean.
   - If a test reports `test setup error`, the problem is in the test, not in the solution: say so.
   - If something looks like an error but is correct, say so. Do not blame correct logic.
6. **Give hints** from the least to the most direct:
   - First, the question that points to the problem ("what is the value of `x` when...?").
   - Then, the idea behind the fix, in words and with a small example. No complete code. Short snippets of the Go syntax or functions the user may not know are fine.
   - The user can ask for a more direct hint afterwards.
7. **Never modify** the solution or the tests, and do not leave temporary files in the repository.

## Final answer

Reply in the language of the user's message. If there is no clear signal (for example, only the command), use English. Keep it short:
- the result: how many tests pass and fail, and which ones fail (a short list),
- the cause of the failures, with the trace of one failing input,
- the hints,
- how to run the tests again (`go test -v ./<folder>/`).
