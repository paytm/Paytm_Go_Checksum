/* More Details: https://developer.paytm.com/docs/checksum/#go */

package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"./paytm"
)

func main() {

	/* Generate Checksum via Map */
	/* initialize an map */	
	paytmParams := make(map[string]string)
	
	paytmParams = map[string]string{
		"MID": "YOUR_MID_HERE",
		"ORDER_ID": "YOUR_ORDER_ID_HERE",
	}

	/**
	* Generate checksum by parameters we have
	* Find your Merchant Key in your Paytm Dashboard at https://dashboard.paytm.com/next/apikeys 
	*/
	paytmChecksum := PaytmChecksum.GenerateSignature(paytmParams, "YOUR_KEY_HERE")
	verifyChecksum := PaytmChecksum.VerifySignature(paytmParams, "YOUR_KEY_HERE", paytmChecksum)

	fmt.Printf("GenerateSignature Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignature Returns: %t\n\n", verifyChecksum)

	/* Generate Checksum via String */
	/* initialize JSON String */  
	body := "{\"mid\":\"YOUR_MID_HERE\",\"orderId\":\"YOUR_ORDER_ID_HERE\"}"

	/**
	* Generate checksum by parameters we have
	* Find your Merchant Key in your Paytm Dashboard at https://dashboard.paytm.com/next/apikeys 
	*/
	paytmChecksum = PaytmChecksum.GenerateSignatureByString(body, "YOUR_KEY_HERE")
	verifyChecksum = PaytmChecksum.VerifySignatureByString(body, "YOUR_KEY_HERE", paytmChecksum)

	fmt.Printf("GenerateSignatureByString Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignatureByString Returns: %t\n\n", verifyChecksum)
}