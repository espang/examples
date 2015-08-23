def median(arr):
	sarr = sorted(arr)
	n = len(arr)
	if n%2 == 0:
		return (sarr[n//2-1] + sarr[n//2]) // 2
	return sarr[n//2]
	
def rolling_median_trivial(h, arr):
	n = len(arr)
	res = []
	for i in range(h):
		res.append(float('nan'))
	for i in range(h, n-h):
		res.append(median(arr[i-h:i+h+1]))
	for i in range(h):
		res.append(float('nan'))
	return res
