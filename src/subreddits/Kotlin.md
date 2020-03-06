# Kotlin
## [1][I wrote an article about making Kotlin, Jersey, Jetty, and MongoDB work together to create an easy to maintain RESTful API](https://www.reddit.com/r/Kotlin/comments/febsfa/i_wrote_an_article_about_making_kotlin_jersey/)
- url: https://blog.gikken.co/kotlin-jersey-jetty-mongodb-creating-a-restful-api/
---

## [2][What's the reasoning behind Kotlins lack of enforcement of checked exceptions?](https://www.reddit.com/r/Kotlin/comments/fe7tvx/whats_the_reasoning_behind_kotlins_lack_of/)
- url: https://www.reddit.com/r/Kotlin/comments/fe7tvx/whats_the_reasoning_behind_kotlins_lack_of/
---
I always thought java was "good" in the way that it forced you to check exceptions. Kotlin being a bit more modern would likely have some proof that their way is "better", no? note, I've never fully understood exceptions and such in java. I've worked in codebases where they were heavily used, some where they weren't really used at all, and others where we were forced to create our own custom exceptions for like anything that could go wrong.

From a language design aspect, I'm curious if kotlin does this well, or is there some other language that does this better? PS. I'm sure most of this stuff is subjective, but I'm just really curious to hear thoughts and discuss.
## [3][Kotlin Illustrated Guide, Chapter 1 - Variables, Expressions, and Types](https://www.reddit.com/r/Kotlin/comments/fedekh/kotlin_illustrated_guide_chapter_1_variables/)
- url: https://typealias.com/start/kotlin-variables-expressions-types/
---

## [4][Anyone know why I'm unable to bind to my service in this code?](https://www.reddit.com/r/Kotlin/comments/febti7/anyone_know_why_im_unable_to_bind_to_my_service/)
- url: https://www.reddit.com/r/Kotlin/comments/febti7/anyone_know_why_im_unable_to_bind_to_my_service/
---
I'm a kotlin/android noob, so I'm still getting my head wrapped around how all the parts of android development work together (specifically, the bluetooth API). I'm trying to adapt [this half-finished example project](https://github.com/Nithinjith/LEKotlin-Android) to get it working on a BLE heart rate peripheral (the stock example from Espressif, running on an ESP32 dev board).

Here's how I'm managing the BLE connection:

    object BLEConnectionManager {
    
        private val TAG = "BLEConnectionManager"
        private var mBLEService: BLEService? = null
        private var isBind = false
        private val mServiceConnection = object : ServiceConnection {
            override fun onServiceConnected(componentName: ComponentName, service: IBinder) {
                mBLEService = (service as BLEService.LocalBinder).getService()
                Log.i(TAG, "BLEConnectionManager.onServiceConnected mBLEService = $mBLEService")
                if (!mBLEService?.initialize()!!) {
                    Log.e(TAG, "Unable to initialize")
                }
            }
            override fun onServiceDisconnected(componentName: ComponentName) {
                mBLEService = null
            }
        }
    
        fun initBLEService(context: Context) {
            try {
                if (mBLEService == null) {
                    val gattServiceIntent = Intent(context, BLEService::class.java)
                    if (context != null) {
                        isBind = context.bindService(gattServiceIntent, mServiceConnection,
                            Context.BIND_AUTO_CREATE)
                        Log.i(TAG, "BLEConnectionManager.initBLEService isBind = $isBind")
                    }
                }
            } catch (e: Exception) {
                Log.e(TAG, e.message)
            }
        }
    
        fun connect(deviceAddress: String): Boolean {
            var result = false
            Log.i(TAG, "BLEConnectionManager.connect (to $deviceAddress) and mBLEService is $mBLEService")
            if (mBLEService != null) result = mBLEService!!.connect(deviceAddress)
            return result
        }
    // ...etc

And call this in the main activity onCreate:

    if (!BLEDeviceManager.isEnabled()) {
        val enableBtIntent = Intent(BluetoothAdapter.ACTION_REQUEST_ENABLE)
        startActivityForResult(enableBtIntent, REQUEST_ENABLE_BT)
    }
    BLEConnectionManager.initBLEService(this@MainActivity)

The BLEService class is unchanged from the [original code](https://github.com/Nithinjith/LEKotlin-Android/blob/master/app/src/main/java/com/np/lekotlin/blemodule/BLEService.kt).

When I run it, `$mBLEService` is null in `connect`, and I note that when `initBLEService` is called, `$isBind` is always null too. I therefore think I don't properly understand the mechanics of how services work and what they do in this instance. So, why am I unable to bind the service, and is the binding problem the cause of the service being null?
## [5][Kotlin Game Development questions](https://www.reddit.com/r/Kotlin/comments/fdwz5s/kotlin_game_development_questions/)
- url: https://www.reddit.com/r/Kotlin/comments/fdwz5s/kotlin_game_development_questions/
---
So I was curious as to how to approach this. I love exploring new stuff and game development was my love for a while. I then switched and went into android dev because it is a bit more employable but I'd like to come back a bit into game dev. I am planning of using Libgdx ( or ktx), any other recommendations for the libraries/ frameworks that I can use I'll gladly hear it out. On the other note, how would you recommend going about networking and multiplayer? I am not too familiar with this and this has been daunting me forever. I get it that you can use basic sockets, but how would for example a gRPC server work? I am most comfortable with Kotlin so I was hoping to do everything in it. Both networking and the clients. I am looking for your experiences and any tips / recommendations. I am willing to learn more technologies.  
Edit: Can you use practices like dependancy injcetion in game dev?
## [6][Any Kotlin WASM examples?](https://www.reddit.com/r/Kotlin/comments/fea4cb/any_kotlin_wasm_examples/)
- url: https://www.reddit.com/r/Kotlin/comments/fea4cb/any_kotlin_wasm_examples/
---
I know Kotlin can be compiled to WASM via the Kotlin-native target, but I can't find any examples in the wild...
## [7][From RxJava to Kotlin Flow: Testing](https://www.reddit.com/r/Kotlin/comments/fduwyu/from_rxjava_to_kotlin_flow_testing/)
- url: https://proandroiddev.com/from-rxjava-to-kotlin-flow-testing-42f1641d8433
---

## [8][Newcomer Tutorial: Setting up IDEA and Creating "hello world" Application](https://www.reddit.com/r/Kotlin/comments/fdyqdg/newcomer_tutorial_setting_up_idea_and_creating/)
- url: https://marcuseisele.com/pages/learningKotlin/setup-intellij-hello-kotlin
---

## [9][[QUESTION] Async function fire task in new thread](https://www.reddit.com/r/Kotlin/comments/fe1bjn/question_async_function_fire_task_in_new_thread/)
- url: https://www.reddit.com/r/Kotlin/comments/fe1bjn/question_async_function_fire_task_in_new_thread/
---
Hey, I need to write a function that starts some process in another thread, in a fire and forget way, and returns without waiting for said process to finish.

I tried searching online, but I find many examples for cases where the program needs to wait for a response or does thread pooling etc., and none for my case. 

Does anybody know a neat way to do this in Kotlin?
I would really apreciate a code example, pseudocode or the right term that I can use to look up examples of that.
## [10][What is the IDE experience for Kotilin compared to Java](https://www.reddit.com/r/Kotlin/comments/fdy4dm/what_is_the_ide_experience_for_kotilin_compared/)
- url: https://www.reddit.com/r/Kotlin/comments/fdy4dm/what_is_the_ide_experience_for_kotilin_compared/
---
I would expect it to be similar but is Kotlin on par with Java in terms of IDE features in IntelliJ or Eclipse ?

&amp;#x200B;

Edit: I am aware that IntelliJ is made by JetrBrains and that they also make Kotlin. That does not however mean that the IDE experience is perfect. IntelliJ has had support for Java in over 10+ years probably more. That is why i ask this question. 

&amp;#x200B;

Thanks
