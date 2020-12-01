from read_input import *

def accounting():
    nums = read_line_input_int("01")
    matches = {}
    pairs = []
    for n in nums:
        pair = 2020-n
        if matches.get(pair) is not None:
            print("Matched!")
            pairs = [n, pair]
            break
        matches[n] = 0

    print(pairs)
    print(pairs[0] * pairs[1])

def accounting_three():
    # if 2 values are > 2020 it's not the answer
    # otherwise, search the remainder of the numbers to find the 'pair'
    nums = read_line_input_int("01")
    # possibly sort?
    matches = {}
    for n in nums:
        # key is num, value is the sum of the other 2 numbers
        matches[n] = 2020-n
    
    for i, n1 in enumerate(nums):
        j = i+1
        for j in range(len(nums)-1):
            n2 = nums[j]
            # print(j, n2)
            if n1 + n2 >= 2020:
                # print(f"over, skipping {n1} {n2}")
                j+=1
                continue
            
            # is the compliment present?
            compliment = 2020 - n1 - n2
            if matches.get(compliment) is not None:
                print("found!")
                print(n1, n2, compliment)
                return (n1*n2*compliment)
            else:
                print("compliment not found")

            j+=1


# accounting()
print(accounting_three())
    
