# golang
## [1][Draft GORM V2 Documents Online, suggestions, feedback welcome!](https://www.reddit.com/r/golang/comments/ho0ajo/draft_gorm_v2_documents_online_suggestions/)
- url: http://v2.gorm.io/docs/
---

## [2][Go bindings to QuickJS: a fast, small, and embeddable ES2020 JavaScript interpreter.](https://www.reddit.com/r/golang/comments/ho0o9e/go_bindings_to_quickjs_a_fast_small_and/)
- url: https://github.com/lithdew/quickjs
---

## [3][Let's take a moment to appreciate the cryto/ssh package](https://www.reddit.com/r/golang/comments/hnkhkt/lets_take_a_moment_to_appreciate_the_crytossh/)
- url: https://www.reddit.com/r/golang/comments/hnkhkt/lets_take_a_moment_to_appreciate_the_crytossh/
---
It allowed me to create some, well... at least to me, exciting pieces of software, and I'm looking for more inspiration. 

I don't know exactly what it is about SSH, but it may very well be the fact that it is available everywhere these days, just like the browser, so services created for it are readily accessible for everyone.

The crypto/ssh package allows you to create services that would otherwise be quite tedious to configure with the regular OpenSSH tools

What have you guys built with it?

I'll get the party started:

[https://github.com/fasmide/schttp](https://github.com/fasmide/schttp) \- A daemon that accepts file transfers from SCP clients, on the fly, zip or tgz compresses the transfer and makes the result available to others through shareable HTTP URLs.

&amp;#x200B;

[https://github.com/fasmide/remotemoe](https://github.com/fasmide/remotemoe) \- a fresh take on "tunnels to localhost" software - Using the SSH client, you simply forward your local services, and remotemoe takes care of the rest.

&amp;#x200B;

Should you happen to be one of the creators of crypto/ssh - please tell me - how can I buy you a beer?
## [4][FileServer and Vue Router - problem with dynamic route matching](https://www.reddit.com/r/golang/comments/ho34j5/fileserver_and_vue_router_problem_with_dynamic/)
- url: https://www.reddit.com/r/golang/comments/ho34j5/fileserver_and_vue_router_problem_with_dynamic/
---
Greetings to all,

I am using standard Go approach from net/http for serving static files

`mux := http.NewServeMux()`  
`mux.Handle("/", http.FileServer(http.Dir("/path/to/my/folder")))`

All standard Vue routes work, but ones using dynamic matching, for example:

`{`  
`path: '/:key([a-zA-Z0-9]{32})',`  
`name: 'Foo',`  
`component: Foo`  
`},`

simply fail with error 404.

What is the proper way to handle such routes?

TIA!
## [5][Goc - A new open source coverage collection tool for Golang](https://www.reddit.com/r/golang/comments/hnzem6/goc_a_new_open_source_coverage_collection_tool/)
- url: https://medium.com/@dacarl.ji/goc-a-new-open-source-coverage-collection-tool-for-golang-663d8c4343e9
---

## [6][Building a meetup application leveraging Dgraph and Twitter's API](https://www.reddit.com/r/golang/comments/hnplew/building_a_meetup_application_leveraging_dgraph/)
- url: https://www.youtube.com/watch?v=u_SitrVcwr8
---

## [7][Web Service with Go and Go-Web](https://www.reddit.com/r/golang/comments/ho0z71/web_service_with_go_and_goweb/)
- url: https://medium.com/swlh/web-service-with-go-and-go-web-576bd9a91205
---

## [8][Any good beginner books in Go that is updated covering from basics to intermediate concepts?](https://www.reddit.com/r/golang/comments/hneda0/any_good_beginner_books_in_go_that_is_updated/)
- url: https://www.reddit.com/r/golang/comments/hneda0/any_good_beginner_books_in_go_that_is_updated/
---
Very new to Go language. I know basics like conditionals,arrays,slice,loops,etc. But I see there are lots of terminologies like Buffered I/o, bufio, using command lines, flags, etc. I am finding it hard to understand all these things. Is there any beginner books,resources or any video tutorials/courses to understand these things from Scratch?
## [9][[GoCV] Lack of DrawMatches method in this package? Any way to come around it?](https://www.reddit.com/r/golang/comments/hnndvz/gocv_lack_of_drawmatches_method_in_this_package/)
- url: https://www.reddit.com/r/golang/comments/hnndvz/gocv_lack_of_drawmatches_method_in_this_package/
---
Hello everyone!

I am writing a small program to learn feature matching, and I came to a point where I am basically lost. My current code is working, I got result, nothing hangs, result DMatch slice has length of over 900. However I wanted to confirm that the result is correct. In Python / C++ implementations of OpenCV (at least according to [docs](https://docs.opencv.org/3.4/d4/d5d/group__features2d__draw.html)) I have those functions allowing me to do that:

**Python**:
```python
outImg = cv.drawMatches(
        img1, keypoints1, 
        img2, keypoints2, 
        matches1to2, 
        outImg[, matchColor[, singlePointColor[, matchesMask[, flags]]]])
```

**C++**:
```c++
void cv::drawMatches     (
        InputArray img1,
        const std::vector&lt;KeyPoint&gt; &amp;eypoints1,
        InputArray img2,
        const std::vector&lt;KeyPoint&gt; &amp;keypoints2,
        const std::vector&lt;std::vector&lt;DMatch&gt;&gt; &amp;matches1to2,
        InputOutputArray outImg,
        const Scalar &amp;matchColor = Scalar::all(-1),
        const Scalar &amp;singlePointColor = Scalar::all(-1),
        const std::vector&lt;std::vector&lt;char&gt;&gt; &amp;matchesMask = std::vector&lt;std::vector&lt;char&gt;&gt;(),
        int flags = DrawMatchesFlags::DEFAULT 
    )
```

However GoCV (OpenCV bindings for Go) doesn't have such a function. Did anyone tried to implement that OR has any option to check if SIFT matching outcome is correct?

Here is my function, I appreciate any suggestions what can I improve or change.
```go
func BruteSIFTMatching(query string, train string) {
	img1 := gocv.IMRead(query, gocv.IMReadGrayScale)
	img2 := gocv.IMRead(train, gocv.IMReadGrayScale)

	sift := contrib.NewSIFT()

	kp1, des1 := sift.DetectAndCompute(img1, gocv.NewMat())
	kp2, des2 := sift.DetectAndCompute(img2, gocv.NewMat())

	bf := gocv.NewBFMatcher()
	matches := bf.KnnMatch(des1, des2, 2)

	var good []gocv.DMatch

	for _, m := range matches {
		if len(m) &gt; 1 {
			if m[0].Distance &lt; 0.75 * m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	gocv.DrawKeyPoints(img1, kp1, &amp;img1, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)

	gocv.DrawKeyPoints(img2, kp2, &amp;img2, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}, gocv.DrawDefault)


	window1 := gocv.NewWindow("Query")
	window1.IMShow(img1)

	window2 := gocv.NewWindow("Train")
	window2.IMShow(img2)

	fmt.Println(len(good))

	window1.WaitKey(0)
	window2.WaitKey(0)
}
```
## [10][I'm a Golang beginner and I wrote a Brainfuck interpreter. Check it out](https://www.reddit.com/r/golang/comments/hnpq1e/im_a_golang_beginner_and_i_wrote_a_brainfuck/)
- url: https://mihajlonesic.gitlab.io/archive/brainfuck-in-go/
---

