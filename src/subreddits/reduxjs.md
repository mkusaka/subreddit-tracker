# reduxjs
## [1][Redux Saga: How to use eventChannel to listen to Auth.currentAuthenticatedUser()](https://www.reddit.com/r/reduxjs/comments/iqloll/redux_saga_how_to_use_eventchannel_to_listen_to/)
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
## [2][How to implement Redux Saga with ReactJS and Redux?](https://www.reddit.com/r/reduxjs/comments/iouo4e/how_to_implement_redux_saga_with_reactjs_and_redux/)
- url: https://medium.com/fabisiakradoslaw/how-to-implement-redux-saga-with-reactjs-and-redux-d6116efcc90?source=friends_link&amp;sk=1a4f940ea36c4e6ae12770040f1d10bc
---

## [3][The “Container” Pattern for Better State Management in React.](https://www.reddit.com/r/reduxjs/comments/io1fg3/the_container_pattern_for_better_state_management/)
- url: https://medium.com/@spencerpauly/the-container-pattern-for-better-state-management-in-react-9351fe4381d1
---

## [4][Try Mobx, I'm sure you will love it.](https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/)
- url: https://www.reddit.com/r/reduxjs/comments/imz2xj/try_mobx_im_sure_you_will_love_it/
---
Hey Guys!

Redux is great, it helped react community to form the way we dispatch actions ;)  
Looking back 3 years ago one's we tried mobx it was a no brainer to switch make the switch.  
I know some times is hard to see the value hope this course will show you the light.   
feel free to enroll and share with your friends, free for the next 3 days.  


[https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507](https://www.udemy.com/course/mobx-in-depth-with-react/?couponCode=68D474E01D1CECEA3507)
## [5][Redux or Mobx for larger applications](https://www.reddit.com/r/reduxjs/comments/il6ix6/redux_or_mobx_for_larger_applications/)
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
## [6][Libraries for highly reusable components?](https://www.reddit.com/r/reduxjs/comments/ijt44r/libraries_for_highly_reusable_components/)
- url: /r/reactjs/comments/ijsvwg/libraries_for_highly_reusable_components/
---

## [7][Handling large Json data ?](https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/)
- url: https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/
---
Greetings,   
So I'm implementing a Json tree view for a technical test.  
My redux state has the json in it but I was wondering about optimizing the rendering.

Probably it's not a very redux specific question.  
But I'm thinking of using redux-select as a way to optimize the render but I dont know how exactly  
Im also thinking about rendering the first 20 children of the json with a load more button   


It's more a question on how would you approach optimizing rendering for large data and I'd appreciate the ideas
## [8][AWS Amplify + Redux: How can I use currentAuthenticatedUser() with Redux Saga?](https://www.reddit.com/r/reduxjs/comments/ie3grt/aws_amplify_redux_how_can_i_use/)
- url: https://www.reddit.com/r/reduxjs/comments/ie3grt/aws_amplify_redux_how_can_i_use/
---
How can I use \`currentAuthenticatedUser()\` with Redux Saga? I want to use to keep track if a user is logged in or not.

&amp;#x200B;

Ideally, I want it running as a channel so I know immediately log a user out when the channel returns false/the user isn't authenticated. I haven't been able to find much on keeping track if a user is logged in or out using AWS Amplify and Redux Saga. I hold the state of the user in redux and clear it when a user is logs out.

&amp;#x200B;

**Links:**

[currentAuthenticatedUser](https://aws-amplify.github.io/amplify-js/api/classes/authclass.html#currentauthenticateduser)

[Redux Saga Event Channel](https://redux-saga.js.org/docs/advanced/Channels.html)
## [9][Writing integration tests by exporting data from redux dev tool?](https://www.reddit.com/r/reduxjs/comments/ie4brv/writing_integration_tests_by_exporting_data_from/)
- url: https://www.reddit.com/r/reduxjs/comments/ie4brv/writing_integration_tests_by_exporting_data_from/
---
I'm looking to write integration tests for a project by dispatching a series of actions that are connected in some way and verifying end state is correct.  Ideal scenario would be to use the redux dev tool and export actions and initial state from there, transform that data to remove excess with a small script and into correct format, then manually add asserting logic. That would be a single integration test. In my head, this seems super quick to write tests and with as close to production data payloads for actions as possible.

To accomplish this, I'm thinking each test would be a json file containing an array of actions, along with an object or serialized function if possible to verify state.

Is there a project out there that already does this or something similar before I re-invent the wheel? If not, does anyone know why something like hasn't been implemented yet and what may be some challenges that might become blockers?
## [10][Creating fake store with initial state using RTK](https://www.reddit.com/r/reduxjs/comments/idxj5t/creating_fake_store_with_initial_state_using_rtk/)
- url: https://www.reddit.com/r/reduxjs/comments/idxj5t/creating_fake_store_with_initial_state_using_rtk/
---
Hi,

Earlier for Jest unit testing containers we use to create fake store with “redux” createStore(rootReducer, initialState)

Because I want to test my container after some initialState set up.

How do I create a fake store in RTK with initialState???
configureStore doesn’t seem to have a initialState option?

Thanks!
