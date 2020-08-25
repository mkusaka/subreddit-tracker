# golang
## [1][Anyone interested in a book about graphics in Go?](https://www.reddit.com/r/golang/comments/ig7a0h/anyone_interested_in_a_book_about_graphics_in_go/)
- url: https://www.reddit.com/r/golang/comments/ig7a0h/anyone_interested_in_a_book_about_graphics_in_go/
---
This is more of a general inquiry at this point. I love both Go and graphics programming. At some point I thought, I might actually write a short book combining both topics. As we all know, Go is much more than just a language for writing Web services, yet most materials teach only this side of it. I thought that by using Go to teach a different context (e.g. computer graphics), I might give it a different spin and bring in some fresh minds into the community.

I will soon share a link to my first scratch. It is not really a full chapter at the moment. At this point, I am combining a wiki of things, from basic GP concepts to more advanced, like game programming. 

What do you think? Will this be something useful for the general Go community? Feel free to share feedback as well as resources (e.g. libraries, tutorials, etc) which you want to see featured there.
## [2][Moul • The minimalist publishing tool for photographers](https://www.reddit.com/r/golang/comments/ig5s8x/moul_the_minimalist_publishing_tool_for/)
- url: https://www.reddit.com/r/golang/comments/ig5s8x/moul_the_minimalist_publishing_tool_for/
---
TLDR; Demo: [https://demo.moul.app](https://demo.moul.app/) | GitHub: [https://github.com/moulco/moul](https://github.com/moulco/moul)

If you are familiar with static site generator, Moul is a very similar in a sense that it generates you a static website. But rather content oriented, its focus on photo collection instead.

It's written in Go!

Hope someone else might find it useful other than me. (I just made it for myself more or less 😅)
## [3][Echelon - hierarchical progress in terminals](https://www.reddit.com/r/golang/comments/ifo15p/echelon_hierarchical_progress_in_terminals/)
- url: https://www.reddit.com/r/golang/comments/ifo15p/echelon_hierarchical_progress_in_terminals/
---
Hey everyone! For a recent project we wanted to show progress of a CI build execution in a terminal in a pretty way. There are several libraries for showing progress bars interactively in terminals but I didn't find one which has UI similar to what build systems like Bazel, Buck and Gradle has. So I wrote [one  from scratch called Echelon](https://github.com/cirruslabs/echelon). Here is how it looks:

[Example on macOS](https://i.redd.it/cygiuzanyxi51.gif)

The library is customizable and can work with any VT100-compatible terminal (I don't know any terminal that doesn't support ASCII escape symbols 😅). Echelon also has a simplified renderer for dump terminals in case you pipe the output or run it as part of CI that doesn't have an interactive terminal. 

With  VT100-compatible terminals Echelon does incremental redraws: figures out which lines have changed from the last "frame" and redraws them with minimal amount of ASCII escape symbols. This makes Echelon very smooth.

And it also works on Windows! But doesn't look as pretty since Windows doesn't support emojis by default.

[Example on Windows](https://i.redd.it/36olmbtf0yi51.gif)

I hope Echelon will be useful for someone else. It's [open sourced](https://github.com/cirruslabs/echelon) under MIT license so please give it a try. 🙌
## [4][Just published GLab v1.10.0, with speed improvements, bug fixes, and additional features: https://github.com/profclems/glab/releases/latest](https://www.reddit.com/r/golang/comments/igc1o3/just_published_glab_v1100_with_speed_improvements/)
- url: https://i.redd.it/5mplv0axf5j51.png
---

## [5][Package sortedset provide sorted set, with strings/binary comparator, backed by arrays](https://www.reddit.com/r/golang/comments/igc0ua/package_sortedset_provide_sorted_set_with/)
- url: https://github.com/recoilme/sortedset
---

## [6][Upgrading CockroachDB from dep to Go Modules](https://www.reddit.com/r/golang/comments/iftikw/upgrading_cockroachdb_from_dep_to_go_modules/)
- url: https://www.cockroachlabs.com/blog/dep-go-modules
---

## [7][Writing Business Logic in a testable and clean manner, advice requested](https://www.reddit.com/r/golang/comments/ig77l0/writing_business_logic_in_a_testable_and_clean/)
- url: https://www.reddit.com/r/golang/comments/ig77l0/writing_business_logic_in_a_testable_and_clean/
---
Hi,

We're currently building an app which we want to be testable (as the service we are implementing into has no test environment). We have build the separate packages for the third party services we are integrating, but we want to "combine" these packages in a clean way to move data from Service A to Service B.

In short the app does the following:

\- Integrate with Service A (we have built Service A into a separate package)

\- Integratie with Service B (we have built Service B into a separate package)

\- On request, fetch data from multiple resources by ID from Service A into Struct A (within Service A), map this data into a single struct from Service B and submit this to service B.

Are there any examples on how to do this clean and testable? Should we create an interface for the passed data struct so that the Package B "Struct builder" has a contract with a lot of functions/methods to retrieve the needed data? And have a function that fills a struct based on that interface from the Service A side?

Beforehand, we would just create a large function with retrieves all data from Service A and then just map it all, but with testing in mind this doesn't feel really clean.

Service A needs to call \~5-6 resources to retrieve all the data needed to fill the required fields of Service B.

The Service B struct has around \~100 properties to fill (with nested structs) so we want to keep the code as readable as possible.

Hope anyone can point me in the right direction/examples (or design pattern that fits this).
## [8][Clivern/Poodle - A fast and beautiful command line tool to build API requests.](https://www.reddit.com/r/golang/comments/igbgdh/clivernpoodle_a_fast_and_beautiful_command_line/)
- url: https://github.com/Clivern/Poodle
---

## [9][Weird random number results](https://www.reddit.com/r/golang/comments/iga4zm/weird_random_number_results/)
- url: https://www.reddit.com/r/golang/comments/iga4zm/weird_random_number_results/
---
I've written a program that generates a 2D array to create noise and then an image. It starts out by creating a random value in the top left then completes the top row of the array, adding or subtracting a random value from the previous value. Then it does the same for the first column, adding or subtracting a random value from the one above. It then fills in the rest of the array by averaging the cell to the top and to the left, then adding or subtracting a random value from it. 

This creates a pseudo noise array that I then convert into an image.

A weird thing I've noticed is that the values of the array get smaller and smaller as you get closer to the bottom right of the array, resulting in the image getting darker towards the bottom right. This happens every time and I dont know why, is there something I'm missing out for the random function?

Code: 

    package main
    
    import (
    	"fmt"
    	"image"
    	"image/color"
    	"image/png"
    	"math"
    	"math/rand"
    	"os"
    	"time"
    )
    
    func main() {
    	// set random seed and image resolution
    	rand.Seed(time.Now().UnixNano())
    	resolution := 256
    
    	matrix := createMatrix(resolution)
    
    	// print in matrix format
    	for i:=0; i&lt; resolution; i++ {
    		fmt.Println(matrix[i])
    	}
    
    	createImage(matrix, resolution)
    }
    
    func createImage(matrix [256][256]int, resolution int) {
    	// create image parameters
    	upLeft := image.Point{0,0}
    	lowRight := image.Point{resolution, resolution}
    	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
    
    	// iterate through pixels
    	for y := 0; y &lt; resolution; y++ {
    		for x := 0; x &lt; resolution; x++ {
    			// set colour to colour of matrix pixel
    			col := color.RGBA{uint8(matrix[y][x]), uint8(matrix[y][x]), uint8(matrix[y][x]), 0xff}
    			img.Set(x, y, col)
    		}
    	}
    
    	// save image
    	f, _ := os.Create("image.png")
    	png.Encode(f, img)
    }
    
    func createMatrix(resolution int) [256][256]int {
    	// create main array
    	noise := [256][256]int{}
    
    	// create possible offsets
    	offset := [11]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
    
    	// set first position to random
    	noise[0][0] = rand.Intn(255)
    
    	// create first row and column based on first position
    	for x := 1; x &lt; resolution; x++ {
    		noise[0][x] = int(math.Abs(float64(noise[0][x-1] + offset[rand.Intn(11)])))
    		if noise[0][x] &gt; 255 {
    			noise[0][x] = 255
    		}
    	}
    	for y := 1; y &lt; resolution; y++ {
    		noise[y][0] = int(math.Abs(float64(noise[y-1][0] + offset[rand.Intn(11)])))
    		if noise[y][0] &gt; 255 {
    			noise[y][0] = 255
    		}
    	}
    
    	// fill out rest of table
    	for y := 1; y &lt; resolution; y++ {
    		for x := 1; x &lt; resolution; x++ {
    			noise[y][x] = int(math.Abs(float64((noise[y-1][x]+noise[y][x-1])/2 + offset[rand.Intn(11)])))
    			if noise[y][x] &gt; 255 {
    				noise[y][x] = 255
    			}
    		}
    	}
    
    	return noise
    }

It creates images like this:

&amp;#x200B;

https://preview.redd.it/bvbwvwizt4j51.png?width=256&amp;format=png&amp;auto=webp&amp;s=36c0e48bd8e9d9919de10245d0bb9a8d63715486

&amp;#x200B;

https://preview.redd.it/ay4v2t42u4j51.png?width=256&amp;format=png&amp;auto=webp&amp;s=c4d8e5c9183be34fed833a426ad10fbf188799a2

&amp;#x200B;

https://preview.redd.it/qlr3z9g4u4j51.png?width=256&amp;format=png&amp;auto=webp&amp;s=26c47427f2fbd385adb7f9767eef25c7b809ac5a
## [10][What is Go's memory model?](https://www.reddit.com/r/golang/comments/ig59lp/what_is_gos_memory_model/)
- url: https://www.reddit.com/r/golang/comments/ig59lp/what_is_gos_memory_model/
---
I read [The Go Memory Model](https://golang.org/ref/mem) as well as [these slides](http://nil.csail.mit.edu/6.824/2016/notes/gomem.pdf), but I'm still confused about what Go's memory model actually guarantees and doesn't guarantee.

In particular, the former link only specifies a "happens before" relation, but as pointed out in the second link, this is not enough to provide any meaningful guarantees about the behavior of racy programs. In particular, the specification is silent on the matter of "thin air" values. It's also unclear what exactly counts as reads and writes for the purpose of the memory model.

It seems to me like the specification as written makes data races only one step removed from Undefined Behavior. (Of course, data races on many types like slices and maps _are_ undefined behavior, but people treat machine sized types as if they're atomic.)
