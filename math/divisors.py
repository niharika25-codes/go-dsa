import math

def get_divisors(n):
    divisors = set()
    for i in range(1, int(math.sqrt(n)) + 1):
        if n % i == 0:
            divisors.add(i)
            divisors.add(n // i)
    return sorted(list(divisors))

# Example usage
number = 36
print(f"Divisors of {number}: {get_divisors(number)}")