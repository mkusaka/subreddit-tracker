# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][I've been using Typescript (migrated after 15 years of Java) for about 1.5 years. Here's what I've learned and how my programming style has changed.](https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/)
- url: https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/
---
- I've found I really like inner functions.  I'll put functions within functions to group similar code.  In Java this wasn't possible (at least 8 ... haven't checked lately).  It's better to keep things close together especially if the function isn't reusable.

- I'm calling all interfaces as IFoo rather than just Foo. Interfaces are a bit different than in Java and are far more important.

- Almost all my code is migrating to namespaces and I'm migrating almost entirely from static functions and static classes.  Webpack can't tree shake them and I think this is just cleaner either way.

- I'm moving to a lot more functions, maps, streams, etc.  

- I use undefined a lot more than null.  

- Types and generics work pretty well but every once in a while I get stuck and they can kind of get confusing with edge cases.

Curious how it's changed your programming style.
## [3][HTML 5 Canvas, React Refs and TypeScript](https://www.reddit.com/r/typescript/comments/itucpy/html_5_canvas_react_refs_and_typescript/)
- url: https://hashnode.blainegarrett.com/html-5-canvas-react-refs-and-typescript-ckf4jju8r00eypos1gyisenyf
---

## [4][Assign dynamic properties from constructor to class (typechecking)](https://www.reddit.com/r/typescript/comments/itcssj/assign_dynamic_properties_from_constructor_to/)
- url: https://www.reddit.com/r/typescript/comments/itcssj/assign_dynamic_properties_from_constructor_to/
---
In short, what I'd like to do:

```
class Person {

  name: string

  // dynamic properties from constructor param


  constructor(name: string, dynProps: any) {

    this.name = name

    Object.keys(dynProps).forEach(key =&gt; this[key] = dynProps[key]) 

 } 

}


const dynamicProperties = { age: 25 } 


const someone = new Person('Greg', dynamicProperties)


someone.name // 'Greg'

someone.age // it should aknowledge it as number type and return 25
```
## [5][How to add custom types to the TypeScript project with examples](https://www.reddit.com/r/typescript/comments/it5n5p/how_to_add_custom_types_to_the_typescript_project/)
- url: https://drag13.io/posts/custom-typings/index.html
---

## [6][What is the baseUrl in compilerOptions used for?](https://www.reddit.com/r/typescript/comments/itb8b2/what_is_the_baseurl_in_compileroptions_used_for/)
- url: https://www.reddit.com/r/typescript/comments/itb8b2/what_is_the_baseurl_in_compileroptions_used_for/
---
    $ npm run cypress:open
    
    &gt; react@1.0.0 cypress:open C:\src\project\React
    &gt; cypress open
    
    It looks like this is your first time using Cypress: 5.2.0
    
    [11:28:42]  Verifying Cypress can run C:\Users\admin\AppData\Local\Cypress\Cache\5.2.0\Cypress [started]
    [11:28:45]  Verified Cypress!       C:\Users\admin\AppData\Local\Cypress\Cache\5.2.0\Cypress [title changed]
    [11:28:45]  Verified Cypress!       C:\Users\admin\AppData\Local\Cypress\Cache\5.2.0\Cypress [completed]
    
    Opening Cypress...
    
    Missing baseUrl in compilerOptions. tsconfig-paths will be skipped

I am wondering if it's something I should have or not.
## [7][Strange type error](https://www.reddit.com/r/typescript/comments/itafg4/strange_type_error/)
- url: https://www.reddit.com/r/typescript/comments/itafg4/strange_type_error/
---
    return (
            &lt;div className="form-group"&gt;
              &lt;label for="text"&gt;Name&lt;/label&gt;
            &lt;/div&gt;
        )
    

&gt;Type '{ children: string; for: string; }' is not assignable to type 'DetailedHTMLProps&lt;LabelHTMLAttributes&lt;HTMLLabelElement&gt;, HTMLLabelElement&gt;'.  
&gt;  
&gt;  Property 'for' does not exist on type 'DetailedHTMLProps&lt;LabelHTMLAttributes&lt;HTMLLabelElement&gt;, HTMLLabelElement&gt;'. Did you mean 'htmlFor'?

Hmm, why am I getting a type error on a html attribute?
## [8][I need help with setting dynamic types](https://www.reddit.com/r/typescript/comments/isq6ed/i_need_help_with_setting_dynamic_types/)
- url: https://www.reddit.com/r/typescript/comments/isq6ed/i_need_help_with_setting_dynamic_types/
---
Hey everyone, I'm trying to type a map like class and I'm not sure if it's possible to do it the way I want it to work. Essentially what I want is this:

```
factory.define('user', UserModel, { /* ... default values for the UserModel */ } )

factory.create('user', { /* ... UserModel fields to override */ }): UserModel
```

In the `define` method you specify a key string, a type it matches to and some default values. Then in `create` TS understand that the `'user'` key means that the second parameter must be of type `UserModel`, and it returns a `UserModel`.

I don't want to have to define all the keys (`'user'` in this case) as a type ahead of time, I want to define the methods such that typescript understands that the `define` parameters are setting keys and values. Anyone know if this is possible, and if so, how?

Many thanks for your help!

Edit 1: updated to a better explaining example.
## [9][Small SaaS looking to add Developer Cofounder](https://www.reddit.com/r/typescript/comments/ist6rr/small_saas_looking_to_add_developer_cofounder/)
- url: https://www.reddit.com/r/typescript/comments/ist6rr/small_saas_looking_to_add_developer_cofounder/
---
Hi group. This is the small business we're building one customer at a time - [https://podsquad.app/](https://podsquad.app/). We're looking to add another developer cofounder to the team. Frontend is React Native and backend is an Express JS server running on Heroku using Parse for storage. Please reach out of interested!
## [10][Effective Typescript | A Non-Technical Guide](https://www.reddit.com/r/typescript/comments/is5d0m/effective_typescript_a_nontechnical_guide/)
- url: https://matthewmullin.io/effective-typescript/
---

## [11][How do I set minimum and maximum elements for an array?](https://www.reddit.com/r/typescript/comments/isi7dn/how_do_i_set_minimum_and_maximum_elements_for_an/)
- url: https://www.reddit.com/r/typescript/comments/isi7dn/how_do_i_set_minimum_and_maximum_elements_for_an/
---
basically, I want to limit the length of the `coordinates` parameter in the `getDirection` function.

how do I force the caller to pass an array of `LatLng` with a minimum length of 2 and a maximum elements of 25?

  
please take a look at my codes below

&amp;#x200B;

    interface LatLng {
     lat: number
     lng: number
    }
    
    const getDirection = async (coordinates: LatLng[]) {
    
    }

Iv'e seen tuple examples but it looks like it's only for setting a minimum or fixed length of elements.
