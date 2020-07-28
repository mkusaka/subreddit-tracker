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
## [2][A Mental Model to think in Typescript](https://www.reddit.com/r/typescript/comments/hz93ap/a_mental_model_to_think_in_typescript/)
- url: https://leandrotk.github.io/tk/2020/07/a-mental-model-to-think-in-typescript/
---

## [3][Help with axios method type safety inside function](https://www.reddit.com/r/typescript/comments/hzcyxu/help_with_axios_method_type_safety_inside_function/)
- url: https://www.reddit.com/r/typescript/comments/hzcyxu/help_with_axios_method_type_safety_inside_function/
---
Hello guys, I am trying to type safe a react hook that uses axios for making requests, I am not being able to use axios\[method\] inside the function it gives me the error, I don't get it. It works if I only put the get in the enum though...

https://preview.redd.it/0htw7ftslld51.jpg?width=1098&amp;format=pjpg&amp;auto=webp&amp;s=f0b86b90225b75976c38fe7f5081e7d1be84aa8e
## [4][Learning TypeScript from JavaScript?, and framework support](https://www.reddit.com/r/typescript/comments/hzcpsg/learning_typescript_from_javascript_and_framework/)
- url: https://www.reddit.com/r/typescript/comments/hzcpsg/learning_typescript_from_javascript_and_framework/
---
Hi, I'm just writing here to ask about learning TypeScript. I'm currently in the process of learning Vue.js and want to learn Electron.js. I've understood they have TypeScript support? Does this mean that I would not have to deal with compiling to plain JavaScript myself and that they would handle that for me? 

I also want to know where and how I should go about learning TypeScript coming from someone who knows JavaScript (okay-ish) and is familiar with types (knowing some C#).

Sorry if this is the wrong place to ask.
## [5][Inferring the type of a constrained generic property](https://www.reddit.com/r/typescript/comments/hyxodm/inferring_the_type_of_a_constrained_generic/)
- url: https://www.reddit.com/r/typescript/comments/hyxodm/inferring_the_type_of_a_constrained_generic/
---
Hi,

I'm having an issue while trying to build an abstract DAO class on top of the \`@types/mongodb\` typings, which boils down to this piece of code:

        type ExtractId&lt;T&gt; = T extends { id: infer U } ? U : never;
        function test&lt;T extends { id: string }&gt;(param: T) {
            // error TS2322: Type 'string' is not assignable to type 'ExtractId&lt;T&gt;'.
            const paramId: ExtractId&lt;T&gt; = param.id;
        }

I don't really understand why `string` is not assignable to `ExtractId&lt;T&gt;` given that `T extends { id: string }`.

Are generics constraints (I'm talking about the `&lt;T extends { id: string }&gt;`) not accounted for while inferring the type of a property of this generic ?
## [6][Limit the types passed to a generic](https://www.reddit.com/r/typescript/comments/hyzynt/limit_the_types_passed_to_a_generic/)
- url: https://www.reddit.com/r/typescript/comments/hyzynt/limit_the_types_passed_to_a_generic/
---
I'm trying to create a generic function that can return values of a different type, depending on a parameter and the type used to invoke the generic. The problem is, it fails with this error:

    Type 'First' is not assignable to type 'T'.
      'T' could be instantiated with an arbitrary type which could be unrelated to 'First'.(2322)
 
I understand the concern, for ex. the user of this function might pass a type that is not valid, but what's the proper way to do this kind of thing in Typescript? Is there a way to limit which types can be passed to the generic?

Here's a [playground link](https://www.typescriptlang.org/play/#code/JYOwLgpgTgZghgYwgAgGLCgZzMg3gWAChlkYNsAuZbKUAcwG4iBfI0SWRFAZQgQHsQAEzxESmPoKFUa9JoVaEiMAK4gEYYIOR0IYAGpwANiogAeACoA+ABQ0ZYWiDoBKKhdHFqAd2BgEABbIdo4uniQkCHASyABEcLEUYhEkUHoqUCA6euhYYDYu8ilRMbEARonJKWlgGVm6YLwCwgVFEUIQ8CpGYEleKWABUPzewQDkALbAmJgAhGOFyYqKymoaWvU55PluaNvhyM3YyABuVLnHALx4pNtUY3BjyMxtNXWn8iuEquqa2g1NKQFKiA4QHI44M7IUEia64aiSYT3MpPF7JN6ZD4sIhECGnYwARmQ1wahhM5guYFs8Vii0IR34RggADojPw6DYToTmWQ8i4cfTBMcuUYAEzE7IGYymMww6kVOkMpms9mc4yi5kSZpCfmEIA)

and this is the code:

    interface First {
      first: string;
    }
    interface Second {
      second: string;
    }
    
    function getValue&lt;T&gt;(str: string): T {
      switch (str) {
        case "a":
          return getFirst();
        case "b":
          return getSecond();
        default:
          throw ('miss!');
      }
    }
    
    function getFirst(): First {
      const v: First = { first: 'a' };
      return v;
    }
    
    function getSecond(): Second {
      const v: Second = { second: 'b' };
      return v;
    }
    
    const val1 = getValue&lt;First&gt;("a");
    console.log(val1.first)
    
    const val2 = getValue&lt;Second&gt;("b");
    console.log(val2.second)    



Thanks for your help!
## [7][Dynamically extend a generic class](https://www.reddit.com/r/typescript/comments/hz2k0w/dynamically_extend_a_generic_class/)
- url: https://www.reddit.com/r/typescript/comments/hz2k0w/dynamically_extend_a_generic_class/
---
This is my use case (I have renamed everything, but the relationship is the same).

I have a basic class hierarchy, essentially a `View` class that has several subclasses, and itself inherits from `Shape`.

    abstract class LibraryItem {
      protected dueBy: number | undefined;
    
      abstract pushDueDate(numberOfDays: number): void;
    }

&amp;#x200B;

    abstract class Book extends LibraryItem {
    }

&amp;#x200B;

    class HardcoverBook extends Book {
      pushDueDate(numberOfDays: number): void {
        this.dueDate += numberOfDays;
      }
      
      // ... unique implementations of the other abstract methods
    }

&amp;#x200B;

    class PaperbackBook extends Book {
      pushDueDate(numberOfDays: number): void {
        this.dueDate += numberOfDays % 7; // wtv, something different here
      }
    
      // ... unique implementations of the other abstract methods
    }

&amp;#x200B;

I then have a 3 final classes that exploit the above hierarchy:

    class HighDemandBook extents HardcoverBook {
      public calculation: number;
    
      constructor() {
        super();
        this.calculation = 0;
      }
    
      pushDueDate(numberOfDays: number) { 
        super.pushDueDate(numberOfDays);
            
        calculation = 42;
        
        // custom code here to augment the parent implementation
        this.dueDate = this.dueDate / 2; // access protected props as a subclass
      }
      
      // this is actually all the methods and props in this class
    }

&amp;#x200B;

    class Bookshelf {
      book: Book | undefined; // yes, a bookshelf with only one Book
    
      setBook(book: Book): void {
        this.book = book;
        // ...
      }
    }

&amp;#x200B;

    class Showcase extends Bookshelf {
      constructor() {
        super();
        this.setBook(new HighDemandBook());
        
        // ...
      }
    }

Everything works well. The issue is that I want the `NodeFlashPainter` class to not just extend `RectangleView` but to be able to extend any subclass of `View` dynamically. So ideally I would be able to do:

&amp;#x200B;

    class HighDemandBook&lt;T extends Book&gt; extends T { // HighDemandBook can extend any Book subclass
     // ... everything else the same
    }

&amp;#x200B;

    class Showcase extends Bookshelf {
      constructor() {
        super();
        this.setBook(new HighDemandBook&lt;PaperbackBook&gt;());
        
        // ...
      }
    }

And this would make it so that it effectively translated to:

    class HighDemandBook extends PaperbackBook {
      // ... everything else the same
    }

I really have no idea how to do this. I am open to alternatives also since this seems like an anti-pattern, however it seems pretty much imposed that `HighDemandBook` must be a subclass of at least `Book` since it must access `protected` props. If there is really no other way then I could make the `protected` props public, but that's really undesirable for me.
## [8][Can you add types to custom setState?](https://www.reddit.com/r/typescript/comments/hyywfr/can_you_add_types_to_custom_setstate/)
- url: https://www.reddit.com/r/typescript/comments/hyywfr/can_you_add_types_to_custom_setstate/
---
I struggle with kind of challange. The react's \`useState\` accepts functions which returns prevState. The goal is it to implemented by myself to make a nice way to store something to browser storage.

Everything works for me fine, except type checking of mutating state. See the simplifed example. Do you know how to type this please?  


    interface IState {
      hello: string
    }
    
    type Input = (prevState: IState) =&gt; IState
    
    const setState = (input: Input) =&gt; {
      const state = { hello: "..." }
      return input(state)
    }
    
    // usage
    setState((input) =&gt; ({
      ...input,
      hello: "✅ this is ok",
      notHere: "❌ this should throw error", // 👈 should throw error
    }))
## [9][Select from discriminated union with discriminant](https://www.reddit.com/r/typescript/comments/hz0hmg/select_from_discriminated_union_with_discriminant/)
- url: https://www.reddit.com/r/typescript/comments/hz0hmg/select_from_discriminated_union_with_discriminant/
---
I'm wondering if it's possible to select a type from a discriminated union using its discriminant value?

This is the data structure I'm working with. I'd like to compose the ThemeTypeMap using generics.


    type ThemeColor = {
        type: 'color'
        name: string
        value: string
    }

    type ThemeSize = {
        type: 'size'
        value: number
    }

    type ThemeValue = ThemeColor | ThemeSize

    type ThemeTypeMap = {
        color: ThemeColor
        size: ThemeSize
    }
## [10][Type Assertion on JSON data import?](https://www.reddit.com/r/typescript/comments/hyxcoy/type_assertion_on_json_data_import/)
- url: https://www.reddit.com/r/typescript/comments/hyxcoy/type_assertion_on_json_data_import/
---
Is there a way to annotate or assert a data import like below?

    // data.json
    {
      "x": { "a": "val1", "b": "1", "c": "" },
      "y": { "a": "val2", "b": "2", "c": "" },
      "z": { "a": "val3", "b": "3", "c": "" }
    }

This doesn't work in the import file

    import data as DataShape from './data.json'
    // import data: DataShape from './data.json'
## [11][Where is the best place to learn typescript?](https://www.reddit.com/r/typescript/comments/hym86c/where_is_the_best_place_to_learn_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hym86c/where_is_the_best_place_to_learn_typescript/
---

