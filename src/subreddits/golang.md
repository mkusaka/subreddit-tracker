# golang
## [1][Super Excited - Writing my first Golang book](https://www.reddit.com/r/golang/comments/ivdt43/super_excited_writing_my_first_golang_book/)
- url: https://www.reddit.com/r/golang/comments/ivdt43/super_excited_writing_my_first_golang_book/
---
Hey all :)

Super pumped about this and I want to ensure I keep the community up to date. I'm writing my first Golang book and Apress is the publisher. I'll be posting stories and updates about my journey.

If you'd like to follow along, check out my Twitter:  [https://twitter.com/TheNJDevOpsGuy](https://twitter.com/TheNJDevOpsGuy)
## [2][How can I combine multiple files to one for different clients?](https://www.reddit.com/r/golang/comments/ivqv5d/how_can_i_combine_multiple_files_to_one_for/)
- url: https://www.reddit.com/r/golang/comments/ivqv5d/how_can_i_combine_multiple_files_to_one_for/
---
Hello,

&amp;#x200B;

I have multiple files (around 500-2000) from 5kb to 10gb in size. The average file is about 250mb. Each user Needs to download different files. So I can't just pack it into a Zip.

I don't really want to download each file separately because the connection closing/opening is taking a lot of time and the time adds up for such amounts of files.

&amp;#x200B;

I was wondering if I could dynamically pack it into some sort of file format (Just storage properly) and sent this file to the users that need the files.

The max amount of users that want to download files at once is about 50-80.
## [3][Learning Go and Git together with "gogit", an implementation of git from scratch in Golang.](https://www.reddit.com/r/golang/comments/ivg4a8/learning_go_and_git_together_with_gogit_an/)
- url: https://www.reddit.com/r/golang/comments/ivg4a8/learning_go_and_git_together_with_gogit_an/
---
[https://github.com/ssrathi/gogit](https://github.com/ssrathi/gogit)

I have always liked Git, and have recently started using Golang, both in my personal and professional projects. So the most logical thing to do is to implement parts of Git internals in Golang as a learning exercise.

This is my first large scale Go project. Feedback/comments welcome!
## [4][golang-dev : Implementing Generics](https://www.reddit.com/r/golang/comments/iv4idd/golangdev_implementing_generics/)
- url: https://groups.google.com/forum/#!topic/golang-dev/OcW0ATRS4oM
---

## [5][Is there a better way to do this?](https://www.reddit.com/r/golang/comments/ivdoat/is_there_a_better_way_to_do_this/)
- url: https://www.reddit.com/r/golang/comments/ivdoat/is_there_a_better_way_to_do_this/
---
TLDR; [is this code shit?](https://github.com/oze4/godaddygo)

Lately, I have been learning Go. I come from a mostly Node background with some C#. So far, I have been absolutely loving it.

I suppose I am getting a little "imposter syndrome" and was hoping someone would be kind enough to "review" some code I wrote?

It is an SDK for the GoDaddy API (don't focus on what it is, rather how it is implemented) and is currently rather small.

I just wanted to make sure I am using the correct project structure, using interfaces/structs correctly, and that I am using Go correctly.

[You can find the repo, godaddygo, here](https://github.com/oze4/godaddygo)

edit: if this is the wrong place for this, please let me know!

edit2: why down-vote without commenting? I don't mind criticism, but I do mind lack of feedback. I'm trying to improve here...and if you don't like that ----&gt; there's the door.
## [6][cockroachdb/pebble: RocksDB/LevelDB inspired key-value database in Go](https://www.reddit.com/r/golang/comments/iv67vp/cockroachdbpebble_rocksdbleveldb_inspired/)
- url: https://github.com/cockroachdb/pebble
---

## [7][The Most Popular Programming Languages - 1965/2020](https://www.reddit.com/r/golang/comments/ivp0pd/the_most_popular_programming_languages_19652020/)
- url: https://youtu.be/UNSoPa-XQN0
---

## [8][How do I get started with golang web development](https://www.reddit.com/r/golang/comments/ivctja/how_do_i_get_started_with_golang_web_development/)
- url: https://www.reddit.com/r/golang/comments/ivctja/how_do_i_get_started_with_golang_web_development/
---
Hi guys,

I am a Nodejs developer, highly interested in Golang backend development

can you recommend any books/courses ?

someone told me this [https://www.udemy.com/course/go-programming-language/](https://www.udemy.com/course/go-programming-language/)

but I am not sure if 2017 is too old
## [9][Decoding JSON, the... a way](https://www.reddit.com/r/golang/comments/iv2zpd/decoding_json_the_a_way/)
- url: https://www.reddit.com/r/golang/comments/iv2zpd/decoding_json_the_a_way/
---
 [https://twitter.com/caarlos0/status/1306447785831673867?s=19](https://twitter.com/caarlos0/status/1306447785831673867?s=19)
## [10][Structs in REST API framework](https://www.reddit.com/r/golang/comments/iv67s1/structs_in_rest_api_framework/)
- url: https://www.reddit.com/r/golang/comments/iv67s1/structs_in_rest_api_framework/
---
I used to maintain Ruby on Rails backend and now I have to use Go-Gin + MongoDB and I find it very confusing cos I'm used to the magic in rails

Let's say I have a POST Person API to create a person and PUT Person API to update a person's details.

The request for POST Person API is:

    {
      "name": "John",
      "age": 21,
      "fav_color": "Green",
      "current_pet": {
         "name": Birdie",
         "species": "parrot"
      }
    }

The response of POST and PUT API is:

    {
      "id": "232dwdwd2-231232",
      "name": "John",
      "age": 21,
      "fav_color": "Green",
      "current_pet": {
         "id": "43324dwdwd3d-23edwwd",
         "name": Birdie",
         "species": "parrot"
      }
    }

The request for PUT Person API (PUT person/232dwdwd2-23123) is:

    {
      "fav_color": "Green" // any field that needs to be updated
    }

Although the JSON fields have the same name, for POST Person API, the person's name is required but for PUT Person API, it is not, only the fields to be updated should be supplied. To tackle this problem, I'm writing the structs 3 times in 3 diff files:

    // models/person_post.go
    
    type PetPostResponse struct {
      Name string `json:"name"`
      Species string `json:"species"`
    }
    
    type PersonPostResponse struct {
      Name string `json:"name" binding:"required"`
      Age int `json:"age"`
      FavColor string `json:"fav_color"`
      CurrentPet PetPostResponse `json:"current_pet"`
    }

The PUT Person API structs need BSON omit empty so that the fields that are not updated are not included in the update query of collection.FindOneAndUpdate() so that it doesn't override the old values:

    // models/person_put.go
    
    type PetPutResponse struct {
      Name string `bson:",omitempty" json:"name"`
      Species string `bson:",omitempty" json:"species"`
    }
    
    type PersonPutResponse struct {
      Name string `bson:",omitempty" json:"name" binding:"required"`
      Age int `bson:",omitempty" json:"age"`
      FavColor string `bson:",omitempty" json:"fav_color"`
      CurrentPet PetPutResponse `bson:",omitempty" json:"current_pet"`
    }

&amp;#x200B;

    // models/person.go
    
    type Pet struct {
      Id string `json:"id"`
      Name string `json:"name"`
      Species string `json:"species"`
    }
    
    type Person struct {
      Id primitive.ObjectID `bson:"_id,omitempty" json:"id"`
      Name string `json:"name"`
      Age int `json:"age"`
      FavColor string `bson:fav_color" json:"fav_color"`
      CurrentPet Pet `json:"current_pet"`
    }

I convert the API structs to DB structs using this helper method:

    func ConvertApiToDbStruct(i interface{}) ds.Person {
      apiStruct, ok := i.(ds.PersonPostRequest)
      if !ok {
        return ds.Person{} 
      } 
      p := ds.Person{ 
         Name: apiStruct.Name,
         Age: apiStruct.Age,
         CurrentPet: ds.Pet{ 
           Name: apiStruct.Pet.Name, 
           Species: apiStruct.Pet.Species 
         } 
      } 
      return p
    }

I would like to ask:

1. Is this normal to define the struct 3 times because only the DB struct have IDs and the json and bson tags differ among them? In reality I have 17 fields but I don't know other better way to do this
2. Is there anyway to make the ConvertApiToDbStruct more generic ? So that if I add another Person field, I don't need to add another line of code in that function

Thanks so much and sorry if its confusing
