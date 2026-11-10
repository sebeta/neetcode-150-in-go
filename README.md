# NeetCode 150 in Go

I'm simply working through the [NeetCode](https://neetcode.io) 150 problem list using Go. I'm following this
[video](https://www.youtube.com/watch?v=T0u5nwSA0w0&t=1061s) from [freeCodeCamp YouTube channel](https://www.youtube.com/@freecodecamp),
first solving each exercise on my own, and then comparing my solution with the "best" answer in the video.

I use AI to create the test files and as a helper in some cases, but my goal here is to learn.

I leave comments where relevant so I can understand the solutions when I come back to them in the future.

## Solved exercises

Each exercise has its own folder with the problem statement (`.md`), the solution (`.go`) and its tests (`_test.go`).
The categories follow the ones used by NeetCode. To run all the tests: `go test ./...`

### Arrays & Hashing

| LeetCode | Problem | Difficulty | Solution |
|---|---|---|---|
| 217 | [Contains Duplicate](01_LC217_containsduplicate/leetcode217.md) | Easy | [containsduplicate.go](01_LC217_containsduplicate/containsduplicate.go) |
| 242 | [Valid Anagram](02_LC242_validanagram/leetcode242.md) | Easy | [validanagram.go](02_LC242_validanagram/validanagram.go) |
| 1 | [Two Sum](03_LC01_twosum/leetcode01.md) | Easy | [twosum.go](03_LC01_twosum/twosum.go) |
| 49 | [Group Anagrams](04_LC49_groupanagrams/leetcode49.md) | Medium | [groupanagrams.go](04_LC49_groupanagrams/groupanagrams.go) |
| 347 | [Top K Frequent Elements](05_LC347_topkfrequentelements/leetcode347.md) | Medium | [topkfrequent.go](05_LC347_topkfrequentelements/topkfrequent.go) |
| 238 | [Product of Array Except Self](06_LC238_productofarrayexceptself/leetcode238.md) | Medium | [productofarrayexceptself.go](06_LC238_productofarrayexceptself/productofarrayexceptself.go) |
| 36 | [Valid Sudoku](07_LC36_validsudoku/leetcode36.md) | Medium | [validsudoku.go](07_LC36_validsudoku/validsudoku.go) |
| 271 | [Encode and Decode Strings](08_LC271_encodeanddecodestrings/leetcode271.md) | Medium | [encodeanddecodestrings.go](08_LC271_encodeanddecodestrings/encodeanddecodestrings.go) |
| 128 | [Longest Consecutive Sequence](09_LC128_longestconsecutivesequence/leetcode128.md) | Medium | [longestconsecutive.go](09_LC128_longestconsecutivesequence/longestconsecutive.go) |

### Two Pointers

| LeetCode | Problem | Difficulty | Solution |
|---|---|---|---|
| 125 | [Valid Palindrome](10_LC125_validpalindrome/leetcode125.md) | Easy | [validpalindrome.go](10_LC125_validpalindrome/validpalindrome.go) |
| 167 | [Two Sum II - Input Array Is Sorted](11_LC167_twosum2/leetcode167.md) | Medium | [twosum2.go](11_LC167_twosum2/twosum2.go) |
| 15 | [3Sum](12_LC15_3sum/leetcode15.md) | Medium | [3sum.go](12_LC15_3sum/3sum.go) |
| 11 | [Container With Most Water](13_LC11_containerwithmostwater/leetcode11.md) | Medium | [containerwithmostwater.go](13_LC11_containerwithmostwater/containerwithmostwater.go) |
| 42 | [Trapping Rain Water](14_LC42_trappingrainwater/leetcode42.md) | Hard | [trappingrainwater.go](14_LC42_trappingrainwater/trappingrainwater.go) |

### Sliding Window

| LeetCode | Problem | Difficulty | Solution |
|---|---|---|---|
| 121 | [Best Time to Buy and Sell Stock](15_LC121_besttimetobuyandsellstock/leetcode121.md) | Easy | [besttimetobuyandsellstock.go](15_LC121_besttimetobuyandsellstock/besttimetobuyandsellstock.go) |
| 3 | [Longest Substring Without Repeating Characters](16_LC03_longestsubstring/leetcode03.md) | Medium | [longestsubstring.go](16_LC03_longestsubstring/longestsubstring.go) |
| 424 | [Longest Repeating Character Replacement](17_LC_424_longestrepeatingcharacter/leetcode424.md) | Medium | [longestrepeatingcharacter.go](17_LC_424_longestrepeatingcharacter/longestrepeatingcharacter.go) |
