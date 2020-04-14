# Kotlin
## [1][[side-project] Daft Punk - Da Funk Quarantine Remix in Kotlin](https://www.reddit.com/r/Kotlin/comments/g12l51/sideproject_daft_punk_da_funk_quarantine_remix_in/)
- url: https://www.youtube.com/watch?v=OdQQJPpL6Lo&amp;t=138s
---

## [2][Going Full Kotlin Multiplatform - Talking Kotlin](https://www.reddit.com/r/Kotlin/comments/g14tij/going_full_kotlin_multiplatform_talking_kotlin/)
- url: https://talkingkotlin.com/going-full-kotlin-multiplatform/
---

## [3][Rounding doubles in Kotlin, Android studio](https://www.reddit.com/r/Kotlin/comments/g11xbf/rounding_doubles_in_kotlin_android_studio/)
- url: https://www.reddit.com/r/Kotlin/comments/g11xbf/rounding_doubles_in_kotlin_android_studio/
---
How can i round doubles in Kotlin? I've tried a lot of things i found on google but I'm probably just not filling it  in properly. So I would like an example of a double being rounded to a certain amount of decimals.
## [4][Question about RESTful API with DTO's](https://www.reddit.com/r/Kotlin/comments/g0wbmd/question_about_restful_api_with_dtos/)
- url: https://www.reddit.com/r/Kotlin/comments/g0wbmd/question_about_restful_api_with_dtos/
---
I'm trying to implement a RESTful API abstracting my entities using DTO's.

I've been trying to keep my entities immutable but I'm not sure it's really worth the hassle.

    @Entity
    class Exercise(
            val name: String,
            val description: String,
            @JsonIgnore
            @ManyToOne
            val creator: User,
            @OneToMany(fetch = EAGER)
            @Fetch(SUBSELECT)
            val videos: MutableList&lt;Video&gt; = mutableListOf(),
            @Id
            @GeneratedValue(strategy = SEQUENCE)
            val id: Long = 0
    )

The problem arises when I try to implement updating of an object.

PUT /exercises/{id}

My DTO looks as the following

    class ExerciseRequest(
            val name: String,
            val description: String,
            val id: Long?
    )

I convert my DTO to an entity by using the following extension function

    private fun ExerciseRequest.toExercise(user: User, id: Long = 0) = Exercise(name, description, creator = user, id = this.id ?: id)

Which leaves the videos property as an empty list.

So when I update any existing videos associated with the exercise gets removed.

I could solve this by declaring the video property as `val videos: MutableList&lt;Video&gt;? = null` and merge my incoming 'update' object onto the existing object using merging code like [this](https://gist.github.com/josdejong/fbb43ae33fcdd922040dac4ffc31aeaf).

But then when I want to add a video the initial property is null and I can't assign a new list to it unless I change it's declaration to var. Thus breaking immutability.

Having a merge function which would merge the content of lists rather than overwriting them would solve the problem. But I'm not sure if that's a good idea.

How do you guys go about implementing update?
## [5][Is there any site, where I can find different functional programming problems?](https://www.reddit.com/r/Kotlin/comments/g0i0c4/is_there_any_site_where_i_can_find_different/)
- url: https://www.reddit.com/r/Kotlin/comments/g0i0c4/is_there_any_site_where_i_can_find_different/
---

## [6][Problem with a function](https://www.reddit.com/r/Kotlin/comments/g0pzex/problem_with_a_function/)
- url: https://www.reddit.com/r/Kotlin/comments/g0pzex/problem_with_a_function/
---
Hey , I'm a newbie in kotlin I'm doing a program that calculates the number of digits that a number has. This is my function

`fun number(a:Int):Int{`  
`var i = 0`  
`while (a &gt; 0){`  
`a = a/10`  
`i++`  
`}`  
`return i`  
`}`

The problem is that I can't do   a = a / 10 because I get

&gt; val cannot be reassigned 

but I never assigned a as a val , is this a syntax thing ? Sorry if this is like the most basic thing but I can't find information.
## [7][MVIKotlin 2.0.0-beta2 is out with complete documentation!](https://www.reddit.com/r/Kotlin/comments/g0gq90/mvikotlin_200beta2_is_out_with_complete/)
- url: https://www.reddit.com/r/Kotlin/comments/g0gq90/mvikotlin_200beta2_is_out_with_complete/
---
MVIKotlin is a Kotlin Multiplatform framework that provides a way of (not only) writing shared code using MVI pattern. It also includes powerful debug tools like logging and time travel. The main functionality of the framework does not depend on any reactive or coroutines library. Extensions for Reaktive and for coroutines libraries are provided as separate modules.

GitHub: [https://github.com/arkivanov/MVIKotlin](https://github.com/arkivanov/MVIKotlin)

Documentation: [https://arkivanov.github.io/MVIKotlin](https://arkivanov.github.io/MVIKotlin)
## [8][fetching google knowledge panel results](https://www.reddit.com/r/Kotlin/comments/g0uqof/fetching_google_knowledge_panel_results/)
- url: https://www.reddit.com/r/Kotlin/comments/g0uqof/fetching_google_knowledge_panel_results/
---
 How can I fetch Google knowledge panel result using jsoup. I have to fetch author name of the entered book and show it in my android application. I don't know which elements to look for. Would be grateful for guidance in this regard.
## [9][Building a Proof-of-Concept with Kotlin Multiplatform](https://www.reddit.com/r/Kotlin/comments/g0nf15/building_a_proofofconcept_with_kotlin/)
- url: https://www.youtube.com/watch?v=EJVq_QWaWXE
---

## [10][How to Evaluate Kotlin Multiplatform, React Native and Flutter](https://www.reddit.com/r/Kotlin/comments/g0ncvc/how_to_evaluate_kotlin_multiplatform_react_native/)
- url: https://www.youtube.com/watch?v=6RquJJM1jaE
---

