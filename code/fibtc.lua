function fib(n, a, b)
  if n == 0 then
    return a
  elseif n == 1 then
    return b
  end
  return fib(n-1, b, a+b)
end

print(fib(35, 0, 1))
