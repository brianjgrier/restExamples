This is a simple set of code to show using a REST api with three languages.

I created this out of a request to provide a code example of getting the owner
of a MAC ADDRESS from the site macaddress.io.

This site allows you to register and obtain an API Key to access their information.
Because it is bad practice to code an API Key, or secret, into code I chose to place 
the API key in a file in my "HOME" directory. That file is named ".macaddress"

To add to the level of difficulty I decided to make sure that the python3 and go language 
implementations can work on both windows and linux systems. 

As I am relatively new to 'go' I am sure that this could be done much more efficiently.

