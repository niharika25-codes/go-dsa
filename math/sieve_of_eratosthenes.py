def sieve_of_eratosthenes(n):
    if n <= 2: return 0
    primes = [True for _ in range(n)]
    p = 2
    while (p*p) < n:
        if primes[p] == True:
            for i in range(p*p, n, p):
                primes[i] = False
        p += 1
    return len([p for p in range(2, n) if primes[p]])

print(sieve_of_eratosthenes(3))
print(sieve_of_eratosthenes(2))
