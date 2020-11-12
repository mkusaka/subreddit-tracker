# reduxjs
## [1][Content of actions and reducers](https://www.reddit.com/r/reduxjs/comments/jslxjb/content_of_actions_and_reducers/)
- url: https://www.reddit.com/r/reduxjs/comments/jslxjb/content_of_actions_and_reducers/
---
I'm new to React Redux and this maybe a weird question. As far as I understand, we handle data fetching, error handling etc. at actions and state logic (map, filter, find etc.) at reducers. Is this simple explanation correct? And what other concepts we handle at actions and reducers? Thanks
## [2][What type of tests do you have in a Redux Toolkit app?](https://www.reddit.com/r/reduxjs/comments/jrp0lg/what_type_of_tests_do_you_have_in_a_redux_toolkit/)
- url: https://www.reddit.com/r/reduxjs/comments/jrp0lg/what_type_of_tests_do_you_have_in_a_redux_toolkit/
---
I have an app containing a few componenets and now looking to add some tests. I have added a slice test following the Toolkit tutorial, but wondering what other tests might be needed?

If Toolkit is a wrapper around Redux, then looking at [Action Creator](https://redux.js.org/recipes/writing-tests#action-creators) tests I'm not really sure whether this type of tests is actually needed in a Toolkit app as actions are not created manually and in Redux?

Integration might be the best test providing more confidence, but unit tests are easier to write and fix. So I'm trying to understand what types of tests are needed to cover all different parts and not miss on a connection seam where it may break. And also not to have same function tested multiple times. 
## [3][Sorry I am new](https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/)
- url: https://www.reddit.com/r/reduxjs/comments/joydjk/sorry_i_am_new/
---
I was wondering if there is anyone that could help me with redux-promise-middleware. I am in school and a part of my personal project got deleted and I presented it but it was a dumpster fire.. I haven't been able to get a hold of my tutor for 2 days. Im sorry if I posted this in the wrong place.... Please Help!!!!
## [4][How to store initially loaded data in the service and not redux store?](https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/)
- url: https://www.reddit.com/r/reduxjs/comments/joho2d/how_to_store_initially_loaded_data_in_the_service/
---
I have an app that fetches data from the server on applicaton startup, like getAuthors, getBooks, etc.

This data is then used to populate dropdowns on app forms (I have only one form now but planning to add more forms that will use same Authours and Books data).

While Books data is being loaded, I want a spinner to be displayed on the Books dropdown.

How can I notify the app when data load is complete so that spinners could be hidden?

Is there a way to achieve this with the service-like patter rather than storing all inital data in the store? (the application is using Redux Toolkit)
## [5][A great react state management library](https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/)
- url: https://www.reddit.com/r/reduxjs/comments/jnwrbr/a_great_react_state_management_library/
---
[https://github.com/concentjs/concent](https://github.com/concentjs/concent)
## [6][Do people use Redux to manage state data of logged in user's username?](https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/)
- url: https://www.reddit.com/r/reduxjs/comments/jn2i3k/do_people_use_redux_to_manage_state_data_of/
---
Just curious if for authentication purposes I should use Redux on the front end. Currently using firebase OAuth on the backend for managing the actual user authentication. Just wondering for front-end if Redux is a popular option for logged in user state management.
## [7][[Code Review] Is there an easier way to conditionally dispatch actions according to itself](https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/)
- url: https://www.reddit.com/r/reduxjs/comments/jkhtwo/code_review_is_there_an_easier_way_to/
---
I'm practicing Redux with Toolkit and trying to use it everywhere I can (even if it's a little over-engineered). I have several values I'm storing in Redux that are booleans with actions to toggle them (i.e isModalOpen, setIsModalOpen).

However I find myself writing a lot of logic to confirm if I should dispatch the action or not based on it's own current state. Is there a less cumbersome way to do this? Can I access state within dispatch? I'd love to refactor my excessive use of useSelector.

[https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56](https://github.com/seandonn-boston/seandonnio/blob/02516fef5a79e1e4a6d2bbed6541623ba4f729d5/src/app/Navigation/Navigation.js#L56)
## [8][Forget About Redux Boilerplate — Now It is Just One Little Hook](https://www.reddit.com/r/reduxjs/comments/jiegy0/forget_about_redux_boilerplate_now_it_is_just_one/)
- url: https://medium.com/parzhitsky/forget-about-redux-boilerplate-now-it-is-just-one-little-hook-bd95a7a44d6f?source=friends_link&amp;sk=b177046bbefc35bc872c5722ad4c5184
---

## [9][React Native with Redux](https://www.reddit.com/r/reduxjs/comments/jg5shh/react_native_with_redux/)
- url: https://www.imaginarycloud.com/blog/react-native-redux/
---

## [10][Is Redux Better Than Mobx?](https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/)
- url: https://www.reddit.com/r/reduxjs/comments/jevwwj/is_redux_better_than_mobx/
---
  In the case of application development that makes use of the component-based library, like ***ReactJs***, every component manages one’s own local state, and the root component usually manages the application state.Several components communicate with the state data and modify it before having it to send back to the server. Components may also manage their state through component life cycle methods. As the complexity of the application increases, the component state becomes unstable and it becomes a long process to debug these applications.The ***state management*** systems such as *Redux* and *MobX* are used to transfer state from components as a stand-alone testable unit to prevent certain situations as well as provide methods for synchronizing application state with react components. When data comes directly from a common source, the logic of the application code becomes clearer.

# Comparison Between  Redux and MobX

The Redux library is somehow very limited. A lot of the time, you’re just working with basic ***JavaScript objects and arrays***. In MobX, the objects and arrays are wrapped in *observable objects* that conceal much of the boilerplate. It develops on hidden abstractions. It’s simpler to deal with plain JavaScript in Redux. It makes it quite simpler to test and debug the application.In comparison, one would have to think again about where we come from in single-page applications. A couple of single-page application frameworks and libraries had the same state management issues that were ultimately overcome by the overarching ***flux pattern*** and Redux is the successor to the pattern.Now in the case of MobX, it appears to go in the opposite direction. Once more, we begin to explicitly mutate the state without reaping the benefits of ***functional programming***. Some individuals feel connected to two-way data binding again. Pretty quickly, individuals may face similar challenges again before a state management library like Redux was implemented. The state management is spread around the components and ends up in chaos.MobX is *less opinionated*. But it will be appropriate to follow the best practices in MobX. Individuals ought to know how to structure state management in order to strengthen their logic about it. Or else, people prefer to directly mutate state into components.To conclude the comparison of ***Redux Vs MobX***, both the libraries are amazing. Although Redux is very renowned, MobX has become a suitable option for state management. 

# Check This Article For Better Overview - [Redux Vs MobX](https://codersera.com/blog/redux-vs-mobx/)
