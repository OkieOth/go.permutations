
i = 11


def fakultaet(number):
	result = 1
	for i in range(1, number+1):
		result = result * i
	return result

def testFakultaet():
	f1 = fakultaet(1)
	f2 = fakultaet(2)
	f3 = fakultaet(3)
	f4 = fakultaet(4)
	f5 = fakultaet(5)
	print("f1: {}´".format(f1))
	print("f2: {}".format(f2))
	print("f3: {}".format(f3))
	print("f4: {}".format(f4))
	print("f5: {}".format(f5))


def createPermutationsArray(numberCount):
	permutationsCount = fakultaet(numberCount)
	permutations = []
	for i in range(permutationsCount):
		tmp = []
		for j in range(numberCount):
			tmp.append(-1)
		permutations.append(tmp)
	return permutations


def getRestSlice(input, usedIndex):
	ret = []
	for i in range(len(input)):
		if i != usedIndex:
			ret.append(input[i])
	return ret


def permutations(input):
	inputLen = len(input)
	if inputLen == 2:
		ret = []
		ret.append((input[0], input[1]))
		ret.append((input[1], input[0]))
		return ret
	else:
		ret = createPermutationsArray(inputLen)
		permCount_0 = fakultaet(inputLen - 1)
		for i in range(inputLen):
			restSlice = getRestSlice(input, i)
			restPermutations = permutations(restSlice)
			for j in range(permCount_0):
				index = (i * permCount_0) + j
				ret[index][0] = input[i]
				for k in range(1, inputLen):
					ret[index][k] = restPermutations[j][k-1]
		return ret


def printPermutations(numberCount):
	input = []
    # input number to array
	for i in range(numberCount):
		input.append(i)

    # get result
	permutationSlice = permutations(input)
	print("Number of permutations: {}".format(len(permutationSlice)))

	# i = 0
	# for v in permutationSlice:
	# 	print("{}: {}".format(i, v))
	# 	i = i +1


#testFakultaet()
printPermutations(i)
