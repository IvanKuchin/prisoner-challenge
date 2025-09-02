# The 100 Prisoners Problem

## Scenario:


- There are 100 prisoners, each assigned a unique number from 1 to 100.
- There is a room with 100 boxes, each containing a slip of paper with one of the numbers (from 1 to 100), randomly distributed.
- The prisoners are not allowed to communicate with each other once the challenge begins.

## Objective:
Each prisoner must find the box containing their own number. If every prisoner succeeds, they all go free. If even one fails, none escape.


## Rules:
- Each prisoner is allowed to enter the room one at a time.
- Each prisoner may open up to 50 boxes (half of the total).
- After a prisoner leaves, all boxes are closed again as before.
- Prisoners may discuss and agree on a strategy before the process begins, but cannot communicate once the challenge starts.

## Challenge:
On the surface, if each prisoner randomly opens 50 boxes, the chance that all prisoners find their own number is astronomically low (about 1 in 2^100). However, there is a clever strategy that increases the probability of all prisoners escaping to over 30%.



Key Points of the Optimal Strategy (The "Cycle" Strategy)

## Pre-Agreement:
All prisoners agree that each will start by opening the box labeled with their own number.

## Follow the Chain:
If their number is not inside, they open the box labeled with the number they found in the previous box, and repeat this up to 50 times.

## Why This Works:
This strategy exploits the properties of permutation cycles, ensuring that, as long as all cycles in the permutation are 50 or fewer in length, all prisoners will find their number.



## Example

- Prisoner #17 starts by opening Box 17.
- If Box 17 contains slip #42, they next open Box 42.
- If Box 42 contains slip #99, they next open Box 99.
…continue up to 50 times or until their own number is found.


## Probability

The probability all prisoners succeed is equal to the probability that the largest cycle in the permutation of the 100 numbers is 50 or less.
This probability is roughly 31%—dramatically better than random guessing.


Note: This problem is a popular illustration of probability, permutations, and group strategy, and has been featured in mathematical puzzles and computer science discussions.

## Code

Calculates following parameters from 4 prisoners / boxes / stickers:
- total possible number of allocations stickers per boxes (4 !)
- number of combinations that set prisoners free
- percerntage of combinations that set prosoners free
- also write out all combinations that sets prisoners free

To run the code:
```bash
go run cmd/main.go
```
Output:
```bash
0123 : longest cycle = 1
0132 : longest cycle = 2
0213 : longest cycle = 2
0321 : longest cycle = 2
1023 : longest cycle = 2
1032 : longest cycle = 2
2103 : longest cycle = 2
2301 : longest cycle = 2
3120 : longest cycle = 2
3210 : longest cycle = 2
Total permutations:  24
Set free:  10
Set free percentage:  41.67 %
```

To change number of prisoners, edit file main.go line 
```go
func main() {
   crates := 4
```
