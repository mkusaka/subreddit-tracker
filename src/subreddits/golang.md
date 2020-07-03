# golang
## [1][[Q&amp;A] //go:build draft design](https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hitexe/qa_gobuild_draft_design/
---
I posted a draft design today for updating the // +build lines to fix some usability problems. 

Video: [https://golang.org/s/go-build-video](https://golang.org/s/go-build-video)\
Design: [https://golang.org/s/go-build-design](https://golang.org/s/go-build-design)

As an experiment, let's try doing Q&amp;A about the design here in Reddit.
My hope is that the threading support will help keep questions and answers matched.

**Please start a new top-level comment for each new question.**
## [2][Github-like calendar heatmap in plain Go](https://www.reddit.com/r/golang/comments/hkgcz9/githublike_calendar_heatmap_in_plain_go/)
- url: https://github.com/nikolaydubina/calendarheatmap
---

## [3][Started learning Go today. Is my understanding correct that GOPATH is no longer required in favour of Go modules?](https://www.reddit.com/r/golang/comments/hkishj/started_learning_go_today_is_my_understanding/)
- url: https://www.reddit.com/r/golang/comments/hkishj/started_learning_go_today_is_my_understanding/
---
My [Go Programming Language](https://www.gopl.io/) book mentions the use of GOPATH as a means to store Go code related to source files, packages and a bin directory. I'm slightly confused about this, does that mean that the only place I can write go code would be in the $GOPATH ?

Also, [go w/tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go#go-modules) states that GOPATH will be deprecated soon. So should I just forget about this concept and move on with Go modules?

Truth be told, I haven't researched into the benefits of Go modules too much as of yet as I'm just trying to set up my local Go environment before I start coding. It seems that I should just forget about GOPATH and just get cracking with code. Would anyone be able to ease my mind off this issue?

Much appreciated :)
## [4][Size doesn't matter](https://www.reddit.com/r/golang/comments/hjuyfv/size_doesnt_matter/)
- url: https://i.redd.it/m5tjssbydf851.jpg
---

## [5][How do you deal with private dependencies in CI?](https://www.reddit.com/r/golang/comments/hkiyfo/how_do_you_deal_with_private_dependencies_in_ci/)
- url: https://www.reddit.com/r/golang/comments/hkiyfo/how_do_you_deal_with_private_dependencies_in_ci/
---
What are people doing for private dependencies these days in your CI?

Vendoring? Do you put credentials into your builder nodes? SSH key? Access token?

What's considered standard these days?
## [6][Is it possible to get multiple AWS S3 object tags in batch using Go?](https://www.reddit.com/r/golang/comments/hkipr3/is_it_possible_to_get_multiple_aws_s3_object_tags/)
- url: https://www.reddit.com/r/golang/comments/hkipr3/is_it_possible_to_get_multiple_aws_s3_object_tags/
---
Hello everyone, I asked the following question on StackOverflow too but did not receive a response. I was wondering if anyone here could help me.

Suppose I have a given bucket name `bucket` and a slice of (string) item keys `keys`. I can get the tags for all of these objects by doing the following:

    sess, _ := session.NewSession(&amp;aws.Config{
    	Region: aws.String("eu-west-2"),
    })
    
    svc := s3.New(sess)
    
    for _, key := range(keys) {
        response, _ := svc.GetObjectTagging(
            &amp;s3.GetObjectTaggingInput{
    	        Bucket: aws.String(bucket),
    	        Key:    aws.String(key),
    	    }
        )
    	fmt.Println(response.TagSet)
    }

In my program, I am finding this quite slow. Is there a way to get all the tags in one go, instead of calling this multiple times?

Any help would be much appreciated, thank you.
## [7][WaitGroup in Go](https://www.reddit.com/r/golang/comments/hkibg3/waitgroup_in_go/)
- url: https://chenpy.com/article/87235493714501274
---

## [8][A Journey building a fast JSON parser and full JSONPath, Oj for Go](https://www.reddit.com/r/golang/comments/hk1hm6/a_journey_building_a_fast_json_parser_and_full/)
- url: https://github.com/ohler55/ojg/blob/master/design.md
---

## [9][How to substitute io.Reader for os.Stdin](https://www.reddit.com/r/golang/comments/hkgahy/how_to_substitute_ioreader_for_osstdin/)
- url: https://www.reddit.com/r/golang/comments/hkgahy/how_to_substitute_ioreader_for_osstdin/
---
I wanted to pass an `io.Reader` to a function to replace `os.Stdin` for testing. Unfortunately my function started behaving differently when using `strings.NewReader()`. Below is a simplified reproduction.

The first line of `INPUT` gives the number of records I want to read from stdin. Each record consists of two lines: the number of integers and the integers separated by spaces.

    package main
    
    import (
    	"fmt"
    	"io"
    	"os"
    	"strings"
    )
    
    func main() {
    	INPUT := `2
    2
    1 2
    3
    1 2 3
    `
    	fmt.Println(ReadInput(strings.NewReader(INPUT)))
    	fmt.Println(ReadInput(os.Stdin))
    }
    
    func ReadInput(reader io.Reader) [][]int {
    	var countRecords int
    	_, _ = fmt.Fscanln(reader, &amp;countRecords)
    
    	records := make([][]int, countRecords)
    	for i := 0; i &lt; countRecords; i++ {
    		var countNumbers int
    		_, _ = fmt.Fscanln(reader, &amp;countNumbers)
    
    		numbers := make([]int, countNumbers)
    		for j := 0; j &lt; countNumbers; j++ {
    			fmt.Fscan(reader, &amp;numbers[j])
    		}
    
    		records[i] = numbers
    		// fmt.Fscanln(reader) // FIXME
    	}
    	return records
    }

If I `go run` this code, I get the following:

    [[1 2] []]
    ... [copy INPUT to Stdin]
    [[1 2] [1 2 3]]

Here I get the expected result (`[[1 2] [1 2 3]]`) when reading from `os.Stdin` but not from `strings.NewReader`.

If I uncomment the last `Fscanln` marked with FIXME, I get the opposite:

    [[1 2] [1 2 3]]
    ... [copy INPUT to Stdin]
    [[1 2] []]

I would have expected this code to produce the same result with both `strings.NewReader` and `os.Stdin`. What is happening here?

What can I substitute for `os.Stdin` in my test function?
## [10][Go driven rpc code generation tool for right now.](https://www.reddit.com/r/golang/comments/hk6u9a/go_driven_rpc_code_generation_tool_for_right_now/)
- url: https://github.com/pacedotdev/oto
---

## [11][Which is the best way to convert int to string in golang ?](https://www.reddit.com/r/golang/comments/hkcgms/which_is_the_best_way_to_convert_int_to_string_in/)
- url: https://www.reddit.com/r/golang/comments/hkcgms/which_is_the_best_way_to_convert_int_to_string_in/
---
There are some solutions:

1. stringValue = string(intValue)

1. stringValue = strconv.Itoa(intValue)

1. stringValue = fmt.Sprintf("%d", intValue)

I checked a project by a Guru at Google Engineers (link below), they use the third one. Could someone explain why ?

[https://github.com/google/exposure-notifications-server/blob/main/internal/database/connection.go](https://github.com/google/exposure-notifications-server/blob/main/internal/database/connection.go)
