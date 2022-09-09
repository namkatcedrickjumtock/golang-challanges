# golang-tutorials

## Arrays    / Slices
````
var ages [3]int = [3]int {1, 2 , 40} // strongely  styped
var ages = [3]int {1, 2 , 40}  // typed infered
ages = [3]int{1, 2, 3} // short hand assignmnt

<!-- both arrays and slices can be accessed via indexes -->
names[1] = 30

types like int, string applies for other types.
````

````
slices uses arrays under the hood

var score = []int{23 , 45 ,6}

newScores := append(score, 50) // doesn't append to the original slice but return a new slice of names

ftmt.printl(newScores) // print new score sith 50 added
ftmt.printl(score) // print original score
````
````
Slice Ranges

 names := []String {'cedrick', 'Junior', 'prince', 'Prosper'}

<!-- ranges return new slices -->
rangeOne :=names[1:3]

ftmt.printl(rangeOne) // output junior and prince. starts at pos 1 but doesn't pos 3 
````




