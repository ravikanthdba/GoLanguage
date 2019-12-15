/* Searches the value of the key really fast 
   Most efficient
   unordered search - ideally takes a long time, since it linearly searches all the keys.
   improves efficiency by sorting.
   break sections into categories - like 1 category - letter a, another category - letter b etc. if we do this, we don't need to search for where are the alphabets starting with some character. we can directly hit that bucket and fetch results. - uneven distribution of buckets. some buckets can have more words and other buckets have less words.
   hence hash function. It takes in an letter, runs some algorithm and produces a hash number and based on the hash number, determine which bucket it will fall under. - evenly distributes all our buckets.
   input changes even by a small ",", output is compeltely different.
*/