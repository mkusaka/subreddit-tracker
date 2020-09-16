# reduxjs
## [1][Advanced Express JS REST API [#1] Introduction | Building REST API Node JS | Full Course - Please subscribe](https://www.reddit.com/r/reduxjs/comments/itgx6u/advanced_express_js_rest_api_1_introduction/)
- url: https://m.youtube.com/watch?v=CLdkGgv9Miw&amp;list=PLs1waz0ZKTGO7agN0cntpe6ro6TIka0ow&amp;index=2&amp;t=0s
---

## [2][What's your approach to the size of redux state slices?](https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/)
- url: https://www.reddit.com/r/reduxjs/comments/iswhn1/whats_your_approach_to_the_size_of_redux_state/
---
I like to have my redux state slices small and focused on specific parts of the interface.

But at the same time, I am feeling a growing resistance from the library to splitting the state, because some slice reducers end up wanting information from other slices of the state in order to decide how to update their portion of the state.

I've examined the [docs](https://redux.js.org/recipes/structuring-reducers/beyond-combinereducers#sharing-data-between-slice-reducers) to see what they advise. The first suggestion (passing parent state as the third argument) [does not seem to be supported by the typescript definitions](https://github.com/reduxjs/redux/pull/3465). The second suggestion — just use thunks — is something I am currently doing, but this both feels like a hack (thunks were intended to handle async code; that's why they are called thunks, or long-running computations), and ruins most of the elegance of redux-toolkit.

What has your experience been? How large are your state slices? Do you often need to look up data from other slices? How do you manage this?
## [3][AWS Amplify + Redux Saga: Adding Amazon Cognito Attributes on Auth.signUp?](https://www.reddit.com/r/reduxjs/comments/it14y6/aws_amplify_redux_saga_adding_amazon_cognito/)
- url: https://www.reddit.com/r/reduxjs/comments/it14y6/aws_amplify_redux_saga_adding_amazon_cognito/
---
**My Goal:**

I'm trying to sign up a user for the first time after setting up the \`required attributes\` in the AWS Amplify/Cognito console. However, I'm getting an error (See Below) when trying to add attributes when using `Auth.signUp` when using Redux Saga. The docs state that it can be done (See Below), but I think I'm missing some sort of syntax with Redux Saga.

&amp;#x200B;

**What I've tried:**

I'm passing all of the attributes in `yield call([Auth, 'signUp'], email, password, email, phone_number, given_name, family_name, updated_at);`, but the error is still coming up.

Is this the correct way to pass the attributes in the saga or does it need to be passed in an attributes object??

&amp;#x200B;

**AWS Amplify Console:**

https://preview.redd.it/gs1yb9r8h8n51.png?width=1376&amp;format=png&amp;auto=webp&amp;s=bfb1d1b6a0dd874ddc8015a1fec95821b2f548c1

**Error:**

    Attributes did not conform to the schema:
    given_name: The attribute is required.
    family_name: The attribute is required.
    updated_at: The attribute is required.

&amp;#x200B;

**The AWS Amplify Docs:**

Link: [https://docs.amplify.aws/lib/auth/emailpassword/q/platform/js#sign-up](https://docs.amplify.aws/lib/auth/emailpassword/q/platform/js#sign-up)

**The docs state that it can be done such as the following**, but I can't get the saga to work by passing the each attribute or also trying to pass an attributes object with each attribute as the key/values.

    async function signUp() {
        try {
            const { user } = await Auth.signUp({
                username,
                password,
                attributes: {
                    email,          // optional
                    phone_number,   // optional - E.164 number convention
                    // other custom attributes 
                }
            });
            console.log(user);
        } catch (error) {
            console.log('error signing up:', error);
        }
    }

&amp;#x200B;

**Redux Saga (Sign Up):**

&amp;#x200B;

          try {
            // Credentials
            const { email, password, given_name, family_name } : Credentials = action.credentials;
            let { phone_number } : Credentials = action.credentials;
        
            // Updated At
            const updatedAt: Date = new Date();
        
            // Format Phone Number For AWS To E.164 Standard ('(949) 123-4567' To '+19491234567890')
            // Remove Characters From Phone Number
            phone_number = phone_number.replace(/\D/g,'');
        
            // Add +1 To Beginning To Phone Number
            phone_number = ('+1').concat(phone_number);
        
            // AWS: Sign Up
            yield call([Auth, 'signUp'], email, password, email, phone_number, given_name, family_name, updated_at);
        
            // Redux: Sign Up Success
            yield put(signUpSuccess());
        
            // React Navigation: Sign Up Confirm
            yield ReactNavigation.navigate('Sign Up Confirm');
          }
## [4][How can I run an API call AFTER the results from my useSelector() hook?](https://www.reddit.com/r/reduxjs/comments/iry4g4/how_can_i_run_an_api_call_after_the_results_from/)
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
## [5][Using Redux with very large flat arrays](https://www.reddit.com/r/reduxjs/comments/irzdsw/using_redux_with_very_large_flat_arrays/)
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
## [6][Using an action type as a value in state](https://www.reddit.com/r/reduxjs/comments/irkgi9/using_an_action_type_as_a_value_in_state/)
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
## [7][Redux Saga: How to use eventChannel to listen to Auth.currentAuthenticatedUser()](https://www.reddit.com/r/reduxjs/comments/iqloll/redux_saga_how_to_use_eventchannel_to_listen_to/)
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
## [8][How to implement Redux Saga with ReactJS and Redux?](https://www.reddit.com/r/reduxjs/comments/iouo4e/how_to_implement_redux_saga_with_reactjs_and_redux/)
- url: https://medium.com/fabisiakradoslaw/how-to-implement-redux-saga-with-reactjs-and-redux-d6116efcc90?source=friends_link&amp;sk=1a4f940ea36c4e6ae12770040f1d10bc
---

## [9][The “Container” Pattern for Better State Management in React.](https://www.reddit.com/r/reduxjs/comments/io1fg3/the_container_pattern_for_better_state_management/)
- url: https://medium.com/@spencerpauly/the-container-pattern-for-better-state-management-in-react-9351fe4381d1
---

## [10][Try Mobx, I'm sure you will love it.](https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/)
- url: https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/
---
Hey Guys!

Redux is great, it helped react community to form the way we dispatch actions ;)  
Looking back 3 years ago one's we tried mobx it was a no brainer to switch make the switch.  
I know some times is hard to see the value hope this course will show you the light.   
feel free to enroll and share with your friends, free for the next 3 days.  


[https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507](https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507)
