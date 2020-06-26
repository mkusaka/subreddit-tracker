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
## [2][Introduction React Native, Typescript, Redux, Rxjs, Node.js, Mongo, Mongoose, Nexus Aurora Platform](https://www.reddit.com/r/typescript/comments/hg5g8z/introduction_react_native_typescript_redux_rxjs/)
- url: https://www.youtube.com/watch?v=06wsg2f76hQ
---

## [3][Stryker Plugin in Webstorm for easier Mutation Testing](https://www.reddit.com/r/typescript/comments/hfrnm0/stryker_plugin_in_webstorm_for_easier_mutation/)
- url: https://plugins.jetbrains.com/plugin/14482-stryker
---

## [4][Requiring property on a type based on the presence of another property.](https://www.reddit.com/r/typescript/comments/hg52p4/requiring_property_on_a_type_based_on_the/)
- url: https://www.reddit.com/r/typescript/comments/hg52p4/requiring_property_on_a_type_based_on_the/
---
So, I am not really sure what to search to find out more about this concept, and therefore I apologise in advance if this is a duplicate, but I have been stuck for a day now and need some help.

I am using TS 3.6.4 in a React project, and I am using an interface to validate props passed to a component.

So, let's say I have an interface called IButton as follows:

    interface IButton {
      label: string;
      iconName?: string;
      iconColor?: string;
    }

Is there a way I can only require `iconColor` be supplied only if `iconName` is present, and if `iconName` is not present, then have both required to be undefined?

For some context, the reason I want to do this is because the button may not always have an icon, but when an icon is being used, the color for said icon is required alongside it.
## [5][A naive ECS implementation in typescript!](https://www.reddit.com/r/typescript/comments/hfnjsn/a_naive_ecs_implementation_in_typescript/)
- url: https://github.com/Avokadoen/ts-ecs-waveshoot
---

## [6][Best way to handle family members in People objects?](https://www.reddit.com/r/typescript/comments/hfu5cz/best_way_to_handle_family_members_in_people/)
- url: https://www.reddit.com/r/typescript/comments/hfu5cz/best_way_to_handle_family_members_in_people/
---
Let's say I have a list of patrons, People&lt;&gt;. What is the best pattern to have a single source of truth for who their family members are? (Parties of people, no need for family trees.) Do I need to create Family objects?
## [7][Intersección de tipos en Typescript](https://www.reddit.com/r/typescript/comments/hg42h0/intersección_de_tipos_en_typescript/)
- url: https://emanuelpeg.blogspot.com/2020/06/interseccion-de-tipos-en-typescript.html#.XvW3595FkMw.reddit
---

## [8][What is New in Angular 10?](https://www.reddit.com/r/typescript/comments/hfmox0/what_is_new_in_angular_10/)
- url: https://volosoft.com/blog/what-is-new-in-angular-10
---

## [9][Type 'Promise&lt;X&gt;' is not a valid async function return type in ES5/ES3 because it does not refer to a Promise-compatible constructor value.ts(1055)](https://www.reddit.com/r/typescript/comments/hfq8co/type_promisex_is_not_a_valid_async_function/)
- url: https://www.reddit.com/r/typescript/comments/hfq8co/type_promisex_is_not_a_valid_async_function/
---
Method was looking good until I had to convert it to async because I needed to await another internal method that makes an API call.

I still have trouble with return types on async functions and this current error totally confuses me, hopefully someone knows what it means:

    /* 
       Asks for the foreign word
       
       Gets a definition, asks for confirmation or an adjustment.
       Returns an object with the foreign word and definition pair
       
       Returns false if user marks "done" commands
       */
       async protected getForeignWordAndDefinition()
          : Promise&lt;ForeignWordDefinitionPair&gt; | Promise&lt;false&gt; {
        // (alias) interface ForeignWordDefinitionPair
        // import ForeignWordDefinitionPair
        // Type 'Promise&lt;ForeignWordDefinitionPair&gt; | Promise&lt;false&gt;' 
        // is not a valid async function return type in ES5/ES3 because 
        // it does not refer to a Promise-compatible constructor value. ts(1055)
    
        // Originally it was   : ForeignWordDefinitionPair | false
        // same error occured
    
          console.log("prompt text")
          const userInput = readLine.question();
    
          const userHasExited = this.isDone(userInput);
          if (userHasExited) return false;
    
          // translate foreign to english
          const foreignWord = userInput;
    
          const googleOfferedDefinition: string = await this.textTranslate(
             foreignWord, translationDirection.toEnglish
          );
    
          console.log("prompt text")
          console.log(googleOfferedDefinition);
          const userDefinition: string = readLine.question();
    
          let acceptedDefinition = googleOfferedDefinition;
          if (userDefinition !== "") {
             acceptedDefinition = userDefinition;
          } 
    
          // shape the object and return it
          const foreignWordDefinitionPair: ForeignWordDefinitionPair = {
             foreignWord
             , englishDefinition: acceptedDefinition
          }
    
          return foreignWordDefinitionPair;
       }
## [10][simple algo visualizer -- reverse an array](https://www.reddit.com/r/typescript/comments/hfcnnp/simple_algo_visualizer_reverse_an_array/)
- url: https://www.reddit.com/r/typescript/comments/hfcnnp/simple_algo_visualizer_reverse_an_array/
---
A simple little app to visualize reversing an array.   


Want to start doing more algorithms as well as visualizing them.   


Starting off very basic! 

&amp;#x200B;

 [https://github.com/risingBirdSong/visualizeReverseArray](https://github.com/risingBirdSong/visualizeReverseArray)
## [11][8 Visual Studio Code Assistant rules for nasty TypeScript / Angular bugs](https://www.reddit.com/r/typescript/comments/hfhu2t/8_visual_studio_code_assistant_rules_for_nasty/)
- url: https://medium.com/@tomaszs2/8-visual-studio-code-assistant-rules-for-nasty-angular-bugs-9f186277e0ab
---

