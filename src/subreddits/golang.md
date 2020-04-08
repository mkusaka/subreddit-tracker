# golang
## [1][GO Modules Behind The Corporate Firewall](https://www.reddit.com/r/golang/comments/fx648a/go_modules_behind_the_corporate_firewall/)
- url: https://www.dudley.codes/posts/2020.04.02-golang-behind-corporate-firewall/
---

## [2][VmTest, a library to simplify QEMU emulation in your integration tests](https://www.reddit.com/r/golang/comments/fwupcc/vmtest_a_library_to_simplify_qemu_emulation_in/)
- url: https://www.reddit.com/r/golang/comments/fwupcc/vmtest_a_library_to_simplify_qemu_emulation_in/
---
Hello everyone,

I would like to share my small contribution to this wonderful golang community.

Before going to the project description I would like to give some background. I’ve been working on a project that involved a custom Linux kernel + initramfs image. To verify my project functionality I initially used QEMU and ran it from a command-line. While manual testing kinda worked it became clear that I need an automated way to run the tests. I did not find a good library that allows me to setup &amp; run QEMU from the tests hence I wrote my own library.

Please welcome VmTest project. It allows one to set up and run a QEMU instance programmatically. The library has a number of options (kernel params, disks, architecture, timeout..) that allow to configure QEMU depending on the tests needs. VmTest also handles console input and output which makes it possible to implement test cases that interact with the emulated application.

Please find more information and examples at the project github page [https://github.com/anatol/vmtest](https://github.com/anatol/vmtest) Any feedback and contributions are welcome.
## [3][#Goyave v2.9.0](https://www.reddit.com/r/golang/comments/fx52nx/goyave_v290/)
- url: https://www.reddit.com/r/golang/comments/fx52nx/goyave_v290/
---
[\#Goyave](https://twitter.com/hashtag/Goyave?src=hashtag_click) v2.9.0 (Golang Web Framework) released check out the changelog here [https://system-glitch.github.io/goyave/guide/changelog.html#v2-9-0%E2%80%A6](https://system-glitch.github.io/goyave/guide/changelog.html#v2-9-0%E2%80%A6)
## [4][lithdew/recaptcha: Quickly verify reCAPTCHA v2/v3 submissions in Go.](https://www.reddit.com/r/golang/comments/fx0h8n/lithdewrecaptcha_quickly_verify_recaptcha_v2v3/)
- url: https://github.com/lithdew/recaptcha
---

## [5][GoLang Game Development : devlog 04 (pixel lib)](https://www.reddit.com/r/golang/comments/fwimmw/golang_game_development_devlog_04_pixel_lib/)
- url: https://www.youtube.com/watch?v=XBQ6jMGGk_Y
---

## [6][What is a three comma separated expression?](https://www.reddit.com/r/golang/comments/fwzlsx/what_is_a_three_comma_separated_expression/)
- url: https://www.reddit.com/r/golang/comments/fwzlsx/what_is_a_three_comma_separated_expression/
---
How does a three comma separated expression work?
I'm a fairly new programmer and having done Java, Python and Javascript, C++ moderately I have never come across the three comma separated expression that you see inside the go while(for) loop.
`z, t=z-(z*z-x)/(2*z), z`
Could you please provide some documentation on this?
```
func Sqrt(x float64) float64 {
	z := 1.0
	var t float64

	fmt.Printf("Sqrt approximation of %v:\n", x)
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) &lt; 1e-6 {
			break
		}
	}
	return z
}
```
## [7][UniDoc now supports JBIG2 encoding format](https://www.reddit.com/r/golang/comments/fx3vyf/unidoc_now_supports_jbig2_encoding_format/)
- url: https://unidoc.io/news/jbig2-support-in-golang
---

## [8][Persisting cookies after application exits?](https://www.reddit.com/r/golang/comments/fx14wu/persisting_cookies_after_application_exits/)
- url: https://www.reddit.com/r/golang/comments/fx14wu/persisting_cookies_after_application_exits/
---
I have a program that uses an `http.Client` with a standard cookie jar from the `net/http/cookiejar` package. The API I'm consuming gives a couple of cookies at login that, when consumed by a browser, allow the user to remain logged in even after closing and reopening the browser.

Annoyingly, `http.Cookie` doesn't seem to have JSON struct tags, so my naive initial solution was to make my own struct with the same fields, copy each cookie and store it as JSON, then read the cookies back and set them on the client at launch to prevent having to send login requests every time it starts. But when I do this, I get missing fields in the stored JSON (domain is missing from all of them).

I'm curious what solutions you have used to tackle this problem.
## [9][goSSE: (Yet Another) light and humble go SSE server](https://www.reddit.com/r/golang/comments/fx3f63/gosse_yet_another_light_and_humble_go_sse_server/)
- url: https://www.reddit.com/r/golang/comments/fx3f63/gosse_yet_another_light_and_humble_go_sse_server/
---
Does containment mean boredom ? It depends.... Having needs for stream and persistent connexion between my legacy apps and a browser notification system (uni-directional), i wrote this:

 [https://github.com/xefiji/goSSE](https://github.com/xefiji/goSSE) 

It was mostly as a playgroynd to learn more about the language.

No tests yet :(

Chears!
## [10][go-interview: A collection of technical interview problems solved with Go](https://www.reddit.com/r/golang/comments/fwkpeo/gointerview_a_collection_of_technical_interview/)
- url: https://www.reddit.com/r/golang/comments/fwkpeo/gointerview_a_collection_of_technical_interview/
---
I have been collecting various technical interview problems and coding challenges then solving them with Go.

From as simple as reversing a string to as more involving as writing the A Star algorithm. Moreover, implementing from scratch the data structures needed.

Since I started, I have received plenty of great feedback and more fun problems/challenges to solve. I was asked to design an LRU Cache, so I will be doing that next.

Send me some of your favorite ones! I'd like to grow the collection for anyone interested in reading and/or learning.

https://github.com/shomali11/go-interview
