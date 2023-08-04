/* More Details: https://developer.paytm.com/docs/checksum/#go */

package main

import (
	"fmt"

	PaytmChecksum "./paytm"
)

func main() {

	/* Generate Checksum via Map */
	/* initialize an map */
	paytmParams := make(map[string]string)

	paytmParams = map[string]string{
		"MID":      "YOUR_MID_HERE",
		"ORDER_ID": "YOUR_ORDER_ID_HERE",
	}

	/**
	* Generate checksum by parameters we have
	* Find your Merchant Key in your Paytm Dashboard at https://dashboard.paytm.com/next/apikeys
	 */
	paytmChecksum, err := PaytmChecksum.GenerateSignature(paytmParams, "YOUR_MERCHANT_KEY")
	if err != nil {
		fmt.Printf("GenerateSignature return error %s\n", err.Error())
		return
	}
	verifyChecksum, err := PaytmChecksum.VerifySignature(paytmParams, "YOUR_MERCHANT_KEY", paytmChecksum)
	if err != nil {
		fmt.Printf("VerifySignature return error %s\n", err.Error())
		return
	}

	fmt.Printf("GenerateSignature Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignature Returns: %t\n\n", verifyChecksum)

	/* Generate Checksum via String */
	/* initialize JSON String */
	body := "{\"mid\":\"YOUR_MID_HERE\",\"orderId\":\"YOUR_ORDER_ID_HERE\"}"

	/**
	* Generate checksum by parameters we have
	* Find your Merchant Key in your Paytm Dashboard at https://dashboard.paytm.com/next/apikeys
	 */
	paytmChecksum, err := PaytmChecksum.GenerateSignatureByString(body, "YOUR_MERCHANT_KEY")
	if err != nil {
		fmt.Printf("GenerateSignatureByString return error %s\n", err.Error())
		return
	}
	verifyChecksum, err := PaytmChecksum.VerifySignatureByString(body, "YOUR_MERCHANT_KEY", paytmChecksum)
	if err != nil {
		fmt.Printf("VerifySignatureByString return error %s\n", err.Error())
		return
	}

	fmt.Printf("GenerateSignatureByString Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignatureByString Returns: %t\n\n", verifyChecksum)
}
