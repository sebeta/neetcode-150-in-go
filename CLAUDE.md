# NeetCode 150 in Go

Personal study repository: NeetCode 150 problems solved in Go. The goal is to **learn**, so the user writes the solutions. Go module: `github.com/sebeta/neetcode-150-in-go`.

## Language

- Reply in the language of the user's message. If there is no clear signal (for example, only a command), use English.
- Everything written in the repository is in **English**: code, comments, doc comments, `.md` files, test names and error messages, README.

## Layout

One folder per exercise, named `NN_LC<number>_<name>` (for example `17_LC424_longestrepeatingcharacter`). `NN` is the order in which the exercises were done. Inside:

| File | What it is |
|---|---|
| `leetcode<number>.md` | Problem statement |
| `<package>.go` | Solution (package name is the lowercase name, no separators) |
| `<package>_test.go` | Tests |
| images (`.png`, `.jpg`) | Used by the statement, linked with a relative path |

## Automatic rule: new exercise

When a folder has a `leetcode<number>.md` file, create what is missing, **never overwriting what already exists**:

1. **Statement (`.md`)**: if it is still the raw text copied from LeetCode, rewrite it in the format below.
2. **Skeleton (`<package>.go`)**: only if the folder has no solution `.go` file yet.
3. **Tests (`<package>_test.go`)**: only if the folder has no `_test.go` file yet.

In the final answer, mention the best time and space complexities achievable for the problem, but give hints on how to solve it only if the user asks for them.

Do not solve the exercise. Do not add the exercise to the README unless the user asks.

### Statement format (`leetcode<number>.md`)

```markdown
# LeetCode <number>: <Title>

**Difficulty:** Easy|Medium|Hard

<statement, with names and values in `backticks`>

## Examples

### Example 1

- **Input:** `...`
- **Output:** `...`
- **Explanation:** ...

## Constraints

- `1 <= n <= 10^5`

## Follow up
```

- Remove the website noise: the leading `<number>.` in the title, "Topics", "premium lock icon", "Companies", "Hint", extra blank lines and trailing spaces.
- Fix exponents lost when copying (`105` is `10^5`), one `###` section per example, and keep images with a relative link (`![alt](./image.png)`).
- Keep the difficulty written in the file. If it is missing, look it up.
- Omit `## Follow up` if the problem has none.

### Skeleton (`<package>.go`)

```go
// Package <package> solves the "<Problem name>" problem.
package <package>

func <LeetCodeName>(...) ... {
	// TODO: implement the solution.
	return <zero value>
}
```

- Use the LeetCode function name, exported (`twoSum` becomes `TwoSum`), with the same parameters.
- It must compile. No documentation yet: the user asks for it once the solution is done.

### Tests (`<package>_test.go`)

Follow the statement strictly, using the same structure as the existing exercises:

- `checkStatement(t, ...)`: fails with `test setup error` if an input breaks the constraints (sizes, value ranges, sorted input, "exactly one solution", etc.).
- An **independent** oracle (`bruteForce`, a different algorithm) that also checks the expected value of every hand-written case, so a wrong test is found before it blames the user's code.
- Table-driven test with the statement examples, edge cases, repeated values, negatives/zeros when allowed, and the constraint bounds.
- A random test with small inputs and a **fixed seed** (`rand.New(rand.NewSource(n))`) compared with the oracle.
- A large-input test at the **maximum size** allowed, with a fast oracle or a closed-form result, so a slow solution fails.
- When the statement says "in any order", compare after sorting, never an exact order. Keep the statement's index base (for example, 1-indexed answers).
- Do not require the input to stay unmodified unless the problem is clearly about not changing it.

**Always verify the tests before giving them to the user**: copy the test file to a temporary folder under `/tmp`, add a reference solution there, run `go vet` and `go test`, and delete the folder. The reference solution must never stay in the repository.

## Helping with a solution

- If the user asks for **hints**, give only hints and examples. No complete code.
- If the user asks to **review** or says it fails, run `go vet` and `go test` first, and explain the failure with a concrete trace of the failing input.
- If the user asks to **fix or finish** the code, do it and explain what changed.
- Don't blame the user for correct logic: if a suspected error turns out to be correct, say so.

## After the solution works

When the user asks for documentation:

- Add a doc comment to every function with what it does, how the algorithm works (and a short example if it helps), `Time complexity: O(...)` and `Space complexity: O(...)`, saying what `n` is.
- Keep the user's `// NOTE:` comments exactly as they are. Only mention typos, and fix grammar in them if asked.
- Run `gofmt -l`, `go vet` and `go test` after editing.

When the user asks to add the exercise to `README.md`, add a row to the table of its NeetCode category (Arrays & Hashing, Two Pointers, Sliding Window, Stack, ...), creating the section if it does not exist yet. Check that every local link points to an existing file.

## Commands

```bash
go test ./...                       # all the tests
go test -v ./NN_LC<number>_<name>/  # one exercise
go test -bench . -benchmem ./NN_LC<number>_<name>/
```

## Other

- Don't commit or push unless the user asks.
- Don't leave temporary files in the repository.
- If asked only for "the best time complexity", answer with just that (and the space complexity if it matters), no hints.
