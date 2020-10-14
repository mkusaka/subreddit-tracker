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

## [3][Smooth DOC - a premium Gatsby.js documentation theme is now open source](https://www.reddit.com/r/reactjs/comments/javt8b/smooth_doc_a_premium_gatsbyjs_documentation_theme/)
- url: https://github.com/gregberge/smooth-doc
---

## [4][Hey guys, I wanted to share my tutorial on how to make a Socket IO Chat App using ReactJS and NodeJS!](https://www.reddit.com/r/reactjs/comments/jaf269/hey_guys_i_wanted_to_share_my_tutorial_on_how_to/)
- url: https://www.youtube.com/watch?v=XS79S6nmaGM
---

## [5][Is it safe to mutate destructured state variable?](https://www.reddit.com/r/reactjs/comments/jazc53/is_it_safe_to_mutate_destructured_state_variable/)
- url: https://www.reddit.com/r/reactjs/comments/jazc53/is_it_safe_to_mutate_destructured_state_variable/
---
Golden rule of React is not to mutate a state variable directly.

However, if I destructure the state, can I then update this variable safely? For example;

    this.state = {
      arr: [{id: 1}]
    }
    
    fooBar = () =&gt; {
      const { arr } = this.state;
      arr.push({id: 2});
      this.setState({ arr });
    }

Or should I clone this array and work on it?

    fooBar = () =&gt; {
      const { arr } = this.state;
      const clonedArr = [...arr];
      clonedArr.push({id: 2});
      this.setState({ arr: clonedArr });
    }
## [6][chat application using react native](https://www.reddit.com/r/reactjs/comments/jawzuq/chat_application_using_react_native/)
- url: https://www.reddit.com/r/reactjs/comments/jawzuq/chat_application_using_react_native/
---
Hello everyone, I have an idea for my last year college project in react. I want to build a real time mobile chat application. Its about 2 years that I am learning front end development and all my focus was on Javascript and I think now I am in intermediate level. Is it really hard and complex to build such application using react?

and what other technologies should I use I searched and found that I can use firebase or sokcet..

what database and server side programming language should I use, can I use node.js for server side and mangodb for database?
## [7][Back button requires 4 clicks to work](https://www.reddit.com/r/reactjs/comments/jaw969/back_button_requires_4_clicks_to_work/)
- url: https://www.reddit.com/r/reactjs/comments/jaw969/back_button_requires_4_clicks_to_work/
---
I'm using React Router v6 and using the navigate API and passing minus 1 with the intention of returning to the previous page. [This page](https://pastebin.com/19CbQrY4) is landed upon when when an error occurs during an API request. When I'm on this page and try to go back to the previous one, it takes 4 clicks to do that. What could be the reason for this?
## [8][How to code JS in ReactJS?](https://www.reddit.com/r/reactjs/comments/jb0mu0/how_to_code_js_in_reactjs/)
- url: https://www.reddit.com/r/reactjs/comments/jb0mu0/how_to_code_js_in_reactjs/
---
I've read that you always code JS in useEffect, but I've seen other people code it differently. So confused about it, would like some clarity!
## [9][Phosphor Icons, a flexible icon family for everyone](https://www.reddit.com/r/reactjs/comments/jai9gi/phosphor_icons_a_flexible_icon_family_for_everyone/)
- url: https://phosphoricons.com/
---

## [10][Bind an method in the constructor function](https://www.reddit.com/r/reactjs/comments/jayxd9/bind_an_method_in_the_constructor_function/)
- url: https://www.reddit.com/r/reactjs/comments/jayxd9/bind_an_method_in_the_constructor_function/
---
Hi fellows !  
I took some vacations to go for an online React course at [codeacademy.com](https://codeacademy.com) .  
I have a question:  
Why do you need to bind, in the constructor function, a method to the object itself, when that method uses the *this* keyword?  
If the method is from a different class component instance, I understand why:  
When that happens, the *this* would refer to the instance and not the the source class component, that's why its needed to bind the method to the source class component.

But, if the method is from the same class component why do I need that *this* refers to the object itself?  


I hope my text was not confusing.  


Thanks a lot
## [11][How to save,edit and update user profile information in Reactjs with Nodejs](https://www.reddit.com/r/reactjs/comments/jayw0w/how_to_saveedit_and_update_user_profile/)
- url: https://www.reddit.com/r/reactjs/comments/jayw0w/how_to_saveedit_and_update_user_profile/
---
 This is the user profile page: [User Page](https://codesandbox.io/s/gifted-varahamihira-dbhco?file=/src/UserInfo.js)

I would like as soon as a user registers all fields are "readOnly",  then the user clicks the "Edit" button and the fields become editable.  While the fields are editable, the "Edit" button changes to "Save". Once  the user has filled out the form, he clicks on "Save", the data is sent  to the database and becomes visible in "readOnly". Once the data has been saved, from the second "Edit" onwards the "Save"  button will have to update without creating new elements in the  database.

First time: Edit--&gt;Input Data--&gt; Save   
 From second time: Edit--&gt;Change Data--&gt;Update

I know maybe it's a bit too much but I can't solve it myself. The back-end side works, but obviously every time I send the data a new line is created in the table.
## [12][I want to create a video editor along with some text animations](https://www.reddit.com/r/reactjs/comments/jayudf/i_want_to_create_a_video_editor_along_with_some/)
- url: https://www.reddit.com/r/reactjs/comments/jayudf/i_want_to_create_a_video_editor_along_with_some/
---
hey guys,  
I want to create a video editor in ReactJS.  
any recommendation about the libraries I need to use?  
if anyone had already created please allow me to fork.  
I had found these libraries.  any reviews about these?

1: PixiJS  
2: D3JS  
3: FabricJS

&amp;#x200B;

Cheers.
