# reduxjs
## [1][Need help please - how to access data from Object ID reference in state (React Native, Mongo, Redux)](https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/)
- url: https://www.reddit.com/r/reduxjs/comments/iah6yc/need_help_please_how_to_access_data_from_object/
---
I'm using React Native, Mongo and Redux

I have a data model Rounds, Courses and Users.

The Rounds model references Users storing the object ID of each "player" in an array. It also references the Course model to attach a single ID.

`players: [{type: mongoose.Schema.Types.ObjectId,ref: 'User',}],`

I have the players object currently loading into the state. What I'm trying to do is load the details from the Object ID that is referenced, not just the ID. How do I display firstName, lastName etc. from the player ID reference to the User Model and the Course Name rather than the ID from the Course Model?

Any help would be appreciated, I'm stuck and having trouble figuring this part out. Thank you!!!
## [2][My code doesn't function properly when I try to create user](https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/)
- url: https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/
---
https://github.com/calebdockal/CognitoProject

Does anyone know their way around redux to tell me what's going on?
## [3][a better modulized redux solution: module-reaction](https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/)
- url: https://www.reddit.com/r/reduxjs/comments/i41wo6/a_better_modulized_redux_solution_modulereaction/
---
a better redux-based framework which let u manage store modulized   
[https://github.com/swellee/reaction](https://github.com/swellee/reaction)  
or see npmjs: [https://www.npmjs.com/package/module-reaction](https://www.npmjs.com/package/module-reaction)
## [4][A better organizational pattern than /dispatchers, /reducers](https://www.reddit.com/r/reduxjs/comments/i3fadf/a_better_organizational_pattern_than_dispatchers/)
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
## [5][A highly scalable, performance focused React starter template, that focuses on best practices and a great developer experience.](https://www.reddit.com/r/reduxjs/comments/i29n1h/a_highly_scalable_performance_focused_react/)
- url: https://github.com/react-boilerplate/react-boilerplate-cra-template/
---

## [6][How can I separate actionCreator files and import them into the main index.js?](https://www.reddit.com/r/reduxjs/comments/i28bhv/how_can_i_separate_actioncreator_files_and_import/)
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
## [7][Top 8 Commandments for building apps with Redux](https://www.reddit.com/r/reduxjs/comments/hyqegy/top_8_commandments_for_building_apps_with_redux/)
- url: https://blog.logrocket.com/8-definitive-rules-building-apps-redux/
---

## [8][How I Made the Django React and Redux Blog](https://www.reddit.com/r/reduxjs/comments/hyba94/how_i_made_the_django_react_and_redux_blog/)
- url: https://www.codeingschool.com/2020/07/how-i-made-django-react-blog.html
---

## [9][I've been using just one saga file and it is getting nasty, should I separate them into different saga?](https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/)
- url: https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/
---
I've seen a lot of reducers have been separately saved into multiple files, but haven't seen many sagas like that? 

is it alright to separate them to clean up some codes?
## [10][redux-toolkit unit testing strategy?](https://www.reddit.com/r/reduxjs/comments/hvwqc9/reduxtoolkit_unit_testing_strategy/)
- url: https://www.reddit.com/r/reduxjs/comments/hvwqc9/reduxtoolkit_unit_testing_strategy/
---
Hi All,

I am using redux-toolkit for the 1st time, got a solid understanding of the concepts using docs. I have previously written unit tests for traditional redux - actions, reducers.

&amp;#x200B;

Wondering what is the "official" strategy for writing tests for slices which have

1. standard reducers key.
2. includes asyncThunks with extra-reducers.

&amp;#x200B;

update: 

using testing-library.

i checked Mark's reply here on slice reducer [https://stackoverflow.com/a/61921088](https://stackoverflow.com/a/61921088)

but **still need some approach for asyncThunks with extra-reducers.**

&amp;#x200B;

&amp;#x200B;

Regards.
