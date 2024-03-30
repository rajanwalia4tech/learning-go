import time
import math

def is_prime(n):
    if n <= 1:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def count_primes_up_to(n):
    count = 0
    for i in range(n + 1):
        if is_prime(i):
            count += 1
    return count

start_time = time.time()
count = count_primes_up_to(int(math.pow(10,8)))
end_time = time.time()
print()
print("---------- Python Performance -----------")
print()
print("Number of primes up to 1 million:", count)
print("Time taken: ", end_time - start_time, "seconds")
print()