# golang
## [1][[Tutorial] How to setup your first Kibana dashboard for your Go application](https://www.reddit.com/r/golang/comments/i6fd9m/tutorial_how_to_setup_your_first_kibana_dashboard/)
- url: https://pmihaylov.com/kibana-dashboard-tutorial/
---

## [2][Go package and CLI to download files 4x faster than cURL and Wget](https://www.reddit.com/r/golang/comments/i5yokn/go_package_and_cli_to_download_files_4x_faster/)
- url: https://github.com/melbahja/got
---

## [3][Should we run net/http Server ListenAndServe() inside a goroutine and how should handle its returning errors?](https://www.reddit.com/r/golang/comments/i6fopj/should_we_run_nethttp_server_listenandserve/)
- url: https://www.reddit.com/r/golang/comments/i6fopj/should_we_run_nethttp_server_listenandserve/
---
- Should we run `net/http` Server `ListenAndServe()` inside a goroutine?
- What is the correct way to handle `net/http` `Server` `ListenAndServe()` errors?

1. The very first way I tried without graceful shutdowns

```go
s := &amp;http.Server{
	Addr:         ":8080",
	Handler:      appRouter, // appRouter == *chi.Mux
	ReadTimeout:  5*time.Second,
	WriteTimeout: 5*time.Second,
	IdleTimeout:  5*time.Second,
}

if err := s.ListenAndServe(); err != nil &amp;&amp; err != http.ErrServerClosed {
	// log error
}
```

2. With graceful shutdowns

```go
s := &amp;http.Server{
	Addr:         ":8080",
	Handler:      appRouter, // appRouter == *chi.Mux
	ReadTimeout:  5*time.Second ,
	WriteTimeout: 5*time.Second,
	IdleTimeout:  5*time.Second,
}

closed := make(chan struct{})
go func() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)
	&lt;-sigint

	if err := s.Shutdown(ctx); err != nil {
		// log error
	}
        // kill DB like server deps
	close(closed)
}()

if err := s.ListenAndServe(); err != nil &amp;&amp; err != http.ErrServerClosed {
	// log error
}

&lt;-closed
```

3. Running `Server.ListenAndServe()` inside separate goroutine and use select

```go
s := &amp;http.Server{
	Addr:         ":8080",
	Handler:      appRouter, // appRouter == *chi.Mux
	ReadTimeout:  5*time.Second ,
	WriteTimeout: 5*time.Second,
	IdleTimeout:  5*time.Second,
}

sigint := make(chan os.Signal, 1)
signal.Notify(sigint, os.Interrupt)
signal.Notify(sigint, syscall.SIGTERM)

serverErrors := make(chan error, 1)

go func() {
	serverErrors &lt;- s.ListenAndServe()
}()

select {
case err := &lt;-serverErrors:
	// log error

case &lt;-sigint:
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		// log error

		if err = s.Close(); err != nil {
		    // log error
		}
	}

}
```

- May I know any special benefit on 3rd attempt; running `ListenAndServe()` in separate goroutine and which way you suggest to handle server errors without using any 3rd party lib. 

- Or is there any other way you suggest?

Thanks
## [4][Introducing postgrestest, a test harness that starts an ephemeral PostgreSQL server](https://www.reddit.com/r/golang/comments/i65wfr/introducing_postgrestest_a_test_harness_that/)
- url: https://www.zombiezen.com/blog/2020/08/introducing-postgrestest/
---

## [5][slice of string as response net/http](https://www.reddit.com/r/golang/comments/i6gurs/slice_of_string_as_response_nethttp/)
- url: https://www.reddit.com/r/golang/comments/i6gurs/slice_of_string_as_response_nethttp/
---
I want to return slice of string as response.

```
func getEmails(res http.ResponseWriter,req *http.Request){

	// database call
        // sample
	userEmails := []string{"userone@email.com","usertwo@email.com"}
	
       // how to return userEmails
}
```

I have tried using `fmt.Fprintf(res,userEmails)`
But it fails as its not a string.
## [6][Code review: CLI sitemap builder](https://www.reddit.com/r/golang/comments/i6gcxg/code_review_cli_sitemap_builder/)
- url: https://www.reddit.com/r/golang/comments/i6gcxg/code_review_cli_sitemap_builder/
---
Hi everyone! I’m a newbie gopher and for the last month on and off I was working on a small project, and now I think it’s very close to being done, and I want to make it public for everyone to review it.  
[https://github.com/TofuOverdose/WebMapMaker](https://github.com/TofuOverdose/WebMapMaker)

The program itself is pretty simple: it’s a CLI tool which crawls the specified website to build a sitemap of it. Obviously it heavily uses goroutines and channels for that so that’s one major point I’d like you to look through. But besides this, feel free to comment on anything: naming, package structure, adherence to general and go-specific programming principles/best practises. Judge as strict as possible, don’t give anything a pass - I might be new to Go, but not to programming in general.

There’s almost no documentation except CLI usage guide, but I think programs in Go have this natural readability, I only have docs for exposed functions and structs in packages, and also small comments here and there on specific logic and whatnot.You can either raise an issue on GH or write it in the comments, whatever works for you. Thanks everyone in advance!
## [7][Go client library for accessing the KSC (Kaspersky Security Center) Open API](https://www.reddit.com/r/golang/comments/i6fxc7/go_client_library_for_accessing_the_ksc_kaspersky/)
- url: https://www.reddit.com/r/golang/comments/i6fxc7/go_client_library_for_accessing_the_ksc_kaspersky/
---
Hello, to begin with, I'm not a programmer, but at work I had to launch this project, at the moment I want to ask for help in continuing development and an indication of errors in the implementation.

[Github repo](https://github.com/pixfid/go-ksc)
## [8][A simple example on implementing progress bar in GoLang](https://www.reddit.com/r/golang/comments/i6eq5e/a_simple_example_on_implementing_progress_bar_in/)
- url: https://www.pixelstech.net/article/1596946473-A-simple-example-on-implementing-progress-bar-in-GoLang
---

## [9][Make my own library or package and create binary files](https://www.reddit.com/r/golang/comments/i6c0b9/make_my_own_library_or_package_and_create_binary/)
- url: https://www.reddit.com/r/golang/comments/i6c0b9/make_my_own_library_or_package_and_create_binary/
---
I'm having a little bit of troubles with GO. i'm new in this language and it's been hard to find the right documentacion to help myself. i don't want to have all my functions in one file because it's a lot of code. other thing... in c language it's so easy the management of the files, in this case a need to create multiple binary files to simulate a hard disks and write in it a struct(mbr) and read this struct in order to make reports. in C i used fwrite(&amp;disk, sizeof(struct MBR), 1, fp); to write de struct and to move the pointer of the file i used fseek(fp, sizeof(struct MBR) + 1, SEEK\_SET); . but i could not find similar fuctions in GO. i hope you can help me, i'll apreciate a lot
## [10][Let's Fork built with Go and React Native](https://www.reddit.com/r/golang/comments/i5e0t9/lets_fork_built_with_go_and_react_native/)
- url: https://v.redd.it/85paupeg3lf51
---

