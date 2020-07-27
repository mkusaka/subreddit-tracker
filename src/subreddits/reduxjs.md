# reduxjs
## [1][Top 8 Commandments for building apps with Redux](https://www.reddit.com/r/reduxjs/comments/hyqegy/top_8_commandments_for_building_apps_with_redux/)
- url: https://blog.logrocket.com/8-definitive-rules-building-apps-redux/
---

## [2][How I Made the Django React and Redux Blog](https://www.reddit.com/r/reduxjs/comments/hyba94/how_i_made_the_django_react_and_redux_blog/)
- url: https://www.codeingschool.com/2020/07/how-i-made-django-react-blog.html
---

## [3][I've been using just one saga file and it is getting nasty, should I separate them into different saga?](https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/)
- url: https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/
---
I've seen a lot of reducers have been separately saved into multiple files, but haven't seen many sagas like that? 

is it alright to separate them to clean up some codes?
## [4][redux-toolkit unit testing strategy?](https://www.reddit.com/r/reduxjs/comments/hvwqc9/reduxtoolkit_unit_testing_strategy/)
- url: https://www.reddit.com/r/reduxjs/comments/hvwqc9/reduxtoolkit_unit_testing_strategy/
---
Hi All,

I am using redux-toolkit for the 1st time, got a solid understanding of the concepts using docs. I have previously written unit tests for traditional redux - actions, reducers.

&amp;#x200B;

Wondering what is the "official" strategy for writing tests for slices which have

1. standard reducers key.
2. includes asyncThunks with extra-reducers.

&amp;#x200B;

update: 

using testing-library.

i checked Mark's reply here on slice reducer [https://stackoverflow.com/a/61921088](https://stackoverflow.com/a/61921088)

but **still need some approach for asyncThunks with extra-reducers.**

&amp;#x200B;

&amp;#x200B;

Regards.
## [5][Do we need to install redeux as a dependency if we are installing redux-toolkit?](https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/)
- url: https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/
---
Hi All,

Was starting out on redux-toolkit. Thanks to the awesome docs, i was able to set up a sample app with redux-toolkit. but i was wondering do we need to install redux explicitly as a dependency if we are installing redux-toolkit??

&amp;#x200B;

Looking for a definitive answer from the community,  FYI, i have installed only "@reduxjs/toolkit" &amp; "react-redux" and the app is working fine. 

&amp;#x200B;

Regards.
## [6][Building Scalable Redux-First Apps](https://www.reddit.com/r/reduxjs/comments/huc8ok/building_scalable_reduxfirst_apps/)
- url: https://medium.com/@robbiehall26/building-scalable-redux-first-apps-5a8d09e9bd04?sk=23a705bcad8d07e47500bf382213619d
---

## [7][New "Redux Essentials" core docs tutorial is LIVE! Teaches how to use Redux the right way, using our latest recommended APIs and practices](https://www.reddit.com/r/reduxjs/comments/hr3yx1/new_redux_essentials_core_docs_tutorial_is_live/)
- url: https://redux.js.org/tutorials/essentials/part-1-overview-concepts
---

## [8][Are graphs better than normalized state for complex apps ?](https://www.reddit.com/r/reduxjs/comments/hotwnp/are_graphs_better_than_normalized_state_for/)
- url: https://www.reddit.com/r/reduxjs/comments/hotwnp/are_graphs_better_than_normalized_state_for/
---
I have never used redux . But I have spent time looking at the docs of redux and mobx (I have used mobx) . I was reading [this](https://medium.com/hackernoon/becoming-fully-reactive-an-in-depth-explanation-of-mobservable-55995262a254) article about mobx and I stumbled upon the following sentence :

&gt;for any app that is more complex than TodoMVC, you will often need a data graph, instead of a normalized tree, to store the state in a mentally manageable yet optimal way.

I really find this sentence confusing . We can normalize our state [as explained nicely in the redux docs](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/) and we can create relationships tables between the entities with their ids  . I can not understand how can that break in a complex app . Can anyone help me here ?

Edit : Maybe the answer is [here](https://medium.com/@katedoesdev/normalized-vs-denormalized-databases-210e1d67927d) .
## [9][Modern Redux with Redux Toolkit [OC]](https://www.reddit.com/r/reduxjs/comments/hm8bvh/modern_redux_with_redux_toolkit_oc/)
- url: https://wunnle.com/modern-redux-with-redux-toolkit
---

## [10][Do I need Redux if I have Firebase?](https://www.reddit.com/r/reduxjs/comments/hmag1s/do_i_need_redux_if_i_have_firebase/)
- url: https://www.reddit.com/r/reduxjs/comments/hmag1s/do_i_need_redux_if_i_have_firebase/
---
I use Firebase Authentication and my app works fine. I want to implement a way to simply store username, first name, last name and a JSON object after the user is signed in, so that I don’t have to fetch for them on render of each screen (which may get costly).

I read many of articles online and everyone is insisting on Redux, but it is really so much code to simply store 3 string variables and 1 object, globally. I have class based components so I can’t use `React.useContext` either. 

How else could I do this? Perhaps Asyncstorage? Is that a good idea? Any help is much appreciated :)
