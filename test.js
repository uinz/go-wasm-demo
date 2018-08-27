function jsRun() {
  console.time("js fib(30) took");
  fib(30);
  console.timeEnd("js fib(30) took");
}

function fib(n) {
  if (n === 1 || n === 2) {
    return 1;
  }
  return fib(n - 1) + fib(n - 2);
}
