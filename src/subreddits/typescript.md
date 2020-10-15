# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Question about creating an interface out of T](https://www.reddit.com/r/typescript/comments/jbgjrs/question_about_creating_an_interface_out_of_t/)
- url: https://www.reddit.com/r/typescript/comments/jbgjrs/question_about_creating_an_interface_out_of_t/
---
Hello,

&amp;#x200B;

is it possible to create an interface based on a given type?

    // main file
    interface LoginForm {
        username: string
        password: string
    }
    const form = Form&lt;LoginForm&gt;{
        username: ''
        password: ''
    }
    
    // types.ts
    export interface Form&lt;T&gt; {
     // create the Form interface based on the given type 'T'
    }

what I want to create is the Form interface based on the given `T` so that I can use the `keyof` like the following 

    const FormKey = keyof Form

and then in the main file, I can just use the FormKey type on a method like..

    const hasError (formKey: FormKey) =&gt; {
        return this.form[formKey]
    }

&amp;#x200B;

is my approach doable? or are any way around to do this? 

  
any response is appreciated, thanks!
## [3][node file.ts](https://www.reddit.com/r/typescript/comments/jbmps8/node_filets/)
- url: https://www.reddit.com/r/typescript/comments/jbmps8/node_filets/
---
When I use ts-jest I do :

    npx jest aTestFileName.test.ts

which compiles the ts file to JavaScript on the run, and executes it, but also auto source maps the errors. There are no JavaScript file or source map files created on my project folder.

Is there any way to do something similar when I want to execute a ts file that is not a test.ts file for jest?

For the time being I create a fake ts test file and use ts-jest .
## [4][Building a game With TypeScript. New episode! Team and Fleet. Today we are going to talk about camaraderie and opposition and introduce 3 new members of our happy family: Team, Fleet, and, of course, Ship](https://www.reddit.com/r/typescript/comments/jb4h1b/building_a_game_with_typescript_new_episode_team/)
- url: https://medium.com/@gregsolo/building-a-game-with-typescript-team-and-fleet-f223d39e9248?source=friends_link&amp;sk=937227e4118edb7ea3125a641a52422d
---

## [5][Looking for typescript websocket library to communicate between node.js and vue.js](https://www.reddit.com/r/typescript/comments/jbalqc/looking_for_typescript_websocket_library_to/)
- url: https://www.reddit.com/r/typescript/comments/jbalqc/looking_for_typescript_websocket_library_to/
---
I need push communication from the server.

I was trying [socket.io](https://socket.io) but couldn't figure out how to turn of cross-site-scripting protection that was triggering for localhost (also [socket.io](https://socket.io) doesn't seem to be great with typescript).
## [6][Why do interfaces even exist in TypeScript?](https://www.reddit.com/r/typescript/comments/jas9hl/why_do_interfaces_even_exist_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jas9hl/why_do_interfaces_even_exist_in_typescript/
---
I am rather inexperienced in TypeScript, but I am not asking of how to use them, but rather why.

I may be missing something, but as I searched for `interface` vs `type`, I of course got some comparison, but I still wonder.

As I chose TypeScript, I expect to type my variables and then I read, that actually interfaces do overlap or merge? **Does not this mean, that they actually get closer to** `any` **type?**

...

During of writing this question I have realized, that I might have approached TypeScript with some naivety concerning error prevention, while viewing TypeScript as omnipotent tool.

But I realize, that it still gets transpiled back to JavaScript, so its 'runtime' (which is actually JavaScript runtime) cannot offer more tools, then JavaScript offers - unless TypeScript would create absurd constructions, which would allow passing types as strings for example - which might have been much easier if it was a regularly compiled language.

So what I think now is, that we cannot use TypeScript to absolutely prevent runtime errors.

That was quite confusing to me. I understand it now as, that we actually choose protection level, where `interfaces` could be considered as *medium security*, while `any` would be *no security* and `type` or `class` *higher security*. So I view it, that we choose the trade off of security and *human-writeability*.

Correct me, if I am wrong.

...

However, what I just found cool about combination of an `interface` with a `class` is that while it does not guarantee, that `class` **hasn't** *undefined* properties (which would be a real issue, if `interface` suggests, that there is an `object` within our `class` and we would try to access property of that *undefined* `object`), sometimes it is nice to have *undefined* properties, as there may be no need of initializing them in some cases, which would waste performance.

Again, correct me if I am wrong.

...

But maybe I still miss something, so feel free to become transpilers of my text and rise exceptions. Thanks!

...

Also still I found some scary 'tutorials', in one, it was suggested to use `noImplicityAny = false`, which I think, is pretty insane to do, while choosing strictly typed language. Also different tutorial stated, that `type` and `interface` should be used in specific different scenarios, however they haven't mentioned any reason behind - maybe you have some real reasons?
## [7][Strange compilation on down level iteration](https://www.reddit.com/r/typescript/comments/jawgy4/strange_compilation_on_down_level_iteration/)
- url: https://www.reddit.com/r/typescript/comments/jawgy4/strange_compilation_on_down_level_iteration/
---
Hello!

```
const countChildren = &lt;T extends HTMLElement&gt;(el: T) =&gt; {
  return [...el.children].length;
};
```

The idea being, to turn an HTMLCollection to an array. Of course to do this, you need to have these in `tsconfig.json`, `lib: ["dom.iterable", /*w.e. else*/]` and `downlevelIteration:true`.

When I do this just now in the [typescript playground](https://www.typescriptlang.org/play?#code/MYewdgzgLgBApgGxgXhgExMArgWzmKAOmACc4BDKOAUQTjwIAoByLBZgSgG4AoH0SLABmASxLQU6TLnxFSFKrXqyWCEZ37gJEOALSSM2BnLKUadY6vUceiQuQAOD-GgDCACxEI0jUeKg2do7OYG6e3ow6ejaagjAiVCSSANqEaXbA4WhkYAC6vLEQIHSECCAA5owJcCQcQA) it works, though the transpiled code persist with `[...arr]`.

However, and this is perhaps oddly specific, running `tsdx create some-project` and then enabling `dom.iterable` and `downlevelIteration`, compiles `countChildren` to:

```
var countChildren = function countChildren(el) {
  return [].concat(el.children).length;
};
```

Which is wrong. 

&gt; This yields `[HTMLCollection(2)]` instead of  `[child0, child1]`.

It should be something akin to `[].slice.call(el.children)`

Am I missing something here? Some option in `tsconfig` is perhaps playing me foul?
## [8][TS &amp; React.js: Conflict between initial state and type during prop passage on conditionally rendered component](https://www.reddit.com/r/typescript/comments/jap2sv/ts_reactjs_conflict_between_initial_state_and/)
- url: https://www.reddit.com/r/typescript/comments/jap2sv/ts_reactjs_conflict_between_initial_state_and/
---
In react I assign accountData to an empty object using useState. If a session is detected and loads a dataset, accountData is then updated to the shape of an interface Accountdata.

In the parent component:

      const [accountData, setAccountData] = React.useState({});
    
    // inside return()
      } else if (isAuthenticated &amp;&amp; userDataHasLoaded) {
        return &lt;AuthView 
                accountData={accountData}
              /&gt;
      } else {

Inside the child component:

    interface PropsShape {
      accountData: IAccountData;
    }
    
    export default ({ setIsAuthenticated, accountData }: PropsShape) =&gt; {
      return (
                &lt;Originals
                  accountData={accountData}

But there is a conflict initially since the type is initially set to {}.

I was hoping the conditional render would block a type error but it did not. What then, is the proper way to handle this situation and prevent this error:

    Failed to compile
    /home/owner/cp/project/src/ApplicationBase/App.tsx
    TypeScript error in /home/owner/cp/project/src/ApplicationBase/App.tsx(102,13):
    Type '{}' is missing the following properties from type 'IAccountData': projectCompacts, tasksEnhanced  TS2739
    
        100 |     return &lt;AuthView 
        101 |             setIsAuthenticated={setIsAuthenticated} 
      &gt; 102 |             accountData={accountData}
            |             ^
        103 |           /&gt;
        104 |   } else {
        105 |     return (
## [9][Prefer default or named exports?](https://www.reddit.com/r/typescript/comments/jalhr2/prefer_default_or_named_exports/)
- url: https://www.reddit.com/r/typescript/comments/jalhr2/prefer_default_or_named_exports/
---
ESLint AirBnb template wants me to prefer default exports. When I brought up that VsCode would be nice to prefer default exports on "Refactor into new file" I got some replies saying named exports are better in most ways.

What do you guys have to say about this? I'm moving towards named recently as it makes it so nice to write interfaces and classes in the current file and "poof" them away into a support file without another thought. If named is also better, great.
## [10][How to display singular resources from those with variable ID endpoints.](https://www.reddit.com/r/typescript/comments/jak7yj/how_to_display_singular_resources_from_those_with/)
- url: https://www.reddit.com/r/typescript/comments/jak7yj/how_to_display_singular_resources_from_those_with/
---
Hello, I have a list of data coming from a resource all in one endpoint of /users from the backend in a Java controller. To get singular resources, I use /users/{id} on that same controller. I am wondering how to get those singular resources to display using my user.service.ts, and app-routing.module.ts. This is what I currently have:



    public find(): Observable&lt;User&gt;{
        return this.http.get&lt;User&gt;(this.usersUrl + '/1'); //users/{id} = /users/{id}
    } //under user.service.ts
    ...
    const routes: Routes = [
        { path: 'users', component: UserListComponent },
        { path: 'users/1', component: UserComponent },
    ]; //under app-routing.module.ts

TLDR: I just want to pass in users/{id} rather than users/1 but TS doesn't let me. I think it's a formatting thing for automatically pushing stuff into browser fields.
## [11][Need Some Help with the Pipe Operator](https://www.reddit.com/r/typescript/comments/jafad3/need_some_help_with_the_pipe_operator/)
- url: https://www.reddit.com/r/typescript/comments/jafad3/need_some_help_with_the_pipe_operator/
---
Hello,  I need some help figuring this out.

I have a interface User:

    interface User {
      name?: string;
      email: string;
      password?: string;
      newPassword?: string;
      currentPassword?: string;
    }

And a interface Product:

    interface Product {
      title: string;
      description: string;
      url: string;
    }

Then I have my  useFormValidation hook where I pass initialState as a parameter:

    function useFormValidation(initialState: User | Product, ...){
        const [values, setValues] = useState(initialState);
        ...
        return{
            ...
            values
        }
    }

When I go to use the hook and get the values:

    const { ..., values } = useFormValidation(
        INITIAL_STATE,
        ...
    );

And use the values:

    const { title, description, url } = values;

I get this error:  `Property 'title' does not exist on type 'User | Product'` 

Same for description and url. 

If I try setting the type to Product : 

    const { title, description, url }: Product = values;

I get this error:

    Type 'User | Product' is not assignable to type 'Product'.
      Type 'User' is missing the following properties from type 'Product': title, description, url

What can I do to fix this? Thanks.
