The program starts with an Init function in main.go which initializes the tic tac toe array and the players details

Once initilialized, the game will start using StartPlay function

I have kept the main.go neat and the program modularized so that it is less cluttered

inside structs.go there are all the structures and global variables which are used in functions.go

helpers.go contains all the functions which helps the game to move ahead

Inside Init function, I have asked the user about the grid size and the type of play, i.e. single player, multiplayer or computers only.
I guess these should be inside Init as they are used to instantiate the game and will not be used further

Player has a Notation and Symbol, i.e. Symbol on the grid will be either X or 0, Notation will be the player number

ShowGrid function is called after every move by any player so that they are updated with the current situation of the game

When the computer decides its move, it selects a random number from the available slots using the freeslots slice

CheckValidMove function checks the bound of the entered value by the user or generated value by the random function

If the sanity checks are cleared, then we check whether there is a winner already, the code is scalable and can check with any grid size

Once the grid is updated, the freeslots slice gets updated and the value is removed. 
Once the winner is checked, and there is no winner or there is no draw scenario, the players are toggled and the game continues

As we get a winner, or the match is drawn, the game engs gracefully