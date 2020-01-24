#!/usr/bin/python3

#
# A simple python script, NOT A CLASS BASED PROGRAM
# Making a class makes no sense because the information does not meet the definition of an object
#

import os
import requests
import sys

from pprint import pprint

defaultMacAddr = "44:38:39:ff:ef:57"
#defaultMacAddr = "44:38"

baseURL = "https://api.macaddress.io/v1?output=json&search="
myApiKey = ""

def getFromURL(url, macAddr):
    '''
    Get the owner information of a "mac address" from the registration site specified in the variable url
    two values are returned, the success/failure status and the accompanying text for that status.
    There is a lot of information returned from the query, but the problem statement only asked for the owner name
    :param url:
    :param macAddr:
    :return:
    '''

    resultText = ""         # Ensure we will return a string type variable
    resultStatus = -1       # Assume an Error
    response = requests.get(url+macAddr,
                            headers={'X-Authentication-Token': myApiKey})
    resultJson = response.json()
    if response.status_code != 200:
        resultText = "Error: " + resultJson['error']
        resultStatus = 1
    else:
        resultText = "Owner Name: " + resultJson['vendorDetails']['companyName']
        resultStatus = 0
    return resultStatus, resultText

#
# ONLY run if this is being executed as the "main" program...
# DO NOT run if being imported by another module
#
status = 0

if __name__ == "__main__":
    #
    # I could use argparse, or getargs but they are a bit much for this example
    #
    # I am also not using the exit immediately upon error style of programming
    # Why not? I do not generally like it. There are times is improves readability, this is not one of those
    #
    if len(sys.argv) != 2:
        print("usage: command macaddr")
        print("       you must supply exactly one argument")
        macAddr = defaultMacAddr
        status = 1
    else:
        #
        # Now check to see if we can find the file $HOME/.macaddress
        # How we do this depends upon the OS this is running on (Linux or Windows)
        # Linux:    $HOME/.macaddress
        # Windows:  $HOMEPATH/.macaddress
        #
        seperator = '\\' if os.name == 'nt' else '/'
        homeVariable = 'HOMEPATH' if os.name == 'nt' else 'HOME'
        homeDir = os.getenv(homeVariable)
        keyFile = homeDir + seperator + '.macaddress'
        #
        # there are two ways to do this. A try except clause of using 'with'
        # I prefer the 'with' method
        fileOpened = False
        with open(keyFile, 'r') as filePtr:
            fileOpened = True
            #
            # read the key from the file, remove the newline charater and convert to a bytestream
            # There are things I do not like about python3 and strings/bytes
            #
            myApiKey = filePtr.readline().strip('\n').encode()
            if len(myApiKey) == 0:
                print("Key file, {} was empty")
                print("Please register and obtain an API key at macaddress.io")
            else:
                macAddr = sys.argv[1]
                status, text = getFromURL(baseURL, macAddr)
                print(text)
        if not fileOpened:
            print('Could not open file containing your API key: {}'.format(keyFile))

    #
    # The exit needs to be here to prevent the program from exiting IF this
    # file is being included by another python module
    #
    exit(status)