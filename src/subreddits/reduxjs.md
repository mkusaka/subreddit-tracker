# reduxjs
## [1][My code doesn't function properly when I try to create user](https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/)
- url: https://www.reddit.com/r/reduxjs/comments/i85sts/my_code_doesnt_function_properly_when_i_try_to/
---
https://github.com/calebdockal/CognitoProject

Does anyone know their way around redux to tell me what's going on?
## [2][Build a Shopping Cart with React, Redux, and React-DnD](https://www.reddit.com/r/reduxjs/comments/i6stix/build_a_shopping_cart_with_react_redux_and/)
- url: https://www.reddit.com/r/reduxjs/comments/i6stix/build_a_shopping_cart_with_react_redux_and/
---
Let's build a simple shopping cart that takes advantage of React DnD which is a set of React utilities to help build complex drag and drop interfaces while keeping your components decoupled.

[Part 1](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-1-2433558c3f38?source=friends_link&amp;sk=9e2e7192050b93a5ece2b204a6c36550)

[Part 2](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-2-b4cd649e25db?source=friends_link&amp;sk=29b34207bc446ff51c420f34575d968e)

[Part 3](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-3-f1e1e8265d14?source=friends_link&amp;sk=83a5cb1a1b3df19d0db6d48dd4fce14c)

[Part 4](https://itnext.io/build-a-shopping-cart-with-react-redux-and-react-dnd-part-4-6a77f0a39c0c?source=friends_link&amp;sk=f8871c5bb5899dae4347142f5cdaf931)
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
