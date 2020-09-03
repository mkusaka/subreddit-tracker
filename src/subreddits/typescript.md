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
## [2][Is it wrong to use namespaces? ES5 imports feel inferior to me compared to namespaces](https://www.reddit.com/r/typescript/comments/illv8o/is_it_wrong_to_use_namespaces_es5_imports_feel/)
- url: https://www.reddit.com/r/typescript/comments/illv8o/is_it_wrong_to_use_namespaces_es5_imports_feel/
---
I have the following code that works perfectly.

    export namespace Vid {
      export const SPLASH = require('./Splash.mp4')
      export const SEQUENCE_1 = require('./Sequence01.mp4')
      export const SEQUENCE_2 = require('./Sequence02.mp4')
      export const SEQUENCE_3 = require('./Sequence03.mp4')
    }

So if I want to import a Video in my file, I type `Vid.`, and I get autocomplete for all my available videos. This is great for me.

_____

Unfortunately namespaces are deprecated for some reason and I'm supposed to export the const directly.

      export const SPLASH = require('./Splash.mp4')
      export const SEQUENCE_1 = require('./Sequence01.mp4')
      export const SEQUENCE_2 = require('./Sequence02.mp4')
      export const SEQUENCE_3 = require('./Sequence03.mp4')

This is NOT ideal. Because now when I'm importing files, I need to remember the file name itself instead of just typing `Vid.` and being able to loop through my files. There will also be higher namespace collisions with other exported consts.

_____

I've seen someone suggest to do the following to replicate namespaces.

    const SPLASH = require('./Splash.mp4')
    const SEQUENCE_1 = require('./Sequence01.mp4')
    const SEQUENCE_2 = require('./Sequence02.mp4')
    const SEQUENCE_3 = require('./Sequence03.mp4')

    export const Vid = { SPLASH, SEQUENCE_1, SEQUENCE_2, SEQUENCE_3 }

However after using it for awhile, I still find it inferior.

1) `Go to Declaration` no longer works properly. When I try to lookup the declaration, it points me to `export const Vid`, then I have to look up the declaration a second time to jump to my code. This is really annoying.

2) It no longer prompts me which consts are unused. With namespaces, I'm informed which ones are used and unused. With consts, it doesn't seem to know anymore.

- [Namespaces](https://i.imgur.com/ChJwD9K.png)
- [ES5](https://i.imgur.com/4f1qyMP.png)

_____

I've reverted to using namespaces for the time being, because in my opinion they are superior. Can someone tell me what potential problems I could have by continuing to use namespaces? Does anyone know how long more they will be supported for?

**EDIT: Too many people are answering by telling me the pros of ES imports. That is not the question. The question is:**

- **What is wrong with using namespaces?**
## [3][TypeScript ORM with no query builder, supporting full SQL queries](https://www.reddit.com/r/typescript/comments/iltfzj/typescript_orm_with_no_query_builder_supporting/)
- url: https://github.com/Seb-C/kiss-orm#kiss-orm
---

## [4][Make third party package as global visible to all other files](https://www.reddit.com/r/typescript/comments/ilowzn/make_third_party_package_as_global_visible_to_all/)
- url: https://www.reddit.com/r/typescript/comments/ilowzn/make_third_party_package_as_global_visible_to_all/
---
Hello,

I installed momentjs and want to the project automatically access to the function everywhere without explicitly importing. Is this possible in typescript?

Thanks
## [5][Need help to understand how assignments in arrays/objects are done](https://www.reddit.com/r/typescript/comments/ilkrk4/need_help_to_understand_how_assignments_in/)
- url: https://www.reddit.com/r/typescript/comments/ilkrk4/need_help_to_understand_how_assignments_in/
---
Hello,

I have a problem that I think is pretty simple, but I can't get my head around the actual mechanics and why it does not work. Here is my issue:

I have an object (pModel) that is passed (by reference, I assume) in a function. That object contains an array of objects (geometries) and those objects have a member variable called isHighlighted (boolean). So basically, pModel.geometries\[i\].isHighlighted = true or false.

What I am doing is that I am looping through the geometries to toggle the isHighlighted variable. I am currently doing it this way:

    for (let geometry of pModel.geometries)
    {
        console.log(geometry.isHighlighted); // returns false
        console.log(geometry);  // isHighlight inside object is false
        geometry.isHighlighted = !geometry.isHighlighted;
        console.log(geometry.isHighlighted); // returns true
        console.log(geometry);  // isHighlighted inside object is still false
    }

I don't understand why isHightlighted is not changed in the second console.log(geometry). Also, if I write geometry.isHighlighted = true, it works. the second console.log(geometry) will have isHighlighted = true. I don't understand that either...
## [6][Building a game with TypeScript. Drawing Grid 4/5](https://www.reddit.com/r/typescript/comments/il512l/building_a_game_with_typescript_drawing_grid_45/)
- url: https://medium.com/@gregsolo/building-a-game-with-typescript-iii-drawing-grid-4-5-398af1dd638d?sk=49c92b3604c33e633bf6dc4a1e2846ed
---

## [7][What to expect from Typescript 4.0 and more](https://www.reddit.com/r/typescript/comments/il2uaw/what_to_expect_from_typescript_40_and_more/)
- url: https://medium.com/@metodieff.stefan/what-is-new-in-typescript-4-0-and-more-6c5fa72fa1db?source=friends_link&amp;sk=649d4059694facf4f13223b3831da308
---

## [8][learning typescript](https://www.reddit.com/r/typescript/comments/il0ysf/learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/il0ysf/learning_typescript/
---
so, i really wanna learn typescript, for discord.js and maybe a little bit of game dev and website dev, but i honestly have zero clue on how to get started, i learn by doing, and not just watching videos because then i can't remember things, any tips? i would really appreciate the help
## [9][Declaration of types in a separate file](https://www.reddit.com/r/typescript/comments/ikruk6/declaration_of_types_in_a_separate_file/)
- url: https://www.reddit.com/r/typescript/comments/ikruk6/declaration_of_types_in_a_separate_file/
---
Hi everyone!

I have a module which export a class with some methods. Methods returns a complex data.

For example `SomeClass.ts`:

`// some imports`

`export type ComplexType1 = … // generated from json. Contains a lot of subtypes, enums etc.`

`export type ComplexType2= …`

`// and so on`

`class SomeClass {`

`method1 = (): ComplexType1 =&gt; {…}`

`method2 = (): ComplexType1 =&gt; {…}`

`// etc`

`}`

This structure is repeated in several files.

My problem is that type declarations takes several screens. I trying just copy-paste types into SomeClass.d.ts but types became unavailable inside SomeClass.ts

How I can move types in separate file?
## [10][How to define a prop type for an object that is generic](https://www.reddit.com/r/typescript/comments/ikvnvh/how_to_define_a_prop_type_for_an_object_that_is/)
- url: https://www.reddit.com/r/typescript/comments/ikvnvh/how_to_define_a_prop_type_for_an_object_that_is/
---
So I have a React component with maybe some excessive generics. I've got a component that takes in an array of rules, and a separate prop that is an object of arguments to pass to each rule. It's a pattern that I've adopted because it avoids creating a closure over the arguments so the rules can be re-evaluated on each render (since the props being passed get updated).

Anyway, my main problem is my PropTypes. I want to still use PropTypes because this code is for a library that will also be used in JavaScript projects, not just TypeScript ones. I just can't get these to work together, the moment I do it breaks compilation.

I know it has something to do with the fact that I'm typing "ruleProps" as an object, whereas it's generic in the TypeScript interface. I'll also point out that it's the "rules" prop which breaks, because the arguments expected by the function no longer match the type.

Hoping for guidance on how to do this. Thanks.

    export interface Rule&lt;RuleProps extends object&gt; {
        allow: (ruleProps?: RuleProps) =&gt; boolean;
        redirect: string;
    }
    
    interface Props&lt;CompProps extends object, RuleProps extends object&gt; {
        rules?: Array&lt;Rule&lt;RuleProps&gt;&gt;;
        ruleProps?: RuleProps;
    }
    
    ProtectedRoute.propTypes = {
        rules: PropTypes.arrayOf(PropTypes.shape({
            allow: PropTypes.func.isRequired,
            redirect: PropTypes.string.isRequired
        })),
        ruleProps: PropTypes.object
    };

Edit: The error I'm getting is:

    TS2322: Type 'Rule&lt;RuleProps&gt;[] | undefined' is not assignable to type 'Rule&lt;object&gt;[] | undefined'.   Type 'Rule&lt;RuleProps&gt;[]' is not assignable to type 'Rule&lt;object&gt;[]'.     Type 'Rule&lt;RuleProps&gt;' is not assignable to type 'Rule&lt;object&gt;'.       Property 'allow' is missing in type '{}' but required in type 'RuleProps'.

This error occurs in a unit test when I try to actually use the component and pass values to these props. If I remove the PropTypes, everything works just fine.
## [11][If you are using TypeDoc, I created a theme with particles background. Hope you like it.](https://www.reddit.com/r/typescript/comments/ikofcd/if_you_are_using_typedoc_i_created_a_theme_with/)
- url: https://github.com/tsparticles/typedoc-particles-theme
---

