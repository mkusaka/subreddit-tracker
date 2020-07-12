# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][After learning Rust, I was surprised that the concept of "match" (also in Haskell) didn't make it to ES2020/TS4.0. So I created my own little class for it. Thoughts?](https://www.reddit.com/r/typescript/comments/hphpu8/after_learning_rust_i_was_surprised_that_the/)
- url: https://gist.github.com/brianboyko/1decb88a793b92c06ad7d13f73d88854
---

## [3][Importing .ts files without extension](https://www.reddit.com/r/typescript/comments/hps207/importing_ts_files_without_extension/)
- url: https://www.reddit.com/r/typescript/comments/hps207/importing_ts_files_without_extension/
---
Hi, I recently migrated a vue project to typescript manually without cli, and I managed to set everything up except importing ts files. Somehow I need to specify the .ts on the import path.

Can someone please help?
## [4][Q: Using yarn workspace but Typescript isn't aware of it](https://www.reddit.com/r/typescript/comments/hpqnrm/q_using_yarn_workspace_but_typescript_isnt_aware/)
- url: https://www.reddit.com/r/typescript/comments/hpqnrm/q_using_yarn_workspace_but_typescript_isnt_aware/
---
When I import something from an another workspace I get `Cannot find module @org/common/Movie or its corresponding type...` from the TS compiler/Intellisense. How can make TS aware of my workspace structure?
## [5][Type '"X"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)](https://www.reddit.com/r/typescript/comments/hpranq/type_x_does_not_satisfy_the_constraint_keyof/)
- url: https://www.reddit.com/r/typescript/comments/hpranq/type_x_does_not_satisfy_the_constraint_keyof/
---
       protected combineAndSave &lt;NamingOption&gt;(
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , fileNameOptions : 
             Pick&lt;NamingOption, "prefix"&gt; |
             Pick&lt;NamingOption, "name"&gt;
       ) : void {
    
    /*
    Type '"prefix"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)
    Type '"name"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)
    */

I'm still rough with generics, why isn't `NamingOption` being treated as amorphous in shape above? I thought that's the point of a generic type, maybe my concept of them is wrong, or maybe it's a syntax error.

PS: The intent of the above is that argument 3 receive an object with a single key, for either `prefix` or `name` (not both). I know an object union could also do this but I want to learn the generic type way.
## [6][Object argument: Any way to enforce one of two keys must be present?](https://www.reddit.com/r/typescript/comments/hpg1mr/object_argument_any_way_to_enforce_one_of_two/)
- url: https://www.reddit.com/r/typescript/comments/hpg1mr/object_argument_any_way_to_enforce_one_of_two/
---
       protected combineAndSave(
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , fileNameOptions : {
             prefix?: string
             , name?: string
          }
       ) : void {
          const {prefix, name} = fileNameOptions;
          let finalFileSavePath: string;
          
          // used to save sentence files
          if (prefix) {
             finalFileSavePath = `${savePath}/${prefix} - ${this.sentence.folderName}.ogg`;
          } else {
             finalFileSavePath = `${savePath}/${name}`;
          }

I think there's a hole above. Looks like someone may be able to pass an empty object, or even exclude the 3rd argument completely, which would hit the `else` clause with an `undefined` variable and cause a `TypeError` (if this is wrong please let me know).

Is there any way to constrain the `fileNameOptions` object so that one of the keys must be specified?

I tried this

          , fileNameOptions : Partial&lt;{
             prefix?: string
             , name?: string
          }&gt;

But on review of the docs this apparently does nothing the dual `?` don't already do.

I may leave it anyways for code clarity, if anyone has an opinion on that please share.
## [7][[Showoff Saturday] Platform for uploading and tracking Applications build with TypeScript](https://www.reddit.com/r/typescript/comments/hpia6e/showoff_saturday_platform_for_uploading_and/)
- url: https://progressiveapp.store/home
---

## [8][Noob Question: how do I enforce the implementation of a static method on a child class?](https://www.reddit.com/r/typescript/comments/hpj5h2/noob_question_how_do_i_enforce_the_implementation/)
- url: https://www.reddit.com/r/typescript/comments/hpj5h2/noob_question_how_do_i_enforce_the_implementation/
---
I know I can do this

```
abstract class A {
  myMethod();
}

class B extends A {
    myMethod() {
        console.log("implemented");
    }
}
```

and I can do this

```
class B {
    static myProperty = "hello";

    static myMethod() {
        console.log("hello");
    }
}
```

but I can't seem to mix the two.

I want to make sure that all child classes of a parent (abstract) class override a certain static property (or method).

Specifically, I want all descendants of `FromSQL` to have a static property `TABLE_NAME`, so that I can have a method `fetch&lt;T extends FromSQL&gt;` and then use `TABLE_NAME` to fetch from the correct table.

Is there any way to do this? If not, is there a workaround I can use?
## [9]["--isolatedModules flag" in typescript](https://www.reddit.com/r/typescript/comments/hp440a/isolatedmodules_flag_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hp440a/isolatedmodules_flag_in_typescript/
---
&gt; All files must be modules when the '--isolatedModules' flag is provided

I've googled this error but I honestly don't think I'm violating anything. 

    import React from 'react';
    import Hero from '../components/Hero';
    
    const About: React.FC = () =&gt; (
      &lt;Hero /&gt;
    );
    
    export default About;
    
The error occurs at "(1,1)"..so  the first character. 

I have an almost exactly similar file that doesn't produce the error. It is also super strange because when I was working on this file yesterday and saved it in this state, it was rendering fine. But right now as I run the app again..I get that error. Does anyone know what it could be?

One answer from S/O: This error happens when there is no import or export statement in a file (these make a file a module). But I do have an export statement..
## [10][[ Help ] Create new method (wrapping)](https://www.reddit.com/r/typescript/comments/hp5gao/help_create_new_method_wrapping/)
- url: https://www.reddit.com/r/typescript/comments/hp5gao/help_create_new_method_wrapping/
---
Hi you! As the title says I need to create a new method and I know there is a prototype way but isn't encouraged.
I was told about wrapping and then I looked into the documentation and Stack Overflow but I couldn't find any clear explanation/answer.

I want achieve something like this:
numberInstance.myMethod(anyNumber)

Thanks for helping me !
## [11][How to share common typings in a project? [Lerna / Workspaces]](https://www.reddit.com/r/typescript/comments/hor78i/how_to_share_common_typings_in_a_project_lerna/)
- url: https://www.reddit.com/r/typescript/comments/hor78i/how_to_share_common_typings_in_a_project_lerna/
---
I did bit of a searching and couldn't find a real example anywhere.  

I'm about to create a project (both backend and frontend) that might share common typings (especially models etc). I haven't set up the lerna yet (the only share-able components are these typings at the moment). What's the best way to organize this? Can someone shed some light on this?  

Ideally, I guess, this common types would export two modules (if I understand the `module`s correctly) that matches the package names of the frontend and backend packages.

It'd be great if I get some ELI5 on `module` vs `namespace`, and how to organize typings that spans across multiple files.

Thanks!  

Edit: Big thanks to u/intrepidsovereign for creating a [boilerplate](https://github.com/RyanChristian4427/example-monorepo-ts)
