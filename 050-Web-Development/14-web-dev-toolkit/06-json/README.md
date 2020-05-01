GoLang has 2 ways to convert the data:

1) Marshal and Unmarshal
2) Encode and Decode


Difference between the two:

**Golang Marshal and Unmarshal:**

1) Marshal - if we have a Go data structure, convert into JSON and assign it to a variable, Marshal does it.
2) Unmarshal - Exact opposite, takes a JSON, converts into Go data structure and assigns to variable.


**Encode and Decode:**
1) Encode - takes a Go Data Structure, converts into JSON and sends over the wire.
2) Decode - receives data from the wire as JSON, converts into Go Data structure 

