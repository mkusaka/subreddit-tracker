# reduxjs
## [1][Handling large Json data ?](https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/)
- url: https://www.reddit.com/r/reduxjs/comments/ifoavq/handling_large_json_data/
---
Greetings,   
So I'm implementing a Json tree view for a technical test.  
My redux state has the json in it but I was wondering about optimizing the rendering.

Probably it's not a very redux specific question.  
But I'm thinking of using redux-select as a way to optimize the render but I dont know how exactly  
Im also thinking about rendering the first 20 children of the json with a load more button   


It's more a question on how would you approach optimizing rendering for large data and I'd appreciate the ideas
## [2][AWS Amplify + Redux: How can I use currentAuthenticatedUser() with Redux Saga?](https://www.reddit.com/r/reduxjs/comments/ie3grt/aws_amplify_redux_how_can_i_use/)
- url: https://www.reddit.com/r/reduxjs/comments/ie3grt/aws_amplify_redux_how_can_i_use/
---
How can I use \`currentAuthenticatedUser()\` with Redux Saga? I want to use to keep track if a user is logged in or not.

&amp;#x200B;

Ideally, I want it running as a channel so I know immediately log a user out when the channel returns false/the user isn't authenticated. I haven't been able to find much on keeping track if a user is logged in or out using AWS Amplify and Redux Saga. I hold the state of the user in redux and clear it when a user is logs out.

&amp;#x200B;

**Links:**

[currentAuthenticatedUser](https://aws-amplify.github.io/amplify-js/api/classes/authclass.html#currentauthenticateduser)

[Redux Saga Event Channel](https://redux-saga.js.org/docs/advanced/Channels.html)
## [3][Writing integration tests by exporting data from redux dev tool?](https://www.reddit.com/r/reduxjs/comments/ie4brv/writing_integration_tests_by_exporting_data_from/)
- url: https://www.reddit.com/r/reduxjs/comments/ie4brv/writing_integration_tests_by_exporting_data_from/
---
I'm looking to write integration tests for a project by dispatching a series of actions that are connected in some way and verifying end state is correct.  Ideal scenario would be to use the redux dev tool and export actions and initial state from there, transform that data to remove excess with a small script and into correct format, then manually add asserting logic. That would be a single integration test. In my head, this seems super quick to write tests and with as close to production data payloads for actions as possible.

To accomplish this, I'm thinking each test would be a json file containing an array of actions, along with an object or serialized function if possible to verify state.

Is there a project out there that already does this or something similar before I re-invent the wheel? If not, does anyone know why something like hasn't been implemented yet and what may be some challenges that might become blockers?
## [4][Creating fake store with initial state using RTK](https://www.reddit.com/r/reduxjs/comments/idxj5t/creating_fake_store_with_initial_state_using_rtk/)
- url: https://www.reddit.com/r/reduxjs/comments/idxj5t/creating_fake_store_with_initial_state_using_rtk/
---
Hi,

Earlier for Jest unit testing containers we use to create fake store with “redux” createStore(rootReducer, initialState)

Because I want to test my container after some initialState set up.

How do I create a fake store in RTK with initialState???
configureStore doesn’t seem to have a initialState option?

Thanks!
## [5][What is the difference between React-Redux and Redux Toolkit?](https://www.reddit.com/r/reduxjs/comments/icxp0b/what_is_the_difference_between_reactredux_and/)
- url: https://www.reddit.com/r/reduxjs/comments/icxp0b/what_is_the_difference_between_reactredux_and/
---
Hi, so I was learning Redux from this youtube [playlist](https://www.youtube.com/playlist?list=PLC3y8-rFHvwheJHvseC3I0HuYI2f46oAK), which is about a year old. In the aforementioned series, the instructor uses [React Redux](https://react-redux.js.org/) to create the store and use Redux in a React app. 

I was then going through the official [tutorial](https://redux.js.org/tutorials/essentials/part-1-overview-concepts) and I found out about [Redux Toolkit](https://redux-toolkit.js.org/), which they say is "is our recommended approach for writing Redux logic". 

So I am a bit confused now about the difference between them and which library should I use now in 2020.
## [6][Understanding Redux Epics and Rxjs](https://www.reddit.com/r/reduxjs/comments/ic9xyg/understanding_redux_epics_and_rxjs/)
- url: https://www.reddit.com/r/reduxjs/comments/ic9xyg/understanding_redux_epics_and_rxjs/
---
Hello all,  
   I was wondering if someone could help me understand this piece of code (from the offical redux docs)  


    const fetchUserEpic = action$ =&gt; action$.pipe(
      ofType(FETCH_USER),
      mergeMap(action =&gt;
        ajax.getJSON(`https://api.github.com/users/${action.payload}`).pipe(
          map(response =&gt; fetchUserFulfilled(response))
        )
      )
    );

I am aware what epics are (actions in, actions out etc) and understand that the actions will go through via action$.pipe and then you pick the one you want through ofType and they must return another action.   
However, I am having trouble understanding what happens after calling mergeMap.   
From what I understand (which maybe very wrong), mergeMap will flatten and merge the outer observable (in this case action$) with the inner observable (in this case the call to get the json). From the inner observable, we are piping map, which will take the data from the api call and use it to call the next action.   
I feel I am missing something here and not understand the observable flow here.
## [7][check out this drag-and-drop kanban app I made with Redux (demo and live link in the repo)](https://www.reddit.com/r/reduxjs/comments/ibpc7a/check_out_this_draganddrop_kanban_app_i_made_with/)
- url: https://github.com/brietsparks/kanban-dashboard
---

## [8][Need help please - how to access data from Object ID reference in state (React Native, Mongo, Redux)](https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/)
- url: https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/
---
I'm using React Native, Mongo and Redux

I have a data model Rounds, Courses and Users.

The Rounds model references Users storing the object ID of each "player" in an array. It also references the Course model to attach a single ID.

`players: [{type: mongoose.Schema.Types.ObjectId,ref: 'User',}],`

I have the players object currently loading into the state. What I'm trying to do is load the details from the Object ID that is referenced, not just the ID. How do I display firstName, lastName etc. from the player ID reference to the User Model and the Course Name rather than the ID from the Course Model?

Any help would be appreciated, I'm stuck and having trouble figuring this part out. Thank you!!!
## [9][My code doesn't function properly when I try to create user](https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/)
- url: https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/
---
https://github.com/calebdockal/CognitoProject

Does anyone know their way around redux to tell me what's going on?
## [10][a better modulized redux solution: module-reaction](https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/)
- url: https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/
---
a better redux-based framework which let u manage store modulized   
[https://github.com/swellee/reaction](https://github.com/swellee/reaction)  
or see npmjs: [https://www.npmjs.com/package/module-reaction](https://www.npmjs.com/package/module-reaction)
