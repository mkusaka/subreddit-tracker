# reactjs
## [1][Beginner's Thread / Easy Questions (October 2020)](https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/)
- url: https://www.reddit.com/r/reactjs/comments/j31umf/beginners_thread_easy_questions_october_2020/
---
&gt; Previous Beginner's Threads can be found in the [wiki][wiki previous threads].

Ask about React or anything else in its ecosystem :)

Stuck making progress on your app?  
Still Ask away! We’re a friendly bunch.

No question is too simple. 🙂

---

## Want Help with your Code?

1. **Improve your chances** of reply by
   1. adding minimal example with [JSFiddle][jsfiddle], [CodeSandbox][code sandbox], or [Stackblitz][stackblitz] links
   1. describing what you want it to do (ask yourself if it's an [XY problem](https://meta.stackexchange.com/questions/66377/what-is-the-xy-problem))
   1. things you've tried. (Don't just post big blocks of code!)
1. **[Formatting Code][wiki formatting code]** wiki shows how to format code in this thread.
1. **Pay it forward!** Answer questions even if there is already an answer. Other perspectives can be helpful to beginners. Also, there's no quicker way to learn than [being wrong on the Internet][being wrong on the internet].

---

## New to React?

### Check out the sub's **sidebar**! 👉

🆓 Here are great, **free** resources!

- [Read the **official** Getting Started page][official getting started page] on the docs.
- [Microsoft Frontend Bootcamp][microsoft frontend bootcamp]
- [Codecademy's React courses][codecademy's react courses]
- [Scrimba's React Course][scrimba's react course]
- [FreeCodeCamp's React course][freecodecamp's react course]
- [Kent Dodd's Egghead.io course][kent dodd's egghead.io course]
- New to Hooks? Check out [Amelia Wattenberger's Thinking in React Hooks][thinking in react hooks]
- and these React Hook recipes on [useHooks.com][usehooks.com] by [Gabe Ragland](https://twitter.com/gabe_ragland)
- What other updated resources do you suggest?

Any ideas/suggestions to improve this thread - feel free to comment here!

_Finally, thank you to all who post questions and those who answer them. We're a growing community and [helping each other][learn by teaching] only [strengthens it!][learn in public]_

---

[usehooks.com]: https://usehooks.com/
[thinking in react hooks]: https://wattenberger.com/blog/react-hooks
[freecodecamp's react course]: https://www.freecodecamp.org/news/learn-react-course/
[microsoft frontend bootcamp]: https://www.reddit.com/r/reactjs/comments/auu02f/microsoft_has_open_sourced_their_frontend/
[official getting started page]: https://reactjs.org/docs/getting-started.html
[/u/acemarke]: https://www.reddit.com/u/acemarke
[suggested resources for learning react]: http://blog.isquaredsoftware.com/2017/12/blogged-answers-learn-react/
[kent dodd's egghead.io course]: http://kcd.im/beginner-react
[codecademy's react courses]: https://www.codecademy.com/catalog/language/javascript
[scrimba's react course]: https://scrimba.com/g/glearnreact
[wiki formatting code]: https://www.reddit.com/r/reactjs/wiki/index#wiki_formatting_code
[wiki previous threads]: https://www.reddit.com/r/reactjs/wiki/index#wiki_previous_threads
[code sandbox]: https://codesandbox.io/s/new
[jsfiddle]: https://jsfiddle.net/Luktwrdm/
[stackblitz]: https://stackblitz.com/
[being wrong on the internet]: https://xkcd.com/386/
[tweet organization]: https://twitter.com/dan_abramov/status/1027245759232651270?lang=en
[get started with redux]: https://www.reddit.com/r/reactjs/wiki/index#wiki_getting_started_with_redux
[learn by teaching]: https://en.wikipedia.org/wiki/Learning_by_teaching
[learn in public]: https://www.swyx.io/writing/learn-in-public/
## [2][Help the React team write new docs—take the 2020 React Community Survey](https://www.reddit.com/r/reactjs/comments/j5iqj3/help_the_react_team_write_new_docstake_the_2020/)
- url: https://www.surveymonkey.co.uk/r/T58DPNS
---

## [3][Supercharge testing React applications with Wallaby.js](https://www.reddit.com/r/reactjs/comments/jcpnnk/supercharge_testing_react_applications_with/)
- url: https://www.smashingmagazine.com/2020/10/supercharge-testing-react-applications-wallabyjs/
---

## [4][Just a cool dropdown animation that I recently completed.](https://www.reddit.com/r/reactjs/comments/jc515n/just_a_cool_dropdown_animation_that_i_recently/)
- url: https://v.redd.it/ccg4jukklet51
---

## [5][Made a colour mixing tool](https://www.reddit.com/r/reactjs/comments/jcs0ml/made_a_colour_mixing_tool/)
- url: https://jackrhumphrey.github.io/colour
---

## [6][How to start learn react js](https://www.reddit.com/r/reactjs/comments/jcuayx/how_to_start_learn_react_js/)
- url: https://www.reddit.com/r/reactjs/comments/jcuayx/how_to_start_learn_react_js/
---
Hi, recently I started to learn python and I really enjoy that. I have an idea how to simplify my job, but to do that i need frontend application. My friend told me to give up at flask and django and learn react js. The problem is that i never learned basic javascrip and I dont know how to start. I found this [course](https://www.udemy.com/course/react-the-complete-guide-incl-redux/) on udemy, can I start with this? I will appreciate every hint. Thanks.
## [7][React-map-gl // How to filter with buttons and display markers](https://www.reddit.com/r/reactjs/comments/jcri3l/reactmapgl_how_to_filter_with_buttons_and_display/)
- url: https://www.reddit.com/r/reactjs/comments/jcri3l/reactmapgl_how_to_filter_with_buttons_and_display/
---
 

Hello everyone!

React   newbie here. I'm seeking your help because I am currently struggling   with a feature. I managed to display a list of markers from an external   file (list of 3 museums in London) and added popups to them. But now  I'd  like to filter them with 3 buttons: let's say button1 displays  marker  1, etc.

I have no clue on  how to  perform such thing: my first thought would be to add some  specific  characteristic to each of them in my Data.js  
file but then I  do not know how to retrieve this info and based my  buttons on this  filter in my main component, could you please tell me  how to do it? Any  hints?

Here is what I've done so far:

&amp;#x200B;

 

    // Main document
    import React, {PureComponent} from "react";
    import ReactDOM from "react-dom";
    import ReactMapGL, {Marker, Popup} from "react-map-gl";
    import "./index.css";
    import Data from "./Data.js";
    
    
    const MAPBOX_TOKEN = "Mapbox_token";
    
    
    class Mapp extends PureComponent {
        constructor(props) {
            super(props);
            this.state = {
                viewport: {
                    latitude: 51.50985,
                    longitude: -0.11892,
                    zoom: 12,
                },
                markers: Data,
                selectedMarker: null
            };
        }
    
        _renderMarkers() {
            return (
                this.state.markers.map(marker =&gt; (
                    &lt;Marker
                        key={marker.id}
                        longitude={marker.longitude}
                        latitude={marker.latitude}&gt;
                            &lt;button 
                                className="marker-btn"
                                onClick={event =&gt; {
                                    event.preventDefault();
                                    this.setState({selectedMarker: marker})
                                }}&gt;
                                &lt;img 
                                    className="marker-btn"
                                    src="location-icon.svg"
                                    alt="location Icon"
                                /&gt;
                            &lt;/button&gt;
                    &lt;/Marker&gt;
                ))
            )
        }
    
        _renderPopup() {
            return(
                this.state.selectedMarker &amp;&amp; (
                    &lt;Popup
                        longitude={this.state.selectedMarker.longitude}
                        latitude={this.state.selectedMarker.latitude}
                        onClose={() =&gt; {this.setState({selectedMarker: null})}}
                        offsetLeft={22.5}
                    &gt;
                        &lt;div&gt;
                            &lt;h3&gt;Name: {this.state.selectedMarker.site_name}&lt;/h3&gt;
                            &lt;h4&gt;Website: {this.state.selectedMarker.website}&lt;/h4&gt;
                        &lt;/div&gt;
                    &lt;/Popup&gt;
                )
            )
        }
    
        render() {
            return (
                &lt;div className="App"&gt;
                    &lt;div className="Map"&gt;
                        &lt;ReactMapGL
                            {...this.state.viewport}
                                width="100vw"
                                height="100vh"
                                mapStyle="mapbox://styles/mapbox/streets-v11"
                                onViewportChange={viewport =&gt; this.setState({viewport})}
                                mapboxApiAccessToken={MAPBOX_TOKEN}
                            &gt;
    
                                {this._renderMarkers()}
                                {this._renderPopup()}
    
                        &lt;/ReactMapGL&gt;
                    &lt;/div&gt;
                &lt;/div&gt;
            );
        }
    }
    
    document.body.style.margin = 0;
    ReactDOM.render(
        &lt;Mapp /&gt;, 
        document.getElementById("root")
    )

 The file containing the information about my markers: 

 

    //Data.js
    const Data = [
      {
        "id": 1,
        "latitude": 51.5206436238135,
        "longitude": -0.12869802265875,
        "site_name": "Museum of Writing",
        "website": "http://www.museumofwriting.co.uk/"
      },
      {
        "id": 2,
        "latitude": 51.5252701073731,
        "longitude": -0.1216462226263,
        "site_name": "The Foundling Museum",
        "website": "http://www.foundlingmuseum.org.uk/"
      },
      {
        "id": 3,
        "latitude": 51.5484218889309,
        "longitude": -0.17724635343509,
        "site_name": "Freud Museum",
        "website": "https://www.freud.org.uk/"
      }
    ]
    
    export default Data

&amp;#x200B;

Thak you in advance! Cheeeeeers
## [8][Pins elements w/ Framer Motion](https://www.reddit.com/r/reactjs/comments/jcremw/pins_elements_w_framer_motion/)
- url: https://www.reddit.com/r/reactjs/comments/jcremw/pins_elements_w_framer_motion/
---
Hi everybody! Someone have an idea how to pins elements with Framer Motion?

Thank you in advance!
## [9][Has anybody learned anything from the React Summit?](https://www.reddit.com/r/reactjs/comments/jcjqyr/has_anybody_learned_anything_from_the_react_summit/)
- url: https://www.reddit.com/r/reactjs/comments/jcjqyr/has_anybody_learned_anything_from_the_react_summit/
---
I didn't have a chance to look back at the summit videos from today and yesterday but I only have the free base track anyway so just wondering if anyone attended the summit track and has anything interesting to share?
## [10][A livestream where I go through building a spa, answer react questions, start building a store, and a blog. All from scratch in modern reactjs. More to come!](https://www.reddit.com/r/reactjs/comments/jcbvd5/a_livestream_where_i_go_through_building_a_spa/)
- url: https://youtu.be/_1YojDFZEfU
---

## [11][Anyone serving React with Kubernetes?](https://www.reddit.com/r/reactjs/comments/jcq38k/anyone_serving_react_with_kubernetes/)
- url: https://www.reddit.com/r/reactjs/comments/jcq38k/anyone_serving_react_with_kubernetes/
---
Hello all!  


I have made my initial deployment in Kubernetes and I was wondering if someone has done the same. I ask because I am wondering what type of livenessProbe/readinessProbe I can place for the NGINX pod serving the static files.  


Thank you in advance and regards.
## [12][Need help with contact manager app](https://www.reddit.com/r/reactjs/comments/jcsbzv/need_help_with_contact_manager_app/)
- url: https://www.reddit.com/r/reactjs/comments/jcsbzv/need_help_with_contact_manager_app/
---
Here's a simple contact manager app I made with React. The problem is that it only stores one record at a time so when I create a new record, the old one disappears. Does anyone know why? Here's the link:

[https://stackblitz.com/edit/react-wbyplw?file=src/store.js](https://stackblitz.com/edit/react-wbyplw?file=src/store.js)
