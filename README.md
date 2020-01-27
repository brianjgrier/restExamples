This is a simple set of code to show using a REST api with three languages.

I created this out of a request to provide a code example of getting the owner
of a MAC ADDRESS from the site macaddress.io.

This site allows you to register and obtain an API Key to access their information.
Because it is bad practice to code an API Key, or secret, into code I chose to place 
the API key in a file in my "HOME" directory. That file is named ".macaddress"

To add to the level of difficulty I decided to make sure that the python3 and go language 
implementations can work on both windows and linux systems. 

As I am relatively new to 'go' I am sure that this could be done much more efficiently.

Part of the request was to create a docker image and supply the build and test instructions.

I decided to buuld a container image for each option. The 'Go' example created an image that
was over 800MB, which seemed crazy since Go programs are supposed to be able to be compiled
and exist outside of a container.

Investigating this showed that the base image for the container was the consuming the space,
not the application. Switching to a minimal container (Ubuntu 18.04) resulted in a very small
image, but did require building the program with a static build instead of allowing it to use 
the run-time dynamically linked libraries. 

The results for my test builds are:<br>

| REPOSITORY | TAG | IMAGE ID | CREATED | SIZE |
| ---------- | -------- | -------------- | -------------------- | ------ |
| go_test2 | latest | 500318ee8094 | About a minute ago | 114MB |
| go_test1 | latest | b200f0d8975b | 2 minutes ago | 818MB |
| shell_test | latest | 21cacc7136c5 | 2 minutes ago | 197MB | 
| python_test | latest | d0c746dca86c | 2 minutes ago | 183MB |

Another surprise was the the shell version resulted in a larger image than the python version.
This is because the shell version leverages python to parse the results, and I also needed to 
add the utility 'curl' to  the image.

Adding my own json parser in shell would result in a considerably smaller container, but would
not be worth the effort required.

If you want to use the included bash script to build the container images you will need to get your own
API Key from 'macaddressio.com' and place it in a file named '.macaddress' in your home directory.

To test a version use the following linux commands:

docker run -i <image_name> <valid macaddress>
docker run -i <image_name> <invalid macaddress>
docker run -i <image_name> <first three octects of macaddr>


