# Kotlin
## [1][Trying to proxy data class using ModelMapper](https://www.reddit.com/r/Kotlin/comments/f183zk/trying_to_proxy_data_class_using_modelmapper/)
- url: https://www.reddit.com/r/Kotlin/comments/f183zk/trying_to_proxy_data_class_using_modelmapper/
---
So I'm using ModelMapper to convert between DTO and DB objects. I've got a scenario where my DB object has a few fields that I do not want to expose via the API. So I'm running an update operation, and I'm trying to use ModelMapper to overwrite the DB entity with data from the DTO. See below:

&amp;#x200B;

`return videoFileRepo.findById(fileId)`  
`.map { videoFile -&gt;`  
`modelMapper.map(payload, videoFile)`  
`videoFile.fileId = fileId`  
`val savedVideoFile =` [`videoFileRepo.save`](https://videoFileRepo.save)`(videoFile)`  
`modelMapper.map(savedVideoFile, VideoFilePayload::class.java)`  
`}`  
`.orElse(null)`

My problem is that I have nested entities in the VideoFile object. The DTO has DTO versions of those entities, the DB entity has DB entity versions. So a set of CategoryPayload vs a set of Category, basically. Using ModelMapper this way, with a destination instance instead of a class, doesn't map the nested objects properly.

The solution, I would think, is to create a ModelMapper TypeMap for my classes. That way I could provide more explicit instructions. However, my VideoFilePayload and VideoFile entity objects are both data classes, and ModelMapper throws an error about how it needs to be able to proxy the classes for a TypeMap, and it can't do that with data classes.

So, I'm hoping there is some workaround here where I can proxy my data classes and create the necessary TypeMap. Thanks.
## [2][From Java to Kotlin: life without static](https://www.reddit.com/r/Kotlin/comments/f0t306/from_java_to_kotlin_life_without_static/)
- url: https://jelmini.dev/post/from-java-to-kotlin-life-without-static/
---

## [3][Migrating Java to kotlin (4 minute video)](https://www.reddit.com/r/Kotlin/comments/f0trz5/migrating_java_to_kotlin_4_minute_video/)
- url: https://www.youtube.com/watch?v=w9WE0c6fjRQ
---

## [4][Added programmatic test builder API by Fleshgrinder · Pull Request #2177 · junit-team/junit5 · GitHub](https://www.reddit.com/r/Kotlin/comments/f0vra5/added_programmatic_test_builder_api_by/)
- url: https://github.com/junit-team/junit5/pull/2177
---

## [5][In any development these principles will never let you down and decrease the number of WTF per minute](https://www.reddit.com/r/Kotlin/comments/f139za/in_any_development_these_principles_will_never/)
- url: https://www.reddit.com/r/Kotlin/comments/f139za/in_any_development_these_principles_will_never/
---
“SOLID Principles: Explanation and examples” by Simon LH https://link.medium.com/F66860pCV3
## [6][Indentation when operating on the result of a block - what would you do here?](https://www.reddit.com/r/Kotlin/comments/f05m7b/indentation_when_operating_on_the_result_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/f05m7b/indentation_when_operating_on_the_result_of_a/
---
Given the example block of code, which style do you prefer:

&amp;#x200B;

https://preview.redd.it/thhtiauxwff41.png?width=356&amp;format=png&amp;auto=webp&amp;s=05765122c8bc14a8051f2d027588f2e3173f2056

I encounter this a lot, usually when working with RxJava and I'm never sure what the best approach is, none of them feel very neat to me.
## [7][Kotlin overtakes Scala and Clojure, to become the 2nd most popular language on the JVM](https://www.reddit.com/r/Kotlin/comments/ezqfpn/kotlin_overtakes_scala_and_clojure_to_become_the/)
- url: https://snyk.io/blog/kotlin-overtakes-scala-and-clojure-to-become-the-2nd-most-popular-language-on-the-jvm/
---

## [8][It’s Nothing](https://www.reddit.com/r/Kotlin/comments/ezuhgz/its_nothing/)
- url: https://www.zacsweers.dev/its-nothing/
---

## [9][Detekt-hint is a tool for spotting programming principles violations. I need some feedback on what rules to implement and if you think this could be useful. Please have a look! https://github.com/Mkohm/detekt-hint](https://www.reddit.com/r/Kotlin/comments/ezaoda/detekthint_is_a_tool_for_spotting_programming/)
- url: https://i.redd.it/4roaqhr0c4f41.png
---

## [10][Intellisense or auto-completion not working (Kotlin)](https://www.reddit.com/r/Kotlin/comments/ezwpyo/intellisense_or_autocompletion_not_working_kotlin/)
- url: /r/vscode/comments/ezv4f7/intellisense_or_autocompletion_not_working_kotlin/
---

