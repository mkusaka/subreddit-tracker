# Kotlin
## [1][What do you use for Dependency Injection?](https://www.reddit.com/r/Kotlin/comments/epj2p4/what_do_you_use_for_dependency_injection/)
- url: https://www.reddit.com/r/Kotlin/comments/epj2p4/what_do_you_use_for_dependency_injection/
---
I'm working on a project creating AWS Lambdas using Kotlin, and there's currently no dependency injection being used. I've only previously used Spring, which we are not using on this project. So rather than use Spring just to get it's DI abilities (which aren't really suited to cloud functions anyway due to being runtime), I am looking at an alternative.

Dagger seems to be widely used, but looks complicated with a lot of boiler plate. Does anyone have any preference for Koin or Kodein?
## [2][Kotlin 1.3.70 EAP](https://www.reddit.com/r/Kotlin/comments/ep7o7c/kotlin_1370_eap/)
- url: https://discuss.kotlinlang.org/t/kotlin-1-3-70-early-access-preview/15876
---

## [3][LambdaConf 2020 Call for Proposals - Submit Kotlin Proposals by Feb 15 (June 3-5, Estes Park, CO)](https://www.reddit.com/r/Kotlin/comments/epjb9o/lambdaconf_2020_call_for_proposals_submit_kotlin/)
- url: https://lambdaconf.zohobackstage.com/LambdaConf2020#/cfp?lang=en
---

## [4][Beginner Question regarding Tutorial-Code](https://www.reddit.com/r/Kotlin/comments/ephx0z/beginner_question_regarding_tutorialcode/)
- url: https://www.reddit.com/r/Kotlin/comments/ephx0z/beginner_question_regarding_tutorialcode/
---
Hey there,

&lt;tl;dr&gt; Highlighted the actual question because lots of context blabla

I have a master degree in computer science and and been working with Java in the past 7 years.  
I recently decided to get into Kotlin and am currently working my way through their code examples.

I'm currently looking into **Extension Functions and Properties**, which conceptually I understand. I also understand the provided code example, though I don't quite get some of the **syntax**.

Heres a link to the tutorial section:  [https://play.kotlinlang.org/byExample/04\_functional/03\_extensionFunctions](https://play.kotlinlang.org/byExample/04_functional/03_extensionFunctions) 

&amp;#x200B;

I'm referring to these lines of code:

`fun Order.maxPricedItemValue(): Float = this.items.maxBy { it.price }?.price ?: 0F`  
`fun Order.maxPricedItemName() = this.items.maxBy { it.price }?.name ?: "NO_PRODUCTS"`

Specifically, **the part after curly brackets are closed**.

From my understanding now, I define the extension functions as calling a maxBy() function from Collection, using lambda `{it.price}` to define which parameter to sort by and then I'm not sure any more.

At first, this reminded me of ternary operations, but they don't exist in Kotlin.  
Then I recalled that `?` is used to mark nullable objects, so I guess if the Collection doesn't contain any entries or for whatever reason can't find the max priced Item, it returns null. Then I try to call price / name from the nullable result of maxBy? Sounds plausible, but again, not sure. And how `?:` works in the end is beyond me.

I mean, I understand that the first function returns the price of the highest priced item or zero (float) if there is none, and the second returns the name or "NO\_PRODUCTS" if there are none. But I deduct that from context, **not** from **understanding the syntax**.
## [5][Would you recommend Kotlin as a back-end and general purpose language? (In comparison to Go)](https://www.reddit.com/r/Kotlin/comments/ep1pkd/would_you_recommend_kotlin_as_a_backend_and/)
- url: https://www.reddit.com/r/Kotlin/comments/ep1pkd/would_you_recommend_kotlin_as_a_backend_and/
---
I've been doing most of my server-side work with Go, and recently got exposed to Kotlin in a native Android project and I really like the language.

I know that one language is generally not the answer to all problems but, in small teams building expertise in one language is a benefit I'm willing to consider.

So the question here is, how's the Kotlin story outside of Android in things like:

* Microservices
* Serverless Functions
* General purpose data processing
* Daemons
* CLIs

Would you recommend writing such applications in Kotlin? And if you have written things like these, how was the experience?
## [6][Good books to learn Kotlin as Java developer in 2020?](https://www.reddit.com/r/Kotlin/comments/ep133r/good_books_to_learn_kotlin_as_java_developer_in/)
- url: https://www.reddit.com/r/Kotlin/comments/ep133r/good_books_to_learn_kotlin_as_java_developer_in/
---
I'm a Java developer and I would like to learn kotlin? Are there any good books about it?   


On the kotlin website they recommend the book  " Kotlin in Action " but it was released in February 2017, so I don't think its up to date.  


Any recommendations?
## [7][How to create AppBars in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/eozcr8/how_to_create_appbars_in_jetpack_compose/)
- url: https://proandroiddev.com/creating-appbars-in-jetpack-compose-a8b5a5639930
---

## [8][Where programming languages are headed in 2020](https://www.reddit.com/r/Kotlin/comments/eoqoqn/where_programming_languages_are_headed_in_2020/)
- url: https://www.oreilly.com/radar/where-programming-languages-are-headed-in-2020/
---

## [9][Slate Kit: Result&lt;T,E&gt; - Model successes and failures with optional statuses/codes](https://www.reddit.com/r/Kotlin/comments/eowb1n/slate_kit_resultte_model_successes_and_failures/)
- url: https://www.reddit.com/r/Kotlin/comments/eowb1n/slate_kit_resultte_model_successes_and_failures/
---
This is the 1st stable 1.0.0 release of the [Slate Kit: Result&lt;T,E&gt;](https://www.slatekit.com/arch/results/) module. This is an approach to error-handling designed as a **variation** of the **Result&lt;T&gt;** component which exists in Kotlin and in other languages, and is a way to accurately model the success or failure of any operation. This implementation differs in some key design choices which are listed in the guide and summarized in the notes below. It is generally available now and also hoping to get some feedback for potential improvements in the future. Thanks!

* **Guide**: [Slate Kit: Result](https://www.slatekit.com/arch/results/)
* **Purpose**: Modeling of successes and failures with optional statuses/codes
* **Install**: [Gradle setup](https://www.slatekit.com/arch/results/#install)
* **Sources**: [slatekit-result](https://github.com/code-helix/slatekit/tree/master/src/lib/kotlin/slatekit-result/src/main/kotlin/slatekit/results)
* **Example**: [Examples\_Result.kt](https://github.com/code-helix/slatekit/blob/master/src/lib/kotlin/slatekit-examples/src/main/kotlin/slatekit/examples/Example_Results.kt)
* **License**: Apache 2.0
* **Upcoming**: Other 1.0.0 versions of modules of Slate Kit will be released in the coming weeks
* **Notes**: This differs from other implementations via:

1. Allowing for a custom error type ( instead of just Exception )
2. Optionally supports sub-categories of Successes / Failures via **Statuses** ( e.g. Denied, Ignored, Invalid )
3. Provides sensible defaults and ways to construct failures via various builders
## [10][Creating a Kotlin project to log GPS data without Google Play Services?](https://www.reddit.com/r/Kotlin/comments/eovzo1/creating_a_kotlin_project_to_log_gps_data_without/)
- url: https://www.reddit.com/r/Kotlin/comments/eovzo1/creating_a_kotlin_project_to_log_gps_data_without/
---
Hi all,

I'm a long term python / docker hobyist but I'm in the need of creating a personal GSP logging app for my phone.  I'm using LineageOS for my Android and using a Google Pixel 1.  

I have no Google services installed on the phone and want to build myself my first java app to log this data for my phone. Due to this I came across Kotlin.

In my searches I keep seeing people going to google services to get GPS data.  Without google services, what way would you suggest I do this?  

I want this to eventually end up as open source on the FDroid store, so I want to keep things as free/open as possible.  

Thanks so much
