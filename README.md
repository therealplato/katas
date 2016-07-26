https://www.hackerrank.com/challenges/crush

You are given a list of size , initialized with zeroes. You have to perform  operations on the list and output the maximum of final values of all the  elements in the list. For every operation, you are given three integers ,  and  and you have to add value  to all the elements ranging from index  to (both inclusive).

Input Format 
First line will contain two integers  and  separated by a single space.
Next  lines will contain three integers ,  and  separated by a single space.
Numbers in list are numbered from  to .

Output Format 
A single line containing maximum value in the updated list.

Constraints
3 <= N <= 10^7
1 <= M <= 2*10^5
1 <= a <= b <= N
0 <= k <= 10^9


Sample Input #00

5 3
1 2 100
2 5 100
3 4 100
Sample Output #00

200
Explanation

After first update list will be 100 100 0 0 0. 
After second update list will be 100 200 100 100 100.
After third update list will be 100 200 200 200 100.
So the required answer will be 200.
