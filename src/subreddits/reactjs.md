# reactjs
## [1][Beginner's Thread / Easy Questions (October 2020)](https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/)
- url: https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Ask about React or anything else in its ecosystem :)

Stuck making progress on your app?  
Still Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** of reply by
   1. adding minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz] links
   1. describing what you want it to do (ask yourself if it's an [XY problem](https://meta.stackexchange.com/questions/66377/what-is-the-xy-problem))
   1. things you've tried. (Don't just post big blocks of code!)
1. **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
1. **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**! 👉

🆓 Here are great, **free** resources!

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- New to Hooks? Check out [Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- and these React Hook recipes on [useHooks.com][usehooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[usehooks.com]: https://usehooks.com/
[thinking in react hooks]: https://wattenberger.com/blog/react-hooks
[freecodecamp's react course]: https://www.freecodecamp.org/news/learn-react-course/
[microsoft frontend bootcamp]: https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[kent dodd's egghead.io course]: http://kcd.im/beginner-react
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[being wrong on the internet]: https://xkcd.com/386/
[tweet organization]: https://twitter.com/dan_abramov/status/1027245759232651270?lang=en
[get started with redux]: https://www.reddit.com/r/reactjs/wiki/index#wiki_getting_started_with_redux
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Help the React team write new docs—take the 2020 React Community Survey](https://www.reddit.com/r/reactjs/comments/j5iqj3/help_the_react_team_write_new_docstake_the_2020/)
- url: https://www.surveymonkey.co.uk/r/T58DPNS
---

## [3][Microsoft plans to unify Outlook across platforms using web technologies [React / React Native]](https://www.reddit.com/r/reactjs/comments/j7jwis/microsoft_plans_to_unify_outlook_across_platforms/)
- url: https://www.neowin.net/news/microsoft-plans-to-unify-outlook-across-platforms-using-web-technologies?utm_source=feedburner&amp;utm_medium=feed&amp;utm_campaign=Feed%3A+neowin-main+%28Neowin+News%29
---

## [4][Jotai - a simple State Manager for React.js](https://www.reddit.com/r/reactjs/comments/j7wh5o/jotai_a_simple_state_manager_for_reactjs/)
- url: https://medium.com/javascript-in-plain-english/jotai-simple-state-management-for-react-b9318b0f7616?source=friends_link&amp;sk=57d2f0357a97a77ffa77f29132b876e2
---

## [5][Creating a presigned url to authenticate users](https://www.reddit.com/r/reactjs/comments/j7ycx4/creating_a_presigned_url_to_authenticate_users/)
- url: https://www.reddit.com/r/reactjs/comments/j7ycx4/creating_a_presigned_url_to_authenticate_users/
---
I have a react app that has a traditional login page and flow. Once the user enters correct credentials, they are given a jwt which is stored in the browser local storage.

However, I'd like to grant access to certain users by creating a short-term pre-signed URL (same as S3 does). The URL would look like `myapp.com/resource?token=ey6347fdefd....`

I can then add a new route to my app `&lt;Route exact path="/resource" component={Resource}&gt;`. When anyone hits this path, I can define some logic in the `Resource` component to:

1. Extract `token` from query string.
2. Validate `token` and ensure it's not expired using `exp` from token.
3. If successful, render the UI else redirect to `/login`.

Just want to confirm if this is the right approach. 

Also, would there be any case if I redirect users to `/login` in step 3 above? Asking because I will be the person creating those pre-signed URLs so assuming it's always correct (?). If a third party created those pre-signed URLs, then I'd be more concerned about validating them.
## [6][Reset value of Material UI Checkbox components programmatically?](https://www.reddit.com/r/reactjs/comments/j7whwp/reset_value_of_material_ui_checkbox_components/)
- url: https://www.reddit.com/r/reactjs/comments/j7whwp/reset_value_of_material_ui_checkbox_components/
---
Hi Everyone,

I'm using Material UI Checkbox components to allow a user to select toppings for their pizza.  I followed the MUI docs when creating the component, ie the components have a 'checked' prop with a boolean value.  True = checked, false = unchecked.  I've tried to programmatically reset the values of each key/value pair to false onClick.  My function successfully resets the object to all keys having a false value, however, the UI does not update and the checkboxes remain checked.  I need the checkboxes to reset to reflect the state of the objects key/value pairs if, for example the user clicks to add a pizza to their cart and then wants to order another one, they should have a blank slate...the topping selections shouldn't remain checked. I've properly managed state using the useState hook, and the function that resets the object looks like this: 

*const clearSelection = () =&gt; {*  
 *setRegularChecked(regObj)*  
 *setPremiumChecked(premObj)*  
 *setSize(null)*  
 *setVariant(null)*  
 *}*

setRegularChecked receives the regObj which looks like {ham: false, pepperoni: false}..etc for all regular toppings

setPremiumChecked receives the premObj which looks like {prociutto: false, gorgonzola: false}..etc for all premium toppings

setSize resets the pizza size selection to null as expected

setVariant resets the pizza variant (sicilian, regular, etc) back to null as expected

Does anyone know why this is happening?  Is it a bug?  Know of a solution? If you feel up to it, here is the full component code

[https://github.com/Perrottarichard/pizzapizza-client/blob/master/src/components/Pizza.js](https://github.com/Perrottarichard/pizzapizza-client/blob/master/src/components/Pizza.js)
## [7][React 17 rc 3 is out - final call for bugs](https://www.reddit.com/r/reactjs/comments/j7kble/react_17_rc_3_is_out_final_call_for_bugs/)
- url: https://twitter.com/dan_abramov/status/1314293585739681792?s=20
---

## [8][Material UI checkbox false value not updating in Ui](https://www.reddit.com/r/reactjs/comments/j7xgzq/material_ui_checkbox_false_value_not_updating_in/)
- url: https://v.redd.it/ikkg59ks42s51
---

## [9][How to join a Slack Channel for React development?](https://www.reddit.com/r/reactjs/comments/j7xb08/how_to_join_a_slack_channel_for_react_development/)
- url: https://www.reddit.com/r/reactjs/comments/j7xb08/how_to_join_a_slack_channel_for_react_development/
---
I would like to join a Slack channel of React developers.  
I find [https://reactjs.org/community/support.html](https://reactjs.org/community/support.html) but none of the links lead to an opportunity to join a Slack channel.  
This is confusing to me.  
What am I missing?
## [10][React - Keydown event is taking 500ms and the same function via onClick take 50ms](https://www.reddit.com/r/reactjs/comments/j7wog3/react_keydown_event_is_taking_500ms_and_the_same/)
- url: https://www.reddit.com/r/reactjs/comments/j7wog3/react_keydown_event_is_taking_500ms_and_the_same/
---
I am trying to create a keyboard shortcut for my react-redux application. The keydown function takes 500 ms and if I do the same via the onClick function it takes 50ms. I have tried the same by creating a production version but no good results.

A hack I tested was to call an onClick event inside the keydown function which take lesser time than the original function. I also tried using Chrome Memory Profiler to see that's wrong. But no luck.

Can anyone suggest any solution to debug this problem?
## [11][Epic React: Javascript You Need To Know For React](https://www.reddit.com/r/reactjs/comments/j7afir/epic_react_javascript_you_need_to_know_for_react/)
- url: https://blog.bhanuteja.dev/epic-react-javascript-you-need-to-know-for-react
---

## [12][Build a Timeline Component With React and React-Chrono](https://www.reddit.com/r/reactjs/comments/j7vvpz/build_a_timeline_component_with_react_and/)
- url: https://medium.com/better-programming/build-a-timeline-component-with-react-and-react-chrono-fb1b962b020e
---

