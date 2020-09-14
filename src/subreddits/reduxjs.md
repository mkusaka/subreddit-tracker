# reduxjs
## [1][How can I run an API call AFTER the results from my useSelector() hook?](https://www.reddit.com/r/reduxjs/comments/iry4g4/how_can_i_run_an_api_call_after_the_results_from/)
- url: https://www.reddit.com/r/reduxjs/comments/iry4g4/how_can_i_run_an_api_call_after_the_results_from/
---
I am trying to use some data that will come from my `useSelector()` hook in a network API call. However, with the code below, I get the error `TypeError: cannot read property 'query' of undefined.`

I understand this is because the query has not come back from the `useSelector()` yet. Is there a way I can wait for that data THEN call the API?

&amp;#x200B;

`const pageContent = useSelector(getPageContent);`  
  `useEffect(() =&gt; {`  
 `axios`  
 `.get('https://www.googleapis.com/youtube/v3/search', {`  
 `params: {`  
 `key: process.env.API_KEY,`  
 `part: 'snippet',`  
 `type: 'video',`  
 `q: pageContent.data.query,`  
 `},`  
 `})`  
 `.then((res) =&gt; res)`  
 `.then((data) =&gt; console.log(data));`  
 `}, []);`
## [2][Using Redux with very large flat arrays](https://www.reddit.com/r/reduxjs/comments/irzdsw/using_redux_with_very_large_flat_arrays/)
- url: https://www.reddit.com/r/reduxjs/comments/irzdsw/using_redux_with_very_large_flat_arrays/
---
I am developing an app for recording and editing paths on a map. These paths can be large (up to 20000 {x, y} points), and they can be edited by dragging these points and applying automatic pathfinding. Custom info can be attached to any point, and the sidebar computes some statistics on the fly. While I don't need to show all the points simultaneously, the user should be able to view and edit at least 1000-2000 at the same time.

My code for working with just one path currently looks like this (greatly simplified):

    function pathDataReducer(pathData, action) {
        switch(action.type) {
            case 'ADD_POINT':
            case' EDIT_POINT':
                return {...pathData, [action.payload.id]: action.payload};
            case 'DELETE_POINT':
                return deleteKeyImmutably(pathData, action.payload.id);            
        }
        return pathData;
    }

    function pathPointIds(ids, action) {
        switch(action.type) {
            case 'ADD_POINT': return [...ids, action.payload.id];
            case 'DELETE_POINT': return ids.filter(id =&gt; id !== action.payload.id);
        }
        return ids;
    }

    const getPoints = createSelector(
        [state =&gt; state.pathPointIds, state =&gt; state.pathData],
        (ids, data) =&gt; ids.map(id =&gt; data[id])
    );

    funtion Point({x, y}) {
        return &lt;Circle x={x} y={y} /&gt;;
    }

    function Path() {
        const points = useSelector(getPoints);
        return points.map(point =&gt; &lt;Point key={point.id} x={point.x} y={point.y} /&gt;);
    }

And the app is pretty much unusable. Adding or moving any point (via EDIT_POINT action) causes the whole `Path` to be re-rendered. This, in turn, causes 2000 virtualDOM comparisons (or an equivalent amount of prop checks, when using `React.memo`), which causes a 0.5-1 second lag. It seems impossible to achieve a smooth dragging experience (moving just the DOM node directly is not enough, because the point may also affect its neighbors - so the dragging should affect the data in Redux store).

I tried using MobX as an alternative, which gave acceptable performance out of the box, but it has another problem. MobX expects a deeply nested object graph (as opposed to normalized Redux store), and discourages referencing object fields directly. This is an issue for me, because I need to compute a lot of different statistics for paths, and working with denormalized data is a hassle.

I fell like an ideal solution would be a normalized Redux-like store, but without immutability, so that I could reasonably mutate just one object in a huge array and automatically subscribe to its updates. But I don't know of a popular battle-tested library that achieves this.

Is there a way to achieve acceptable performance using Redux in my case? Or is there an alternative state management solution suitable for working with large datasets?

Any help is appreciated!
## [3][Using an action type as a value in state](https://www.reddit.com/r/reduxjs/comments/irkgi9/using_an_action_type_as_a_value_in_state/)
- url: https://www.reddit.com/r/reduxjs/comments/irkgi9/using_an_action_type_as_a_value_in_state/
---
I'm putting together a video playlist and 3 of the actions involve incorporating *repeat* functionality, which can be found on most music and video players. For those not familiar, pressing the button once will repeat an album infinitely, pressing it a second time will repeat the currently playing track infinitely, and a third press will turn off repeating entirely.

Since this can be 3 values I opted not to use a boolean. Setting this as a numeric value just didn't seem implicit either so I thought of using a string value. Is it common to use action-types outside of actions?

**action-types/playlist:**

    export const REPEAT_ALL = 'PLAYLISTS.REPEAT_ALL';
    export const REPEAT_ONE = 'PLAYLISTS.REPEAT_ONE';
    export const REPEAT_OFF = 'PLAYLISTS.REPEAT_OFF';

**reducers/playlist:**

    const initialState = {
      ...
      settings: {
        ...
        repeat: types.REPEAT_OFF,
      },
    };
    
    export default (state = initialState, action) =&gt; {
      switch (action.type) {
        case types.REPEAT_ONE:
          return {
            ...state,
            settings: {
              ...state.settings,
              repeat: types.REPEAT_OFF,
            }
          };
       case types.REPEAT_ALL:
         return {
           ...state,
           settings: {
             ...state.settings,
             repeat: types.REPEAT_ALL
           };
      case types.REPEAT_OFF:
        return {
          ...state,
          settings: {
            ...state.settings,
            repeatTrack: false,
            repeatTracks: false,
          }
        };
      }
    };

&amp;#x200B;
## [4][Redux Saga: How to use eventChannel to listen to Auth.currentAuthenticatedUser()](https://www.reddit.com/r/reduxjs/comments/iqloll/redux_saga_how_to_use_eventchannel_to_listen_to/)
- url: https://www.reddit.com/r/reduxjs/comments/iqloll/redux_saga_how_to_use_eventchannel_to_listen_to/
---
**What I'm Trying To Do:**

I'm trying to implement an auth listener using redux saga using an [Redux Saga eventChannel](https://redux-saga.js.org/docs/advanced/Channels.html). Here is the AWS Cognito [currentAuthenticatedUser](https://aws-amplify.github.io/amplify-js/api/classes/authclass.html#currentauthenticateduser) method. How can I implement an auth channel using these?

&amp;#x200B;

**My Research/What I've Found:**

I've found a Stackoverflow post where they were able to implement an eventChannel with Firebase's `onAuthStateChanged`, but I can't seem to figure out how to get it working with AWS's `Auth.currentAuthenticatedUser();`

    function getAuthChannel() {
      if (!this.authChannel) {
        this.authChannel = eventChannel(emit =&gt; {
          const unsubscribe = firebase.auth().onAuthStateChanged(user =&gt; emit({ user }));
          return unsubscribe;
        });
      }
      return this.authChannel;
    }
    
    function* watchForFirebaseAuth() {
      ...
      // This is where you wait for a callback from firebase
      const channel = yield call(getAuthChannel);
      const result = yield take(channel);
      // result is what you pass to the emit function. In this case, it's an object like { user: { name: 'xyz' } }
      ...
    }
## [5][How to implement Redux Saga with ReactJS and Redux?](https://www.reddit.com/r/reduxjs/comments/iouo4e/how_to_implement_redux_saga_with_reactjs_and_redux/)
- url: https://medium.com/fabisiakradoslaw/how-to-implement-redux-saga-with-reactjs-and-redux-d6116efcc90?source=friends_link&amp;sk=1a4f940ea36c4e6ae12770040f1d10bc
---

## [6][The “Container” Pattern for Better State Management in React.](https://www.reddit.com/r/reduxjs/comments/io1fg3/the_container_pattern_for_better_state_management/)
- url: https://medium.com/@spencerpauly/the-container-pattern-for-better-state-management-in-react-9351fe4381d1
---

## [7][Try Mobx, I'm sure you will love it.](https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/)
- url: https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/
---
Hey Guys!

Redux is great, it helped react community to form the way we dispatch actions ;)  
Looking back 3 years ago one's we tried mobx it was a no brainer to switch make the switch.  
I know some times is hard to see the value hope this course will show you the light.   
feel free to enroll and share with your friends, free for the next 3 days.  


[https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507](https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507)
## [8][Redux or Mobx for larger applications](https://www.reddit.com/r/reduxjs/comments/il6ix6/redux_or_mobx_for_larger_applications/)
- url: https://www.reddit.com/r/reduxjs/comments/il6ix6/redux_or_mobx_for_larger_applications/
---
Hello all,  
   I'm wondering if anyone has experience of using Redux and/or Mobx for larger applications and what their experiences are.  
Currently I have some experience with Redux but it seems like Mobx could be a decent alternative.   
With Redux there is a lot of boilerplate and also a steep learning curve for newer developers.   
Mobx looks to have less boiler plate and will take less time to develop in.  
It's also more OO instead of functional so would be easier for new JS developers to pick up.  
I'm wondering if anyone has any experience of using either in larger applications and what the advantages and pitfalls are?  
Any input would be appreciated.
## [9][Libraries for highly reusable components?](https://www.reddit.com/r/reduxjs/comments/ijt44r/libraries_for_highly_reusable_components/)
- url: /r/reactjs/comments/ijsvwg/libraries_for_highly_reusable_components/
---

## [10][Handling large Json data ?](https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/)
- url: https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/
---
Greetings,   
So I'm implementing a Json tree view for a technical test.  
My redux state has the json in it but I was wondering about optimizing the rendering.

Probably it's not a very redux specific question.  
But I'm thinking of using redux-select as a way to optimize the render but I dont know how exactly  
Im also thinking about rendering the first 20 children of the json with a load more button   


It's more a question on how would you approach optimizing rendering for large data and I'd appreciate the ideas
