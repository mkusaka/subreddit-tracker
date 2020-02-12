# golang
## [1][simdjson-go: Parsing gigabyes of JSON per second in Go](https://www.reddit.com/r/golang/comments/f2cxza/simdjsongo_parsing_gigabyes_of_json_per_second_in/)
- url: https://blog.min.io/simdjson-go-parsing-gigabyes-of-json-per-second-in-go/
---

## [2][Experience report on a large Python-to-Go translation](https://www.reddit.com/r/golang/comments/f2hbzv/experience_report_on_a_large_pythontogo/)
- url: https://gitlab.com/esr/reposurgeon/blob/master/GoNotes.adoc
---

## [3][New to server side](https://www.reddit.com/r/golang/comments/f2pnh6/new_to_server_side/)
- url: https://www.reddit.com/r/golang/comments/f2pnh6/new_to_server_side/
---
Hey guys I was wondering if any one could provide resources for a new back end Engineer I am currently coming from a front end background. I am very green when it comes to back end so any tutorials that teach both golang and the concepts around server side would be highly appreciated.
## [4][Is this the correct way to encrypt/decrypt a file?](https://www.reddit.com/r/golang/comments/f2pk2x/is_this_the_correct_way_to_encryptdecrypt_a_file/)
- url: https://www.reddit.com/r/golang/comments/f2pk2x/is_this_the_correct_way_to_encryptdecrypt_a_file/
---
I'm learning Go and wanted to write a program for encrypting/decrypting a file by password. You can find my project [here](https://github.com/EgidioCaprino/Broccoli). I have a couple of question.

1. Is AES safe enough?
2. Am I using the AES cipher correctly?

Any other advice is welcome ☺
## [5][Birthday reminder with Golang](https://www.reddit.com/r/golang/comments/f2p0b0/birthday_reminder_with_golang/)
- url: https://github.com/fallais/gobirthday
---

## [6][How to download output PDF file from generator into client](https://www.reddit.com/r/golang/comments/f2ooxm/how_to_download_output_pdf_file_from_generator/)
- url: https://www.reddit.com/r/golang/comments/f2ooxm/how_to_download_output_pdf_file_from_generator/
---
i create a service to generate pdf file with \`go-wkhtmltopdf\`

this is the function

    func Pdf() func(c *gin.Context) {
    	return func(c *gin.Context) {
    		// download wkhtmltopdf and set that wkhtmltopdf on below
    		wkhtmltopdf.SetPath("/usr/local/bin/wkhtmltopdf")
    		type ReqHtml struct {
    			Html string `json:"html"`
    		}
    		var html ReqHtml
    		_ = c.BindJSON(&amp;html)
    
    		pdfg, err := wkhtmltopdf.NewPDFGenerator()
    		if err != nil {
    			log.Fatal(err)
    		}
    
    		page := wkhtmltopdf.NewPageReader(strings.NewReader(html.Html))
    		pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
    		page.FooterRight.Set("[page]")
    		page.FooterFontSize.Set(10)
    
    		pdfg.AddPage(page)
    
    		err = pdfg.Create()
    		if err != nil {
    			log.Println(err)
    		}
    		err = pdfg.WriteFile("./myPdf.pdf")
    		if err != nil {
    			log.Fatal(err)
    		}
    
    	}
    }

that Writefile writes file into my root folder server, how to make the client download it?

let say i have &lt;button&gt; download file &lt;/button&gt;

should I give a response to the client with the buffer of that result? or what?
## [7][I created a body parser for database models](https://www.reddit.com/r/golang/comments/f2b6uf/i_created_a_body_parser_for_database_models/)
- url: https://www.reddit.com/r/golang/comments/f2b6uf/i_created_a_body_parser_for_database_models/
---
Hi,

This is my first package and wanted to know your opinion especially on the benchmark. I wanted your opinion if the hit in performance is worth it.

Look at the README for the motivation.

[https://github.com/eduarhasanaj/bop](https://github.com/eduarhasanaj/bop)
## [8][Pipeline builder package](https://www.reddit.com/r/golang/comments/f2hhw1/pipeline_builder_package/)
- url: https://www.reddit.com/r/golang/comments/f2hhw1/pipeline_builder_package/
---
Hi goers (or golanders?), this is my first package built on Go and I wanted to share it. Any feedback is greatly appreciated :)

The main reason of it is to build/execute (and visualize!) workflows/pipelines, thus executing complex workflows while keeping the responsibilities as decoupled as possible.

Maybe there’s something already implemented out there, the ones I found where not that developed

https://github.com/saantiaguilera/go-pipeline

Thanks!
## [9][What's New In Go 1.14: Test Cleanup](https://www.reddit.com/r/golang/comments/f2d2vp/whats_new_in_go_114_test_cleanup/)
- url: https://www.gopherguides.com/articles/test-cleanup-in-go-1-14/?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=golang-testing-cleanup-114
---

## [10][Web Framework Question](https://www.reddit.com/r/golang/comments/f2e0wj/web_framework_question/)
- url: https://www.reddit.com/r/golang/comments/f2e0wj/web_framework_question/
---
Hey! :) 

For the coming semester I need to build a REST API in Go and, coming from python, I was looking for a webframework similar to Django or Flask.

I've googled and it seems that 1-2 years ago most ppl were recommending to not use any framework at all, but the standard library. Others were recommending echo, some blog posts or articles recommended Gin or Beego.

Now, with my goal in mind, and since I'm fairly new to Go and hope I could get pointers from more experienced folks, what framework would be the best choice? Or should I go without a framework completely still? 

I appreciate any advice! :)
