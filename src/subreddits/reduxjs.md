# reduxjs
## [1][I need to write an async action creator to set and hide (dispatch) notification messages with setTimeout. I don't understand how to do this.](https://www.reddit.com/r/reduxjs/comments/hjj8ga/i_need_to_write_an_async_action_creator_to_set/)
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
## [2][Opt Redux to Manage the Application States](https://www.reddit.com/r/reduxjs/comments/hjqklc/opt_redux_to_manage_the_application_states/)
- url: https://www.ucodice.com/technology/hire-redux-developer-team
---

## [3][Is this right? Whatsapp Analogy for Redux](https://www.reddit.com/r/reduxjs/comments/hj5ys5/is_this_right_whatsapp_analogy_for_redux/)
- url: https://www.reddit.com/r/reduxjs/comments/hj5ys5/is_this_right_whatsapp_analogy_for_redux/
---
Hey I am trying to learn redux and have written this up. 

Does this make sense? Is it right? What should I change?

[https://medium.com/@acgoff/what-redux-do-do-de4525c6f5d7](https://medium.com/@acgoff/what-redux-do-do-de4525c6f5d7)
## [4][A Complete reference guide to Redux: State of the art state management](https://www.reddit.com/r/reduxjs/comments/hhvew0/a_complete_reference_guide_to_redux_state_of_the/)
- url: https://blog.soshace.com/a-complete-reference-guide-to-redux-state-of-the-art-state-management/
---

## [5][Redux-toolkit + redux-orm](https://www.reddit.com/r/reduxjs/comments/hhhvqz/reduxtoolkit_reduxorm/)
- url: https://www.reddit.com/r/reduxjs/comments/hhhvqz/reduxtoolkit_reduxorm/
---
Hi, I'm trying to use both redux-toolkit and redux-orm in my project, and having troubles marrying the two. If anyone has links to any projects that successfully achieve that, please share!
## [6][useSelector vs Connect (react-redux)](https://www.reddit.com/r/reduxjs/comments/hgi2qi/useselector_vs_connect_reactredux/)
- url: https://www.samdawson.dev/article/react-redux-use-selector-vs-connect
---

## [7][Introduction React Native, Typescript, Redux, Rxjs, Node.js, Mongo, Mongoose. Nexus Aurora Platform](https://www.reddit.com/r/reduxjs/comments/hg5q9y/introduction_react_native_typescript_redux_rxjs/)
- url: https://youtu.be/06wsg2f76hQ
---

## [8][Modify-via-query (alternatives to immutability-helper)](https://www.reddit.com/r/reduxjs/comments/hg54tk/modifyviaquery_alternatives_to_immutabilityhelper/)
- url: https://github.com/wongjiahau/modify-via-query
---

## [9][Should I maintain state in components?](https://www.reddit.com/r/reduxjs/comments/hg1w01/should_i_maintain_state_in_components/)
- url: https://www.reddit.com/r/reduxjs/comments/hg1w01/should_i_maintain_state_in_components/
---
I have a form that has 10 inputs. Is it advised to maintain this state in the component or just pass props and update through redux? What is best practice.
## [10][Create simple POS with React.js, Node.js, and MongoDB #7: Adding redux to other component](https://www.reddit.com/r/reduxjs/comments/hebhol/create_simple_pos_with_reactjs_nodejs_and_mongodb/)
- url: https://blog.soshace.com/create-simple-pos-with-react-js-node-js-and-mongodb-7-adding-redux-to-other-component/
---

