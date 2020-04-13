# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Been building my first TypeScript project for 3 months, was going well. Updated all NPM packages and everything Broke...](https://www.reddit.com/r/typescript/comments/g08pum/been_building_my_first_typescript_project_for_3/)
- url: https://www.reddit.com/r/typescript/comments/g08pum/been_building_my_first_typescript_project_for_3/
---
So I've given most of my Express middleware this type: [https://pastebin.com/kCnjfUex](https://pastebin.com/kCnjfUex)

But now, after updating all of my NPM packages to the latest I have 100+ errors all saying something like this:`TS2322: Type 'RequestHandler&lt;ParamsDictionary, any, any, any&gt;' is not assignable to type 'Middleware'.   Type 'RequestHandler&lt;ParamsDictionary, any, any, any&gt;' provides no match for the signature '(req: RequestExtended, res: Response&lt;any&gt;, next: NextFunction): void | Promise&lt;Response&lt;any&gt;&gt;'`

Any ideas on what I'm doing wrong?  


EDIT: Fixed, I deleted node\_modules and package.lock.json and then used command `npm i`  
@ Types/express needed to be at version 4.17.0, and so changed that in package.json  
Found this information here: [https://github.com/DefinitelyTyped/DefinitelyTyped/issues/40138](https://github.com/DefinitelyTyped/DefinitelyTyped/issues/40138)
## [3][Help understanding type inference and utility types](https://www.reddit.com/r/typescript/comments/g0f5em/help_understanding_type_inference_and_utility/)
- url: https://www.reddit.com/r/typescript/comments/g0f5em/help_understanding_type_inference_and_utility/
---
That's how `Extract` utility type is defined in TS lib:

    /**
     * Extract from T those types that are assignable to U
     */
    type Extract&lt;T, U&gt; = T extends U ? T : never;

I understand that when `T` is say `'a'` and `U` is say `'a' | 'b'` it'll return `'a'` because `'a'` is the subtype (“extends“) of `'a' | 'b'`. But I cannot understand how does it work when `T` is not a subtype of `U` but have common types with it. Here's an example:

    type AB = 'a' | 'b'
    type BC = 'b' | 'c'
    type E = Extract&lt;AB, BC&gt;
    const e: E = 'b'

This works, though `AB` is not a subtype of `BC` (which I don't understand). Okay, but what if I inline `Extract`?

    type AB = 'a' | 'b'
    type BC = 'b' | 'c'
    type E = AB extends BC ? AB : never;
    const e: E = 'b'    // Does not work, E is never

Nope, does not work (which I understand). But how does it happen that, if `Extract` is defined as in lib, it works?
## [4][Yupress: combining Yup and Express](https://www.reddit.com/r/typescript/comments/g07r7l/yupress_combining_yup_and_express/)
- url: https://www.reddit.com/r/typescript/comments/g07r7l/yupress_combining_yup_and_express/
---
The [`yupress` library](https://github.com/Malvolio/yupress) harnesses the power of Yup to make writing Express-based applications in Typescript simpler, more reliable, and more secure.

I'm posting it here before releasing it as an NPM module because I wanted to see whether people wanted significant changes (and even understood it at all).  All comments welcome.
## [5][Moving to Typescript from JS on express site. Do you guys know the Typescript equivalent of module.exports?](https://www.reddit.com/r/typescript/comments/g0475a/moving_to_typescript_from_js_on_express_site_do/)
- url: https://www.reddit.com/r/typescript/comments/g0475a/moving_to_typescript_from_js_on_express_site_do/
---

I use module.exports to export a function in a file and the Typescript compiler is telling me that it isn't exporting the functions in my file.

Here is the file

    // @ts-nocheck
    import db from "../psql-con";
    
    module.exports = {
      // returns true if the input inserted by the user meets the requirements in the form.
      validUser: (email, username, password) =&gt; {
        const emailRegEx = /^(([^&lt;&gt;()\[\]\\.,;:\s@"]+(\.[^&lt;&gt;()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
        const validEmail = emailRegEx.test(email);
    
        const validUsername =
          typeof username == "string" &amp;&amp; username.trim().length &gt; 2;
    
        // minimum eight characters, at least one uppercase letter, one lowercase letter and one number
        const passwordRegEx = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/;
        const validPassword = passwordRegEx.test(password);
        return validEmail &amp;&amp; validUsername &amp;&amp; validPassword;
      },
    
      // returns true if email exists
      emailExists: async (email) =&gt; {
        const { rows } = await db.query(
          "SELECT email FROM users WHERE email = $1",
          [email]
        );
        if (rows.length === 0) {
          return false;
        }
        return true;
      },
    
      // returns true if username exists
      usernameExists: async (username) =&gt; {
        const {
          rows,
        } = await db.query("SELECT username FROM users WHERE username = $1", [
          username,
        ]);
        if (rows.length === 0) {
          return false;
        }
        return true;
      },
    };


Do you guys know the Typescript equivalent of module.exports?
## [6][How to setup Typescript, Eslint, Prettier and React in 5 minutes](https://www.reddit.com/r/typescript/comments/g0e3hw/how_to_setup_typescript_eslint_prettier_and_react/)
- url: https://medium.com/@carljohan.kihl/how-to-setup-typescript-eslint-prettier-and-react-in-5-minutes-44cfe8af5081?source=friends_link&amp;sk=a46f3a9158870daa7887dcf4cf34cc5b
---

## [7][I am developing a new JavaScript library to add keyboard shortcuts. Would you use it?](https://www.reddit.com/r/typescript/comments/g0gsn4/i_am_developing_a_new_javascript_library_to_add/)
- url: https://www.reddit.com/r/typescript/comments/g0gsn4/i_am_developing_a_new_javascript_library_to_add/
---


[View Poll](https://www.reddit.com/poll/g0gsn4)
## [8][Typed function to recursively access object properties?](https://www.reddit.com/r/typescript/comments/fzx6e8/typed_function_to_recursively_access_object/)
- url: https://www.reddit.com/r/typescript/comments/fzx6e8/typed_function_to_recursively_access_object/
---
Is there a good way in typescript to type a function that will recursively access object properties? Here's a hand-coded example of going two levels deep:

```
function deepProp&lt;T, K extends keyof T&gt;(obj: T, prop1: K, prop2: keyof T[K]) {
    return obj[prop1][prop2];
}

// Example use
const obj = {
    user: {
        pets: [{
            toys: [{
                name: "Toughy",
                price: 1999
            }]
        }]
    }
} as const
deepProp(obj, "user", "pets");
```

But I'm looking for a good way to take any number of `props` in the `deepProp` function to dive down as deep as necessary. I imagine that function's signature would be something like `function deepProp(obj, ...props) {  }`. Is there a good way to do this? Thanks!
## [9][Getting correct props on a generic form field in React](https://www.reddit.com/r/typescript/comments/g03equ/getting_correct_props_on_a_generic_form_field_in/)
- url: https://www.reddit.com/r/typescript/comments/g03equ/getting_correct_props_on_a_generic_form_field_in/
---
I'm using Formik to generate some forms. I have a generic `Form.Field` component that has a `component` prop. The component passed in here could be of any type - `Form.Input.Text`, `Form.Input.Select`, `Form.Input.Date`, etc.

When using the `Form.Field` component, I want it to know the required props for the component it's rendering it. How can I do this?

[Here's an example of what I'm asking for in Code Sandbox](https://codesandbox.io/s/quirky-river-b7p35?file=/src/App.tsx) (excuse the handful of shortcuts I made to make this.) See the comments inside of `src/App.tsx`. Intellisense says the `data` property on the 2nd form component doesn't fit there but the page still renders as expected. If you remove the `data` property, the compiler throws an error because `data.map(...)` inside of `src/Form/Select.tsx` doesn't handle when data is null elegantly. I know I can handle that better, but `data` is a required property, so I shouldn't ever need to do that in this simple example.

The reason for this structure is re-usability, so I can import just 1 component and build an entire Form from it. I'm not looking to directly import form fields separately, but I'm willing to change how this code is written. Again, the end goal of the structure I want is to import 1 Form component that can let me build an entire form.
## [10][Help with understanding a function signature](https://www.reddit.com/r/typescript/comments/g01fyg/help_with_understanding_a_function_signature/)
- url: https://www.reddit.com/r/typescript/comments/g01fyg/help_with_understanding_a_function_signature/
---
Hi everyone, 
I'm fairly new to Typescript and I was wondering if you could help me.

In my project im using React and Typescript and I have the following component:
```
&lt;Chip key={chipData.key} label={chipData.label} className={classes.chip} onClick={handleChipClick(chipData)} /&gt;
```
My onClick event handler looks like this: 
```
const handleChipClick = (chipToDelete: ChipData) =&gt; () =&gt; {}
```
My question to you is:
```
= (chipToDelete: ChipData) =&gt; () 
```
what does this function signature mean and why does it work, but this one doesnt?
```
= (chipToDelete: ChipData) =&gt; {}
```
## [11][How to add relay to Create-React-App with Typescript](https://www.reddit.com/r/typescript/comments/fzr1jm/how_to_add_relay_to_createreactapp_with_typescript/)
- url: https://medium.com/@ricardojgonzlez/how-to-add-relay-to-create-react-app-with-typescript-b6daacea21dd
---

