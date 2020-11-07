// Golang program to get the current
// date and time in various format
package main

import (
	"fmt"
	"time"
)

func main() {

	// using time.Now() function
	// to get the current time
	currentTime := time.Now()

	// getting the time in string format
	fmt.Println("Show Current Time in String: ", currentTime.String())

	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2017.09.07 17:06:06"))

	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2017#09#07"))

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("09-07-2017"))

	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2017-09-07"))

	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2017-09-07 17:06:06"))

	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000"))

	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000000"))

	fmt.Println("ShortNum Wedth : ", currentTime.Format("2017-02-07"))

	fmt.Println("ShortYear : ", currentTime.Format("06-Feb-07"))

	fmt.Println("LongWeekDay : ", currentTime.Format("2017-09-07 17:06:06 Wednesday"))

	fmt.Println("ShortWeek Day : ", currentTime.Format("2017-09-07 Wed"))

	fmt.Println("ShortDay : ", currentTime.Format("Wed 2017-09-2"))

	fmt.Println("LongWedth : ", currentTime.Format("2017-March-07"))

	fmt.Println("ShortWedth : ", currentTime.Format("2017-Feb-07"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 PM"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 pm"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5"))

}
