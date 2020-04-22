# golang
## [1][Coming from Javascript and I'm on a Go project now. I'm flummoxed by pointers and references and so on. Anyone got a good blog post that will help?](https://www.reddit.com/r/golang/comments/g5o3k4/coming_from_javascript_and_im_on_a_go_project_now/)
- url: https://www.reddit.com/r/golang/comments/g5o3k4/coming_from_javascript_and_im_on_a_go_project_now/
---
Did the tour, read Effective Go, still perplexed. I'm slow to learn new languages. My background is Node and C#. Last time I did C++ was 15 years ago in college.
## [2][linfangrong/redismodule-ratelimit: Redis module that provides ratelimit in Go](https://www.reddit.com/r/golang/comments/g60012/linfangrongredismoduleratelimit_redis_module_that/)
- url: https://github.com/linfangrong/redismodule-ratelimit
---

## [3][What is the future of the Golang programming language for 2020-2021?](https://www.reddit.com/r/golang/comments/g5zm8u/what_is_the_future_of_the_golang_programming/)
- url: https://www.reddit.com/r/golang/comments/g5zm8u/what_is_the_future_of_the_golang_programming/
---
 

# What is the future of the Golang programming language for 2020-2021? Does it make sense to learn it as the first language for backend development? Can I find a job as a junior in the USA?
## [4][Using 'Go Generate' To Deploy Multi-Process Apps](https://www.reddit.com/r/golang/comments/g613zk/using_go_generate_to_deploy_multiprocess_apps/)
- url: https://qvault.io/2020/04/22/using-go-generate-to-deploy-multi-process-apps/
---

## [5][Image uploading with go?](https://www.reddit.com/r/golang/comments/g5zir7/image_uploading_with_go/)
- url: https://www.reddit.com/r/golang/comments/g5zir7/image_uploading_with_go/
---
Hello guys!

I'm creating a REST API with go for an iOS application and I want the option to upload and download images.

I saw one GitHub repo but it was about 4 years old.

I thought about uploading the image to Google's Firebase from the client, and when I get the url to send a request to the server, and save the url I got in the database so I can have the url whenever I need it.

One problem I got with it is that for one image upload, I need two network calls, and that does not seem so efficient.
## [6][a library for creating ISO disk images in Go](https://www.reddit.com/r/golang/comments/g5c9ut/a_library_for_creating_iso_disk_images_in_go/)
- url: https://github.com/kdomanski/iso9660
---

## [7][Ebook for backend](https://www.reddit.com/r/golang/comments/g5s2jq/ebook_for_backend/)
- url: https://www.reddit.com/r/golang/comments/g5s2jq/ebook_for_backend/
---
Hi everyone! I'm learning go and I wanna use it for backend so I wonder if there's a free ebook for using go with mySQL and mongoDB
## [8][How to handle unsafe pointers when recursively generate a hash key string for caching?](https://www.reddit.com/r/golang/comments/g5vixh/how_to_handle_unsafe_pointers_when_recursively/)
- url: https://www.reddit.com/r/golang/comments/g5vixh/how_to_handle_unsafe_pointers_when_recursively/
---
So maybe I'm doing this wrong, but I need to cache queries to Redis before the actual query is called. At the moment I wrote this function which works completely fine until it's given an unsafe pointer. Triggers `reflect: call of reflect.Value.Elem on unsafe.Pointer Value` when I use the `v.Elem()` with it. I need it to nest through the value type it finds and converts it to a string which then I call MD5 hash function on.

    func GenerateQueryCacheKey(args ...interface{}) string {
    	var argumentString = ""
    
    	for _, arg := range args {
    		argumentString += generateRawValueKey(reflect.ValueOf(arg))
    	}
    
    	h := md5.New()
    	io.WriteString(h, argumentString)
    
    	return fmt.Sprintf("%x", h.Sum(nil))
    }
    
    func generateRawValueKey(v reflect.Value) string {
    
    	fmt.Println("generateRawValueKey", v.Kind())
    	switch v.Kind() {
    	case reflect.Bool:
    		return strconv.FormatBool(v.Bool())
    	case reflect.String:
    		return v.String()
    	case reflect.Int, reflect.Int32, reflect.Int64:
    		return strconv.Itoa(int(v.Int()))
    	case reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint32:
    		return strconv.Itoa(int(v.Uint()))
    	case reflect.Float64, reflect.Float32:
    		return strconv.FormatFloat(v.Float(), 'E', -1, 32)
    	case reflect.Struct:
    		token := ""
    		for i := 0; i &lt; v.NumField(); i++ {
    			token += generateRawValueKey(v.Field(i))
    		}
    
    		return token
    	case reflect.Ptr, reflect.UnsafePointer:
    		if v.IsNil() {
    			return "null"
    		}
    
    		return generateRawValueKey(v.Elem()) // &lt;&lt;&lt; [PANIC RECOVER] reflect: call of reflect.Value.Elem on unsafe.Pointer Value
    	case reflect.Array, reflect.Slice:
    		token := ""
    
    		for i := 0; i &lt; v.Len(); i++ {
    			token += generateRawValueKey(v.Field(i))
    		}
    
    		fmt.Println("Array token:", token)
    		return token
    	case reflect.Map:
    		token := ""
    		for _, key := range v.MapKeys() {
    			token += generateRawValueKey(key)
    			token += generateRawValueKey(v.MapIndex(key))
    		}
    
    		return token
    	case reflect.Func:
    		return ""
    	default:
    		return "nil"
    	}
    }

I'm not sure if it means anything, but the struct that is being passed into this also has a `reflect.Value` inside which keeps that unsafe pointer I assume which is giving me issues. [https://github.com/go-pg/pg/blob/master/types/array.go#L9](https://github.com/go-pg/pg/blob/master/types/array.go#L9)
## [9][How to Manage Database Timeouts and Cancellations in Go](https://www.reddit.com/r/golang/comments/g59fdv/how_to_manage_database_timeouts_and_cancellations/)
- url: https://www.alexedwards.net/blog/how-to-manage-database-timeouts-and-cancellations-in-go
---

## [10][WebRTC For Gophers: Learn to build sub-second decentralized real-time communications software](https://www.reddit.com/r/golang/comments/g4xsws/webrtc_for_gophers_learn_to_build_subsecond/)
- url: https://www.youtube.com/watch?v=FdgoOrJH8ok&amp;feature=youtu.be&amp;t=989
---

