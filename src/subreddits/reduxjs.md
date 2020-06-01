# reduxjs
## [1][Do you connect every component which fetches data to Redux?](https://www.reddit.com/r/reduxjs/comments/guhewr/do_you_connect_every_component_which_fetches_data/)
- url: https://www.reddit.com/r/reduxjs/comments/guhewr/do_you_connect_every_component_which_fetches_data/
---
 It just seems crazy to add so much boilerplate code for every data-fetching component. I've seen this pattern a lot. Any motives for it?It seems that simple ApiCall in componentDidMount / useEffect would work better...
## [2][What happens when Redux is compiled/minified/deployed?](https://www.reddit.com/r/reduxjs/comments/gr46kl/what_happens_when_redux_is/)
- url: https://www.reddit.com/r/reduxjs/comments/gr46kl/what_happens_when_redux_is/
---
This is a random conceptual question but  I was just thinking, at the end of the day when you launch something and you have static files, they're deployed on some website/app. What does Redux/reducer/actions become... just an object/methods?
## [3][[Puzzler] Action creator using function.name](https://www.reddit.com/r/reduxjs/comments/gqze1l/puzzler_action_creator_using_functionname/)
- url: https://www.reddit.com/r/reduxjs/comments/gqze1l/puzzler_action_creator_using_functionname/
---
Today I encountered interesting bug which was caused by action creators like this:

    export function setCartStatus(status) {
        return {
            type: setCartStatus.name,
            cartStatus: status
        }
    }

All creators had unique names, I checked this several times manually across all modules. All other components (store, reducers etc) were also ok — the problem is only in the code given above. Can you spot one?

Note: comments under the post mention solution.

Answer is under the spoiler: &gt;! webpack minimizes function names, and different functions from different modules may happen to have the same minimized name. !&lt;
## [4][Frontend Cache Question](https://www.reddit.com/r/reduxjs/comments/gqqzkc/frontend_cache_question/)
- url: https://www.reddit.com/r/reduxjs/comments/gqqzkc/frontend_cache_question/
---
Hey everyone! So for the last few weeks I’ve been designing and beginning to implement a GraphQL API using Prisma. I’m at a stage where I have enough of an API foundation to think about developing frontend. I’ll be using react-native and I have a few questions regarding fetching data from my API. I want to be able to provide users with an offline experience and therefore require cache. I also want to use subscriptions (real-time functionality in GraphQL). I’ve found Apollo Client and seen that it has a lot of good reviews, however, I’m not a huge fan of the built in React components that display the data. I haven’t used them so I can’t really be sure that I don’t like them, however I don’t think they’ll be great for query testing which isn’t a huge deal since I’ll be thoroughly testing my API and Prisma is tested. On the other hand I’ve used redux in the past and am wondering if it has the possibly of acting as a sort of cache when paired with simple https requests. Any thoughts are appreciated! 🙏
## [5][Red - Type-safe, composable, boilerplateless reducers ; https://github.com/betafcc/red](https://www.reddit.com/r/reduxjs/comments/gq1rj9/red_typesafe_composable_boilerplateless_reducers/)
- url: https://v.redd.it/t0wnti4hdt051
---

## [6][Is the constructor even needed when using Redux?](https://www.reddit.com/r/reduxjs/comments/gp5xfa/is_the_constructor_even_needed_when_using_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/gp5xfa/is_the_constructor_even_needed_when_using_redux/
---
I haven’t yet had any reason to use it since I place all the props I need in mapStateToProps which goes into the connect function. I was pretty good at passing props around in React but now that I have added Redux I’m not too sure how props are passed around. 

I’m making an application with Redux because I want to expand my web development “toolkit”. So even though my application doesn’t really need Redux I’m just curious about how to use it. I feel like I’ve been taught Redux badly or maybe I’m just not getting it but if you know of any good resources for me to learn that would be great too.
## [7][I worked with Redux for two years and I wrote a blog post about tips and tricks.](https://www.reddit.com/r/reduxjs/comments/goqbi3/i_worked_with_redux_for_two_years_and_i_wrote_a/)
- url: https://www.cyprientaque.com/blog/redux-tips
---

## [8][need some help passing an array to a component after actions dispatch](https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/)
- url: https://www.reddit.com/r/reduxjs/comments/gnw73c/need_some_help_passing_an_array_to_a_component/
---
I am trying to pass an array from a container to a component inside of the container. the issue im having is that i dispatch my  actions in componentDidMount of the container. then i try passing the array that is acquired to the comp but it passes the initial state, an empty array, instead of the third and final state
## [9][What to put on the Redux state](https://www.reddit.com/r/reduxjs/comments/gnir1f/what_to_put_on_the_redux_state/)
- url: https://x8lucas8x.com/what-to-put-on-the-redux-state/
---

## [10][Experimental state management lib for React from Facebook](https://www.reddit.com/r/reduxjs/comments/gkahrf/experimental_state_management_lib_for_react_from/)
- url: https://blog.graphqleditor.com/recoil-react-state-management-library/
---

