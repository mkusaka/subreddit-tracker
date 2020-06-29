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
## [2][Should you include build output on GitHub?](https://www.reddit.com/r/typescript/comments/hhru63/should_you_include_build_output_on_github/)
- url: https://www.reddit.com/r/typescript/comments/hhru63/should_you_include_build_output_on_github/
---
I have several projects written in TypeScript that I publish to NPM and GitHub. It seems like bad practice to post the output JavaScript and Definition files on GitHub, however, you cannot include \`out\` in \`.gitignore\` as they will not be included in the NPM package. 

&amp;#x200B;

What is standard practice here?
## [3][[Ionic 5] Typescript error when initializing a string with single quotes!](https://www.reddit.com/r/typescript/comments/hhzhw9/ionic_5_typescript_error_when_initializing_a/)
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
## [4][Is a conversion of this code possible? If not, why not?](https://www.reddit.com/r/typescript/comments/hhz2fc/is_a_conversion_of_this_code_possible_if_not_why/)
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
## [5][TypeDi or Inversify or others?](https://www.reddit.com/r/typescript/comments/hhnqyp/typedi_or_inversify_or_others/)
- url: https://www.reddit.com/r/typescript/comments/hhnqyp/typedi_or_inversify_or_others/
---
Greetings typescripters,

Wanted to know what IoC framework you all are using and why?

P. S. I found TypeDi to be simpler to implement than Inversify.
## [6][Recommendations for server-side tsx](https://www.reddit.com/r/typescript/comments/hhs4ma/recommendations_for_serverside_tsx/)
- url: https://www.reddit.com/r/typescript/comments/hhs4ma/recommendations_for_serverside_tsx/
---
I would like to be able to generate HTML using tsx on the server. I know it's possible to specify my own jsxFactory, but perhaps there's already good libraries out there.

Do people do this? Do you have a recommendation?
## [7][How does testing work with an IoC container like Inversify?](https://www.reddit.com/r/typescript/comments/hhd1pn/how_does_testing_work_with_an_ioc_container_like/)
- url: https://www.reddit.com/r/typescript/comments/hhd1pn/how_does_testing_work_with_an_ioc_container_like/
---
I have a project using [Inversify](http://inversify.io/) for IoC/DI and I want to start writing tests, but I'm very confused on how to properly approach it, and for some reason I can't find any good articles or examples regarding this topic. I thought IoC was specifically made to make testing easier, as it decouples the direct implementation of classes with interfaces instead, but so far I can't find any evidence for this online in Inversify's case.

I have written tests for a simple Logger class, one with the IoC container (`container.get(Logger)`) and one with a direct implementation (`new Logger()`) and both work, but I don't know how to approach it for the more complex parts of my code, where a lot more injection is happening.

Does anyone have some tips on how to approach testing with Inversify or just IoC in general? Would love to read some good articles about it, or hear from your experiences.
## [8][Typing an instance property inside of another class using : syntax](https://www.reddit.com/r/typescript/comments/hhhyoj/typing_an_instance_property_inside_of_another/)
- url: https://www.reddit.com/r/typescript/comments/hhhyoj/typing_an_instance_property_inside_of_another/
---
Inside my main orchestration class I am trying to type a sentence instance that is passed as an argument:

    // inside Orchestrator class
    public async makeFolderAndAudioFile(sentence: Sentence): Promise&lt;void&gt; {
    
          // ; expected
          sentence.foreignWordDefinitionPairs: ForeignWordDefinitionPair[] = await Utilities.loopUntilFalse(this.getForeignWordAndDefinition);

Here is how this property is setup in its class:

    export default class Sentence {
       public foreignWordDefinitionPairs: ForeignWordDefinitionPair[]= []; 

I am able to type the first code block using `as` syntax:

    sentence.foreignWordDefinitionPairs = await Utilities.loopUntilFalse(this.getForeignWordAndDefinition) as ForeignWordDefinitionPair[];

Is this the proper way to do it? I don't use `as` syntax often. I tend to use it when TS isn't able to infer the proper type due to complex logic. normally I use the colon syntax to just make it clear what the expected return type is.
## [9][Why does ESlint think an object is a string?](https://www.reddit.com/r/typescript/comments/hh9i3w/why_does_eslint_think_an_object_is_a_string/)
- url: https://i.redd.it/kcd6itqcel751.png
---

## [10][Any optimizing compilers yet (like the Closure compiler)](https://www.reddit.com/r/typescript/comments/hh1rdt/any_optimizing_compilers_yet_like_the_closure/)
- url: https://www.reddit.com/r/typescript/comments/hh1rdt/any_optimizing_compilers_yet_like_the_closure/
---
Hi all – Do you know of any third-party Typescript compilers, especially optimizing ones? I figured there would be more than one compiler at this point, but I haven't found any.

There seems to be some room for improvement in optimization of the JS output. Google's Closure compiler is probably the best JS optimizer right now, but it doesn't take Typescript input. (It does take type annotations though, in JSDoc form.)

Thanks.
## [11][Exploring possible JS-APIs that mix between classes and inline-styling - Rosebox](https://www.reddit.com/r/typescript/comments/hgtpfg/exploring_possible_jsapis_that_mix_between/)
- url: https://www.reddit.com/r/typescript/comments/hgtpfg/exploring_possible_jsapis_that_mix_between/
---
Hi guys!I'm working on [https://www.rosebox.dev/](https://www.rosebox.dev/) which is an active project for building a complete TS framework for styling and which has the following goals:

1. Design and implement narrower types of properties than the ones in React.CSSProperties. This means help with auto-completion/IntelliSense and fewer typos.
2. Manage UI-state in JS instead of CSS.

So far RB  has been producing inline-styles but currently I'm exploring ways to reduce performance issues that might arise when using inline-styling with big lists which obviously increases the time the browser spends parsing HTML/CSS which is unnecessary and can be avoided using classes. So I want to give the user the ability to decide which parts of the style should go into a class and which parts need to be computed for each element in the list using the same object-based API. This is my first draft [https://codesandbox.io/s/blissful-joliot-4m5iw?file=/src/index.tsx](https://codesandbox.io/s/blissful-joliot-4m5iw?file=/src/index.tsx). I'd love to hear your opinions.

I'd also like to take the opportunity to announce that there is a now a discord-channel for RB https://discord.gg/bd7BHf2
