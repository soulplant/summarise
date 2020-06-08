package algos


type run struct {
	from, to int
}

func (r run) String() string {
	if r.from == r.to {
		return fmt.Sprintf("%d", r.from)
	}
	return fmt.Sprintf("%d-%d", r.from, r.to)
}

func summarise(exclude []int) string {
	sort.Slice(exclude, func(i, j int) bool { return exclude[i] < exclude[j] })
	result := []run{}
	cur := 1
	for _, e := range exclude {
		if cur > e {
			continue
		}
		if cur == e {
			cur++
			continue
		}
		result = append(result, run{cur, e - 1})
		cur = e + 1
	}
	if cur < 99 {
		result = append(result, run{cur, 99})
	}
	parts := []string{}
	for _, r := range result {
		parts = append(parts, r.String())
	}
	return strings.Join(parts, ", ")
}

func TestSummariseGaps(t *testing.T) {
	for tci, tc := range []struct {
		in  []int
		out string
	}{
		{[]int{4, 10}, "1-3, 5-9, 11-99"},
		{[]int{4, 5, 6}, "1-3, 7-99"},
		{[]int{1, 99}, "2-98"},
		{[]int{}, "1-99"},
		{[]int{5}, "1-4, 6-99"},
		{[]int{5, 7}, "1-4, 6, 8-99"},
	} {
		actual := summarise(tc.in)
		if actual != tc.out {
			t.Errorf("[%d] Got '%s', expected '%s' from input %v", tci, actual, tc.out, tc.in)
		}
	}
}

