---
name: submit-exercise
description: Checks that an exercise of this repository is finished (tests pass, the documentation and comments are right, it is listed in README.md), and then commits it and pushes it to main. Use only when the user asks to submit an exercise.
argument-hint: "[folder number or name, e.g. 21]"
disable-model-invocation: true
---

# Submit exercise

Check that an exercise is finished and, only if everything is right, commit it and push it to `main`. The user asked for this by running the command, so committing and pushing is allowed here.

Follow the rules of `CLAUDE.md`. **If any check fails, report the problem and stop: do not stage, commit or push anything.** The only file this skill may edit is `README.md` (step 5).

## Steps

1. **Find the folder.** The argument is the folder number (`21`) or part of its name. Match it with the folders named `NN_LC<number>_<name>`.
   - If there is no argument, use the folder whose solution `.go` file was modified most recently (not the `_test.go`).
   - If more than one folder matches or none does, tell the user and ask.
2. **Check git before doing anything else.**
   - The current branch must be `main` (`git branch --show-current`).
   - Nothing may be staged already (`git diff --cached --name-only` must be empty), so the commit has only this exercise.
   - Run `git fetch origin main`. If the local branch is behind `origin/main`, stop and tell the user. Do not pull, rebase or force anything.
3. **Tests.**
   - The solution must not be the skeleton (it must not have `// TODO: implement the solution.`).
   - `gofmt -l` must print nothing, and `go vet` and `go test -count=1 -timeout 60s ./<folder>/` must pass.
4. **Documentation and grammar** of the solution `.go` file (not the tests):
   - Every function has a doc comment that starts with its name and follows `CLAUDE.md`: what it does, how the algorithm works, and `Time complexity` and `Space complexity` with `n` explained.
   - The doc comments and the other comments have no grammar or spelling errors.
   - Do not edit anything. `// NOTE:` comments are the user's: report their typos as warnings, but they do not stop the submission.
   - List every problem with its file, line and a proposed fix, and stop.
5. **README.** Check that `README.md` has a row for this exercise in the table of its NeetCode category.
   - If it is missing, add it as `CLAUDE.md` says (creating the category section in NeetCode's order if it does not exist).
   - Check that every local link of the README points to an existing file.
6. **Stage** only the exercise folder and the README: `git add -- <folder> README.md`. Never `git add .` or `git add -A`. Then check with `git diff --cached --name-only` that only those paths are staged. If there is nothing to commit, say so and stop (do not push).
7. **Commit** with this one-line message, using the typographic quotes `“` and `”` (not straight quotes):

   ```text
   feat: add “<Problem name>” exercise (LeetCode <number>)
   ```

   - `<number>` is the LeetCode number (from `leetcode<number>.md`).
   - `<Problem name>` is the name written in the package comment of the solution, between quotes (`// Package x solves the "<Problem name>" problem.`). If it is missing, use the title of the statement.
   - No body, and **no `Co-Authored-By` line or any other trailer**.
   - Pass the message with `git commit -F -` and a heredoc, so the quotes are kept.
8. **Push** with `git push origin main`. Never use `--force`. If the push fails, report the error and stop: do not retry or change the remote.

## Final answer

Reply in the language of the user's message. If there is no clear signal (for example, only the command), use English. Keep it short and say:
- the checks that were made and that they passed,
- the files that went into the commit,
- the commit hash and message, and the result of the push.

If the process stopped, say at which step and why, with the list of problems.
