# golang
## [1][Announcing the 2020 Go Developer Survey](https://www.reddit.com/r/golang/comments/jeuosg/announcing_the_2020_go_developer_survey/)
- url: https://blog.golang.org/survey2020
---

## [2][Just published my first desktop application. I noticed that there wasn't any open source and easy to use software for deflickering timelapse footage. So I wrote a tool in Go! It's only version 0.1.0 but I hope it's a good start.](https://www.reddit.com/r/golang/comments/jjblds/just_published_my_first_desktop_application_i/)
- url: https://github.com/StruffelProductions/simple-deflicker
---

## [3][GO-ORA: An Oracle client in Pure Go. No CGO, faster compile time and no extra libraries dependency](https://www.reddit.com/r/golang/comments/jjbvbg/goora_an_oracle_client_in_pure_go_no_cgo_faster/)
- url: https://github.com/sijms/go-ora
---

## [4][How to share test utilities accros packages ?](https://www.reddit.com/r/golang/comments/jjmsoc/how_to_share_test_utilities_accros_packages/)
- url: https://www.reddit.com/r/golang/comments/jjmsoc/how_to_share_test_utilities_accros_packages/
---
Hi,

I have these packages

* db
* api
* main (but not relevant for this question)

&amp;#x200B;

Both packages db and api need to connect to db for running the tests  (api is calls to gin integration test more). The first solution was to dublicate the connection to mysql code in main\_test.go of each package which is very ugly as code is exactly the same.

Then now I am trying to find a way to share testutils accross package

&amp;#x200B;

I have the ConnectToDb method inside `db/main_test.go`

    package db
    
    // imports here
    
    func ConnectToDb()  *sql.DB {
    	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
    	if err != nil {
    		log.Panic(err)
    	}
    	return ConnectToDbOrDie(Config{
    		Driver:   os.Getenv("DB_DRIVER"),
    		Host:     os.Getenv("DB_HOST"),
    		Port:     port,
    		Name:     "", // because want to only access then create the random test db named
    		User:     os.Getenv("DB_USERNAME"),
    		Password: os.Getenv("DB_PASSWORD"),
    	})
    }

&amp;#x200B;

The problem is that I cannot use this method inside `api/main_test.go` because it is not exported

Additionally I don't want to put that function outside the \_test.go file because don't want to export to the clients of my library!

And also this function is unique only for connecting to test database different from other function , so must not get exported, and just be internal and shared accross different packages that need to use it in their tests
## [5][Ebiten (A dead simple 2D game library for Go) v2.0.0 released: no new features, but the API is cleaned up](https://www.reddit.com/r/golang/comments/jj2116/ebiten_a_dead_simple_2d_game_library_for_go_v200/)
- url: https://ebiten.org/blog/v2.0.0.html
---

## [6][A golang library to detect differences/similarities between two string. Detecting if two string is "the same" or "similar". For example, if you wants to detect the user putting bad-word as user name, or to forbid the use of unwanted words. This lib implements, some popular algorithm for you.](https://www.reddit.com/r/golang/comments/jjj04g/a_golang_library_to_detect/)
- url: https://github.com/hyperjumptech/beda
---

## [7][A library for creating and storing idempotency keys.](https://www.reddit.com/r/golang/comments/jjgnzn/a_library_for_creating_and_storing_idempotency/)
- url: https://github.com/catmullet/one
---

## [8][amfora - Gemini client (gopher successor)](https://www.reddit.com/r/golang/comments/jjf4p9/amfora_gemini_client_gopher_successor/)
- url: https://github.com/makeworld-the-better-one/amfora
---

## [9][Looking for an idiomatic Go way to get progress from multiple workers](https://www.reddit.com/r/golang/comments/jjkvnw/looking_for_an_idiomatic_go_way_to_get_progress/)
- url: https://www.reddit.com/r/golang/comments/jjkvnw/looking_for_an_idiomatic_go_way_to_get_progress/
---
Hi,

I have a program that scans the filesystem for files, examines those, and if they match a requirement, copies those files to a new location.

I'm trying to implement this using a producer-consumer pattern. It works fine, but following the progress of the long running operations is troublesome for me.

Here's the problematic part:

    var wg sync.WaitGroup
    jobs := make(chan *scanJob, config.Workers*config.Workers)
    copyOps := make(chan int, config.Workers*config.Workers)
    runningWorkers := config.Workers
    for i := 0; i &lt; config.Workers; i++ {
        wg.Add(1)
        go worker(&amp;wg, jobs, copyOps)
    }
    go scanner(&amp;wg, jobs)
    
    // Follow the progress.
    copyCount := 0
    for value := range copyOps {
        if value &gt; 0 {
            copyCount++
        } else {
            log.Info().Msg("A worker finished")
            runningWorkers--
            if runningWorkers &lt; 1 {
                break
            }
        }
    }
    
    // We're done.
    wg.Wait()
    close(copyOps)
    log.Info().Int("files", copyCount).Msg("Copied")

The current implementation is such that each worker writes integer 1 to the copyOps channel to indicate that a file has been processed. A worker writes integer 0 to the copyOps channel once it is done with its `for job := range jobs` \-loop. When a zero appears on the copyOps channel I decrease the value of `runningWorkers`. Once it reaches zero, I know that all workers are done, and I can break out of the progress loop.

Is there a more idiomatic Go way of achieving this? I know, that `wg.Wait` will wait for workers to finish. But if I wait for that to happen, I would print the progress after all operations are done, and that does not seem like a good solution. I would like to report the progress *while* it's happening.
## [10][text editor library ?](https://www.reddit.com/r/golang/comments/jjl3wj/text_editor_library/)
- url: https://www.reddit.com/r/golang/comments/jjl3wj/text_editor_library/
---
The javascript ecosystem has many libraries that are designed to build custom text editors. However, I failed to find something similar for go, a library designed to build custom, terminal based text editors for a specific purpose.

On one hand, I found multiple tui toolkits, but they are not really designed with text editing in mind. It is certainly possible to use them to build a text editor, but that would require quite a lot of plumbing.

On the other end, I found [micro](https://github.com/zyedidia/micro), which has pretty much everything necessary, but isn't designed as a library.

Did I miss something? How would you build a custom text editor?
## [11][I need an advice for setting an interface{} value](https://www.reddit.com/r/golang/comments/jjkmlo/i_need_an_advice_for_setting_an_interface_value/)
- url: https://www.reddit.com/r/golang/comments/jjkmlo/i_need_an_advice_for_setting_an_interface_value/
---
Hello i have a Tag struct which has data type, value and bytes fields. I am setting the value and bytes of value with SetValue and SetBytes fuctions. When i set the value i also set bytes and vice versa.

If there is a better or idiomatic way, please advice.

Here is the code:
```go
//data types
const (
	Bit     uint8 = 1
	Uint8   uint8 = 2
	Int8    uint8 = 3
	//...
)

type Tag struct {
	//...
	DataType     uint8
	bytes        []byte
	value        interface{}
}

func (t *Tag) SetBytes(bytesOfValue []byte) error {
	//try to set value
	var value interface{}
	b := bytes.NewReader(bytesOfValue)
	var err error
	switch t.DataType {
	case Bit:
		var v bool
		err = binary.Read(b, binary.BigEndian, &amp;v)
		value = v
	case Uint8:
		var v uint8
		err = binary.Read(b, binary.BigEndian, &amp;v)
		value = v
	case Int8:
		var v int8
		err = binary.Read(b, binary.BigEndian, &amp;v)
		value = v
		//...
	}
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("reading bytes failed")
	}
	t.value = value
	//set bytes
	t.bytes = bytesOfValue
	return nil
}

func (t *Tag) SetValue(value interface{}) error {
	//try to set bytes
	var v interface{}
	var ok bool
	switch t.DataType {
	case Bit:
		v, ok = value.(bool)
	case Uint8:
		v, ok = value.(uint8)
	case Int8:
		v, ok = value.(int8)
	//...
	}
	if !ok {
		return errors.New("value is not valid")
	}
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.BigEndian, v)
	if err != nil {
		return errors.New("writing bytes failed")
	}
	t.bytes = b.Bytes()
	//set value
	t.value = value
	return nil
}
```
And i set the value and bytes of value like this:
```go
//set value
t.SetValue(int8(-128))

//set bytes
t.SetBytes([]byte{128})
```
Playground link: https://play.golang.org/p/hHpTIw9Fxx5

Best regards
