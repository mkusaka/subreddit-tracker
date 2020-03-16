# reduxjs
## [1][React Context API with async hooks as an alternative to state management](https://www.reddit.com/r/reduxjs/comments/fjie4o/react_context_api_with_async_hooks_as_an/)
- url: https://medium.com//when-you-finally-decided-to-rid-yourself-of-redux-8fff0624d2fb?source=friends_link&amp;sk=812065aa580515002a12a051554404e2
---

## [2][reduxform question on "paths"](https://www.reddit.com/r/reduxjs/comments/fig13b/reduxform_question_on_paths/)
- url: https://www.reddit.com/r/reduxjs/comments/fig13b/reduxform_question_on_paths/
---
Hey all, wasn't sure where else to go so I decided to try my luck with Reddit. I am currently trying to make an application for my school assignment, which is where I have to meet the demands of a client in terms of writing a program. 

What I am currently trying to do is make a "glorified spreadsheet".

Anyway, I am trying to make a series of input fields and display said data. Where I am currently stuck is that I want to create a so-called path when 1 option is selected.

 &lt;select *name*="category" *defaultValue*={todo.category} *onChange*={handleChange}&gt;  
                    &lt;option *value*="option1"&gt;option 1&lt;/option&gt;  
                    &lt;option *value*="option2"&gt;option 2&lt;/option&gt;  
 &lt;/select&gt;

So as you can see above, in the input, after a few other things have been inputed by the client, I am trying to get a "path" (not sure what to call it), where selecting 1 of these options provides specific options.

i.e.: selecting option1 (fruit, for example), and i then can display another select of possible fruits. and if option2 is selected then i will do another drop down menu, but of veggies etc.

Hope I am making sense and hope to get some input on what to do. I tried if statements, but that didn't work. I am quite new to programming and the ultimate goal of this is to give my client and my school a fully functioning program (using react, redux, thunk) that fulfils both their criteria.

Sorry for being a dummy, I just got lost :((
## [3][Reducers : Difference between state and action](https://www.reddit.com/r/reduxjs/comments/fhklg8/reducers_difference_between_state_and_action/)
- url: https://www.reddit.com/r/reduxjs/comments/fhklg8/reducers_difference_between_state_and_action/
---
So i was learning my way through redux by reading a lot of projects. It accures to me that in some projects the use .state instead of .action. if somebody can explain me differences and other way to use reducers more efficiently thanks. Here is an example:

[https://github.com/omrihaviv/react-native-boilerplate/blob/master/js/reducers/auth.js](https://github.com/omrihaviv/react-native-boilerplate/blob/master/js/reducers/auth.js)
## [4][How Redux toolkit can reduce your setup of Redux in your React app](https://www.reddit.com/r/reduxjs/comments/fget4u/how_redux_toolkit_can_reduce_your_setup_of_redux/)
- url: https://medium.com/@gstvribs/how-redux-toolkit-can-reduce-your-setup-of-redux-in-your-react-app-d87baab59268
---

## [5][Two ways to skin a form. Design Pattern Discussion: How would you guys implement the following?](https://www.reddit.com/r/reduxjs/comments/fe1thu/two_ways_to_skin_a_form_design_pattern_discussion/)
- url: https://www.reddit.com/r/reduxjs/comments/fe1thu/two_ways_to_skin_a_form_design_pattern_discussion/
---
Let’s say I’m posting/putting an object on a form that is essentially an add/edit screen for "car".  When i submit, I obviously wanna know what happens so I can navigate away or show errors on the form if necessary.   I could do this one of two ways:

1. I could treat the form and action/saga as a closed environment that never goes to the store at all.  I simply dispatch an action POST_VERSION and the payload is the body AND callbacks for success/fail/statuschange. Then the saga calls those callbacks “oh i started…..oh i succeeded…..oh i got validation errors”.  These callbacks are handled on the front end and do what is appropriate.  The store neither knows nor cares that the form exists or is doing anything.  I use this pattern for really strict UI interaction sometimes.

2. I could make the store have a slice called “current_car_being_edited" or something, and it has the statuses and the errors and uses the classic redux reducers etc to notify the front end via state change.   I know this is the "redux way" but something about it feels....loose and incorrect.  Is my DB brain overthinking it?   Solution #1 is so snappy and elegant, maybe this case is just too simple to need it?  Would you use it anyways?
## [6][state is an object tree ? What is that supposed to mean ?](https://www.reddit.com/r/reduxjs/comments/fdvrge/state_is_an_object_tree_what_is_that_supposed_to/)
- url: https://www.reddit.com/r/reduxjs/comments/fdvrge/state_is_an_object_tree_what_is_that_supposed_to/
---
What does the tree supposed to mean ?
## [7][Don't Waste Your Ducking Time: An opinionated guide on how to test Redux ducks](https://www.reddit.com/r/reduxjs/comments/f9n9j2/dont_waste_your_ducking_time_an_opinionated_guide/)
- url: https://github.com/tophat/dont-waste-your-ducking-time
---

## [8][How to normalize data ?](https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/)
- url: https://www.reddit.com/r/reduxjs/comments/f6rj80/how_to_normalize_data/
---
Hi 👋!Can someone help me with normalizr?

How you actually organize your stored data and where you store meta info about domain entity (like loading status and etc) ?

[https://codesandbox.io/s/cool-kare-6z09b](https://codesandbox.io/s/cool-kare-6z09b)
## [9][Sending post data React](https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/)
- url: https://www.reddit.com/r/reduxjs/comments/f61s41/sending_post_data_react/
---
 I have a really noobish question I'm working with Redux for the first time and I'm able to get data from a get request but I was hoping if you guys would be able to send me any links or any pointers on how post-data would be sent to the background using Redux when I look at the tutorials some people use mapStateToProps others use proptypes and others just don't pass anything at all any help would be greatly appreciated


 also what I want to do is register a user and when the registration is successful or not I will use that status code to determine whether they should go to a new component. I figured I would be able to do this in the  action.JS file but I don't understand how I would be able to send my state data on the component to Redux and then give me back my I response
## [10][Implementing Undo-Redo with NgRx or Redux](https://www.reddit.com/r/reduxjs/comments/f3q1rp/implementing_undoredo_with_ngrx_or_redux/)
- url: https://nils-mehlhorn.de/posts/angular-undo-redo-ngrx-redux
---

