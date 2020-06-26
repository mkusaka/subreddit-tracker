# Kotlin
## [1][Kotlin Flows and Channels for Android](https://www.reddit.com/r/Kotlin/comments/hg6940/kotlin_flows_and_channels_for_android/)
- url: https://youtu.be/xch4aw7hNcY?list=PLEx5khR4g7PL-JwckuOkkc5cR6X5hn6ug
---

## [2][How do I bring array information from one intent to another?](https://www.reddit.com/r/Kotlin/comments/hfzlpy/how_do_i_bring_array_information_from_one_intent/)
- url: https://www.reddit.com/r/Kotlin/comments/hfzlpy/how_do_i_bring_array_information_from_one_intent/
---
As in I have a home page and a game page for my app. On the home screen I get the list of names of players. I want to know how to transfer that array to the next page of the app. I’ve tried writing to a file, but File() crashes the app. This is the missing link between both pages.

Might be a rookie mistake, but I’ve tried for the past two days trying.
## [3][Possible bug with generic where clause and star projection.](https://www.reddit.com/r/Kotlin/comments/hg6npk/possible_bug_with_generic_where_clause_and_star/)
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
## [4][Single line if statement without brances](https://www.reddit.com/r/Kotlin/comments/hft79x/single_line_if_statement_without_brances/)
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
## [5][Ideas for naming a Kotlin/Common Telegram library (the MTProto API, not the HTTP bot API)](https://www.reddit.com/r/Kotlin/comments/hft5bb/ideas_for_naming_a_kotlincommon_telegram_library/)
- url: https://www.reddit.com/r/Kotlin/comments/hft5bb/ideas_for_naming_a_kotlincommon_telegram_library/
---
Ideally with a K in it. Kotlogram is already taken.
## [6][Problem reaching a piece of code outside of it's method](https://www.reddit.com/r/Kotlin/comments/hfr08k/problem_reaching_a_piece_of_code_outside_of_its/)
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
## [7][Announcing Lingua 1.0.0: The most accurate natural language detection library for the JVM](https://www.reddit.com/r/Kotlin/comments/hf59dc/announcing_lingua_100_the_most_accurate_natural/)
- url: https://github.com/pemistahl/lingua
---

## [8][How do I deep copy a Kotlin object?](https://www.reddit.com/r/Kotlin/comments/hfcrzu/how_do_i_deep_copy_a_kotlin_object/)
- url: https://www.reddit.com/r/Kotlin/comments/hfcrzu/how_do_i_deep_copy_a_kotlin_object/
---
I have a method that takes a MutableList of objects and needs return a MutableList of the same objects but with significant modifications. However, the original versions of the objects are still needed. Running the method on objects that have already been modified would generate invalid data. WHat's the best way to deep copy the list object so I can make modifications without affecting the originals?
## [9][Kotlin/Native extra imports](https://www.reddit.com/r/Kotlin/comments/hf59n1/kotlinnative_extra_imports/)
- url: https://www.reddit.com/r/Kotlin/comments/hf59n1/kotlinnative_extra_imports/
---
I'm trying to import an existing swift library of ours, and while it exports to obj-c just fine, when kotlin goes to generate headers it adds in some classes related to UIPointer and an import for UIAxis, which then cant be resolved. These aren't used anywhere in our swift library, and I'm not sure how to get rid of them. Using Kotlin 1.3.72, Xcode 11.5, targeting iOS 13.4 (i've tried targeting older to see if that would fix it. it did not). Current best guess is that kotlin/native is importing the fields from the UIKit header in the swift library, but not mapping these particular ones across to the provided platform library correctly. Is there a way I can force that to regenerate?
## [10][Discussion: Do you think this would be a cool feature?](https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/)
- url: https://www.reddit.com/r/Kotlin/comments/hexe8l/discussion_do_you_think_this_would_be_a_cool/
---
TL;DR let the Kotlin compiler generate the boilerplate code for reifying typeargs in order to circumvent type erasure. Or, in other words, a `reified` keyword for class / non-inline function typeargs.

Working with Java (and reflection) a lot, I frequently feel the urge of banging my head against the table because of type erasure. Kotlin having a `reified` keyword for inline function typeargs is already a helpful thing, but I feel like there could be done more. In order to manually reify typeargs, I found myself using this boilerplate pattern:

1. Store a `Class&lt;?&gt;` instance in the class (this usually requires passing that instance to the constructor)
2. Perform instance checks using `clazz.isInstance(obj)`
3. Perform an unchecked cast to the requested typearg.

Kotlin is already known as a language that mostly uses syntactic sugar to improve Java, while still ensuring interoperability. Therefore it would be quite fitting if Kotlin provided an universally usable `reified` keyword that hides the boilerplate pattern from you and lets you use your typearg as if it wasn't subject to type erasure. So you could do

```
var inner: T
// ...
fun set(newval: Any) {
    if (newval is T) {
        inner = newval // smartcast
    }
}
```

Which would actually store the KClass representation of T and performs an isInstance check in the background.

What are your thoughts? Is this a good idea? Is it useful? How would you improve it? And ultimately, do you think it should be a feature request on YouTrack?
