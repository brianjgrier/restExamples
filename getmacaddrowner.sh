#!/bin/bash

if [[ "$#" != 1 ]]; then
    echo "Please supply one MAC adderss"
    exit 1
fi

    apiKeyFile=~/.macaddress

if [ ! -f "$apiKeyFile" ]; then
	
    echo "Cannot find API Key file for maccaddress.io"
    echo "Please place your API Key in the file ~/.macaddress"
    echo
    exit 1
fi

while IFS= read -r line
do
    mytoken=$line
done < $apiKeyFile

requestedMacAddr=$1

authHeader="X-Authentication-Token:"${mytoken}
urlInfo="https://api.macaddress.io/v1?output=json&search="${requestedMacAddr}
cmd="curl -s -H ${authHeader} ${urlInfo}"
output=`$cmd`

if [[ "$output" == *"error"* ]]; then
    errorMessage=`echo $output | python3 -c "import sys; import json; print(json.load(sys.stdin)['error'])"`
    echo "Request error:" $errorMessage
    exit 1
fi

companyName=`echo $output| python3 -c "import sys; import json; print(json.load(sys.stdin)['vendorDetails']['companyName'])"`
    
echo "Owner Name =" $companyName
