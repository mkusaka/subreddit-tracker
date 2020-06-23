# Kotlin
## [1][Kotlin Coroutines vs Java Threads](https://www.reddit.com/r/Kotlin/comments/hecrel/kotlin_coroutines_vs_java_threads/)
- url: https://piotrminkowski.com/2020/06/23/kotlin-coroutines-vs-java-threads/
---

## [2][SQLDelight 1.4.0 released](https://www.reddit.com/r/Kotlin/comments/hdvn5g/sqldelight_140_released/)
- url: https://github.com/cashapp/sqldelight/releases/tag/1.4.0
---

## [3][variable name vs type position](https://www.reddit.com/r/Kotlin/comments/hedgeg/variable_name_vs_type_position/)
- url: https://www.reddit.com/r/Kotlin/comments/hedgeg/variable_name_vs_type_position/
---
Hey everyone, trying to learn Kotlin, coming from Java.

Why did Kotlin decide to switch the variable name and the type positions around? Is there a good reason for this?  
I find it so much easier to write code like this:

Dog dog; (Something like D&lt;ctrl-space&gt;&lt;space&gt;&lt;ctrl-space&gt;&lt;enter&gt;)

vs.

dog: Dog (Where the autocomplete is only available towards the end of the line)

What advantages are there to do it the Kotlin way?
## [4][What's the best way to check if something is null in Kotlin?](https://www.reddit.com/r/Kotlin/comments/he5792/whats_the_best_way_to_check_if_something_is_null/)
- url: https://www.reddit.com/r/Kotlin/comments/he5792/whats_the_best_way_to_check_if_something_is_null/
---
Should I use `variable == null`, `variable === null` or something else?
## [5][Where can I see a list of the built-in exceptions in Kotlin?](https://www.reddit.com/r/Kotlin/comments/he6nv6/where_can_i_see_a_list_of_the_builtin_exceptions/)
- url: https://www.reddit.com/r/Kotlin/comments/he6nv6/where_can_i_see_a_list_of_the_builtin_exceptions/
---
I like to throw exceptions built into the language with a custom message explaining the error when possible, instead of making my own exceptions every time. Howver, I couldn't find a definitive list of all the exceptions Kotlin defines in the standard library and their usages. Does such a list exist? Is this way of throwing exceptions even a good idea?
## [6][New Dokka - Designed for Fearless Creativity](https://www.reddit.com/r/Kotlin/comments/hdxgj4/new_dokka_designed_for_fearless_creativity/)
- url: https://www.youtube.com/watch?v=OvFoTRhqaKg
---

## [7][Kotlin REPL?](https://www.reddit.com/r/Kotlin/comments/hdsppd/kotlin_repl/)
- url: https://www.reddit.com/r/Kotlin/comments/hdsppd/kotlin_repl/
---
Does Kotlin have a REPL? Something like JShell or Ipython?
## [8][Android | Having a strange WrongThread error that I can't figure out](https://www.reddit.com/r/Kotlin/comments/hdzap2/android_having_a_strange_wrongthread_error_that_i/)
- url: https://www.reddit.com/r/Kotlin/comments/hdzap2/android_having_a_strange_wrongthread_error_that_i/
---
SOLVED :D

\-------

Hey all,

Just quickly wanted to say thanks again to the guys that helped me out a few days ago, thanks this sub :D

I'm getting the following error message when I try to add some new values to a text view on my app.

    android.view.ViewRootImpl$CalledFromWrongThreadException: Only the original thread that created a view hierarchy can touch its views.

The strangest thing is that this problem sometimes happens, and sometimes doesnt.

I'm running it within an overriden onResponse() method within a OkHttpClient. Everything has been working great, but now it's gotten to a point where I can't add any more values to my weather app without the app crashing.

Some examples of code that works currently:

    threeDay.text = getNextDay(twoDay.text.toString())
    threeWeather.text = weatherData.daily[2].weather[0].main
    cityText.text = getLocalLocation(weatherData.timezone)
    weatherSummaryText.text = weatherData.current.weather[0].description

And then some error code, which when I test in a println method gives me exactly the value I want:

    nowTemp.text = kelvinToCelsius(weatherData.hourly[0].temp).toString()
    oneTemp.text = kelvinToCelsius(weatherData.hourly[1].temp).toString()

Where kelvinToCelsius is just a converter method I made.

&amp;#x200B;

Can anyone give me any insight into this strange problem?

Cheers

&amp;#x200B;

EDIT:

&amp;#x200B;

the onResponse method where I am doing inputting all my information from the API.

    private fun fetchJson() {
    
            val url = "https://api.openweathermap.org/data/2.5/onecall?lat=51.51&amp;lon=-0.19&amp;exclude=minutely&amp;appid=MYKEY"
    
            val request = Request.Builder().url(url).build()
    
            client.newCall(request).enqueue(object : Callback {
    
                override fun onFailure(call: Call, e: IOException) {
                    println("Failed to execute request")
                }
    
                override fun onResponse(call: Call, response: Response) {
                    val body = response.body?.string()
                    println(body)
    
                    val gson = GsonBuilder().create()
    
                    val weatherData = gson.fromJson(body, WeatherData::class.java)
    
                    /**
                     * ERROR: High/Low temps not registering.
                     * android.view.ViewRootImpl$CalledFromWrongThreadException: Only         the original thread that created a view hierarchy can touch its views.
                     * Sometimes works, sometimes doesn't???
                     */
    
                    //Fill out info from top to bottom.
                    //Top
                    cityText.text = getLocalLocation(weatherData.timezone)
                    weatherSummaryText.text = weatherData.current.weather[0].description
                    tempText.text = kelvinToCelsius(weatherData.current.temp).toString() + "°"
    
                    /////////////////////////Hourly
                    currentDayText.text = getDateDay(weatherData.current.dt)
    
    //                nowTemp.text = kelvinToCelsius(weatherData.hourly[0].temp).toString()
    //
    //                oneTemp.text = kelvinToCelsius(weatherData.hourly[1].temp).toString()
    //
    //                twoTemp.text = kelvinToCelsius(weatherData.hourly[2].temp).toString()
    //
    //                threeTemp.text = kelvinToCelsius(weatherData.hourly[3].temp).toString()
    //
    //                fourTemp.text = kelvinToCelsius(weatherData.hourly[4].temp).toString()
    //
    //                fiveTemp.text = kelvinToCelsius(weatherData.hourly[5].temp).toString()
    //
    //                sixTemp.text = kelvinToCelsius(weatherData.hourly[6].temp).toString()
    //
    //                sevenTemp.text = kelvinToCelsius(weatherData.hourly[7].temp).toString()
    //
    //                eightTemp.text = kelvinToCelsius(weatherData.hourly[8].temp).toString()
    //
    //                nineTemp.text = kelvinToCelsius(weatherData.hourly[9].temp).toString()
    //
    //                tenTemp.text = kelvinToCelsius(weatherData.hourly[10].temp).toString()
    //
    //                elevenTemp.text = kelvinToCelsius(weatherData.hourly[11].temp).toString()
    //
    //                twelveTemp.text = kelvinToCelsius(weatherData.hourly[12].temp).toString()
    
    
                    /////////////////////////Daily
                    oneDay.text = getNextDay(currentDayText.text.toString())
                    oneWeather.text = weatherData.daily[0].weather[0].main
    //                oneHighTemp.text = kelvinToCelsius(weatherData.daily[0].temp.max).toString()
    //                oneLowTemp.text = kelvinToCelsius(weatherData.daily[0].temp.min).toString()
    
                    twoDay.text = getNextDay(oneDay.text.toString())
                    twoWeather.text = weatherData.daily[1].weather[0].main
    //                twoHighTemp.text = kelvinToCelsius(weatherData.daily[1].temp.max).toString()
    //                twoLowTemp.text = kelvinToCelsius(weatherData.daily[1].temp.min).toString()
    
                    threeDay.text = getNextDay(twoDay.text.toString())
                    threeWeather.text = weatherData.daily[2].weather[0].main
    //                threeHighTemp.text = kelvinToCelsius(weatherData.daily[2].temp.max).toString()
    //                threeLowTemp.text = kelvinToCelsius(weatherData.daily[2].temp.min).toString()
    
                    fourDay.text = getNextDay(threeDay.text.toString())
                    fourWeather.text = weatherData.daily[3].weather[0].main
    //                fourHighTemp.text = kelvinToCelsius(weatherData.daily[3].temp.max).toString()
    //                fourLowTemp.text = kelvinToCelsius(weatherData.daily[3].temp.min).toString()
    
                    fiveDay.text = getNextDay(fourDay.text.toString())
                    fiveWeather.text = weatherData.daily[4].weather[0].main
    //                fiveHighTemp.text = kelvinToCelsius(weatherData.daily[4].temp.max).toString()
    //                fiveLowTemp.text = kelvinToCelsius(weatherData.daily[4].temp.min).toString()
    
                    sixDay.text = getNextDay(fiveDay.text.toString())
                    sixWeather.text = weatherData.daily[5].weather[0].main
    //                sixHighTemp.text = kelvinToCelsius(weatherData.daily[5].temp.max).toString()
    //                sixLowTemp.text = kelvinToCelsius(weatherData.daily[5].temp.min).toString()
    
                    sevenDay.text = getNextDay(sixDay.text.toString())
                    sevenWeather.text = weatherData.daily[6].weather[0].main
    //                sevenHighTemp.text = kelvinToCelsius(weatherData.daily[6].temp.max).toString()
    //                sevenLowTemp.text = kelvinToCelsius(weatherData.daily[6].temp.min).toString()
    
                }
            })
    
    
        }

&amp;#x200B;
## [9][Within 24 hours, kotlin will overtake the scala subreddit in subscriber count](https://www.reddit.com/r/Kotlin/comments/hdbqs1/within_24_hours_kotlin_will_overtake_the_scala/)
- url: https://www.reddit.com/r/Kotlin/comments/hdbqs1/within_24_hours_kotlin_will_overtake_the_scala/
---
Enjoy!
## [10][Low-level api interop in Kotlin/JVM?](https://www.reddit.com/r/Kotlin/comments/hdndw9/lowlevel_api_interop_in_kotlinjvm/)
- url: https://www.reddit.com/r/Kotlin/comments/hdndw9/lowlevel_api_interop_in_kotlinjvm/
---
Is there some guide to get started on pulling info from some low level apis such as wireless information, battery information, etc?

I know in linux its easy by running some commands and filtering outputs, but its a mess in platforms like windows where we need to write some sort of bindings in like jni, but I still don't know much about JNI and want some guides to work.

And is it possible to interact them (low-level api) with help of Kotlin/Native to write bindings for Kotlin/JVM instead of writing in direct C++?
