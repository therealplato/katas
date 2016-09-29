Given a chart []element
Start with window size 3
Sweep the window across the chart from left to right
  - dam = max(xWinLeft, xWinRight)
	- for all xWinMid:
	  if height < dam
		  accumulate (height-dam)
			fill in hole
