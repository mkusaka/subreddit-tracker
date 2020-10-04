# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][Object is possibly null - after null check?](https://www.reddit.com/r/typescript/comments/j4wwx9/object_is_possibly_null_after_null_check/)
- url: https://www.reddit.com/r/typescript/comments/j4wwx9/object_is_possibly_null_after_null_check/
---
Hi, I have following code, and there's something that I don't quite understand:

in method `setupEvents` as you can see I am check whether `this.drawingContext` exists as it can be null, and even after this check in the same method I am getting `Object is possibly 'null'.ts(2531)
`. Why is it happening? Adding the check directly into event handler does work but I don't really want to do it as there going to be more event listeners added, and that would require repeating same part of code in each of them.

    class Canvas {
    private canvasElement: HTMLCanvasElement;
    private drawingContext: CanvasRenderingContext2D | null;
    private isDrawing: boolean = false;

    constructor(lineWidth: number, lineCap: CanvasLineCap) {
        this.canvasElement = document.getElementById('canvas') as HTMLCanvasElement;
        this.drawingContext = this.canvasElement.getContext('2d');

        if (!this.drawingContext) {
            return;
        }

        this.drawingContext.lineWidth = lineWidth;
        this.drawingContext.lineCap = lineCap;

        this.setupEvents();
    }

    setupEvents() {
        if (!this.drawingContext) {
            return;
        }

        this.canvasElement.addEventListener('mousedown', (e) =&gt; {
            this.drawingContext.beginPath();
            this.drawingContext.lineTo(e.offsetX, e.offsetY);
            this.drawingContext.stroke();

            // webSocket.send(JSON.stringify({ x: e.offsetX, y: e.offsetY }));
            this.isDrawing = true;
        });
    }
}
## [3][Need to create and object, with nested object only containing an interface type : Button. How?](https://www.reddit.com/r/typescript/comments/j4x1i3/need_to_create_and_object_with_nested_object_only/)
- url: https://www.reddit.com/r/typescript/comments/j4x1i3/need_to_create_and_object_with_nested_object_only/
---
The following syntax doesn't work. How can I do it?



Const app Data = {
  buttons: {
      Button1: Button : {

      }

      Button2 :Button : {


      }
   }
}
## [4][explicit module boundaries show up when I export a function](https://www.reddit.com/r/typescript/comments/j4x16g/explicit_module_boundaries_show_up_when_i_export/)
- url: https://www.reddit.com/r/typescript/comments/j4x16g/explicit_module_boundaries_show_up_when_i_export/
---
So I have a function like this :

    export function myFactory() {
    	return function myFn(x : number , y : number):number {
    		return x+y;
    	}
    }
    myFactory();

and eslint is happy and does not warn me about explicit module boundaries for `myFactory`. If I replace `myFn` with a class definition I get the explicit module boundaries warning which disappears when I remove export :

    export function myOtherFactory() {
    	return class myClass {
    		x : number;
    		y : number;
    		constructor() {
    			this.x = 1;
    			this.y =2;
    		}
    	}
    }
    myOtherFactory();

Why ?
## [5][Why libraries and frameworks should not have a composition root ?](https://www.reddit.com/r/typescript/comments/j4cxnh/why_libraries_and_frameworks_should_not_have_a/)
- url: https://www.reddit.com/r/typescript/comments/j4cxnh/why_libraries_and_frameworks_should_not_have_a/
---
So I am reading [this](https://blog.ploeh.dk/2011/07/28/CompositionRoot/) link and I stumble upon this sentence :

&gt;Only applications should have Composition Roots. Libraries and frameworks shouldn't.

Why libraries and frameworks should not have a composition root ?

I am currently coding from a scratch a mobx like library , and I have created a composition root . That is why I am asking .
## [6][Why does { [key: string]: string } allow an unlimited number of k/v pairs?](https://www.reddit.com/r/typescript/comments/j3xzm9/why_does_key_string_string_allow_an_unlimited/)
- url: https://www.reddit.com/r/typescript/comments/j3xzm9/why_does_key_string_string_allow_an_unlimited/
---
I feel like something about my understanding of that syntax isn't complete. On one hand I understand that the generic key and string name means the key could be anything, as can the value. 

On the other hand I do not see any indication that there can be more than one of those pairs.
## [7][Should I use enums or "|" strings?](https://www.reddit.com/r/typescript/comments/j3vp9d/should_i_use_enums_or_strings/)
- url: https://www.reddit.com/r/typescript/comments/j3vp9d/should_i_use_enums_or_strings/
---
Example using "|":

    interface Person {
        name: String,
        age: Number,
        favoriteMeal: 'breakfast' | 'lunch' | 'dinner'
    }
    
    const me: Person = {
        name: "name",
        age: 1,
        favoriteMeal: 'breakfast'
    }

Same example with enums:

    enum Meal {
        breakfast,
        lunch,
        dinner
    }
    
    interface Person {
        name: String,
        age: Number,
        favoriteMeal: Meal
    }
    
    const me: Person = {
        name: "name",
        age: 1,
        favoriteMeal: Meal.breakfast
    }

1. Is there are a hard and fast rule on which one I should use? If I'm not wrong, VScode will autocomplete both methods.
2. If I want to prepare the data for a NoSQL database (Firestore, for example), which one would you recommend? With the enum method, by default I will see integers in the database (0 for breakfast, 1 for lunch, ...), but I can easily change this to strings if I want to.

**Edit**

Based on the other answers, I see that looping through enums is a little (not sure how to describe it)

    enum Meal {
        breakfast, 
        lunch, 
        dinner
    }
    
    for (const meal of Object.values(Meal)) {
        console.log(meal)
    }
    
    [LOG]: "breakfast" 
    [LOG]: "lunch" 
    [LOG]: "dinner" 
    [LOG]: 0 
    [LOG]: 1 
    [LOG]: 2 
    
    for (const meal of Object.values(Meal)) {
        if (isNaN(Number(meal))){
            console.log(meal)
        }
    }
    
    [LOG]: "breakfast" 
    [LOG]: "lunch" 
    [LOG]: "dinner" 

The reason I want to iterate over enums is to create the UI (for checkboxes and radio buttons, for example).

I could use string enums:

    enum Meal {
        breakfast = "breakfast", 
        lunch = "lunch", 
        dinner = "dinner"
    }
    
    for (const meal of Object.values(Meal)) {
        console.log(meal)
    }
    
    [LOG]: "breakfast" 
    [LOG]: "lunch" 
    [LOG]: "dinner"

But I lose some advantages. What if I have to change "dinner" to "supper" across all meals?

While u/nagarian_r's solution looked a little weird at first, I feel it is probably the best solution:

    const Meals = [ 'breakfast', 'lunch', 'dinner'] as const
    type Meal = typeof Meals[number]
    
    console.log(Meals)
    
    [LOG]: ["breakfast", "lunch", "dinner"] 

I can loop through, reference them string, or by their numbers:

    const b1: Meal = "breakfast"
    const b2: Meal = Meals[0]
    
    // b1 === b2 = true

So if you're lopping through Meals to make a checkbox form element, you use the index as the value to save in the database, but the value (Meal\[index\]  
) to show to the user:

    for (let index = 0; index &lt; Meals.length; index++) {   const meal: Meal = Meals[index];   console.log(`${index}) ${meal}`) }  [LOG]: "0) breakfast"  [LOG]: "1) lunch"  [LOG]: "2) dinner"
## [8][Unsure about typecasting functionality.](https://www.reddit.com/r/typescript/comments/j449om/unsure_about_typecasting_functionality/)
- url: https://www.reddit.com/r/typescript/comments/j449om/unsure_about_typecasting_functionality/
---
&amp;#x200B;

    type Day = 'mon' | 'tue' | 'wed' | 'thu' | 'fri' | 'sat' | 'sun';
    let input: string = readlineSync('Enter a day of the week'); // 
    let day: Day = input as Day;

If I input a value of "January", I don't get any errors, and I can output the day and see that its is "January"? Why is that the case? Shouldn't I get some sort of type error?   


Any help would be appreciated.
## [9][I could use some help with properly typing an event emitter function](https://www.reddit.com/r/typescript/comments/j3s87i/i_could_use_some_help_with_properly_typing_an/)
- url: https://www.reddit.com/r/typescript/comments/j3s87i/i_could_use_some_help_with_properly_typing_an/
---
**Answered**

```
interface EventChannel&lt;API extends {}&gt; {
  emit: &lt;key extends keyof API&gt;(eventName: key, payload: API[key]) =&gt; void;
}
```

----

I have a use case for an event emitter like API, but I can't seem to properly type one of the functions.

I have this interface

```
interface EventChannel&lt;API extends {}&gt; {
  emit: (eventName: keyof API, payload: API[typeof eventName]) =&gt; void
}
```

However, that doesn't feel right, because we get into this sort of situation

```
interface ExampleInterface {
  onNumber: number;
  onString: string;
}

type MyInterface = EventChannel&lt;ExampleInterface&gt;;
```
## [10][A library to manipulate Value Objects in Typescript](https://www.reddit.com/r/typescript/comments/j3xg0s/a_library_to_manipulate_value_objects_in/)
- url: https://www.reddit.com/r/typescript/comments/j3xg0s/a_library_to_manipulate_value_objects_in/
---
As you all know, Javascript compares two objects (with \`==\` or \`===\`) by reference and not by value.

This is really annoying when you work with Value Objects when you use DDD in your project.

I was wondering how do you make them work your Typescript/Javascript projects? Do you use a library? Or do you create your ValueObject base class and utils each time yourself?
## [11][Is there an [] equivalent for Array&lt;string | undefined&gt; ?](https://www.reddit.com/r/typescript/comments/j3zyit/is_there_an_equivalent_for_arraystring_undefined/)
- url: https://www.reddit.com/r/typescript/comments/j3zyit/is_there_an_equivalent_for_arraystring_undefined/
---
I think the closest thing is either

* string\[\] | undefined\[\]
* string | undefined\[\]

I don't think either are stipulating that any element can be either, as the title syntax is, though. The 2nd one appears to just be typed to a string or array of undefined
