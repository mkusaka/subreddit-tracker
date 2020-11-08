# reduxjs
## [1][Sorry I am new](https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/)
- url: https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/
---
I was wondering if there is anyone that could help me with redux-promise-middleware. I am in school and a part of my personal project got deleted and I presented it but it was a dumpster fire.. I haven't been able to get a hold of my tutor for 2 days. Im sorry if I posted this in the wrong place.... Please Help!!!!
## [2][How to store initially loaded data in the service and not redux store?](https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/)
- url: https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/
---
I have an app that fetches data from the server on applicaton startup, like getAuthors, getBooks, etc.

This data is then used to populate dropdowns on app forms (I have only one form now but planning to add more forms that will use same Authours and Books data).

While Books data is being loaded, I want a spinner to be displayed on the Books dropdown.

How can I notify the app when data load is complete so that spinners could be hidden?

Is there a way to achieve this with the service-like patter rather than storing all inital data in the store? (the application is using Redux Toolkit)
## [3][A great react state management library](https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/)
- url: https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/
---
[https://github.com/concentjs/concent](https://github.com/concentjs/concent)
## [4][Do people use Redux to manage state data of logged in user's username?](https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/)
- url: https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/
---
Just curious if for authentication purposes I should use Redux on the front end. Currently using firebase OAuth on the backend for managing the actual user authentication. Just wondering for front-end if Redux is a popular option for logged in user state management.
## [5][[Code Review] Is there an easier way to conditionally dispatch actions according to itself](https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/
---
I'm practicing Redux with Toolkit and trying to use it everywhere I can (even if it's a little over-engineered). I have several values I'm storing in Redux that are booleans with actions to toggle them (i.e isModalOpen, setIsModalOpen).

However I find myself writing a lot of logic to confirm if I should dispatch the action or not based on it's own current state. Is there a less cumbersome way to do this? Can I access state within dispatch? I'd love to refactor my excessive use of useSelector.

[https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56](https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56)
## [6][Forget About Redux Boilerplate — Now It is Just One Little Hook](https://www.reddit.com/r/reduxjs/comments/jiegy0/forget_about_redux_boilerplate_now_it_is_just_one/)
- url: https://medium.com/parzhitsky/forget-about-redux-boilerplate-now-it-is-just-one-little-hook-bd95a7a44d6f?source=friends_link&amp;sk=b177046bbefc35bc872c5722ad4c5184
---

## [7][React Native with Redux](https://www.reddit.com/r/reduxjs/comments/jg5shh/react_native_with_redux/)
- url: https://www.imaginarycloud.com/blog/react-native-redux/
---

## [8][Is Redux Better Than Mobx?](https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/)
- url: https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/
---
  In the case of application development that makes use of the component-based library, like ***ReactJs***, every component manages one’s own local state, and the root component usually manages the application state.Several components communicate with the state data and modify it before having it to send back to the server. Components may also manage their state through component life cycle methods. As the complexity of the application increases, the component state becomes unstable and it becomes a long process to debug these applications.The ***state management*** systems such as *Redux* and *MobX* are used to transfer state from components as a stand-alone testable unit to prevent certain situations as well as provide methods for synchronizing application state with react components. When data comes directly from a common source, the logic of the application code becomes clearer.

# Comparison Between  Redux and MobX

The Redux library is somehow very limited. A lot of the time, you’re just working with basic ***JavaScript objects and arrays***. In MobX, the objects and arrays are wrapped in *observable objects* that conceal much of the boilerplate. It develops on hidden abstractions. It’s simpler to deal with plain JavaScript in Redux. It makes it quite simpler to test and debug the application.In comparison, one would have to think again about where we come from in single-page applications. A couple of single-page application frameworks and libraries had the same state management issues that were ultimately overcome by the overarching ***flux pattern*** and Redux is the successor to the pattern.Now in the case of MobX, it appears to go in the opposite direction. Once more, we begin to explicitly mutate the state without reaping the benefits of ***functional programming***. Some individuals feel connected to two-way data binding again. Pretty quickly, individuals may face similar challenges again before a state management library like Redux was implemented. The state management is spread around the components and ends up in chaos.MobX is *less opinionated*. But it will be appropriate to follow the best practices in MobX. Individuals ought to know how to structure state management in order to strengthen their logic about it. Or else, people prefer to directly mutate state into components.To conclude the comparison of ***Redux Vs MobX***, both the libraries are amazing. Although Redux is very renowned, MobX has become a suitable option for state management. 

# Check This Article For Better Overview - [Redux Vs MobX](https://codersera.com/blog/redux-vs-mobx/)
## [9][Toolkit in Production?](https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/)
- url: https://www.reddit.com/r/reduxjs/comments/ja2ssv/toolkit_in_production/
---
I see that [https://redux.js.org/tutorials/essentials/part-2-app-structure](https://redux.js.org/tutorials/essentials/part-2-app-structure) Redux is pushing toolkit now. How many of you are using Redux Toolkit in production?
## [10][Where to add hooks after dynamically adding reducers](https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/)
- url: https://www.reddit.com/r/reduxjs/comments/j6v8zd/where_to_add_hooks_after_dynamically_adding/
---
Hey all,

I've seen a few similar questions to this but thought I might go ahead and ask mine.

I'm dynamically adding reducers and middleware to my App but now I'm to the point where I want to tie those reducers I added to the Apps store for reference (Essentially, I'm trying to create self-contained widgets that adds what they need to the Redux store so that I may eventually add/remove components)

I'm not sure how to do it though. I'm using Redux Observables... I'm trying to useStore... or useSelector but I can't use them inside useEffect but I also can't try and access it before the reducers have been dynamically added in the use effect. I've done a ton of googling on this and I feel like I'm missing something key here....
