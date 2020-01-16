# reduxjs
## [1][Can I see an example of redux-thunk callback hell?](https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/)
- url: https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/
---
I am looking into using redux in a real enterprise app for the first time. In my team we have been discussing using either redux-thunk or redux-saga for side effect management.

I have read that you might end up in "callback hell" with redux-thunk. Is there an example of this or can I get an explanation? Having not seen it makes it difficult to judge how big of an issue it is.
## [2][React Ninjas Newsletter #86](https://www.reddit.com/r/reduxjs/comments/eoakib/react_ninjas_newsletter_86/)
- url: https://reactninjs.com/post/react-ninjas-newsletter-86
---

## [3][Implementing Undo/Redo Functionality in Redux using Immer](https://www.reddit.com/r/reduxjs/comments/en4siw/implementing_undoredo_functionality_in_redux/)
- url: https://techinscribed.com/implementing-undo-redo-functionality-in-redux-using-immer/
---

## [4][Why should I use Redux instead of Cookies or LocalStorage](https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/)
- url: https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/
---
Hi folks. 

I have a question about the redux and that is: Why should I use Redux instead of Cookies of LocalStorage. 

Is there any benefit to it?
## [5][Redux &amp; React Router](https://www.reddit.com/r/reduxjs/comments/ektc6g/redux_react_router/)
- url: https://www.reddit.com/r/reduxjs/comments/ektc6g/redux_react_router/
---
I have recently been doing some digging in how to get a SPA up and running within Laravel using React &amp; Redux and noticed a great deal of tutorials and information on the internet using React Router to handle routing within the application.

While this is quite fine it made me wonder what purpose routing serves in a Redux project where I can just as easily dispatch an action to update my store and change the application view.

If we consider a basic example of a web application with two views `A` and `B`, I could implement this in one of two ways:

1) Routing

Create two routes.

    &lt;Route path="/route-a"&gt;
    &lt;Route path="/route-b"&gt;

When I wish to switch views I can do so quite easily by changing routes (via Link, Redirect etc).

2) Redux Store

Track my applications current view within the redux store.

    type View = 'route-a' | 'route-b'
    
    type StoreData = {
        currentView:View
    }

When I wish to switch views I simply dispatch an action to change the current view (`currentView`).

From what I can see the the main difference between the two implementations is that:

* When we change routes the current state of the application is lost (unless we use redux-persist)
* Using routing allows us to build a history stack (e.g. users can navigate back within browsers with the back button)
* Assuming we use redux-persist we introduce an overhead with rehydrating the state on route changes (and perhaps some complexity regarding the merging of state when this occurs).

Personally the overall design and flow of the application feels much simpler with a purely store based solution given I need not worry about about state rehydration on route changes (but can still leverage this behaviour for users who navigate to a different site then come back).

Given the sheer volume of information floating around the internet using React Router I am left wondering ... what am I missing?
## [6][redux-saga style-guide [x-post]](https://www.reddit.com/r/reduxjs/comments/eik2ak/reduxsaga_styleguide_xpost/)
- url: https://www.reddit.com/r/reactjs/comments/eik1od/reduxsaga_styleguide
---

## [7][Redux in 30 seconds](https://www.reddit.com/r/reduxjs/comments/eiktfu/redux_in_30_seconds/)
- url: https://medium.com/@pulkitsingh01/redux-in-30-seconds-4527722aa068
---

## [8][Understand and implement your own Redux](https://www.reddit.com/r/reduxjs/comments/ehoea2/understand_and_implement_your_own_redux/)
- url: https://buttercms.com/blog/understand-and-implement-your-own-redux
---

## [9][Simplifying WebSockets in RxJS.](https://www.reddit.com/r/reduxjs/comments/ehbvdf/simplifying_websockets_in_rxjs/)
- url: https://medium.com//simplifying-websockets-in-rxjs-a177b887f3b8?source=friends_link&amp;sk=a932c9f76ea91c8a763cc1963dad337e
---

## [10][Working with Nested Objects](https://www.reddit.com/r/reduxjs/comments/eghyy3/working_with_nested_objects/)
- url: https://www.reddit.com/r/reduxjs/comments/eghyy3/working_with_nested_objects/
---
Relatively new to React-Redux here, and wondering if someone can direct me to examples of how to work with updating nested objects. I'm aware that \`combineReducers\` can be used, but I've not yet seem examples of how Redux can be coded to deal with levels-within-levels of data storage.

Thanks for any help on this!
