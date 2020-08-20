# reduxjs
## [1][What is the difference between React-Redux and Redux Toolkit?](https://www.reddit.com/r/reduxjs/comments/icxp0b/what_is_the_difference_between_reactredux_and/)
- url: https://www.reddit.com/r/reduxjs/comments/icxp0b/what_is_the_difference_between_reactredux_and/
---
Hi, so I was learning Redux from this youtube [playlist](https://www.youtube.com/playlist?list=PLC3y8-rFHvwheJHvseC3I0HuYI2f46oAK), which is about a year old. In the aforementioned series, the instructor uses [React Redux](https://react-redux.js.org/) to create the store and use Redux in a React app. 

I was then going through the official [tutorial](https://redux.js.org/tutorials/essentials/part-1-overview-concepts) and I found out about [Redux Toolkit](https://redux-toolkit.js.org/), which they say is "is our recommended approach for writing Redux logic". 

So I am a bit confused now about the difference between them and which library should I use now in 2020.
## [2][Understanding Redux Epics and Rxjs](https://www.reddit.com/r/reduxjs/comments/ic9xyg/understanding_redux_epics_and_rxjs/)
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
## [3][check out this drag-and-drop kanban app I made with Redux (demo and live link in the repo)](https://www.reddit.com/r/reduxjs/comments/ibpc7a/check_out_this_draganddrop_kanban_app_i_made_with/)
- url: https://github.com/brietsparks/kanban-dashboard
---

## [4][Need help please - how to access data from Object ID reference in state (React Native, Mongo, Redux)](https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/)
- url: https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/
---
I'm using React Native, Mongo and Redux

I have a data model Rounds, Courses and Users.

The Rounds model references Users storing the object ID of each "player" in an array. It also references the Course model to attach a single ID.

`players: [{type: mongoose.Schema.Types.ObjectId,ref: 'User',}],`

I have the players object currently loading into the state. What I'm trying to do is load the details from the Object ID that is referenced, not just the ID. How do I display firstName, lastName etc. from the player ID reference to the User Model and the Course Name rather than the ID from the Course Model?

Any help would be appreciated, I'm stuck and having trouble figuring this part out. Thank you!!!
## [5][My code doesn't function properly when I try to create user](https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/)
- url: https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/
---
https://github.com/calebdockal/CognitoProject

Does anyone know their way around redux to tell me what's going on?
## [6][a better modulized redux solution: module-reaction](https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/)
- url: https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/
---
a better redux-based framework which let u manage store modulized   
[https://github.com/swellee/reaction](https://github.com/swellee/reaction)  
or see npmjs: [https://www.npmjs.com/package/module-reaction](https://www.npmjs.com/package/module-reaction)
## [7][A better organizational pattern than /dispatchers, /reducers](https://www.reddit.com/r/reduxjs/comments/i3fadf/a_better_organizational_pattern_than_dispatchers/)
- url: https://www.reddit.com/r/reduxjs/comments/i3fadf/a_better_organizational_pattern_than_dispatchers/
---
Starting with a small app:

    src
      nav
        nav.tsx
      contacts
        contacts.tsx
      settings
        settings.tsx
      App.tsx

Normally I would add `src/dispatchers` and `src/reducers`

But other devs are saying it's more ideal to split by functionality.

In that case where would you guys put your redux stuff? What I am theorizing is this:

    src
      shared
        dispatchers.tsx
        reducer.tsx
      nav
        nav.tsx   
      contacts
        contacts.tsx
        dispatchers.tsx
        reducer.tsx
      settings
        settings.tsx
        dispatchers.tsx
        reducer.tsx
      App.tsx
## [8][A highly scalable, performance focused React starter template, that focuses on best practices and a great developer experience.](https://www.reddit.com/r/reduxjs/comments/i29n1h/a_highly_scalable_performance_focused_react/)
- url: https://github.com/react-boilerplate/react-boilerplate-cra-template/
---

## [9][How can I separate actionCreator files and import them into the main index.js?](https://www.reddit.com/r/reduxjs/comments/i28bhv/how_can_i_separate_actioncreator_files_and_import/)
- url: https://www.reddit.com/r/reduxjs/comments/i28bhv/how_can_i_separate_actioncreator_files_and_import/
---
`import * as actionTypes from './actionTypes';`  
`export const addCategory = (name) =&gt; ({`  
 `type: actionTypes.REQUEST_ADD_CATEGORY,`  
 `name,`  
`});`  
`export const addCategorySuccess = (category) =&gt; ({`  
 `type: actionTypes.REQUEST_ADD_CATEGORY_SUCCESS,`  
 `category,`  
`});`  
`export const getCategories = () =&gt; ({`  
 `type: actionTypes.REQUEST_CATEGORIES,`  
`});`

I have this actionCreator file and I want to separate all three into three different creators. How can I import them to my actions/index.js?
## [10][Top 8 Commandments for building apps with Redux](https://www.reddit.com/r/reduxjs/comments/hyqegy/top_8_commandments_for_building_apps_with_redux/)
- url: https://blog.logrocket.com/8-definitive-rules-building-apps-redux/
---

