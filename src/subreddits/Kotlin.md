# Kotlin
## [1][Deep recursion with coroutines](https://www.reddit.com/r/Kotlin/comments/g819ai/deep_recursion_with_coroutines/)
- url: https://medium.com/@elizarov/deep-recursion-with-coroutines-7c53e15993e3
---

## [2][Kotlin ArrayList Tutorial with Examples](https://www.reddit.com/r/Kotlin/comments/g8dojp/kotlin_arraylist_tutorial_with_examples/)
- url: https://kotlin-android.com/kotlin-arraylist/
---

## [3][Preview of Harmony: A multi-process safe SharedPreference](https://www.reddit.com/r/Kotlin/comments/g838so/preview_of_harmony_a_multiprocess_safe/)
- url: /r/androiddev/comments/g835oy/preview_of_harmony_a_multiprocess_safe/
---

## [4][Generic repository/service method](https://www.reddit.com/r/Kotlin/comments/g8ae4n/generic_repositoryservice_method/)
- url: https://www.reddit.com/r/Kotlin/comments/g8ae4n/generic_repositoryservice_method/
---
For this application my weapon of choice is Micronaut but I'm pretty sure the issue would be the same had I used Spring.

I have a lot entities that I want to synchronize with various clients.

All the entities are exposed through a repository and on top of that a service.

Each repository implements a method as illustrated below

    interface EntityRepository {
        fun findByUpdatedAfter(updatedDateTime: LocalDateTime): Iterable&lt;Entity&gt;
    }

And then in the service a similar (entity specific) method is implemented.

Since this is duplicated for all repositories and services I thought I'd make it a bit more generic.

So I created an interface for my entities to ensure that a property called `updated` would exists

    interface DateUpdatedEntity&lt;T&gt; {
        val updated: T
    }

A generic repository

    interface UpdatedAfterRepository&lt;E : DateUpdatedEntity&lt;T&gt;, T&gt; {
        fun findByUpdatedAfter(updatedDateTime: T): Iterable&lt;E&gt;
    }

And a service interface with a default implementation of a method called `findAllAfter` as illustrated below

    interface UpdatedAfterService&lt;E : DateUpdatedEntity&lt;T&gt;, T&gt; {
        val updatedAfterRepository: UpdatedAfterRepository&lt;E, T&gt;
        fun findAllAfter(updatedDateTime: T): Iterable&lt;E&gt; = updatedAfterRepository.findByUpdatedAfter(updatedDateTime)
    }

However when I try to implement the above in one of my services as illustrated below

    @Singleton
    @Transactional
    class ExerciseServiceImpl(
            override val updatedAfterRepository: UpdatedAfterRepository&lt;Exercise, LocalDateTime&gt;,
            ...
    ) : ExerciseService {
    ...
    }

I get the following error message telling me that dependency injection has failed.

    Message: Multiple possible bean candidates found: [dk.fitfit.fitlog.repository.WorkoutRepository$Intercepted, dk.fitfit.fitlog.repository.VideoRepository$Intercepted, dk.fitfit.fitlog.repository.SessionExerciseRepository$Intercepted, dk.fitfit.fitlog.repository.PictureRepository$Intercepted, dk.fitfit.fitlog.repository.SessionRepository$Intercepted, dk.fitfit.fitlog.repository.SessionRoundRepository$Intercepted, dk.fitfit.fitlog.repository.RoundRepository$Intercepted, dk.fitfit.fitlog.repository.RoundExerciseRepository$Intercepted, dk.fitfit.fitlog.repository.ExerciseRepository$Intercepted]

I'm guessing it's because of type erasure but is there anything I can do about it? Or can anyone here suggest me an alternative approach? So I don't have to have repeating code in all my repositories and services
## [5][Github Template for starting an Android app project with: 100% Kotlin + Github Actions + ktlint + Detekt + Gradle Kotlin DSL + buildSrc dependencies already set up.](https://www.reddit.com/r/Kotlin/comments/g7rhne/github_template_for_starting_an_android_app/)
- url: https://github.com/cortinico/kotlin-android-template
---

## [6][When I see a Java dev write Kotlin in a Java-esque way](https://www.reddit.com/r/Kotlin/comments/g7iskp/when_i_see_a_java_dev_write_kotlin_in_a_javaesque/)
- url: https://i.redd.it/btt1mk0aluu41.png
---

## [7][Multicasting for server-side using Kotlin and Project Reactor](https://www.reddit.com/r/Kotlin/comments/g7zequ/multicasting_for_serverside_using_kotlin_and/)
- url: https://www.reddit.com/r/Kotlin/comments/g7zequ/multicasting_for_serverside_using_kotlin_and/
---
I'd like to share our helper class which we use to organize multicasted Project Reactor streams.

In our crypto-currency exchange we pretty often use a multicasting when we broadcast messages from the server-side to our clients.

As an example, our gateway listens to a kernel which sends us all the information we need, but that information is streamed in one stream regardless of a trader it belongs to. So we listen to it, separate on different streams and broadcast them to our clients - such a process is called a multicasting.

For the sake of simplicity we moved all the multicast logic into a standalone class called `Multicast`. And here it is:

    class Multicast&lt;in ReqT, RespT&gt;(
        private val meterRegistry: MeterRegistry,
        private val callerClassName: String,
        private val numberOfCachedItemsAtStart: Int = 0,
        private val subscribe: (ReqT) -&gt; Flux&lt;RespT&gt;
    ) {
    
        private val subscriptions = ConcurrentHashMap&lt;ReqT, FluxProcessor&lt;RespT, RespT&gt;&gt;()
    
        init {
            addMeter("multicast_total_upstreams") { totalUpstreams().toDouble() }
            addMeter("multicast_active_upstreams") { activeUpstreams().toDouble() }
        }
    
        private fun addMeter(name: String, function: (Multicast&lt;ReqT, RespT&gt;) -&gt; Double) {
            meterRegistry.gauge(name, Tags.of("caller_class", callerClassName), this, function)
        }
    
        operator fun invoke(request: ReqT): Flux&lt;RespT&gt; {
            return subscriptions.compute(request) { req, latest -&gt;
                if (latest != null &amp;&amp; !latest.isTerminated) {
                    latest
                } else {
                    val upstream = subscribe(req)
                    val processor = ReplayProcessor.create&lt;RespT&gt;(numberOfCachedItemsAtStart)
                    upstream.subscribe(processor)
                    processor
                }
            }!!
        }
    
        private fun totalUpstreams(): Int = subscriptions.values.size
        private fun activeUpstreams(): Int = subscriptions.values.count { !it.isTerminated }
        internal fun downstreamsNumber(request: ReqT): Long = subscriptions.getValue(request).downstreamCount()
    }

Here you see that we store a map of subscriptions where the keys are requests which received from the consumers and the values are Project Reactor FluxProcessors which help us to shift messages from an upstream subscription to downstream subscriptions.

Also we create some meters to have an opportunity to monitor number of alive and dead upstreams we have in the `Multicast`.

In the main `invoke` method, which essentially is a kotlin `Operator`, we check wether we have an existing `FluxProcessor` related to the specified request. If we have - we reuse it. Otherwise - we create a new upstream connection provided through the `subscribe` lambda in the class constructor, and then we create new `ReplayProcessor` which essentially is a Project Reactor `FluxProcessor` supporting a history of messages with a manageable history depth. This `ReplayProcessor` allows us to set a `numberOfCachedItemsAtStart` which we often use to store snapshot data. The created `ReplayProcessor` immediately subscribes to the upstream we initialized before, and then we return a `Flux` from the prepared `FluxProcessor`.

Also you can see that when we reuse an existing `FluxProcessor` we check wether it's terminated or not. It's necessary for the case when a client subscribed to our upstream, then unsubscribed (hence the `FluxProcessor` terminated) and then another client tries to subscribe. Of course in this case we should create a brand new `FluxProcessor` and we do it.

&amp;#x200B;

And what approaches do you use to manage multicasting in your projects?
## [8][Implementing Retrofit in Android | Paging | Part 17](https://www.reddit.com/r/Kotlin/comments/g7pd0x/implementing_retrofit_in_android_paging_part_17/)
- url: https://www.youtube.com/watch?v=9jibu3RZ1S4
---

## [9][Kobot - configurable conversational bot framework, open source](https://www.reddit.com/r/Kotlin/comments/g77jye/kobot_configurable_conversational_bot_framework/)
- url: https://www.reddit.com/r/Kotlin/comments/g77jye/kobot_configurable_conversational_bot_framework/
---
Hi folks,

In these days of quarantine, I started an open source project. I'd like to share it with you and gently ask, to everyone interested, feedback or critiques.

The project is hosted here: [https://github.com/besil/kobot](https://github.com/besil/kobot).

It is a framework to create conversational bots, using a simple configuration file.

The goal is to help non technical users to create conversational channels in an easy way. I know some proprietary software already do this, but I didn't find any open source project about it.

Main feature include:

* multiple conversational channels: telegram, http, cli, ...
* Database interaction through sql language, in order to send dynamic messages or to store information
* External HTTP API integration, in order to interact with web services
* Published as docker image to simplify devops

Framework is implemented 100% in Kotlin, with Spring Boot integration.

Any critique or feedback is welcome!

Thank you!
## [10][Broker: Publish-Subscribe (a.k.a Pub/Sub, EventBus) library for Android and JVM built with Coroutines](https://www.reddit.com/r/Kotlin/comments/g7gv2k/broker_publishsubscribe_aka_pubsub_eventbus/)
- url: https://github.com/adrielcafe/Broker
---

