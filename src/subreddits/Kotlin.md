# Kotlin
## [1][How to properly define REST resources in kottpd](https://www.reddit.com/r/Kotlin/comments/g9kbpd/how_to_properly_define_rest_resources_in_kottpd/)
- url: https://www.reddit.com/r/Kotlin/comments/g9kbpd/how_to_properly_define_rest_resources_in_kottpd/
---
I'm working on a REST-Server in Kotlin using [kottpd](https://github.com/gimlet2/kottpd)

Here's an excerpt of the resources / URIs I've defined:

    get("/patients") {req, res -&gt; PatientHandler.doGetAll(req, res)} 
    post("/patients") {req, res -&gt; PatientHandler.doPost(req, res)} 
    get("/patients/.*") {req, res -&gt; PatientHandler.doGetOne(req, res)} 
    post("/patients/.*") {req, res -&gt; PatientHandler.doPatch(req, res)} 
    get("/patients/.*/cases") {req, res -&gt; CaseHandler.doGetAll(req, res)} 
    post("/patients/.*/cases") {req, res -&gt; CaseHandler.doPost(req, res)}

`/patients` and `/patients/.*` have been working fine so far: e.g.`POST: /patients` adds a patient to the list and`GET: /patients/1` retrieves the patient with ID 1.

However, when calling `POST: /patients/1/cases` to add a case to patient 1, the functions attached to `POST: /patients/.*` are executed.

I assume(!) that `.*` is a wildcard and thus there's no distinction between there being an ID, some more path or just whatever.

How can I define my resources / URIs to distinguish between these cases? Or, in case my assumptions isn't correct: what am i doing wrong?Also asked that on [stackoverflow](https://stackoverflow.com/questions/61462028/how-to-properly-define-rest-resources-in-kottpd), but no answer
## [2][Understanding hash map basics by building one in Kotlin](https://www.reddit.com/r/Kotlin/comments/g91k9g/understanding_hash_map_basics_by_building_one_in/)
- url: https://medium.com/@ocallaghan.donal/building-a-hash-map-in-kotlin-73c9b23d19ba
---

## [3][Should I use lambda here or not ?](https://www.reddit.com/r/Kotlin/comments/g8wroh/should_i_use_lambda_here_or_not/)
- url: https://i.redd.it/xe9hw3csmbv41.png
---

## [4][Pass web request context in Ktor](https://www.reddit.com/r/Kotlin/comments/g98k2s/pass_web_request_context_in_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/g98k2s/pass_web_request_context_in_ktor/
---
In Go it is possible to pass the context of a web request to functions and go routines. This context may contain data such as authorization details but also a deadline which can be used by the receiver to decide whether there is enough time left to start a long operation which may complete after the deadline expires. It is also possible to cancel the context in the event such as client closing the connection before the request finishes processing which results in wasted resources because the client is no longer there to receive the result. Database query functions in Go’s standard library accept context. I was wondering if there is anything comparable in Ktor or elsewhere in the Kotlin universe?
## [5][Kotlin Multiplatform from the android project?](https://www.reddit.com/r/Kotlin/comments/g95asx/kotlin_multiplatform_from_the_android_project/)
- url: https://www.reddit.com/r/Kotlin/comments/g95asx/kotlin_multiplatform_from_the_android_project/
---
Hi All,

I have a completed android project, I have seen that Kotlin multiplatform can share business logic across platforms.

So I decided to use a multiplatform concept to share my network layer from my android app to the IOS app.

We haven't started working on the IOS app. 

 I'm aware that I need to do some rework on replacing Retrofit with Ktor-Client. I can do that.

I want to know is it possible to **add multiplatform concept in my existing Android project**?

If we can do this, can anyone share your thoughts with me?
## [6][Detailed information about OOP](https://www.reddit.com/r/Kotlin/comments/g94ksc/detailed_information_about_oop/)
- url: https://www.reddit.com/r/Kotlin/comments/g94ksc/detailed_information_about_oop/
---
I want to learn more about kotlin oop to better understand. Is there any document you can share?
## [7][JS/TS package for using Kotlin's Scope Function Extensions (with Optional Chaining support)](https://www.reddit.com/r/Kotlin/comments/g8o2c1/jsts_package_for_using_kotlins_scope_function/)
- url: https://github.com/TheDavidDelta/scope-extensions-js
---

## [8][Flutter-web-Music-UI](https://www.reddit.com/r/Kotlin/comments/g984x0/flutterwebmusicui/)
- url: https://www.reddit.com/r/Kotlin/comments/g984x0/flutterwebmusicui/
---
**Flutter-web-Music-UI**

[Github](https://github.com/imSanjaySoni/Flutter-web-Music-UI)

https://preview.redd.it/t9qq8f3x1fv41.png?width=2560&amp;format=png&amp;auto=webp&amp;s=8cd461e6fce334e5cb60ca1415da6a62181943bb
## [9][Reactor Flux timeouts](https://www.reddit.com/r/Kotlin/comments/g8zp0k/reactor_flux_timeouts/)
- url: https://www.reddit.com/r/Kotlin/comments/g8zp0k/reactor_flux_timeouts/
---
Sometimes it's necessary to listen to a Reactor `Flux`, handle messages in it and then propagate them to a downstream consumer. Besides you probably want to send some 'timeout' messages if your upstream keeps silence for a too long time. Project Reactor has special series of methods for such timeouts which predictably called `timeout`, but in my case those methods turned out to be a little bit useless. I wished to propagate timeouts when the upstream is silent for a too long period, but then I wished to continue listening. Moreover, I wanted to be ready for further timeouts when I returned to a listening of the upstream.

Existing Project Reactor `timeout` methods don't consider such a requirement - they switch me to a brand new flux when a timeout occures. I attempted to use such a behavior to implement my requirements using an approach similar to this:

    fun subscriber() {
      upstream
        .timeout(timeoutDuration, timeoutEventPublisher, { subscriber() } )		
        ...
    }

Probably I mixed up the arguments now but the main idea is right - I attempted to call the same stream with the same timeouts handler when a timeout occurs. Unfortunately, it doesn't work - only the very first timeout event would be propagated, all the further timeout events would be swallowed somewhere. Maybe I work incorrectly with the `.timeout` method - I don't know.

Anyway, I successfully used another approach. I just merge my regular stream and the stream of timeout events which I emit occasionally, and then somewhere at downstream I handle a received event. I look whether the event is instance of timeout event or not. Further handling depends on my business logic. In short:

    provider
    	.getUpstream()
    	.mergeWith(getHeartBeats())
    	.map {
    		...
    	}
    	
    /**
     * Occasionally emits empty lists as heartbeat-signals
     */
    private fun getHeartBeats(): Flux&lt;List&lt;Moment&gt;&gt; {
    	val heartbeat = emptyList&lt;Moment&gt;()
    	return Flux.interval(Duration.ofMillis(500)).map { heartbeat }
    }
## [10][Android Custom Toast in Kotlin with Example](https://www.reddit.com/r/Kotlin/comments/g8zi66/android_custom_toast_in_kotlin_with_example/)
- url: https://kotlin-android.com/android-custom-toast-kotlin/
---

