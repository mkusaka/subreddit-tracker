# reactjs
## [1][Beginner's Thread / Easy Questions (Jan 2020)](https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eictui/beginners_thread_easy_questions_jan_2020/
---
[Previous threads][wiki previous threads] can be found in the Wiki.

Got questions about React or anything else in its ecosystem? Stuck making progress on your app?  
Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## 🆘 Want Help with your Code? 🆘

- **Improve your chances** by putting a minimal example to either [JSFiddle][jsfiddle], [Code Sandbox][code sandbox] or [StackBlitz][stackblitz].
  - Describe what you want it to do, and things you've tried. Don't just post big blocks of code!
  - **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
- **Pay it forward!** Answer questions even if there is already an answer - multiple perspectives can be very helpful to beginners. Also there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].
- **[Learn by teaching][learn by teaching]** &amp; **[Learn in public][learn in public]** - It not only helps the asker but also the answerer.

---

## New to React?

### Check out the sub's sidebar!

🆓 Here are great, **free** resources! 🆓

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp](https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/)
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [Robin Wieruch's Road to React][robin wieruch's road to react]
- [FreeCodeCamp's React course](https://www.freecodecamp.org/news/learn-react-course/)
- [Flavio Copes' React handbook](https://reacthandbook.com/)
- New to Hooks? Check Amelia Wattenberger's [Thinking in React Hooks](https://wattenberger.com/blog/react-hooks)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[official getting started page]: https://reactjs.org/docs/getting-started.html
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[robin wieruch's road to react]: https://roadtoreact.com/
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Who's Available? [Jan 2020]](https://www.reddit.com/r/reactjs/comments/eouupz/whos_available_jan_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eouupz/whos_available_jan_2020/
---
We alternate between hirers (on the 1st of the month) and agencies/freelancers/jobseekers (on the 15th).  
If you are looking to post or reply to React job postings, please check [this month's Who's Hiring post here.][hiring:this month]

---

Top Level comments must be Agencies and React Devs available for contract/permanent work.

Please include Location or any other Requirements in your comment. You can choose to use this format if it helps:

## (Fulltime | Contract | USA | Remote)

or

## (Agency | Europe | Remote)

Then we recommend adding a 2-3 sentence bio as well.

Not required, but may help:

- Link to Github/Portfolio
- Notable [r/reactjs][r/reactjs] submissions
- Preferred stack
- Former companies or clients
- Design or backend dev experience
- anything else you consider relevant. Put on your best show!
- Listing years of experience NOT required, it's a poor metric

If you are looking to hire, you can send a PM, or reply so that others might see your job opening.  
**Note**: Due to the sensitive nature of availability while currently in a job, users may be using alternate accounts.

For more ideas on what to include, look at the [last Who's Available posts][available:last month].

If you just want some portfolio feedback, check the stickied post below.

Good luck! #WriteOnceApplyEverywhere

[r/reactjs]: https://www.reddit.com/r/reactjs/
[available:last month]: https://www.reddit.com/r/reactjs/comments/dxxqdn/whos_available_nov_2019/
[hiring:this month]: https://www.reddit.com/r/reactjs/comments/eidci5/whos_hiring_jan_2020/
## [3][My work on Enatega is coming to shape finally a food delivery solution if someone wants their own individual app. (More info and demo links in comments)](https://www.reddit.com/r/reactjs/comments/ep0ia9/my_work_on_enatega_is_coming_to_shape_finally_a/)
- url: https://www.youtube.com/watch?v=5rAI8FoyG8k&amp;feature=youtu.be
---

## [4][How do I push an array from a form on a post request?](https://www.reddit.com/r/reactjs/comments/ep01a0/how_do_i_push_an_array_from_a_form_on_a_post/)
- url: https://www.reddit.com/r/reactjs/comments/ep01a0/how_do_i_push_an_array_from_a_form_on_a_post/
---
Hello!

I'm stuck, I do not know how to push an array with JavaScript / React, I'm stuck.

I have a server which accepts post requests. With Postman I can make new entries with

```json
{
 "username" : "marios",
 "name" : "John Doe",
 "email": "johndoe@example.com",
 "address": "Somewhere in the World",
 "telephones": [
  {
   "telType": "Home",
   "telNumber": 1234567891
  },
  {
   "telType": "Mobile",
   "telNumber": 2101111111
  }
  ]
 }
```

And my Code in React is 
```react
import React, { Component } from 'react';
import axios from 'axios';

export default class AddContacts extends Component {
  constructor(props) {
    super(props);

    this.onChangeUsername   = this.onChangeUsername.bind(this);
    this.onChangeName       = this.onChangeName.bind(this);
    this.onChangeEmail      = this.onChangeEmail.bind(this);
    this.onChangeAddress    = this.onChangeAddress.bind(this);
    this.onChangeTelType    = this.onChangeTelType.bind(this);
    this.onChangeTelNumber  = this.onChangeTelNumber.bind(this);
    this.onSubmit           = this.onSubmit.bind(this);

    this.state = {
      users: [],
      username: '',
      name: '',
      email: '',
      address: '',
      telephones: [ ]
    }
  }

  componentDidMount() {
    axios.get('http://localhost:5000/users/')
      .then(response =&gt; {
        if (response.data.length &gt; 0) {
          this.setState({
            users: response.data.map(user =&gt; user.username),
            username: response.data[0].username
          })
        }
      })
      .catch((error) =&gt; {
        console.log(error);
      })

  }

  onChangeUsername(e) {
    this.setState({
      username: e.target.value
    })
  }

  onChangeName(e) {
    this.setState({
      name: e.target.value
    })
  }

  onChangeEmail(e) {
    this.setState({
      email: e.target.value
    })
  }

  onChangeAddress(e) {
    this.setState({
      address: e.target.value
    })
  }

  onChangeTelType (e) {
    this.setState({
      telephones: {
        telType: e.target.value
      }
    })
  }

  onChangeTelNumber (e) {
    this.setState({
      telephones: {
        telNumber: e.target.value
      }
    })
  }

  onSubmit(e) {
    e.preventDefault();

    const contact = {
      username: this.state.username,
      name: this.state.name,
      email: this.state.email,
      address: this.state.address,
      telephones: this.state.telephones
    }

    console.log(contact);

    axios.post('http://localhost:5000/contacts/add', contact)
      .then(res =&gt; console.log(res.data));

  }

  render() {
    return (
    &lt;div&gt;
      &lt;h3&gt;Create New Contact&lt;/h3&gt;
      &lt;form onSubmit={this.onSubmit}&gt;
        &lt;div className="form-group"&gt; 
          &lt;label&gt;Username: &lt;/label&gt;
          &lt;select 
            ref="userInput"
            required
            className="form-control"
            value={this.state.username}
            onChange={this.onChangeUsername}&gt;
            {
              this.state.users.map(function(user) {
              return (
                &lt;option 
                  key={user}
                  value={user}&gt;{user}
                &lt;/option&gt;
                )
              })
            }
          &lt;/select&gt;
        &lt;/div&gt;
        &lt;div className="form-group"&gt; 
          &lt;label&gt;Name: &lt;/label&gt;
          &lt;input  
              type="text"
              required
              className="form-control"
              value={this.state.name}
              onChange={this.onChangeName}
          /&gt;
        &lt;/div&gt;
        &lt;div className="form-group"&gt;
          &lt;label&gt;Email: &lt;/label&gt;
          &lt;input 
              type="text" 
              className="form-control"
              value={this.state.email}
              onChange={this.onChangeEmail}
          /&gt;
        &lt;/div&gt;
        &lt;div className="form-group"&gt;
          &lt;label&gt;Address: &lt;/label&gt;
          &lt;input 
              type="text" 
              className="form-control"
              value={this.state.address}
              onChange={this.onChangeAddress}
          /&gt;
        &lt;/div&gt;
        &lt;div className="form-group"&gt;
          &lt;label&gt;Telephones: &lt;/label&gt;
          &lt;div className="container"&gt;
            &lt;div className="row"&gt;
              &lt;div className="col-4"&gt;
                &lt;label&gt;Type: &lt;/label&gt;
                &lt;select 
                  required 
                  className="form-control"
                  value={this.state.telephones.telType}
                  onChange={this.onChangeTelType}
                &gt;
                  &lt;option key="landline" value="landline"&gt;Landline&lt;/option&gt;
                  &lt;option key="mobile" value="mobile"&gt;Mobile&lt;/option&gt;
                  &lt;option key="bussiness" value="bussiness"&gt;Bussiness&lt;/option&gt;
                &lt;/select&gt;
              &lt;/div&gt;
              &lt;div className="col-4"&gt;
                &lt;label&gt;Number: &lt;/label&gt;
                &lt;input 
                  type="string" 
                  className="form-control"
                  value=    {this.state.telephones.telNumber}
                  onChange= {this.onChangeTelNumber}
                /&gt;
              &lt;/div&gt;
            &lt;/div&gt; &lt;br /&gt;
            &lt;button className="btn btn-success"&gt;Add new Number&lt;/button&gt;
          &lt;/div&gt;
        &lt;/div&gt;

        &lt;div className="form-group"&gt;
          &lt;input 
          type="submit" 
          value="Add New Contact" 
          className="btn btn-primary" /&gt;
        &lt;/div&gt;
      &lt;/form&gt;
    &lt;/div&gt;
    )
  }
}
```

Can somebody help? 
Thanks for your time.
## [5][How can my page react to being scrolled?](https://www.reddit.com/r/reactjs/comments/eozusm/how_can_my_page_react_to_being_scrolled/)
- url: https://www.reddit.com/r/reactjs/comments/eozusm/how_can_my_page_react_to_being_scrolled/
---
How do I get the URL of a page (as displayed in the browser location bar) to change as I scroll down the page? For example, [like on this page](https://stripe.com/docs/api)?

I'd guess it's using the history object behind the scenes, but it's specifically doing it in reaction to the scroll reaching a certain part of the page that I don't know how to do.

(I don't know what keyword to search for. My searches for scrolling find pages about how to scroll from javascript, not how to react to the user scrolling.)

(I'd like to know how to do this from react, but information about how to do it in pure javascript also welcome.)
## [6][Storybook 5.3 release: stories &amp; docs in MDX, first-class React support, Declarative config, and Design tool integrations](https://www.reddit.com/r/reactjs/comments/eom889/storybook_53_release_stories_docs_in_mdx/)
- url: https://medium.com/storybookjs/storybook-5-3-83e114e8797c
---

## [7][Storybook 5.3 video summary, features and changes](https://www.reddit.com/r/reactjs/comments/ep10qp/storybook_53_video_summary_features_and_changes/)
- url: https://www.youtube.com/watch?v=Oy86bR1Enrs
---

## [8][Polymorphism in React and Javascript](https://www.reddit.com/r/reactjs/comments/ep0lg6/polymorphism_in_react_and_javascript/)
- url: https://www.reddit.com/r/reactjs/comments/ep0lg6/polymorphism_in_react_and_javascript/
---
Hello. Currently I am writing my first react project, it is a quiz website and I stumbled upon a problem I could not find any elegant way to solve.

I have several question types, and so I write a lot of switch cases in my code. I think this should be avoided, because adding new question type would require rewriting almost whole code instead of one file. I tried storing type as string, and storing questions in different arrays. Is there any correct or nicer way to handle this situation? I think about maybe storing functions inside my objects together with components
## [9][How to call useEffect only on mount while satisfying EsLint missing dependencies warning?](https://www.reddit.com/r/reactjs/comments/ep1ez2/how_to_call_useeffect_only_on_mount_while/)
- url: https://www.reddit.com/r/reactjs/comments/ep1ez2/how_to_call_useeffect_only_on_mount_while/
---
Hi All,

Below is an example of a useEffect with an empty dependency list:

      React.useEffect(() =&gt; {
        console.log("useEffect called")
        window.addEventListener('keydown', (e) =&gt; {
          if (keyCodes[e.keyCode] === props.pad.key.toLowerCase()) {
            playSound();
          }
        });
      }, []);

If I leave it as-is I seem to get the desired behaviour of useEffect only being called on mount, however, EsLint gives me a `react-hooks/exhaustive-deps` and wants me to add the following dependencies: `[playSound, props.pad.key])`.

I've added these dependencies with `playSound` wrapped in useCallback (see below) but doing this causes the `useEffect` to get called every time `props` is updated.

      const playSound = React.useCallback(() =&gt; {
        props.setNowPlaying(props.pad.id)
        const sound = document.getElementById(props.pad.key) 
        sound.onended = () =&gt; {
            props.setNowPlaying("")
            setActive(false);
          };
        setActive(true);
        sound.play();
      }, [props])

I can't inline the `playSound` function into the `useEffect as it's used elsewhere in the component. 

Can anyone provide some guidance on how best to change this code so as to only add the event listeners on mount whilst satisfying the EsLint warnings?

Please let me know if you need more information. Thank you for your time.
## [10][Liquid Swipe (Hover) implementation for React](https://www.reddit.com/r/reactjs/comments/eorjzw/liquid_swipe_hover_implementation_for_react/)
- url: https://www.reddit.com/r/reactjs/comments/eorjzw/liquid_swipe_hover_implementation_for_react/
---
I came across this really cool youtube tutorial for a liquid swipe implementation on React Native. I was wondering if anyone could point me to any tutorials or resources that could show me how to implement this using React for web, where the bezier curve follows the cursor. Thanks in advance!

Video: [Liquid Swipe](https://www.youtube.com/watch?v=gLopy2MCAqM&amp;list=PL-lHSNZUNZvm6JkdhYdvSuL4Uq9f-PmnZ&amp;index=3&amp;t=0s)
## [11][ReasonML / ReasonReact in 2020](https://www.reddit.com/r/reactjs/comments/eopmrp/reasonml_reasonreact_in_2020/)
- url: https://www.reddit.com/r/reactjs/comments/eopmrp/reasonml_reasonreact_in_2020/
---
Hi guys,

I've been aware of ReasonML and ReasonReact for a while now, but only recently decided to do a deep dive. My impression after a few days using it in my spare time is that it's brilliant, it really makes sense when they say it's ES2030 JavaScript.

My concern is that Facebook is going to let Reason die before it has a chance to really take off. Checking the ReasonML blog there wasn't a single update last year and there were only 3 updates on the ReasonReact blog.

So my question is: is it worth learning Reason in 2020?

As far as the tool itself goes, I would love to invest my time in learning how to use it. However the community seems small and there are next to no jobs advertising ReasonML (at least in London), so I'm a bit hesitant and was hoping someone in the know could shed some light on Reason and where it's heading this year.
## [12][RxJS Facades in React: Push-Based Architecture with Less BS](https://www.reddit.com/r/reactjs/comments/eos4xu/rxjs_facades_in_react_pushbased_architecture_with/)
- url: https://medium.com/@thomasburlesonIA/react-facade-best-practices-1c8186d8495a
---

