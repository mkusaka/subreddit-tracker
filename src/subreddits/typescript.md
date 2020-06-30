# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][Import from package with dependencies](https://www.reddit.com/r/typescript/comments/hijylv/import_from_package_with_dependencies/)
- url: https://www.reddit.com/r/typescript/comments/hijylv/import_from_package_with_dependencies/
---
Hello,

I am working an an application that is deployed as some microservices that are mostly written with Typescript and run with Node. We have split some functionality that is used in many backends to small NPM packages that those backends consume. Some of these packages provide abstractions for others and all is well so far.

But we encountered a problem on some packages that is keeping us busy for a while now. [I'll asked on stack overflow](https://stackoverflow.com/questions/61363752/typescript-import-type-from-package-without-dependencies-from-other-types-in-th) about this some months ago with no avail. So I hope somebody here can help us or provide us with some ideas.

&amp;#x200B;

To outline the problem a little bit:

Consider a package that provides a class A and class B. Class A is using an external dependency that also some other packages use. Class B is not using this dependency. To avoid having the dependency installed multiple times and to not having to install the dependency if you don't need it, it is set as peerDependency.

Now when I consume class A in a project and also make use of the dependency I install it and everything works fine. But if in another project I only need class B one could argue, that the dependency won't bring anything of value here, thus should not be installed. However if we try to use class B in a project, Typescript would always complain about the missing dependency.

&gt;error TS2307: Cannot find module 'x' or its corresponding type declarations.

&amp;#x200B;

Note: we use an index.ts for all our packages, where all classes, types and interfaces that should be used outside of the package itself are exported.

&amp;#x200B;

Now, is there any way, apart from splitting class B in it's own package or having to set "skipLibCheck": true in every tsconfig file, to avoid the typescript error?.
## [3][Is it worth setting up an enum for a single use?](https://www.reddit.com/r/typescript/comments/hilkz9/is_it_worth_setting_up_an_enum_for_a_single_use/)
- url: https://www.reddit.com/r/typescript/comments/hilkz9/is_it_worth_setting_up_an_enum_for_a_single_use/
---
I have a method that should process text/audio differently depending on whether the input was a sentence or word. I could setup a boolean like `isWord` but that is the least clear or extensible (phrase support may be added). 

Probably the JS option would be to use string option values.

In Typescript I have a file `minorTypes.ts` that I have used to import little enums. But I'm not sure if this is a good pattern or not, if the enum is only used in one method. So I wanted to ask about it. What do you guys think?
## [4][Publishing Your Deno Modules using GitHub](https://www.reddit.com/r/typescript/comments/hi6lny/publishing_your_deno_modules_using_github/)
- url: https://blog.bitsrc.io/publishing-your-deno-modules-using-github-f2bd86173392
---

## [5][Resource for big collection of Typescript extensions?](https://www.reddit.com/r/typescript/comments/hibh2v/resource_for_big_collection_of_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hibh2v/resource_for_big_collection_of_typescript/
---
I wrote an Array extension today so that I could get a random item from any array

    interface Array&lt;T&gt; {
        randomItem(): T;
    }
    
    Array.prototype.randomItem = function () {
        return this[Math.floor(Math.random() * (this.length) ) + 0];
    }

To avoid writing code that other programmers have already written, do you know of a resource which contains a long list of extensions that would come in handy?

Thank you
## [6][Is a conversion of this code possible? If not, why not?](https://www.reddit.com/r/typescript/comments/hhz2fc/is_a_conversion_of_this_code_possible_if_not_why/)
- url: https://www.reddit.com/r/typescript/comments/hhz2fc/is_a_conversion_of_this_code_possible_if_not_why/
---
is it possible to convert the following code to classes? I've been trying for hours as a way to try to understand what's happening behind the scenes but I can't seem to get anything to work. 

export interface IUser extends mongoose.Document {  
 name: string;   
 somethingElse?: number;   
  };  
   
 export const UserSchema = new mongoose.Schema({  
 name: {type:String, required: true},  
 somethingElse: Number,  
  });  
   
 const User = mongoose.model&lt;IUser&gt;('User', UserSchema);  
 export  {User};

&amp;#x200B;

I would like to convert the code to something like:

Class User... 

&amp;#x200B;

Is it possible? Please help. I've been at this for 8 hours now and I can't seem to understand a thing with the code above. I've even tried lynda and udemy. I figured in order for me to really understand, it would be best to see some sort of transition. I'd really appreciate any help I could get. Thank you.
## [7][Can't assign a object ref by string cause giving 'index' error](https://www.reddit.com/r/typescript/comments/hi5br0/cant_assign_a_object_ref_by_string_cause_giving/)
- url: https://www.reddit.com/r/typescript/comments/hi5br0/cant_assign_a_object_ref_by_string_cause_giving/
---
I have two different scenarios of errors.  


    type NType = {
      type: 'blue' | 'green' | 'red';
      message: string;
    };
    
      // ERROR HERE: (see below)
      const { type, message }: NType = useSelector(noteSelector);
    
      if (type &amp;&amp; message) {
        notify[type](message);
      }

**ERROR**: TS2322: Type '{ type: string; message: string; level: string; }' is not assignable to type 'NType'.   Types of property 'type' are incompatible.     Type 'string' is not assignable to type ''blue' | 'green' | 'red'.  


I feel I have to case type as string, but don't know how to do it within the destructuring. Any help?  


SECOND ISSUE --   
**ERROR**: TS7015: Element implicitly has an 'any' type because index expression is not of type 'number'.

    // error here
    const eventedAuth = (type: string, listener: Function) =&gt; {
      window[`${type}EventListener`](CHANGED_EVENT, listener);
    };
    
    
    // I solved it like so, but seems real sloppy
    const eventedAuth = (type: string, listener: Function) =&gt; {
      const win = (window as { [key: string]: any });
      const t = `${type}EventListener` as string;
      win[t](AUTH_TOKEN_CHANGED_EVENT, listener);
    };
## [8][Should you include build output on GitHub?](https://www.reddit.com/r/typescript/comments/hhru63/should_you_include_build_output_on_github/)
- url: https://www.reddit.com/r/typescript/comments/hhru63/should_you_include_build_output_on_github/
---
I have several projects written in TypeScript that I publish to NPM and GitHub. It seems like bad practice to post the output JavaScript and Definition files on GitHub, however, you cannot include \`out\` in \`.gitignore\` as they will not be included in the NPM package. 

&amp;#x200B;

What is standard practice here?
## [9][[Ionic 5] Typescript error when initializing a string with single quotes!](https://www.reddit.com/r/typescript/comments/hhzhw9/ionic_5_typescript_error_when_initializing_a/)
- url: https://www.reddit.com/r/typescript/comments/hhzhw9/ionic_5_typescript_error_when_initializing_a/
---
I have this code:

`export class ItemDetailsPage {`

  `connectionError: string = '';`

  `...`

It doesn't work unless I use double quotes.

Screenshot of the first part of the error:

https://preview.redd.it/kd5ncd8clu751.png?width=1366&amp;format=png&amp;auto=webp&amp;s=228a2d9238004368d6e61678ec2c9e71d4e7f970

I  am new to Ionic and this totally freaked me out.
## [10][TypeDi or Inversify or others?](https://www.reddit.com/r/typescript/comments/hhnqyp/typedi_or_inversify_or_others/)
- url: https://www.reddit.com/r/typescript/comments/hhnqyp/typedi_or_inversify_or_others/
---
Greetings typescripters,

Wanted to know what IoC framework you all are using and why?

P. S. I found TypeDi to be simpler to implement than Inversify.
## [11][Recommendations for server-side tsx](https://www.reddit.com/r/typescript/comments/hhs4ma/recommendations_for_serverside_tsx/)
- url: https://www.reddit.com/r/typescript/comments/hhs4ma/recommendations_for_serverside_tsx/
---
I would like to be able to generate HTML using tsx on the server. I know it's possible to specify my own jsxFactory, but perhaps there's already good libraries out there.

Do people do this? Do you have a recommendation?
