# Task 1

## What happens if you remove the `go-command` from the `Seek` call in the `main` function?

* **Hypothesis**: The program will not create a goroutine, I.E it will not run parallell and the loop will run sequntial.

* **Result**: When removing the `go keyword` the program outputs the same output:

`Anna sent a message to Bob.`\
`Cody sent a message to Dave.`\
`No one received Evaâ€™s message.`

This indicates that the program is ran sequntial instead of running parallell. When `go keyword` is present, each time the program is ran, the output is different.

## What happens if you switch the declartion `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` amnd the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?

* **Hypothesis**: The `WaitGroup` is passed by value and not by reference. Any changes done to `wg` within `Seek` is not applied in `main` and program deadlocks.

* **Result**: As hypothesis states, program deadlocks.

## What happens if you remove the buffer on the channel match?

* **Hypothesis**: A buffered channel allows for the channel to accept limited number of values without a correpsonding receiver. If the buffered channel is removed the program will deadlock as the channel has no receiver.

* **Result**: As hypothesis states, the program deadlock as it has no corresponding receiver to receiver the data and the value cannot be accepted.

## What happens if you remove the default-case from the case-statement in the `main` function?

* **Hypothesis**: If we remove the default-case and no goroutine is writing to the channel, the program can be deadlocked but this will not happen in this program because there will always be something written to the channel.

* **Result**: Nothing happens, as hypothesis stated.
