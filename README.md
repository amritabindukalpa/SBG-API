# github.com/amritabindukalpa/sbg-api

The bets are placed as parameters to the api. At the moment it allows for 3 types of bets to be placed: - number,  odd or even (parity) and the colour of the resulting number

e.g. http://localhost:8080/?n=29&c=red&p=odd

The result for this is a direct display of the text message whether the user has won or lost a particular category of bet.

I have also attempted to create another api which will take moneys on the bets and generate the output
e.g. http://localhost:8080/nb=1&cb=1&pb=1 where the value is the amount being betted

The result does not perform any mathematical calculation but just returns the value being betted. 
The result generator is an which calls the random number generator. The random number generator is an api in itself which takes a parameter and returns a random number within a specified range. The parameter could be set to generate a random number for different types of betting games.

To run the application :-

1. Extract the repository git clone <bundle-file-directory>/sbg-api.bundle <directory-to-extract-to>
2. Open the folder <directory-to-extract-to> in VS Code
3. Start a new terminal( Ctrl + Shift + ')
4. Run docker-compose up
5. Browse to http://localhost:8080/?n=29&c=red&p=odd with the parameters with n, c and p being changed

Things I would have liked to explore :-

1. I would have liked to implemented more error handling in the application and input validation in the application
2. I would have liked to explore having middleware for logging application flow
3. The application takes the input parameters as simple string. I would have liked to explore json objects