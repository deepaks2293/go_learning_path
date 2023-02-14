Go provides two keywords that control the flow of a loop. The first, continue, immediately skips to the next iteration of a loop, without running any further code in the loop block.


for x:=1 ; x<=3 ; x ++ {
	fmt.Println("before continue")
	continue
	fmt.Println("after continue")
}

before continue
before continue
before continue


In the above example, the string "after continue" never gets printed, because the continue keyword always skips back to the top of the loop before the second call to Println can be run.





The second keyword, break, immediately breaks out of a loop. No further code within the loop block is executed, and no further iterations are run. Execution moves to the first statement following the loop.

for x:=1 ; x<=3 ; x ++ {
        fmt.Println("before break")
        break
        fmt.Println("after break")
}
fmt.Println("after loop")

Here, in the first iteration of the loop, the string "before break" gets printed, but then the break statement immediately breaks out of the loop, without printing the "after break" string, and without running the loop again (even though it normally would have run two more times). Execution instead moves to the statement following the loop.

