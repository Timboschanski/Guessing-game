# Guessing-game

In this Guessing Game you will name 2 numbers, a Random number will be chosen inbetween these. You will guess the number untill you got it right.

## How it works

Write into the Command line: "go run main.go 1 100"
You get: "Guess a number that is between your params. Please input your guess"
Now you guess: "5"
you get one of three options: "Your guess is too low. Try again", "Your guess is too high. Try again" or "Woohoo thats correct. Good Job! That took you # attempts"

When you want to Exit the Program while it's running you have to press: "ctrl + C"

## Errors

Incase your parameters are wrong you get one of three Error Massages:

You write not the right amount of parameters: "error: 2 command line arguments expected"
You dont write numbers as parameters: "error: Please use integers as command line parameters."
When your second parameter is lower then the first: "error: From must be <= to. Please use valid values."

Incase your guesses are wrong you also get an Error: "error: Invalid input. Please use a number"
When there is another Problem with the Input you get a different Error: "error: An error occured while reading input. Please try again"