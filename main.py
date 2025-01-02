import time as t
import math

list_size = 1_000_000


def append_list(n: int, l: list) -> list:
    for i in range(n):
        l.append(i)
    return l


# Calculate time to append items to a list
start = t.time()
append_list(list_size, [])
end = t.time()
# convert to milliseconds
end = (end - start) * 1000
print("Time to append items to a list of size {0}: {1}".format(list_size, end))


print("Happy new year", sum(i**3 for i in range(10)))

sum = 0
for i in range(11):
    sum += i**4
    print(i**3)


print(sum)
math.sqrt(sum)
print(list(i for i in range(10)))


def sum_of_nums_from_zero(n: int) -> int:
    sum = n * (n + 1) / 2
    return int(sum)


def diff_to_zero(n: int) -> int:
    diff = n - ((n - 1) * (n) / 2)
    return int(diff)


print(sum_of_nums_from_zero(5_000_000))
