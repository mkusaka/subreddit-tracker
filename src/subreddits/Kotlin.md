# Kotlin
## [1][Kotlin Multiplatform Mobile (KMM) goes Alpha](https://www.reddit.com/r/Kotlin/comments/j79eyu/kotlin_multiplatform_mobile_kmm_goes_alpha/)
- url: https://thetechjourno.blogspot.com/2020/10/kotlin-multiplatform-mobile-kmm-goes.html
---

## [2][Calling REST API from Kotlin?](https://www.reddit.com/r/Kotlin/comments/j7binh/calling_rest_api_from_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/j7binh/calling_rest_api_from_kotlin/
---
Hi all,

No matter what method I try (googling like a demon), I cannot get this to work.

Nearly every single example I have found online doesn't compile - and the couple that do, simply don't work, or crash the app out with no error the moment I attempt to connect.

Someone here must know the easiest way to make a simple GET request to a third-party API and iterate through the JSON result?

I'm starting to wonder if it's possible - despite, I should think, nearly every app out there needing to do something similar.

So far, this is the closest I've got:

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val text = "Hello toast!"
        val duration = Toast.LENGTH_SHORT
        val toast = Toast.makeText(applicationContext, text, duration)
        toast.show()

        val button = findViewById&lt;Button&gt;(R.id.button)
        button?.setOnClickListener()
        {
            Toast.makeText(
                this@MainActivity,
                "hello again", Toast.LENGTH_LONG
            ).show()
            sentGet()
        }
    }

    private fun sentGet() {
        val url = URL("http://www.google.com/")

        with(url.openConnection() as HttpURLConnection) {
            requestMethod = "GET"  // optional default is GET

            println("\nSent 'GET' request to URL : $url; Response Code : $responseCode")

        }
    }

But when I click the button, the app simply disappears without a word.
## [3][Interested in writing Kotlin articles, get paid to write tutorials, follow the link.](https://www.reddit.com/r/Kotlin/comments/j7b5lp/interested_in_writing_kotlin_articles_get_paid_to/)
- url: https://developers.decoded.africa/get-paid-to-write-tutorials/
---

## [4][Help! Can't even get a "toast" pop up working.](https://www.reddit.com/r/Kotlin/comments/j79qd5/help_cant_even_get_a_toast_pop_up_working/)
- url: https://www.reddit.com/r/Kotlin/comments/j79qd5/help_cant_even_get_a_toast_pop_up_working/
---
Hi all,

I've been going at this for over an hour now. I'm just running through tutorials etc..., but nothing seems to be working at all.

I found this:

https://developer.android.com/guide/topics/ui/notifiers/toasts

And then tried this:

    view.findViewById&lt;Button&gt;(R.id.button_toast).setOnClickListener {
            val text = "Hello toast!"
            val duration = Toast.LENGTH_SHORT
            val toast = Toast.makeText(this,text, duration)
            toast.show()
        }


Which then moaned:

    one of the following functions can be called with the arguments supplied: 
    public open fun makeText(p0: Context!, p1: CharSequence!, p2: Int): Toast! defined in 
        android.widget.Toast
    public open fun makeText(p0: Context!, p1: Int, p2: Int): Toast! defined in android.widget.Toast


So, I tried this:

    fun Context.toast(message: CharSequence) =
            Toast.makeText(this, "hello", Toast.LENGTH_SHORT).show()

Which doesn't moan about anything, but just doesn't work. I can step over/into it and there's no error or anything, just no popup.

What do you have to do to get a popup working in Kotlin?

TIA.
## [5][Kotlin JSON to Data class - JSON Serialization with Jackson Library](https://www.reddit.com/r/Kotlin/comments/j6vf5s/kotlin_json_to_data_class_json_serialization_with/)
- url: https://rrtutors.com/tutorials/Kotlin-JSON-Serialization-with-Jackson
---

## [6][Can Ktor deal with JSON-API?](https://www.reddit.com/r/Kotlin/comments/j6u9o8/can_ktor_deal_with_jsonapi/)
- url: https://www.reddit.com/r/Kotlin/comments/j6u9o8/can_ktor_deal_with_jsonapi/
---
I have someone elses API I want to interact with and am looking to build a client using Ktor for it.

Their api is [JSON-API conform data](https://jsonapi.org/). There is already an existing client based on Spring Boot and looking through the code it looks like Spring can get POJOs from JSON-API conforming data basically out of the box.

I've done all manner of google searches and haven't found any example of Ktor doing this. Can someone here give me some pointers? I want to serialize/deserialize POJOS with the JSON-API format.
## [7][Make Conversational AI Work](https://www.reddit.com/r/Kotlin/comments/j6qu4j/make_conversational_ai_work/)
- url: https://www.reddit.com/r/Kotlin/comments/j6qu4j/make_conversational_ai_work/
---
Junction, a huge [hackathon](https://connected.hackjunction.com/) with a challenge for Kotlin, Android, and mobile devs is to be held on 6-8 of November

Hackers have to create a conversational AI solution, using Kotlin-based framework

The challenge prize is €1.500, and the main hackathon prize is €10.000

If you are interested, apply to the challenge – online/offline hubs are available!

[https://just-ai.com/junction-challenge/](https://just-ai.com/junction-challenge/)
## [8][Make Conversational AI Work with Kotlin](https://www.reddit.com/r/Kotlin/comments/j6qs7s/make_conversational_ai_work_with_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/j6qs7s/make_conversational_ai_work_with_kotlin/
---
Junction, a huge [hackathon](https://connected.hackjunction.com/) with a challenge for Kotlin, Android, and mobile devs is to be held on 6-8 of November

Hackers have to create a conversational AI solution, using Kotlin-based framework

The challenge prize is €1.500, and the main hackathon prize is €10.000

If you are interested, apply to the challenge – online/offline hubs are available!

[https://just-ai.com/junction-challenge/](https://just-ai.com/junction-challenge/)
## [9][Gradle Kotlin jvm vs multiplatform plugin](https://www.reddit.com/r/Kotlin/comments/j6p0ux/gradle_kotlin_jvm_vs_multiplatform_plugin/)
- url: https://www.reddit.com/r/Kotlin/comments/j6p0ux/gradle_kotlin_jvm_vs_multiplatform_plugin/
---
I wanna create a Gradle build for my Kotlin project. I saw there is a Kotlin jvm plugin and a multiplatform plugin.

What's the difference? Which one should I use if I intend to target the JVM?
## [10][Doks: Search for your distributed documentation in one place (written in Kotlin)](https://www.reddit.com/r/Kotlin/comments/j6b6gl/doks_search_for_your_distributed_documentation_in/)
- url: https://github.com/wlezzar/doks
---

