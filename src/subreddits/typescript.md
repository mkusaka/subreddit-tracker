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
## [2][Get union type from property type shared by different types](https://www.reddit.com/r/typescript/comments/hezem0/get_union_type_from_property_type_shared_by/)
- url: https://www.reddit.com/r/typescript/comments/hezem0/get_union_type_from_property_type_shared_by/
---
Sorry for the weird title, but I don't really know how to describe it in a short way :D So here is an example:

    type T1 = {
        prop: 'v1';
    }
    type T2 = {
        prop: 'v2';
    }

Now, I want to have something like that:

    type PropType = 'v1' | 'v2';

BUT without explicitly writing v1 and v2 again. Is there some way to extract the type from "prop" off T1 and T2? Something like:

    type PropType = ExtractPropertyType&lt;T1 | T2, 'prop'&gt;;

Thanks for any help!
## [3][Best approach for video streaming web app](https://www.reddit.com/r/typescript/comments/heqpdn/best_approach_for_video_streaming_web_app/)
- url: https://www.reddit.com/r/typescript/comments/heqpdn/best_approach_for_video_streaming_web_app/
---
Hey there! 

I have a project in mind that would require broadcasting video from user to user, with some features that allow for applying filters to video frames in-between (in node/typescript backend service). 

I've read about webrtc a bit, but there aren't a ton of resources out there and the examples on webrtc site are a bit vague. 

Any ideas / suggestions / alternatives? 

Thank you!
## [4][Help passing generic callback arguments to a utility method's callback](https://www.reddit.com/r/typescript/comments/hf10mx/help_passing_generic_callback_arguments_to_a/)
- url: https://www.reddit.com/r/typescript/comments/hf10mx/help_passing_generic_callback_arguments_to_a/
---
    export default class Utilities {
    
       static loopUntilFalse&lt;Args, CallbackReturn&gt;(callback: (args: Args) =&gt; boolean): unknown {
          let continueLooping = true;
          let returnValue: CallbackReturn;
    
          while (continueLooping === true) {
    
             let returnValue = callback(args);
             if (returnValue === false) { continueLooping = false }
          }
    
          return returnValue;
       }
    }

How should an unknown shape of arguments be passed into the method here? I'm a little rough on generics and also function argument syntax in TS, so I suspect the problem is on one of those, or a combination of both.

&amp;#x200B;
## [5][Generic type defined at class-level](https://www.reddit.com/r/typescript/comments/hewb4z/generic_type_defined_at_classlevel/)
- url: https://www.reddit.com/r/typescript/comments/hewb4z/generic_type_defined_at_classlevel/
---
I'm trying to approach something like this:

```
class Document {}

class User extends Document {
  email: string;
}

export class Service&lt;T extends Document&gt; {
  constructor() {}
  find(): T {
    const user = new User();
    user.email = 'info@example.com';
    return user;
  }
}

export class UserService extends Service&lt;User&gt; {
  constructor() {
    super();
  }
}

const userService = new UserService();
const user = userService.find();
console.log("user &gt;&gt;&gt;&gt;", user);
```

this gives me the error: 

```
test.ts:12:5 - error TS2322: Type 'User' is not assignable to type 'T'.
  'User' is assignable to the constraint of type 'T', but 'T' could be instantiated with a different subtype of constraint 'Document'.

12     return user;
       ~~~~~~~~~~~~
```

how can I fix this?
## [6][(Beta) I've written a guide to architecting Node TypeScript apps - still WIP but has a runnable companion repo and full documentation](https://www.reddit.com/r/typescript/comments/hed665/beta_ive_written_a_guide_to_architecting_node/)
- url: https://jbreckmckye.gitbook.io/node-ts-architecture/
---

## [7][Is there any way to prevent type collision?](https://www.reddit.com/r/typescript/comments/henw82/is_there_any_way_to_prevent_type_collision/)
- url: https://www.reddit.com/r/typescript/comments/henw82/is_there_any_way_to_prevent_type_collision/
---
I have two libs, Electron and react-dropzone that conflict because Electron decide it was going to add 'path' into File in the global namespace.

So now if I use both projects together, they conflict.  

Is there a way to force some types to be ignored or somehow force them to be compatible?
## [8][Assigning Buffer to string issue](https://www.reddit.com/r/typescript/comments/hej52r/assigning_buffer_to_string_issue/)
- url: https://www.reddit.com/r/typescript/comments/hej52r/assigning_buffer_to_string_issue/
---
    // doesn't work
          const configDataRaw = fs.readFileSync(
             path.join(__dirname, "../../../configuration.json", "utf8")
          );
    
          const configData = JSON.parse(configDataRaw);
    // const configDataRaw: Buffer
    // Argument of type 'Buffer' is not assignable to parameter of type 'string'.ts(2345)
    
    // works
          const configData2 = JSON.parse(path.join(__dirname, 
            "../../../configuration.json", "utf8")
          );

All of the following failed:

     const configDataRaw as string = fs.readFileSync(
    
     const configDataRaw&lt;string&gt; = fs.readFileSync(
    
    
          const configData = JSON.parse((configDataRaw as string));
    
          const configData = JSON.parse(configDataRaw&lt;string&gt;);

What is the correct way to assert this is not a buffer, but a string? 

In the first code block, code sample #2, it seems to infer properly, I don't understand why the more iterative version acts differently here.
## [9][Pool about strict mode](https://www.reddit.com/r/typescript/comments/her1r3/pool_about_strict_mode/)
- url: https://www.reddit.com/r/typescript/comments/her1r3/pool_about_strict_mode/
---
Most projects that I start, I start in strict mode.

Strict mode avoid a lot of simples bugs, avoid run time errors, but some programmers consider ir more difficult...

Do you prefer to program in strict mode?

[View Poll](https://www.reddit.com/poll/her1r3)
## [10][Help reating this type](https://www.reddit.com/r/typescript/comments/heehno/help_reating_this_type/)
- url: https://www.reddit.com/r/typescript/comments/heehno/help_reating_this_type/
---
     translationServiceStub?: Promise&lt;{
            [name: string]: Function;
        }&gt;;

I see this is an optional type. It is a returned promise. The promise contains an object. The object has a key `name` that is a string. the value of `name` is a function (in other words `name` is a method).

Is that accurate?

If so can someone tell me why it wasn't written more concisely like this:

     translationServiceStub?: Promise&lt;{
            name: Function;
        }&gt;;

Or even:

     translationServiceStub?: Promise&lt;{
            name: () =&gt; void // or whatever the expected return type is
        }&gt;;
## [11][Typescript interfaces for i18 translations?](https://www.reddit.com/r/typescript/comments/heh3ou/typescript_interfaces_for_i18_translations/)
- url: https://www.reddit.com/r/typescript/comments/heh3ou/typescript_interfaces_for_i18_translations/
---
I want to have our app translated into multiple languages but I also want to do it the 'typescript way' where our strings aren't just in an 'object' but are in a TS interface.

Has anyone done anything like this before?

I've only seen these mapped to basic JSON objects and not typescript interfaces.

I think the following issues exist for Typescript:

 - how do you do hierarchy?  I'd like to avoid excessive nesting and possibly without the need to constantly name and define a new interface for each component.  For example a key of onboarding.hello could/should be used so that the 'onboarding' component can have multiple keys.

- Is it better to isolate internationalizations per component rather than have one big internationalization blob with all of the internationalizations?  So for example you could have a React component like Onboarding.tsx and then a translation of Onboarding.i18 that is a properties file with each translation inside it...  That might be cleaner rather than having one big object. One challenge here is it would have to survive refactoring and still work with the translation service.

- How do you handle interpolation with Typescript strings and get compiler safety?  For example, an interpolation could take 'name' and 'date' properties and these would need to be applied at runtime. If we forget one we would need to get a compiler error.
