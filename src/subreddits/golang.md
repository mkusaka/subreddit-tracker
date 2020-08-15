# golang
## [1][How to find a job as a Go developer?](https://www.reddit.com/r/golang/comments/i9ma30/how_to_find_a_job_as_a_go_developer/)
- url: https://i.redd.it/7ndi0qt33zg51.png
---

## [2][How to update field in mongodb of type map[string]interface{} , Golang ?](https://www.reddit.com/r/golang/comments/ia6rtc/how_to_update_field_in_mongodb_of_type/)
- url: /r/mongodb/comments/ia6foo/how_to_update_field_in_mongodb_of_type/
---

## [3][Swimming in Datalakes with Go](https://www.reddit.com/r/golang/comments/i9oxm8/swimming_in_datalakes_with_go/)
- url: https://fraugster.com/blog/post/future-proofing-our-storage-needs-with-aws-athena-apache-parquet
---

## [4][I also push directly to prod!](https://www.reddit.com/r/golang/comments/i91ynq/i_also_push_directly_to_prod/)
- url: https://i.redd.it/ql3zpki3gsg51.png
---

## [5][Upload images to cloud storage service](https://www.reddit.com/r/golang/comments/i9y426/upload_images_to_cloud_storage_service/)
- url: https://www.reddit.com/r/golang/comments/i9y426/upload_images_to_cloud_storage_service/
---
Hello everyone, i want to upload an image received from a form input to a cloud service like cloudinary, but i am unable to find any suitable library for doing that.
What library do you use to perform file uploads or which cloud storage service do you use with your golang applications.
Thanks
## [6][How do you manage the DB object instance in your microservice while handling concurrency?](https://www.reddit.com/r/golang/comments/i9r7j8/how_do_you_manage_the_db_object_instance_in_your/)
- url: https://www.reddit.com/r/golang/comments/i9r7j8/how_do_you_manage_the_db_object_instance_in_your/
---
Not sure if the subject makes sense.. so more explanation here. Basically I have a simple microservice.. it uses the DB in CRUD fashion. But.. it has to handle potentially hundreds of requests at the same time.

I turned to PGX as it is supposedly faster (for Postgresql), and it has a connection pool for handling concurrency. That's great.

What confuses me is how to maintain the DB object I use to GET connections from, across threads. If it's not obvious, I am using this in an API back end.. so each request is running in it's own thread (or go func).

So let me ask this.. in my code, main.go, I have something like this:

 pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE\_URL")) 

First question I have is.. the context.Background()... shouldn't that be set to the context that my request is using? The issue I see here is.. the above call should be done once (on service start), where as each request comes in randomly and has a new context for it. So I am not sure how I pass my request context to the DB so that if some other routine cancels the request, the DB code that may be running ALSO cancels. There is this notion that maybe a long running request (e.g. one that the socket to the client stays open for some period of time like web socket) can be cancelled. If the DB is doing a billion record search..but the request is no longer of use.. I assume cancelling it via context Cancel() should propagate to the DB code to just terminate the search..thus reducing the use of resources when not needed. Is this accurate? I am unclear HOW you make this happen though.

Anyway, so, I have this in my code:

    type Datasource struct {
      ConPool *Pool
    }
    
    var db Datasource = Datasource{}
    
    func main() {
        // load env
        pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL")) 
        // no err
        db.ConPool = pool
    }
    
    func GetPool() (Datasource, error) {
      return db, nil
    }

I realize the code above is not perfect.. e.g. I would handle errors, check for nil/empty, etc. Trying to reduce the amount.

Anyway, so in my api handler/router code, I am not sure exactly how I do things to pass the context that is part of the request, to the db.. but I assume it is something like this:

 

    func myHandler(w http.ResponseWriter, r *http.Request) {
      conn, err := GetPool().ConPool.Acquire(r.Context())
      // no error
      defer conn.Release()
      // do something with request level connection
    }

  So the two questions I am trying to understand/solve.. first.. is using a global var for the Datasource struct the right way.. and OK in that it holds the one time setup for the connection pool? In other words, I only have to connect to the DB one time during the existence of the running service (short of some connection issue.. which I should ALSO account for and reconnect should it fail, etc)? Or do I have to.. on every single incoming request, establish another connection to the DB? I assume its the former.. as it doesnt seem to make sense that you would have a connection pool to pull from if you have to do the heavier connection to DB every time?

Second, when acquiring a connection or getting one.. is using r.Context() the right way to pass the request context.. and WILL that allow the DB code to "stop" mid stream should the context Cancel() be called.. thus properly freeing up DB resources/connection pool (con returns to pool), etc?

&amp;#x200B;

Thanks.
## [7][From Python to Go: migrating our entire API](https://www.reddit.com/r/golang/comments/i9fo8f/from_python_to_go_migrating_our_entire_api/)
- url: https://www.repustate.com/blog/migrating-entire-api-go-python/
---

## [8][Go 1.15's interface optimization for small integers is invisible to Go programs](https://www.reddit.com/r/golang/comments/i9ll67/go_115s_interface_optimization_for_small_integers/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/Go115InterfaceSmallIntsII
---

## [9][Use Go or Java for back-end machine learning processes to post results to React front-end?](https://www.reddit.com/r/golang/comments/i9ug2k/use_go_or_java_for_backend_machine_learning/)
- url: https://www.reddit.com/r/golang/comments/i9ug2k/use_go_or_java_for_backend_machine_learning/
---
Hi Everyone, I have wrote a back-end machine learning process in Python but due to the global interpreter lock (GIL) , it isn't going to scale for production purposes.

Been looking at  converting the Python code into either Go or Java.  From what I have seen so far, both languages can carry out the processes that I need.

Which of the two languages is faster in Tesseract OCR and linear algebra processes and makes it easy to print the results in a React framework?

Thank you for your time.
## [10][SSH Port Forwarding Tool with Resiliency and UX](https://www.reddit.com/r/golang/comments/i9tjzy/ssh_port_forwarding_tool_with_resiliency_and_ux/)
- url: https://davrodpin.github.io/mole/
---

