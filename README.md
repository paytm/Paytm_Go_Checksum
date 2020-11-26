# Checksum - GO Language

## Example

```go
package main

import (
	"fmt"

	"github.com/dilip640/Paytm_Go_Checksum/paytmchecksum"
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
	paytmChecksum := paytmchecksum.GenerateSignature(paytmParams, "YOUR_MERCHANT_KEY")
	verifyChecksum := paytmchecksum.VerifySignature(paytmParams, "YOUR_MERCHANT_KEY", paytmChecksum)

	fmt.Printf("GenerateSignature Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignature Returns: %t\n\n", verifyChecksum)

	/* Generate Checksum via String */
	/* initialize JSON String */
	body := "{\"mid\":\"YOUR_MID_HERE\",\"orderId\":\"YOUR_ORDER_ID_HERE\"}"

	/**
	* Generate checksum by parameters we have
	* Find your Merchant Key in your Paytm Dashboard at https://dashboard.paytm.com/next/apikeys
	 */
	paytmChecksum = paytmchecksum.GenerateSignatureByString(body, "YOUR_MERCHANT_KEY")
	verifyChecksum = paytmchecksum.VerifySignatureByString(body, "YOUR_MERCHANT_KEY", paytmChecksum)

	fmt.Printf("GenerateSignatureByString Returns: %s\n", paytmChecksum)
	fmt.Printf("VerifySignatureByString Returns: %t\n\n", verifyChecksum)
}

```

* More Details: **https://developer.paytm.com/docs/checksum/#go**