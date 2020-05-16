# Kotlin
## [1][Problem exporting Typescript definitions in 1.4-M1 (from commonMain)](https://www.reddit.com/r/Kotlin/comments/gkq2uc/problem_exporting_typescript_definitions_in_14m1/)
- url: https://www.reddit.com/r/Kotlin/comments/gkq2uc/problem_exporting_typescript_definitions_in_14m1/
---
I'm having trouble with the new typescript definition export from Kotlin in 1.4-M1 😓

1. Adding @ JsExport outside jsMain source directory in multiplatform Gradle project does not work\*? So this means that one cannot export type definitions for types that are shared between platform targets?
2. When creating a regular kotlin/js project in IntelliJ, I cannot find a way to activate the new  IR back-end compiler and the definition export - I assume this is not supported?

Typescript definition export would simplify my current project dramatically, so I keep hoping it would be possible to get at least one of these options working? 🤞

*\* this image illustrates the problem - is this the expected behavior?*

&amp;#x200B;

https://preview.redd.it/n7trf6uo64z41.png?width=1812&amp;format=png&amp;auto=webp&amp;s=9f8ac3bc3434290cc7ec870949d0bed29edaf397
## [2][Why this type mismatch](https://www.reddit.com/r/Kotlin/comments/gkqmob/why_this_type_mismatch/)
- url: https://www.reddit.com/r/Kotlin/comments/gkqmob/why_this_type_mismatch/
---
Calculating the length of a track in a GPX file:

    override fun startElement(uri: String, lName: String, qName: String, att: Attributes) {
        if (qName != "trkpt")
            return
        att.getValue("lat")?.toDouble().let { lat -&gt;
            att.getValue("lon")?.toDouble().let { long -&gt;
                if (lastLat != null &amp;&amp; lastLong != null)
                    _dist += haversine(lat, long)
                lastLat = lat
                lastLong = long
            }
        }
    }

    private fun haversine(lat: Double, long: Double): Double {
        /* omitted for brevity */
    }

Kotlinc complains about a "type mismatch: inferred type is Double? but Double was expected" in the call to `haversine`. I can fix it easily enough by changing the `let` calls to be like:

        att.getValue("lat")?.toDouble()?.let { lat -&gt;
            att.getValue("lon")?.toDouble()?.let { long -&gt;

But why is it giving me a Double? and not a Double in the first place? If I look at the documentation for `String.toDouble()` it's clearly listed as returning type Double (not nullable). Makes no sense I should have to type the extra question mark.
## [3][Universal JVM GC Log Analyzer free tool](https://www.reddit.com/r/Kotlin/comments/gkp4v0/universal_jvm_gc_log_analyzer_free_tool/)
- url: https://gceasy.io/
---

## [4][Retrofit2 API request never returns. Why?](https://www.reddit.com/r/Kotlin/comments/gkrea4/retrofit2_api_request_never_returns_why/)
- url: https://www.reddit.com/r/Kotlin/comments/gkrea4/retrofit2_api_request_never_returns_why/
---
I have an API service:

    private const val BASE_URL = "https://www.random.org/integers/?num=1&amp;min=1&amp;max=6&amp;col=1&amp;base=10&amp;format=plain&amp;rnd=new"
    interface FooApiService {
      @GET
      suspend fun getInt():
        Call&lt;Int&gt;
    }
    object FooApi {
        val retrofitService : FooApiService by lazy {
            retrofit.create(FooApiService::class.java)
        }
    }

The server returns a normal HTTP200 text/html body response in the browser (I'm using a dummy API here). However I can't get it using the retrofit2 service (the function never returns):

    private var job = Job()
    private val fooScope = CoroutineScope(job + Dispatchers.IO)
    private fun doStuff() {
      fooScope.async {
        FooApi.retrofitService.getInt().execute().body()
        Log.i(TAG, "We never reach here! Why?")
      }
    }

Why?
## [5][What is the point of companion objects?](https://www.reddit.com/r/Kotlin/comments/gkbswd/what_is_the_point_of_companion_objects/)
- url: https://www.reddit.com/r/Kotlin/comments/gkbswd/what_is_the_point_of_companion_objects/
---
Just wander what was the reason for adding such feature to the language. Why were they needed, what problem do they solve?
## [6][A pragmatic Koin crash course](https://www.reddit.com/r/Kotlin/comments/gkfxh7/a_pragmatic_koin_crash_course/)
- url: https://www.rockandnull.com/koin-android-example/
---

## [7][Help with starting out in Kotlin](https://www.reddit.com/r/Kotlin/comments/gkdkyl/help_with_starting_out_in_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/gkdkyl/help_with_starting_out_in_kotlin/
---
Hello, so i was thinking of getting into Android development and chose Kotlin as my preferred language. I don't know Kotlin or Java (just know about Python) and was wondering what resources you guys recommend. I'm currently looking at two books i want to bug. Android Apprentice and Kotlin Apprentice ([these ones](https://imgur.com/a/5PpVZwk)). Which one would be the best or which other books should i go with? I find that books are more systematic and streamlined compared to YouTube videos that's why i leaning towards books instead of tutorials on YouTube.

Any advice on what resources for someone who doesn't know Java or Kotlin would be appreciated. Thanks
## [8][Judge0 IDE adds Kotlin support](https://www.reddit.com/r/Kotlin/comments/gk03pm/judge0_ide_adds_kotlin_support/)
- url: https://ide.judge0.com/?YZyR
---

## [9][Less clunky way of declaring byte literals?](https://www.reddit.com/r/Kotlin/comments/gjt5iv/less_clunky_way_of_declaring_byte_literals/)
- url: https://www.reddit.com/r/Kotlin/comments/gjt5iv/less_clunky_way_of_declaring_byte_literals/
---
I work with a decent amount of byte data, and I'm writing some example byte arrays for unit tests. The default way of doing this in Kotlin is pretty cumbersome:

    val array = byteArrayOf(
        0xDE.toByte(), 0xAD.toByte(), 0xBE.toByte(), 0xEF.toByte()
    )

Is there any more elegant way of doing this? In general I love Kotlin, but this byte declaration syntax and the use of infix functions instead of bitshift operators have been something of a downgrade from Java. Or am I missing something? I've thought of trying to make some kind of convenience function, though I'm not sure if it's any better:

    fun b(value: Int): Byte {
        return value.toByte()
    }

Which then lets me do:

    val array = byteArrayOf(
        b(0xDE), b(0xAD), b(0xBE), b(0xEF)
    )

Am I missing something here or are byte literals just kinda clunky in Kotlin?
## [10][Code a (mini) Twitter clone in Kotlin using the Hexagon Toolkit](https://www.reddit.com/r/Kotlin/comments/gjq174/code_a_mini_twitter_clone_in_kotlin_using_the/)
- url: https://twitter.com/hexagon_kt/status/1260973355362107392
---

