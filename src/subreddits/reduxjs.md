# reduxjs
## [1][Do anyone know an example to get action button, to reveal / hide password in a password input with action and reducer?](https://www.reddit.com/r/reduxjs/comments/es08m7/do_anyone_know_an_example_to_get_action_button_to/)
- url: https://www.reddit.com/r/reduxjs/comments/es08m7/do_anyone_know_an_example_to_get_action_button_to/
---
hi everyone,

i have a code, where i render a password input with redux-form, but inside this input, i want to get a buttom to hide/reveal the password, i am trying to do it with action and reducer to keep this function globally, so i can use it in anothers passwords inputs too.

Do you know any example to do this? i am new using react and redux, i would like to find an example that would be usefull for me as a guide.
## [2][How to access redux store from across components??](https://www.reddit.com/r/reduxjs/comments/er8db0/how_to_access_redux_store_from_across_components/)
- url: https://www.reddit.com/r/reduxjs/comments/er8db0/how_to_access_redux_store_from_across_components/
---
So I have an online store project. 
On the homepage I'm getting the data for each book from mapStateToProps.

What I want is to have a buttons on homepage for each book that when clicked will add that book to the cart. 

The cart is a separate page I have using reactRouter. 

I'm so fkn lost and have tried everything, I tried doing  mapDispatchToProps but its quite mind boggling for me how I can make the button on the homepage to send the data from the store to the cart page. I'm lost and want to die because its been 3 weeks working on this shit and I'm not getting ANYWHERE. pls help me pls
## [3][redux-devtools-extension as a webpage component](https://www.reddit.com/r/reduxjs/comments/epyvml/reduxdevtoolsextension_as_a_webpage_component/)
- url: https://www.reddit.com/r/reduxjs/comments/epyvml/reduxdevtoolsextension_as_a_webpage_component/
---
Is redux-devtools-extension able to be used as a component on the page so that a user can see the state, actions, diff, etc? Or a library that allow you to do this?
## [4][Can I see an example of redux-thunk callback hell?](https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/)
- url: https://www.reddit.com/r/reduxjs/comments/ep8oc5/can_i_see_an_example_of_reduxthunk_callback_hell/
---
I am looking into using redux in a real enterprise app for the first time. In my team we have been discussing using either redux-thunk or redux-saga for side effect management.

I have read that you might end up in "callback hell" with redux-thunk. Is there an example of this or can I get an explanation? Having not seen it makes it difficult to judge how big of an issue it is.
## [5][React Ninjas Newsletter #86](https://www.reddit.com/r/reduxjs/comments/eoakib/react_ninjas_newsletter_86/)
- url: https://reactninjs.com/post/react-ninjas-newsletter-86
---

## [6][Implementing Undo/Redo Functionality in Redux using Immer](https://www.reddit.com/r/reduxjs/comments/en4siw/implementing_undoredo_functionality_in_redux/)
- url: https://techinscribed.com/implementing-undo-redo-functionality-in-redux-using-immer/
---

## [7][Why should I use Redux instead of Cookies or LocalStorage](https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/)
- url: https://www.reddit.com/r/reduxjs/comments/en7eh7/why_should_i_use_redux_instead_of_cookies_or/
---
Hi folks. 

I have a question about the redux and that is: Why should I use Redux instead of Cookies of LocalStorage. 

Is there any benefit to it?
## [8][Redux &amp; React Router](https://www.reddit.com/r/reduxjs/comments/ektc6g/redux_react_router/)
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
## [9][redux-saga style-guide [x-post]](https://www.reddit.com/r/reduxjs/comments/eik2ak/reduxsaga_styleguide_xpost/)
- url: https://www.reddit.com/r/reactjs/comments/eik1od/reduxsaga_styleguide
---

## [10][Redux in 30 seconds](https://www.reddit.com/r/reduxjs/comments/eiktfu/redux_in_30_seconds/)
- url: https://medium.com/@pulkitsingh01/redux-in-30-seconds-4527722aa068
---

