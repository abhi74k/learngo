package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func WordFreq(s string) map[string]int {

	var buf strings.Builder
	ret := make(map[string]int)

	flush := func() {
		if buf.Len() > 0 {
			ret[buf.String()]++
			buf.Reset()
		}
	}

	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			buf.WriteRune(c)
		} else {
			flush()
		}
	}

	flush()

	return ret
}

func UniqueTrimmed(xs []string) []string {

	seen := make(map[string]struct{})
	ret := make([]string, 0, len(xs))

	for _, s := range xs {

		trimmed_str := strings.TrimSpace(s)

		if _, ok := seen[trimmed_str]; ok {
			continue
		}

		if s == "" {
			continue
		}

		seen[trimmed_str] = struct{}{}
		ret = append(ret, trimmed_str)
	}

	return ret
}

func GroupAnagrams(xs []string) map[string][]string {

	AnagramSignature := func(anagram string) string {

		runes := []rune(anagram)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		return string(runes)
	}

	ret := make(map[string][]string)

	for _, s := range xs {
		key := AnagramSignature(s)
		fmt.Println("string: ", s, "signature:", key)
		ret[key] = append(ret[key], s)
	}

	return ret
}

func RLE(s string) string {

	var buf strings.Builder

	var current_char rune
	count := 0

	flush := func() {
		buf.WriteRune(current_char)
		if count > 1 {
			buf.WriteString(strconv.Itoa(count))
		}
	}

	for idx, c := range s {

		if idx == 0 {
			current_char = c
			count = 1
			continue
		}

		if current_char == c {
			count++
		} else {
			flush()

			current_char = c
			count = 1
		}
	}

	flush()

	return buf.String()
}

func TwoSum(nums []int, target int) (int, int, bool) {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i := 0
	j := len(nums) - 1

	for i < j {

		if nums[i]+nums[j] == target {
			return i, j, true
		}

		if nums[i]+nums[j] < target {
			i++
			continue
		}

		if nums[i]+nums[j] > target {
			j--
			continue
		}
	}

	return 0, 0, false
}

type Interval struct {
	Start int
	End   int
}

func MergeIntervals(intervals []Interval) []Interval {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		} else {
			return intervals[i].Start < intervals[j].Start
		}
	})

	var merged_intervals []Interval

	for idx, new_interval := range intervals {

		if idx == 0 {
			merged_intervals = append(merged_intervals, new_interval)
			continue
		}

		last_merged_interval := &merged_intervals[len(merged_intervals)-1]

		if new_interval.Start <= last_merged_interval.End {

			new_end := max(last_merged_interval.End, new_interval.End)

			fmt.Printf("Merging intervals: [%d,%d] with [%d,%d] ==> [%d,%d]\n",
				last_merged_interval.Start, last_merged_interval.End,
				new_interval.Start, new_interval.End,
				last_merged_interval.Start, new_end)

			last_merged_interval.End = new_end

		} else {
			fmt.Printf("Adding new interval: [%d,%d]\n", new_interval.Start, new_interval.End)
			merged_intervals = append(merged_intervals, new_interval)
		}
	}

	return merged_intervals
}

func NetBalances(xs []string) map[string]int {

	ret := make(map[string]int)

	for _, s := range xs {
		parts := strings.Split(s, ":")
		money, err := strconv.Atoi(parts[1])
		if err == nil {
			parts = strings.Split(parts[0], "->")
			from := parts[0]
			to := parts[1]

			ret[from] -= money
			ret[to] += money
		}
	}

	return ret
}

type Freq struct {
	val   int
	count int
}

func TopK(nums []int, k int) []int {

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	freq := make([]Freq, 0, len(counts))
	for k, v := range counts {
		freq = append(freq, Freq{val: k, count: v})
	}

	sort.Slice(freq, func(i, j int) bool {
		if freq[i].count == freq[j].count {
			return freq[i].val < freq[j].val
		} else {
			return freq[i].count > freq[j].count
		}
	})

	k = min(k, len(counts))
	topk := freq[:k]

	out := make([]int, 0, k)
	for _, item := range topk {
		out = append(out, item.val)
	}

	return out
}

func IsBalanced(s string) bool {

	stack := make([]rune, 0, len(s))

	push := func(c rune) {
		stack = append(stack, c)
	}

	pop := func() (rune, bool) {
		if len(stack) == 0 {
			return 0, false
		}

		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		return last, true
	}

	match := map[rune]rune{')': '(', '}': '{', ']': '{'}

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			push(c)
		case ')', '}', ']':
			popped, ok := pop()
			if !ok || popped != match[c] {
				return false
			}
		}
	}

	return len(stack) == 0
}
