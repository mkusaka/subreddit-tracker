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
## [2][simple algo visualizer -- reverse an array](https://www.reddit.com/r/typescript/comments/hfcnnp/simple_algo_visualizer_reverse_an_array/)
- url: https://www.reddit.com/r/typescript/comments/hfcnnp/simple_algo_visualizer_reverse_an_array/
---
A simple little app to visualize reversing an array.   


Want to start doing more algorithms as well as visualizing them.   


Starting off very basic! 

&amp;#x200B;

 [https://github.com/risingBirdSong/visualizeReverseArray](https://github.com/risingBirdSong/visualizeReverseArray)
## [3][8 Visual Studio Code Assistant rules for nasty TypeScript / Angular bugs](https://www.reddit.com/r/typescript/comments/hfhu2t/8_visual_studio_code_assistant_rules_for_nasty/)
- url: https://medium.com/@tomaszs2/8-visual-studio-code-assistant-rules-for-nasty-angular-bugs-9f186277e0ab
---

## [4][Getting Map value type](https://www.reddit.com/r/typescript/comments/hf5g3a/getting_map_value_type/)
- url: https://www.reddit.com/r/typescript/comments/hf5g3a/getting_map_value_type/
---
Let's say I have a custom Map type, like this:

```
type shapeMap = Map&lt;string, [number, number][]&gt;
```

How do I get the value type of this Map back (`[number, number][]`)? I can do something stupid like this, but it seems very incorrect:

```
type shapeMapValue = Exclude&lt;ReturnType&lt;shapeMap["get"]&gt;, undefined&gt;
```
## [5][Get union type from property type shared by different types](https://www.reddit.com/r/typescript/comments/hezem0/get_union_type_from_property_type_shared_by/)
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
## [6][An alternative to Google Geocoder API (in Node.js)](https://www.reddit.com/r/typescript/comments/hf2jyh/an_alternative_to_google_geocoder_api_in_nodejs/)
- url: https://www.reddit.com/r/typescript/comments/hf2jyh/an_alternative_to_google_geocoder_api_in_nodejs/
---
Yesterday I started writing a few web scrapers in Node.js to gather some data for a personal project of mine (wait for it). One of the things I had to deal, is how to convert addresses to geolocation (latitude and longitude, basically.) So I started doing what we do best as Javascript developers: 

*Hm… there must be a package for that.*

&amp;#x200B;

[https://medium.com/@armand1m\_/an-alternative-to-google-geocoder-api-in-node-js-78728c7b9faa?source=friends\_link&amp;sk=914c1b3acee09686d740675b1e8424d7](https://medium.com/@armand1m_/an-alternative-to-google-geocoder-api-in-node-js-78728c7b9faa?source=friends_link&amp;sk=914c1b3acee09686d740675b1e8424d7)
## [7][Help passing generic callback arguments to a utility method's callback](https://www.reddit.com/r/typescript/comments/hf10mx/help_passing_generic_callback_arguments_to_a/)
- url: https://www.reddit.com/r/typescript/comments/hf10mx/help_passing_generic_callback_arguments_to_a/
---
    export default class Utilities {
    
       /* take a generic callback argument. Store all return values 
          of the callback until the callback returns false
    
       Then, end looping. Return an array of all prior return values
       */
       static loopUntilFalse&lt;Args, CallbackReturn&gt;(callback: (args: Args) =&gt; CallbackReturn|false): unknown {
          let continueLooping = true;
          const returnValues: Array&lt;CallbackReturn&gt; = [];
    
          while (continueLooping === true) {
    
             // any
             // Cannot find name 'args'.ts(2304)
             const latestReturnValue = callback(args);
    
             if (latestReturnValue === false) { 
                continueLooping = false 
             } else if (typeof latestReturnValue !== "boolean") {
                returnValues.push(latestReturnValue);
             }
          }
    
          return returnValues;
       }
    }

How should an unknown shape of arguments be passed into the method here? I'm a little rough on generics and also function argument syntax in TS, so I suspect the problem is on one of those, or a combination of both.

Also if my generic return syntax looks off any feedback there is also welcome.

&amp;#x200B;
## [8][Compile-time null/undefined check confusion](https://www.reddit.com/r/typescript/comments/hf39oc/compiletime_nullundefined_check_confusion/)
- url: https://www.reddit.com/r/typescript/comments/hf39oc/compiletime_nullundefined_check_confusion/
---
I'm having trouble understanding the static analysis in TypeScript, can anyone tell me why this works:

    const myFunc = (aParam?: string, upper?: boolean): string =&gt; {
      if (aParam &amp;&amp; upper) {
        return aParam.toUpperCase();
      }
      return 'Hello';
    }

but not this:

    const myFunc = (aParam?: string, upper?: boolean): string =&gt; {
      const x = aParam &amp;&amp; upper;
      if (x) {
        return aParam.toUpperCase();
      }
      return 'Hello';
    }

The bottom code results in "Object is possibly 'undefined'" for aParam
## [9][Best approach for video streaming web app](https://www.reddit.com/r/typescript/comments/heqpdn/best_approach_for_video_streaming_web_app/)
- url: https://www.reddit.com/r/typescript/comments/heqpdn/best_approach_for_video_streaming_web_app/
---
Hey there! 

I have a project in mind that would require broadcasting video from user to user, with some features that allow for applying filters to video frames in-between (in node/typescript backend service). 

I've read about webrtc a bit, but there aren't a ton of resources out there and the examples on webrtc site are a bit vague. 

Any ideas / suggestions / alternatives? 

Thank you!
## [10][Generic type defined at class-level](https://www.reddit.com/r/typescript/comments/hewb4z/generic_type_defined_at_classlevel/)
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
## [11][(Beta) I've written a guide to architecting Node TypeScript apps - still WIP but has a runnable companion repo and full documentation](https://www.reddit.com/r/typescript/comments/hed665/beta_ive_written_a_guide_to_architecting_node/)
- url: https://jbreckmckye.gitbook.io/node-ts-architecture/
---

