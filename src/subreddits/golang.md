# golang
## [1][Virtual Go Meetup - Come learn about WebRTC and how you can build sub-second decentralized real-time communications software!](https://www.reddit.com/r/golang/comments/fs6lrc/virtual_go_meetup_come_learn_about_webrtc_and_how/)
- url: https://www.meetup.com/golang/events/269676725/
---

## [2][An engineer who uses Go and Rust details when he likes to use each](https://www.reddit.com/r/golang/comments/frs5av/an_engineer_who_uses_go_and_rust_details_when_he/)
- url: https://dmv.myhatchpad.com/insight/choosing-between-rust-or-go/
---

## [3][unmarshal and anonymus strut in struct](https://www.reddit.com/r/golang/comments/fscqz0/unmarshal_and_anonymus_strut_in_struct/)
- url: https://www.reddit.com/r/golang/comments/fscqz0/unmarshal_and_anonymus_strut_in_struct/
---
Hi.

I struggle to understand how to deserialize json to my types

Consider the flowing:

```
type Car struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type Entry struct {
	Car
	Color string `json:"color"`
}

```
The folowing json is deserialized fine:

```
{"brand":"Nissan","model":"Sunny","color":"white"}
```

However, when Car have a custom unmarshalJSON function, the field Color wont be set.

Here is an example: https://play.golang.org/p/dEIVnJctykV

I can probably find ways around this, but I am curios why this not works as expected. (or why I should expect this not to work :)

br.
## [4][Go Micro v2.4.0 release is out - The Go microservices development framework](https://www.reddit.com/r/golang/comments/fsb58y/go_micro_v240_release_is_out_the_go_microservices/)
- url: https://github.com/micro/go-micro/releases/tag/v2.4.0
---

## [5][Marshal structs the right way: Golang](https://www.reddit.com/r/golang/comments/fs8fkx/marshal_structs_the_right_way_golang/)
- url: https://www.reddit.com/r/golang/comments/fs8fkx/marshal_structs_the_right_way_golang/
---
Ever got stuck with marshalling structs in [\#Golang](https://twitter.com/hashtag/Golang?src=hashtag_click)? 

Penned a new blog on this problem!!  This will surely be useful for you someday.

Do give it a read :)

Feedbacks and improvements are welcome.

[https://mohitkhare.me/blog/marshal-structs-golang/](https://mohitkhare.me/blog/marshal-structs-golang/)
## [6][hashicorp/go-connlimit: A simple library that allows a network server to limit how may concurrent connections it supports from each client IP.](https://www.reddit.com/r/golang/comments/fsaugv/hashicorpgoconnlimit_a_simple_library_that_allows/)
- url: https://github.com/hashicorp/go-connlimit
---

## [7][Minimal push-pull service based on worker pool](https://www.reddit.com/r/golang/comments/fs9ly7/minimal_pushpull_service_based_on_worker_pool/)
- url: https://github.com/vardius/pushpull
---

## [8][Split a text file into two text files of equal size in order](https://www.reddit.com/r/golang/comments/fsbq89/split_a_text_file_into_two_text_files_of_equal/)
- url: https://www.reddit.com/r/golang/comments/fsbq89/split_a_text_file_into_two_text_files_of_equal/
---
[https://stackoverflow.com/questions/60950014/split-a-text-file-into-two-text-files-of-equal-size-in-order](https://stackoverflow.com/questions/60950014/split-a-text-file-into-two-text-files-of-equal-size-in-order) 

`func WordbyWordScan() { //A function for dividing a text file into two file texts.file, err := os.Open("file.txt.txt") //Open a file for doing operation on it.if err != nil {log.Fatal(err)}file1, err := os.Create("file1.txt.txt") //Create a fileif err != nil {panic(err)}file2, err := os.Create("file2.txt.txt") //Create a fileif err != nil {panic(err)}defer file.Close()                //close file at the end.defer file1.Close()               //close file at the end.defer file2.Close()               //close file at the end.file.Seek(0, 0)                   // "Seek sets the offset for the next Read or Write on file to offset."scanner := bufio.NewScanner(file) // "NewScanner returns a new Scanner to read from file"scanner.Split(bufio.ScanWords)    // "Set the split function for the scanning operation."w := 0for scanner.Scan() { //writing to file1 and file2var outfile *os.Fileif w%2 == 0 {outfile = file1} else {outfile = file2}fmt.Fprintln(outfile, scanner.Text()) //"Fprintln formats using the default formats for its operands and writes to outfile."w++}if err := scanner.Err(); err != nil {log.Fatal(err)}}`

This code above, split the file.txt word by word instead of splitting the first half of words and then the second half in order. For example, this sentence “Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object.” should be split like this:

First half: “Package bufio implements buffered I/O. It wraps an io”

Second half: “.Reader or io.Writer object, creating another object.”

If anyone has any idea to fix this, I look forward to hearing from you.
## [9][Goro: A High-level Machine Learning Library](https://www.reddit.com/r/golang/comments/frrz42/goro_a_highlevel_machine_learning_library/)
- url: https://github.com/aunum/goro
---

## [10][Learn Golang Step By Step](https://www.reddit.com/r/golang/comments/fs7qqd/learn_golang_step_by_step/)
- url: https://sagarjaybhay.net/category/google-go/
---

