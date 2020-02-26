# reduxjs
## [1][Don't Waste Your Ducking Time: An opinionated guide on how to test Redux ducks](https://www.reddit.com/r/reduxjs/comments/f9n9j2/dont_waste_your_ducking_time_an_opinionated_guide/)
- url: https://github.com/tophat/dont-waste-your-ducking-time
---

## [2][How to normalize data ?](https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/)
- url: https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/
---
Hi 👋!Can someone help me with normalizr?

How you actually organize your stored data and where you store meta info about domain entity (like loading status and etc) ?

[https://codesandbox.io/s/cool-kare-6z09b](https://codesandbox.io/s/cool-kare-6z09b)
## [3][Sending post data React](https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/)
- url: https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/
---
 I have a really noobish question I'm working with Redux for the first time and I'm able to get data from a get request but I was hoping if you guys would be able to send me any links or any pointers on how post-data would be sent to the background using Redux when I look at the tutorials some people use mapStateToProps others use proptypes and others just don't pass anything at all any help would be greatly appreciated


 also what I want to do is register a user and when the registration is successful or not I will use that status code to determine whether they should go to a new component. I figured I would be able to do this in the  action.JS file but I don't understand how I would be able to send my state data on the component to Redux and then give me back my I response
## [4][Implementing Undo-Redo with NgRx or Redux](https://www.reddit.com/r/reduxjs/comments/f3q1rp/implementing_undoredo_with_ngrx_or_redux/)
- url: https://nils-mehlhorn.de/posts/angular-undo-redo-ngrx-redux
---

## [5][React — state management without libraries (with hooks) PART 2](https://www.reddit.com/r/reduxjs/comments/f2ofaw/react_state_management_without_libraries_with/)
- url: https://medium.com//react-state-management-without-libraries-with-hooks-part-2-d087278185a9?source=friends_link&amp;sk=a184f029f305821c4c57d2c249d0042b
---

## [6][Redux Controllers -&gt; Less verbose Redux for people who love typescript - State Observables | State Subscriptions | State Watchers | Async Actions | Inbuilt Caching | React | Angular | Node | Redux Effects | Provided States](https://www.reddit.com/r/reduxjs/comments/f2cesn/redux_controllers_less_verbose_redux_for_people/)
- url: https://www.npmjs.com/package/redux-controllers
---

## [7][Do I need a Store?](https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/)
- url: https://www.reddit.com/r/reduxjs/comments/f1061x/do_i_need_a_store/
---
I am watching a tutorial atm, and he said you can set up a Store if you like.

But he appears to be putting, what would be the Store, in the index.js file (in the client side).

Is this common practice? Has anyone seen this 'style' of organization before?
## [8][Redux - useSelector empty array equality](https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/)
- url: https://www.reddit.com/r/reduxjs/comments/f0m1l4/redux_useselector_empty_array_equality/
---
Would this be a sane way to ensure that for \[\] each time is true on reference equality. E.g. if that state.items remain undefined each time I use a fallback \[\] would fail reference equality - is useRef the right approach here? (rather not use shallowEquality)

    const items = useSelector((state) =&gt; state.items || useRef([]).current);
## [9][Conditional update in reducer based on other state. Best Practices?](https://www.reddit.com/r/reduxjs/comments/f05rab/conditional_update_in_reducer_based_on_other/)
- url: https://www.reddit.com/r/reduxjs/comments/f05rab/conditional_update_in_reducer_based_on_other/
---
So currently I have an event called ITEM\_DELETED.

When this happens, I update the items part of the state to filter out the item that was deleted.

But I also want to reset a variable (in **another** part of the state) to 0, IF the item that was deleted is currently selected

What's the best practice to do this when I'm using **combineReducers** and the state is separated out, so the reducer doesn't necessarily know if the item was originally selected.

I'm currently thinking:

1. Probably best, the action creator puts a boolean in the payload if the item deleted was selected, that way the reducers can still respond to the same action type.
2. I can update the action creator to use thunk and dispatch two actions, the second with a different type and only being if the item was currently selected
3. Something else? Make the state combined so one reducer can handle the action?

Thanks!
## [10][Which is more performant for useSelector hook with multiple values?](https://www.reddit.com/r/reduxjs/comments/ezsw20/which_is_more_performant_for_useselector_hook/)
- url: https://www.reddit.com/r/reduxjs/comments/ezsw20/which_is_more_performant_for_useselector_hook/
---
As you know, you can use object destructuring for an object to extract values with the same name:

    const values = { 'name': 'John', 'age': 20, 'country': 'USA' }
    const { name, age } = values

But when using Redux's useSelector hook, is it better to have them in their own calls or use the same logic?

For example, let's look at the initial state which is part of a reducer that will eventually be called userInfo in a combineReducers call:

    const initialState = { 'name': '', 'age': 0, 'country': '' }

If I am in a component, is it better to do

    const { name, age } = useSelector(state =&gt; state.userInfo)

or

    const name = useSelector(state =&gt; state.name)
    const age = useSelector(state =&gt; state.age)
    
The reason I ask is that the first example is one call but may initially bring in all the values for that userInfo state while the second example is more direct, but then also calls useSelector an additional time.

To be more performant and cut down on re-renders on data change, which one is the better method?
