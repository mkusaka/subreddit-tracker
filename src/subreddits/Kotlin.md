# Kotlin
## [1][Should I use lambda here or not ?](https://www.reddit.com/r/Kotlin/comments/g8wroh/should_i_use_lambda_here_or_not/)
- url: https://i.redd.it/xe9hw3csmbv41.png
---

## [2][JS/TS package for using Kotlin's Scope Function Extensions (with Optional Chaining support)](https://www.reddit.com/r/Kotlin/comments/g8o2c1/jsts_package_for_using_kotlins_scope_function/)
- url: https://github.com/TheDavidDelta/scope-extensions-js
---

## [3][Reactor Flux timeouts](https://www.reddit.com/r/Kotlin/comments/g8zp0k/reactor_flux_timeouts/)
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
## [4][Android Custom Toast in Kotlin with Example](https://www.reddit.com/r/Kotlin/comments/g8zi66/android_custom_toast_in_kotlin_with_example/)
- url: https://kotlin-android.com/android-custom-toast-kotlin/
---

## [5][Best way to learn kotlin?](https://www.reddit.com/r/Kotlin/comments/g8qv91/best_way_to_learn_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/g8qv91/best_way_to_learn_kotlin/
---
Just trying to get into something new and was wondering the best way to learn kotlin. Would it be books or courses or what? Any advice would be great.
## [6][Fastest Matrix Library for Android](https://www.reddit.com/r/Kotlin/comments/g8vq0e/fastest_matrix_library_for_android/)
- url: https://www.reddit.com/r/Kotlin/comments/g8vq0e/fastest_matrix_library_for_android/
---
Hello all,

I was working on an Android app that requires some linear algebra with matrices. The matrices will be somewhat medium-sized as they are not too small or too big. I was originally using jBlas because of the benchmarks I have seen online but after visiting multiple forums, they have all said that jBlas would be nowhere near as fast as what the benchmarks originally state. I have also heard about the Snapdragon Math Library to help improve the speed but I have not heard too much about it. From the general community, what linear-algebra libraries do you recommend that will be the fastest to use on Android? All help is greatly appreciated!
## [7][Learn, practice &amp; master the basics of Kotlin - The Android Language by registering to a FREE course in Sparkle - https://sparkle.adroitcorp.com.au](https://www.reddit.com/r/Kotlin/comments/g8wkc2/learn_practice_master_the_basics_of_kotlin_the/)
- url: https://i.redd.it/57cbc5lrjbv41.jpg
---

## [8][Deep recursion with coroutines](https://www.reddit.com/r/Kotlin/comments/g819ai/deep_recursion_with_coroutines/)
- url: https://medium.com/@elizarov/deep-recursion-with-coroutines-7c53e15993e3
---

## [9][Generic repository/service method](https://www.reddit.com/r/Kotlin/comments/g8ae4n/generic_repositoryservice_method/)
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
## [10][Preview of Harmony: A multi-process safe SharedPreference](https://www.reddit.com/r/Kotlin/comments/g838so/preview_of_harmony_a_multiprocess_safe/)
- url: /r/androiddev/comments/g835oy/preview_of_harmony_a_multiprocess_safe/
---

