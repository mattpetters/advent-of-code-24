package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "sort"
)

func adventOfCodeDayOne(){
  // locations empty
  // locations have a location ID
  // pair up the smallest number in the left list with the smallest number right list, second smallest with second smallest, etc.
  /*
  Example:

For example:

3   4
4   3
2   5
1   3
3   9
3   3
Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the smallest number in the left list 
with the smallest number in the right list, then the second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example, if you pair up a 3 from the left list with 
a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.

In the example list above, the pairs and distances would be as follows:

The smallest number in the left list is 1, and the smallest number in the right list is 3. The distance between them is 2.
The second-smallest number in the left list is 2, and the second-smallest number in the right list is another 3. The distance between them is 1.
The third-smallest number in both lists is 3, so the distance between them is 0.
The next numbers to pair up are 3 and 4, a distance of 1.
The fifth-smallest numbers in each list are 3 and 5, a distance of 2.
Finally, the largest number in the left list is 4, while the largest number in the right list is 9; these are a distance 5 apart.

so basically I need to sort, then add up the absolute difference between the two lists

input format:

34500   97487
39882   93726
12013   47966
42691   18536
57217   14334
31320   58173
90532   42331
65216   54184
78393   42097
48315   29944
37332   48362
40774   98927
16062   28783
73931   39391
54344   21435
17846   56504
61807   42097
93272   97487

  */

  data, err := ioutil.ReadFile("input.txt")
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fmt.Println("Contents of file:", string(data))

  // split the data into two lists
  leftList := []int{}
  rightList := []int{}

  // split the data into two lists
  for _, line := range strings.Split(string(data), "\n") {
    // skip empty lines
    if line == "" {
      continue
    }
    // split the line into two parts
    parts := strings.Split(line, " ")
    // convert the parts to integers using Atoi
    // note: Atoi is short for "ASCII to integer"

    left, err := strconv.Atoi(parts[0])
    if err != nil {
      fmt.Println("Error converting left number", err)
      return
    }
    right, err := strconv.Atoi(parts[1])
    if err != nil {
      fmt.Println("Error converting right number", err)
      return
    }
    // add the numbers to the lists
    leftList = append(leftList, left)
    rightList = append(rightList, right)
  }

  // sort the lists
  // using the sort package
  sort.Ints(leftList)
  sort.Ints(rightList)

  // calculate the distance
  distance := 0
  for i := 0; i < len(leftList); i++ {
    distance += abs(leftList[i] - rightList[i])
  }

  fmt.Println("Distance:", distance)
}
