num = 1000

def fibSeq():
	ind = 1
	i = 0
	j = 1
	def nextTerm():
		nonlocal ind, i, j
		ind += 1
		k = i + j
		i = j
		j = k
		return ind, k

	return nextTerm

nextTerm = fibSeq()

pair = nextTerm()
while len(str(pair[1])) < num:
	pair = nextTerm()

print(pair)