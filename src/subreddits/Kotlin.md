# Kotlin
## [1][Coroutines and callbacks](https://www.reddit.com/r/Kotlin/comments/fvd35o/coroutines_and_callbacks/)
- url: https://www.coroutinedispatcher.com/2020/04/coroutines-and-callbacks.html
---

## [2][Template-Oriented-Programming to Ship Faster | Kotlin, Arrow | KotlinHyderabad](https://www.reddit.com/r/Kotlin/comments/fvd5dh/templateorientedprogramming_to_ship_faster_kotlin/)
- url: https://youtu.be/_QBlKtUY6ac
---

## [3][Boilerplate free proxy fakes](https://www.reddit.com/r/Kotlin/comments/fvb2z1/boilerplate_free_proxy_fakes/)
- url: https://www.reddit.com/r/Kotlin/comments/fvb2z1/boilerplate_free_proxy_fakes/
---
Where do you stand on mocks vs fakes? Would you go for more fakes if you had less boilerplate and noisy code?

Inspired by retrofit, I tried out a proxy based approach to solve this issue and make fakes a bit more pleasant. Have a look below.

“Boilerplate free proxy fakes” by Fanis Veizis https://link.medium.com/Rql07BXVq5
## [4][Can I make a Discord bot using Kotlin?](https://www.reddit.com/r/Kotlin/comments/fux84n/can_i_make_a_discord_bot_using_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/fux84n/can_i_make_a_discord_bot_using_kotlin/
---
Coming from a Python background, I don't really know about all the applications of Kotlin. It seems like a developer friendly language and I want to try out tasks simply done in Python using Kotlin.

How can I go about this?
## [5][Open Source Food Delivery app and Local Marketplace Builder](https://www.reddit.com/r/Kotlin/comments/fuu6of/open_source_food_delivery_app_and_local/)
- url: https://www.reddit.com/r/Kotlin/comments/fuu6of/open_source_food_delivery_app_and_local/
---
Github : [https://github.com/NearbyShops/Nearby-Shops-Android-app](https://github.com/NearbyShops/Nearby-Shops-Android-app)

Hey this is an Android based open source project so i hope you guys may like to check this out. 

In this Covid-19 Crises we are facing a crises of delivery of food, groceries and other essential items. You can use this open-source project to build your own local marketplace for your Local Community.
## [6][How to avoid code repetition?](https://www.reddit.com/r/Kotlin/comments/fuybbx/how_to_avoid_code_repetition/)
- url: https://www.reddit.com/r/Kotlin/comments/fuybbx/how_to_avoid_code_repetition/
---
Hello Kotliners!

I would like to ask your help in some code snippet that haunts me. I have a Spring controller with a few routes. Each of this routes has the lines of code that can be presented like:

`//some code before which is specific for route`  
`collectionOfSomething.forEach {`  
`if (incomingParameterValid) {`  
`service.methodWhichDoingSomething(enumValue)`  
`} else {`  
`service.methodWhichDoingSomething(defaultEnumValue)`  
`}`  
`}`  
`//some code after which is specific for route`

where, 

`incomingParameterValid` \- is the same for all routes.

`service` \- is the same for all routes.

`methodWhichDoingSomething` \- is the same for all routes.

`enumValue` &amp; `defaultEnumValue` \- are specific for each route.

So,  the question is: since I really hate to see all these repetitions, is there a way how to extract duplicated lines of code and call them with different parameters `enumValue` &amp; `defaultEnumValue` for each route?

&gt;!Please note!!&lt;

One possible way that I found is creating of private method and moving repeatable lines into it (but it does not seem like a good variant):

`private fun sharedMethod(collection: List&lt;of sth&gt;, incomingParameterValid: Boolean, service:`   
`Service, enumValue: Enum, enumValueDefault: Enum) {`  
`collection.forEach {`  
`if (incomingParameterValid) {`  
`service.methodWhichDoingSomething(enumValue)`  
`} else {`  
`service.methodWhichDoingSomething(enumValueDefault)`  
`}`  
`}`  
`}`
## [7][GraphQLize Update (#1) - Pagination, Scalar Types, Scala &amp; Kotlin Support](https://www.reddit.com/r/Kotlin/comments/fux6v9/graphqlize_update_1_pagination_scalar_types_scala/)
- url: https://www.graphqlize.org/blog/graphqlize-update-1
---

## [8][Which Kotlin libraries are you missing?](https://www.reddit.com/r/Kotlin/comments/futxtv/which_kotlin_libraries_are_you_missing/)
- url: https://www.reddit.com/r/Kotlin/comments/futxtv/which_kotlin_libraries_are_you_missing/
---
Sitting at home in quarantine I have some extra time on my hands to do some open source work.

1) What Kotlin libraries do you think are missing in the ecosystem?

2) which Kotlin projects are in need of contributions?
## [9][How can I subtract the valor of the last but one item (penultimate) from the last item in a list?](https://www.reddit.com/r/Kotlin/comments/fuye1b/how_can_i_subtract_the_valor_of_the_last_but_one/)
- url: https://www.reddit.com/r/Kotlin/comments/fuye1b/how_can_i_subtract_the_valor_of_the_last_but_one/
---
I need help with this question. I need to subtract the valor of the penultimate item from the last item in a list.
## [10][Coroutines flow with Kotlin](https://www.reddit.com/r/Kotlin/comments/fu7dl3/coroutines_flow_with_kotlin/)
- url: https://youtu.be/CIvjwIfOG5A
---

