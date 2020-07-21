# reduxjs
## [1][Do we need to install redeux as a dependency if we are installing redux-toolkit?](https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/)
- url: https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/
---
Hi All,

Was starting out on redux-toolkit. Thanks to the awesome docs, i was able to set up a sample app with redux-toolkit. but i was wondering do we need to install redux explicitly as a dependency if we are installing redux-toolkit??

&amp;#x200B;

Looking for a definitive answer from the community,  FYI, i have installed only "@reduxjs/toolkit" &amp; "react-redux" and the app is working fine. 

&amp;#x200B;

Regards.
## [2][Building Scalable Redux-First Apps](https://www.reddit.com/r/reduxjs/comments/huc8ok/building_scalable_reduxfirst_apps/)
- url: https://medium.com/@robbiehall26/building-scalable-redux-first-apps-5a8d09e9bd04?sk=23a705bcad8d07e47500bf382213619d
---

## [3][New "Redux Essentials" core docs tutorial is LIVE! Teaches how to use Redux the right way, using our latest recommended APIs and practices](https://www.reddit.com/r/reduxjs/comments/hr3yx1/new_redux_essentials_core_docs_tutorial_is_live/)
- url: https://redux.js.org/tutorials/essentials/part-1-overview-concepts
---

## [4][Are graphs better than normalized state for complex apps ?](https://www.reddit.com/r/reduxjs/comments/hotwnp/are_graphs_better_than_normalized_state_for/)
- url: https://www.reddit.com/r/reduxjs/comments/hotwnp/are_graphs_better_than_normalized_state_for/
---
I have never used redux . But I have spent time looking at the docs of redux and mobx (I have used mobx) . I was reading [this](https://medium.com/hackernoon/becoming-fully-reactive-an-in-depth-explanation-of-mobservable-55995262a254) article about mobx and I stumbled upon the following sentence :

&gt;for any app that is more complex than TodoMVC, you will often need a data graph, instead of a normalized tree, to store the state in a mentally manageable yet optimal way.

I really find this sentence confusing . We can normalize our state [as explained nicely in the redux docs](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/) and we can create relationships tables between the entities with their ids  . I can not understand how can that break in a complex app . Can anyone help me here ?

Edit : Maybe the answer is [here](https://medium.com/@katedoesdev/normalized-vs-denormalized-databases-210e1d67927d) .
## [5][Modern Redux with Redux Toolkit [OC]](https://www.reddit.com/r/reduxjs/comments/hm8bvh/modern_redux_with_redux_toolkit_oc/)
- url: https://wunnle.com/modern-redux-with-redux-toolkit
---

## [6][Do I need Redux if I have Firebase?](https://www.reddit.com/r/reduxjs/comments/hmag1s/do_i_need_redux_if_i_have_firebase/)
- url: https://www.reddit.com/r/reduxjs/comments/hmag1s/do_i_need_redux_if_i_have_firebase/
---
I use Firebase Authentication and my app works fine. I want to implement a way to simply store username, first name, last name and a JSON object after the user is signed in, so that I don’t have to fetch for them on render of each screen (which may get costly).

I read many of articles online and everyone is insisting on Redux, but it is really so much code to simply store 3 string variables and 1 object, globally. I have class based components so I can’t use `React.useContext` either. 

How else could I do this? Perhaps Asyncstorage? Is that a good idea? Any help is much appreciated :)
## [7][Build a Shopping Cart with React, Redux, and React-DnD by Eyong Kevin](https://www.reddit.com/r/reduxjs/comments/hla4vw/build_a_shopping_cart_with_react_redux_and/)
- url: https://www.reddit.com/r/reduxjs/comments/hla4vw/build_a_shopping_cart_with_react_redux_and/
---
Series on building a simple shopping cart that takes advantage of React DnD. A set of React utilities to help build complex drag and drop interfaces while keeping your components decoupled.

[Part 1](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-1-2433558c3f38?source=friends_link&amp;sk=9e2e7192050b93a5ece2b204a6c36550)

[Part 2](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-2-b4cd649e25db?source=friends_link&amp;sk=29b34207bc446ff51c420f34575d968e)

[Part 3](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-3-f1e1e8265d14?source=friends_link&amp;sk=83a5cb1a1b3df19d0db6d48dd4fce14c)
## [8][How to use redux-saga with graphql?](https://www.reddit.com/r/reduxjs/comments/hk26jy/how_to_use_reduxsaga_with_graphql/)
- url: https://www.reddit.com/r/reduxjs/comments/hk26jy/how_to_use_reduxsaga_with_graphql/
---
&amp;#x200B;

Hey guys,

I'm a bit confused about where should I have to call graphql query in react component or redux action?

I wanna use the best way.
## [9][I need to write an async action creator to set and hide (dispatch) notification messages with setTimeout. I don't understand how to do this.](https://www.reddit.com/r/reduxjs/comments/hjj8ga/i_need_to_write_an_async_action_creator_to_set/)
- url: https://www.reddit.com/r/reduxjs/comments/hjj8ga/i_need_to_write_an_async_action_creator_to_set/
---
I have already refactored my axios requests to fit this pattern, like so:


*Edit: Sorry about the code formatting, I tried using backticks but it doesn't work! 



`export const addAnecdote = (payload) =&gt; {`

`return async (dispatch) =&gt; {`

`const newAnecdote = await anecdoteService.createNew(payload)`

`dispatch({ type: 'ADD_ANECDOTE', payload: newAnecdote, })}}`


Then, in the component:

 `const newAnecdote = (e) =&gt; {`

    `e.preventDefault()`

    `const content = e.target.item.value`

    `e.target.item.value = ''`

    `dispatch(addAnecdote(content))`

`}`

So, anything that requires sending a request to the axios service, (with just a local json database) is written this way. 

HOWEVER, now they want me to write the notifications the same way, and I found [this long, comprehensive explanation by Dan Abramov](https://stackoverflow.com/questions/35411423/how-to-dispatch-a-redux-action-with-a-timeout?rq=1) on stackoverflow, but even with that, I cannot get it to work in my app. 

I'm not using 'connect', I'm using 'useDispatch' and 'useSelector' with React functional components. 

At the very least, can someone show me what the reducer and component would look like for this action: 

`store.dispatch({ type: 'SHOW_NOTIFICATION', text: 'You logged in.' })`


`setTimeout(() =&gt; {`


  `store.dispatch({ type: 'HIDE_NOTIFICATION' })`


`}, 5000)`


because I could NOT get it to work in my app. 

I did get a sorta hacky version to work, I'll go ahead and share it with you now so you can have a good laugh: 

`export const displayNotification = (content, duration) =&gt; {`


  `return async (dispatch) =&gt; {`


   `const message = await content`


    `const timeout = duration * 1000`


    `dispatch({`


      `type: 'SET_NOTIFICATION',`


      `payload: message,`


    `})`


    `setTimeout(() =&gt; {`


      `dispatch({`


        `type: 'SET_NOTIFICATION',`


        `payload: '',`


      `})`


    `}, timeout)`


  `}`


`}`

It ***technically*** fulfills the exercise requirements, which call for using the following format to call the action: 
```
  dispatch(displayNotification(content, 3))
```

But I have no idea what good it does to use async/await for the message content! It's dumb, I know! 

Any help is appreciated, even if it is just to help me find some remedial tutorials or examples to help me understand how redux/redux-thunk uses async/await. Without any extra packages, please!
## [10][Opt Redux to Manage the Application States](https://www.reddit.com/r/reduxjs/comments/hjqklc/opt_redux_to_manage_the_application_states/)
- url: https://www.ucodice.com/technology/hire-redux-developer-team
---

