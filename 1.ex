defmodule Euler.Problem1 do
	# If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
	# The sum of these multiples is 23.

	# Find the sum of all the multiples of 3 or 5 below 1000.
	def solve(limit) do
		(1..limit-1)
			|> Enum.filter(fn x -> rem(x, 3) == 0 || rem(x, 5) == 0 end)
			|> Enum.sum
  end

  def print do
  	IO.puts "The answer is: #{solve(1000)}"
	end

end
