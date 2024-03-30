from flask import Flask
import math

app = Flask(__name__)

def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(math.sqrt(n)) + 1):
        if n % i == 0:
            return False
    return True

@app.route('/slow')
def slow_route():
    count = 0
    number = 2
    primes = []
    while count < 1000000:
        if is_prime(number):
            primes.append(number)
            count += 1
        number += 1
    return "Counted 1 million prime numbers!"

@app.route('/fast')
def fast_route():
    return "Hello"

if __name__ == '__main__':
    app.run()
