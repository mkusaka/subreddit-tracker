# reduxjs
## [1][Two ways to skin a form. Design Pattern Discussion: How would you guys implement the following?](https://www.reddit.com/r/reduxjs/comments/fe1thu/two_ways_to_skin_a_form_design_pattern_discussion/)
- url: https://www.reddit.com/r/reduxjs/comments/fe1thu/two_ways_to_skin_a_form_design_pattern_discussion/
---
Let’s say I’m posting/putting an object on a form that is essentially an add/edit screen for "car".  When i submit, I obviously wanna know what happens so I can navigate away or show errors on the form if necessary.   I could do this one of two ways:

1. I could treat the form and action/saga as a closed environment that never goes to the store at all.  I simply dispatch an action POST_VERSION and the payload is the body AND callbacks for success/fail/statuschange. Then the saga calls those callbacks “oh i started…..oh i succeeded…..oh i got validation errors”.  These callbacks are handled on the front end and do what is appropriate.  The store neither knows nor cares that the form exists or is doing anything.  I use this pattern for really strict UI interaction sometimes.

2. I could make the store have a slice called “current_car_being_edited" or something, and it has the statuses and the errors and uses the classic redux reducers etc to notify the front end via state change.   I know this is the "redux way" but something about it feels....loose and incorrect.  Is my DB brain overthinking it?   Solution #1 is so snappy and elegant, maybe this case is just too simple to need it?  Would you use it anyways?
## [2][state is an object tree ? What is that supposed to mean ?](https://www.reddit.com/r/reduxjs/comments/fdvrge/state_is_an_object_tree_what_is_that_supposed_to/)
- url: https://www.reddit.com/r/reduxjs/comments/fdvrge/state_is_an_object_tree_what_is_that_supposed_to/
---
What does the tree supposed to mean ?
## [3][Don't Waste Your Ducking Time: An opinionated guide on how to test Redux ducks](https://www.reddit.com/r/reduxjs/comments/f9n9j2/dont_waste_your_ducking_time_an_opinionated_guide/)
- url: https://github.com/tophat/dont-waste-your-ducking-time
---

## [4][How to normalize data ?](https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/)
- url: https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/
---
Hi 👋!Can someone help me with normalizr?

How you actually organize your stored data and where you store meta info about domain entity (like loading status and etc) ?

[https://codesandbox.io/s/cool-kare-6z09b](https://codesandbox.io/s/cool-kare-6z09b)
## [5][Sending post data React](https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/)
- url: https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/
---
 I have a really noobish question I'm working with Redux for the first time and I'm able to get data from a get request but I was hoping if you guys would be able to send me any links or any pointers on how post-data would be sent to the background using Redux when I look at the tutorials some people use mapStateToProps others use proptypes and others just don't pass anything at all any help would be greatly appreciated


 also what I want to do is register a user and when the registration is successful or not I will use that status code to determine whether they should go to a new component. I figured I would be able to do this in the  action.JS file but I don't understand how I would be able to send my state data on the component to Redux and then give me back my I response
## [6][Implementing Undo-Redo with NgRx or Redux](https://www.reddit.com/r/reduxjs/comments/f3q1rp/implementing_undoredo_with_ngrx_or_redux/)
- url: https://nils-mehlhorn.de/posts/angular-undo-redo-ngrx-redux
---

## [7][React — state management without libraries (with hooks) PART 2](https://www.reddit.com/r/reduxjs/comments/f2ofaw/react_state_management_without_libraries_with/)
- url: https://medium.com//react-state-management-without-libraries-with-hooks-part-2-d087278185a9?source=friends_link&amp;sk=a184f029f305821c4c57d2c249d0042b
---

## [8][Redux Controllers -&gt; Less verbose Redux for people who love typescript - State Observables | State Subscriptions | State Watchers | Async Actions | Inbuilt Caching | React | Angular | Node | Redux Effects | Provided States](https://www.reddit.com/r/reduxjs/comments/f2cesn/redux_controllers_less_verbose_redux_for_people/)
- url: https://www.npmjs.com/package/redux-controllers
---

## [9][Do I need a Store?](https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/)
- url: https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/
---
I am watching a tutorial atm, and he said you can set up a Store if you like.

But he appears to be putting, what would be the Store, in the index.js file (in the client side).

Is this common practice? Has anyone seen this 'style' of organization before?
## [10][Redux - useSelector empty array equality](https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/)
- url: https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/
---
Would this be a sane way to ensure that for \[\] each time is true on reference equality. E.g. if that state.items remain undefined each time I use a fallback \[\] would fail reference equality - is useRef the right approach here? (rather not use shallowEquality)

    const items = useSelector((state) =&gt; state.items || useRef([]).current);
