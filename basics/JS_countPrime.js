function isPrime(num) {
    if (num < 2) {
      return false;
    }
    for (let i = 2; i * i <= num; i++) {
      if (num % i === 0) {
        return false;
      }
    }
    return true;
  }
  
  function countPrimes(start, end) {
    let count = 0;
    for (let num = start; num <= end; num++) {
      if (isPrime(num)) {
        count++;
      }
    }
    return count;
  }
  
  const start = new Date();
  const count = countPrimes(1, Math.pow(10,8));
  const end = new Date();
  console.log();
	console.log("------- NodeJS Performance -------");
  console.log();
  console.log(`Number of prime numbers between 1 and 1 million: ${count}`);
  console.log(`Time taken: ${(end - start)/1000} seconds`);
  console.log();