# Kotlin
## [1][Creating a BitTorrent client in Kotlin](https://www.reddit.com/r/Kotlin/comments/hg7ym2/creating_a_bittorrent_client_in_kotlin/)
- url: https://kcianfarini.dev/corbit-bencoding/
---

## [2][I have worked with Python extensively. Can any one suggest me resources to get started with Kotlin?](https://www.reddit.com/r/Kotlin/comments/hgpao5/i_have_worked_with_python_extensively_can_any_one/)
- url: https://www.reddit.com/r/Kotlin/comments/hgpao5/i_have_worked_with_python_extensively_can_any_one/
---
I would prefer it would direct me somehow toward mobile application development for Android.
## [3][A question about nulls](https://www.reddit.com/r/Kotlin/comments/hgjdy1/a_question_about_nulls/)
- url: https://www.reddit.com/r/Kotlin/comments/hgjdy1/a_question_about_nulls/
---
Hey guys!

I wanted to ask you a (quick?) question about null types in Kotlin. To give you an idea about my background, I'm a little bit of a newbie programmer that started with Python about a year ago. I started because I wanted to automate some of the activities we do at my office and I really liked it, so I dived into it. Now I wanted to try to get into some app programming and I understood that Kotlin is great for Android development.  And It's also pretty fun actually.

I'm taking some courses and trying things step by step but, maybe mostly because I come from Python as first language, I'm having a really hard time with the static type way of creating variables and stating everything. The hardest thing to understand to me is the "null" type. I can't understand what it stands for and why I should use a nullable type in general.

In Python everything is dynamically stated, so I don't have to think about what type of var I'm creating. Here, on the other hand, it seems to be the base of the workflow. And the fact that I don't seem to properly understand this worries me a bit. Even worse with the null type.

Can you please give me an advice about this? Maybe then I can stop to wrap my head around this with your help.

Thanks guys!
## [4][Kotlin Flows and Channels for Android](https://www.reddit.com/r/Kotlin/comments/hg6940/kotlin_flows_and_channels_for_android/)
- url: https://youtu.be/xch4aw7hNcY?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

## [5][Possible bug with generic where clause and star projection.](https://www.reddit.com/r/Kotlin/comments/hg6npk/possible_bug_with_generic_where_clause_and_star/)
- url: https://www.reddit.com/r/Kotlin/comments/hg6npk/possible_bug_with_generic_where_clause_and_star/
---
It appears that star projections only understand the first where clause. I'm not sure how to best articulate it, so have some code.

Here's a playground link: https://pl.kotl.in/a6vtI8tLG

Here's the code:

    interface Foo {
        fun foo()
    }
    
    interface Bar {
        fun bar()
    }
    
    interface FooAndBar&lt;T&gt; where T: Foo, T: Bar {
        val value: T
    }
    
    interface BarAndFoo&lt;T&gt; where T: Bar, T: Foo {
        val value: T
    }
    
    fun doSomething(thing: FooAndBar&lt;*&gt;) {
        thing.value.foo() // fine
        thing.value.bar() // compile error
    }
    
    fun doSomething(thing: BarAndFoo&lt;*&gt;) {
        thing.value.foo() // compile error
        thing.value.bar() // fine
    }
## [6][Spring Security Webflux with Kotlin coroutines explained](https://www.reddit.com/r/Kotlin/comments/hg90t1/spring_security_webflux_with_kotlin_coroutines/)
- url: https://www.reddit.com/r/Kotlin/comments/hg90t1/spring_security_webflux_with_kotlin_coroutines/
---
Hi there!

I just publish a small contribution to the Kotlin and Spring Security Webflux community. If you are struggling to find out how to add security to your reactive project I created this little repository to explain how it works. Also I added a little frontend app and a simple CI/CD pipeline. It's a work in progress but I think that it could give you some hints to you.

[https://github.com/soasada/kotlin-coroutines-webflux-security](https://github.com/soasada/kotlin-coroutines-webflux-security) 

Please feel free to ask me questions, suggestions or contributions. Hope you enjoy it
## [7][How do I bring array information from one intent to another?](https://www.reddit.com/r/Kotlin/comments/hfzlpy/how_do_i_bring_array_information_from_one_intent/)
- url: https://www.reddit.com/r/Kotlin/comments/hfzlpy/how_do_i_bring_array_information_from_one_intent/
---
As in I have a home page and a game page for my app. On the home screen I get the list of names of players. I want to know how to transfer that array to the next page of the app. I’ve tried writing to a file, but File() crashes the app. This is the missing link between both pages.

Might be a rookie mistake, but I’ve tried for the past two days trying.
## [8][Single line if statement without brances](https://www.reddit.com/r/Kotlin/comments/hft79x/single_line_if_statement_without_brances/)
- url: https://www.reddit.com/r/Kotlin/comments/hft79x/single_line_if_statement_without_brances/
---
I was curious what people's thoughts are on not using braces for `if` statements that are single line and have no else if/else?

For example:
```
if (value == null) return
```

I found [Android Kotlin Guidelines](https://developer.android.com/kotlin/style-guide) advocating it's okay but not many other well-known sources. I personally like it and although it isn't 100% foolproof, I believe it is safer/better than using multi-line `if` without braces.
    
    if (value == null)
      return

Also, I feel like that single line if with no braces fits well with other Kotlin conventions:

    if (true) 1 else 0
    val result = value ?: null
    value ?: null
    value ?: return
## [9][Ideas for naming a Kotlin/Common Telegram library (the MTProto API, not the HTTP bot API)](https://www.reddit.com/r/Kotlin/comments/hft5bb/ideas_for_naming_a_kotlincommon_telegram_library/)
- url: https://www.reddit.com/r/Kotlin/comments/hft5bb/ideas_for_naming_a_kotlincommon_telegram_library/
---
Ideally with a K in it. Kotlogram is already taken.
## [10][Problem reaching a piece of code outside of it's method](https://www.reddit.com/r/Kotlin/comments/hfr08k/problem_reaching_a_piece_of_code_outside_of_its/)
- url: https://www.reddit.com/r/Kotlin/comments/hfr08k/problem_reaching_a_piece_of_code_outside_of_its/
---
Hey all

I'm having a small problem trying to grab some info from a method, I seem stuck in a loop of unreachable code. 

I've spent the day creating a bunch of methods handling permissions for getting locations in my little weather forecasting app, and I've finally got it returning correct longitudes and latitudes. So I want to use them in the URL string I have for my weather data API, however I'm having trouble just getting to that info, I feel like I'm stuck in a loop.

The method where I get the local co-ordinates is this:

        private fun returnLocation() {
            var lat : Double? = 0.0
            var lon : Double? = 0.0
            fusedLocationClient = LocationServices.getFusedLocationProviderClient(this)
            fusedLocationClient.lastLocation.addOnSuccessListener { location : Location? -&gt;
                // Got last known location. In some rare situations this can be null.
                lat = location?.latitude
                lon = location?.longitude
    
                //for troubleshooting in console
                println("THE LATITUDE IS: $lat\n\n THE LONGITUDE IS: $lon")
            }
        }

The problem is I want to do something like, val localUrl = "[https://api.openweathermap.org/data/2.5/onecall?lat=***$lat***&amp;lon=***$lon***&amp;exclude=minutely&amp;appid=ee3cc93e43ef00b96a6bb4e56902d020](https://api.openweathermap.org/data/2.5/onecall?lat=$lat&amp;lon=$lon&amp;exclude=minutely&amp;appid=ee3cc93e43ef00b96a6bb4e56902d020)"

and use that in my JSON method where I use 

    val urlLondon = "https://api.openweathermap.org/data/2.5/onecall?lat=51.51&amp;lon=-0.19&amp;exclude=minutely&amp;appid=ee3cc93e43ef00b96a6bb4e56902d020"
    val urlNewYork =
                "https://api.openweathermap.org/data/2.5/onecall?lat=40.785091&amp;lon=-73.968285&amp;exclude=minutely&amp;appid=ee3cc93e43ef00b96a6bb4e56902d020"
    val urlLosAngeles =
                "https://api.openweathermap.org/data/2.5/onecall?lat=34.0522&amp;lon=-118.24&amp;exclude=minutely&amp;appid=ee3cc93e43ef00b96a6bb4e56902d020"
    
    //////My local URL here//////
    
    val request = Request.Builder().url(urlLondon).build()

&amp;#x200B;

I've tried to do things like make returnLocation return a string with the url, but I can't put the Return statement in the fusedLocationClient.*lastLocation*.addOnSuccessListener **{** location : Location? **-&gt;** ...} lamba, and outside of it lon, lat, and even a top level localUrlString are back to their default values.

&amp;#x200B;

Any help would be greatly appreciated, thanks in advance! :D
