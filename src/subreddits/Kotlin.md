# Kotlin
## [1][KotlinConf 2019 Videos](https://www.reddit.com/r/Kotlin/comments/ebd0np/kotlinconf_2019_videos/)
- url: https://www.youtube.com/playlist?list=PLQ176FUIyIUY6SKGl3Cj9yeYibBuRr3Hl
---

## [2][MediaPlayer is finalised when I make separate API call with a lot of data?](https://www.reddit.com/r/Kotlin/comments/em9kti/mediaplayer_is_finalised_when_i_make_separate_api/)
- url: https://www.reddit.com/r/Kotlin/comments/em9kti/mediaplayer_is_finalised_when_i_make_separate_api/
---
Hi guys,

having a strange thing happen that i can't work out. Any help greatly appreciated.

I have a MediaPlayer running as part of an object. It continues playing when i change fragments and but when I change to a fragment that has to load a recyclerView with a lot of data in it being pulled from an API and start to scroll, the MediaPlayer finalises and stops playing. It doesn't do it on fragments with smaller amounts of data in. Any ideas why this might be happening?

    object AudioService {
    
    
     fun playRadio(view: View, url: String) = try {
            var mp: MediaPlayer? = null
    
                mp = MediaPlayer().apply {
    
    
                    setOnPreparedListener{
                        it.start()
    
                    }
                }
    
                mp!!.reset()
                mp!!.setDataSource(url)
                mp!!.prepareAsync()
    
    
        } catch (ex: Exception) {
    
        }
    
    
    
    }
    
    // below is the offending Fragment
    
    class ArchiveFragment : Fragment() {
    
        override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    
    
    
        ): View? {
    
            return inflater.inflate(R.layout.fragment_archive, container, false)
    
    
        }
    
    
        override fun onStart() {
            super.onStart()
            recyclerView_ArchiveMain.layoutManager =  LinearLayoutManager(activity)
            fetchJson()
    
        }
    
        fun fetchJson() {
    
            val url = "https://b9a12dw32wdz137hoj.execute-api.eu-west-2.amazonaws.com/"
            val request = Request.Builder().url(url).build()
            val client = OkHttpClient()
            client.newCall(request).enqueue(object : Callback {
                override fun onResponse(call: Call, response: Response) {
                    val body = response.body?.string()
                    val gson = GsonBuilder().create()
                    val theList: List&lt;Episodes&gt; = gson.fromJson(body, Array&lt;Episodes&gt;::class.java).toList()
    
                    var episodes = theList.reversed()
    
                   activity?.runOnUiThread { recyclerView_ArchiveMain.adapter =
                        ArchiveAdapter(episodes) }
    
                }
    
                override fun onFailure(call: Call, e: IOException) {
                    println("Oh shit it failed")
                }
            })
    
        }
    
    
    }
## [3][Reaktive 1.1.7 introduces DisposableScope](https://www.reddit.com/r/Kotlin/comments/em92eq/reaktive_117_introduces_disposablescope/)
- url: https://www.reddit.com/r/Kotlin/comments/em92eq/reaktive_117_introduces_disposablescope/
---
One of the biggest complaints about RxJava is the manual Subscription and Disposable management. Reaktive 1.1.7 addresses the problem by introducing DisposableScope. Check out the [updated README](https://github.com/badoo/Reaktive#subscription-management-with-disposablescope) to learn more. Also added ReplaySubject and UnicastSubject.
## [4][Converting PyTorch float tensor to Android RGBA Bitmap with Kotlin](https://www.reddit.com/r/Kotlin/comments/elsumw/converting_pytorch_float_tensor_to_android_rgba/)
- url: https://medium.com/@philipplies/converting-pytorch-float-tensor-to-android-rgba-bitmap-with-kotlin-ffd4602a16b6
---

## [5][Kotlin First Impressions](https://www.reddit.com/r/Kotlin/comments/elfqqs/kotlin_first_impressions/)
- url: https://www.reddit.com/r/Kotlin/comments/elfqqs/kotlin_first_impressions/
---
In short, I’m impressed.

I’ve been doing software development for over three decades. Long ago, I had the chance to do some work in Lisp, and thoroughly enjoyed the power and productivity that functional programming enabled. Unfortunately, Lisp tends to turn off many programmers, simply because of its unconventional syntax (which isn’t hard to get used to, but is unconventional, and as the saying goes “you never get a second chance to make a first impression”).

Kotlin is sneaky. It’s a very Scheme-ish functional programming language disguised as a conventional language. That makes it much less likely to inspire the sort of negative gut reactions that the Lisp family of languages does. (Ruby tried to do the same thing, but falling for the anti-patterns of Do What I Mean and There Is More Than One Way to Do It really hurt that language, IMHO.)

The biggest handicap that any new programming language has is libraries. It’s the existing libraries you can leverage that do so much to drive the decision of what programming language is best for a project, so starting out from scratch with a brand-new and mostly empty ecosystem of libraries sets up a simply huge hurdle for any new programming language to overcome. Targeting the JVM was an enormous win. (Getting endorsed by Google as the preferred language for Android app development was another coup, but it would have never happened without the one of targeting the JVM.)

The question is, what happens now. With great power comes great responsibility. Most programmers are mediocre programmers; will they tend to shy away from the more powerful features of Kotlin and mostly write Java code without semicolons, or will they tend to carelessly play with creating new syntax-level features and create unmaintainable code? The answer to that question will, I think, be the determining factor as to how popular and widespread Kotlin becomes.
## [6][Let's Review: Pokedex - zsmb.co](https://www.reddit.com/r/Kotlin/comments/eld6sb/lets_review_pokedex_zsmbco/)
- url: https://zsmb.co/lets-review-pokedex/
---

## [7][What tool do the presenters use in KotlinConf](https://www.reddit.com/r/Kotlin/comments/el742o/what_tool_do_the_presenters_use_in_kotlinconf/)
- url: https://www.reddit.com/r/Kotlin/comments/el742o/what_tool_do_the_presenters_use_in_kotlinconf/
---
I want to give a talk at my company to embrace Kotlin. We are currently a Java shop and I believe that since we are going to be addressing some technical debt, we might as well clean stuff up.

I am supposed to give a presentation on Kotlin soon and I wanted to know how do the presenters at KotlinConf make the presentations. In particular, how do they get the code snippets to fold, expand, and change as they give their talks. 

Thanks in advance
## [8][EvoMaster: Automated Test Generation for Java/Kotlin RESTful APIs](https://www.reddit.com/r/Kotlin/comments/el935m/evomaster_automated_test_generation_for/)
- url: https://github.com/EMResearch/EvoMaster
---

## [9][Android App architecture. Everything in kotlin vs client - server interaction](https://www.reddit.com/r/Kotlin/comments/elc5p4/android_app_architecture_everything_in_kotlin_vs/)
- url: https://www.reddit.com/r/Kotlin/comments/elc5p4/android_app_architecture_everything_in_kotlin_vs/
---
I am trying to build an android app which will include google maps API, some user data storing, and payments.

I have zero experience with kotlin and android development so I would like to know what is the best practice.

Is it safe to put the business logic including payments in the kotlin code? 

As an alternative, I was thinking to create a server in another language, i.e Java using akka-http for example and make REST API calls from the android client. This would be a nice option if for example I want to make an iOS App out of it in the future.

&amp;#x200B;

Any suggestions?
## [10][OpenAPI/Swagger specs (and clients for specs) using Javalin and Maven](https://www.reddit.com/r/Kotlin/comments/ekv8ht/openapiswagger_specs_and_clients_for_specs_using/)
- url: https://javalin.io/tutorials/openapi-example?language=kotlin
---

## [11][Custom Android XML Lint Rules with Kotlin](https://www.reddit.com/r/Kotlin/comments/ekwlfr/custom_android_xml_lint_rules_with_kotlin/)
- url: https://coderamblings.dev/posts/custom-android-xml-lint-rules-with-kotlin/
---

