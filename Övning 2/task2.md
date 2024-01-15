/*
- What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
	- We don't wait for the go routines to send information to the channels before closing
	- We get a message: "panic: send on closed channel"
	- goroutine 20 [running]:
*/

/*
- What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
    - Program panicked and self shut down because we are closing the channel before finishing the go routines
    - panic: send on closed channel
    - goroutine 21 [running]:

*/

/*
- What happens if you remove the statement close(ch) completely?
    - Go will close the channel if there is no Close statement
*/

/*
- What happens if you increase the number of consumers from 2 to 4?
    - The code will perform faster because there will be 4 concurrent goroutines instead of 2 so they will match with amount of producers
*/

/*
- Can you be sure that all strings are printed before the program stops?
    - By creating the waitgroup for consumers, we can be sure that all strings are printed before the program stops.
*/