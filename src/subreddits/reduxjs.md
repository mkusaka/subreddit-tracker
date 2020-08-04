# reduxjs
## [1][A better organizational pattern than /dispatchers, /reducers](https://www.reddit.com/r/reduxjs/comments/i3fadf/a_better_organizational_pattern_than_dispatchers/)
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
## [2][A highly scalable, performance focused React starter template, that focuses on best practices and a great developer experience.](https://www.reddit.com/r/reduxjs/comments/i29n1h/a_highly_scalable_performance_focused_react/)
- url: https://github.com/react-boilerplate/react-boilerplate-cra-template/
---

## [3][How can I separate actionCreator files and import them into the main index.js?](https://www.reddit.com/r/reduxjs/comments/i28bhv/how_can_i_separate_actioncreator_files_and_import/)
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
## [4][Top 8 Commandments for building apps with Redux](https://www.reddit.com/r/reduxjs/comments/hyqegy/top_8_commandments_for_building_apps_with_redux/)
- url: https://blog.logrocket.com/8-definitive-rules-building-apps-redux/
---

## [5][How I Made the Django React and Redux Blog](https://www.reddit.com/r/reduxjs/comments/hyba94/how_i_made_the_django_react_and_redux_blog/)
- url: https://www.codeingschool.com/2020/07/how-i-made-django-react-blog.html
---

## [6][I've been using just one saga file and it is getting nasty, should I separate them into different saga?](https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/)
- url: https://www.reddit.com/r/reduxjs/comments/hwrtg8/ive_been_using_just_one_saga_file_and_it_is/
---
I've seen a lot of reducers have been separately saved into multiple files, but haven't seen many sagas like that? 

is it alright to separate them to clean up some codes?
## [7][redux-toolkit unit testing strategy?](https://www.reddit.com/r/reduxjs/comments/hvwqc9/reduxtoolkit_unit_testing_strategy/)
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
## [8][Do we need to install redeux as a dependency if we are installing redux-toolkit?](https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/)
- url: https://www.reddit.com/r/reduxjs/comments/huvj1t/do_we_need_to_install_redeux_as_a_dependency_if/
---
Hi All,

Was starting out on redux-toolkit. Thanks to the awesome docs, i was able to set up a sample app with redux-toolkit. but i was wondering do we need to install redux explicitly as a dependency if we are installing redux-toolkit??

&amp;#x200B;

Looking for a definitive answer from the community,  FYI, i have installed only "@reduxjs/toolkit" &amp; "react-redux" and the app is working fine. 

&amp;#x200B;

Regards.
## [9][Building Scalable Redux-First Apps](https://www.reddit.com/r/reduxjs/comments/huc8ok/building_scalable_reduxfirst_apps/)
- url: https://medium.com/@robbiehall26/building-scalable-redux-first-apps-5a8d09e9bd04?sk=23a705bcad8d07e47500bf382213619d
---

## [10][New "Redux Essentials" core docs tutorial is LIVE! Teaches how to use Redux the right way, using our latest recommended APIs and practices](https://www.reddit.com/r/reduxjs/comments/hr3yx1/new_redux_essentials_core_docs_tutorial_is_live/)
- url: https://redux.js.org/tutorials/essentials/part-1-overview-concepts
---

