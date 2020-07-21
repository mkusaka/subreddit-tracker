# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [3][Looking for a mentor/coding buddies](https://www.reddit.com/r/rails/comments/hv74j4/looking_for_a_mentorcoding_buddies/)
- url: https://www.reddit.com/r/rails/comments/hv74j4/looking_for_a_mentorcoding_buddies/
---
Hi, everyone. Over the past 5 years I've been working as a solo in-house dev for two companies. In the first role I developed a CMS-backed corporate website, and in the second I developed a web app. Both projects are Rails-based and both are up and running in production (but the CMS and app parts are not public).

While I'm comfortable with getting a Rails app to production, there are many aspects of Rails and webdev in general that I'm not familiar with due to being a solo dev.

If any experienced developers would be willing to give me some guidance on how things are done in bigger companies/teams or in enterprise level apps, I'd really appreciate the help. As some examples, I'd like to know more about user authentication (e.g. a Rails API backend with a SPA frontend), error handling and TDD.

I'd also really enjoy chatting with anyone who is in a similar position as me, as I don't really know anyone else who codes with Rails and I know how frustrating and lonely it can be as a solo dev.

A bit more about me: I'm an Australian living in Hong Kong and apart from webdev, I research computational linguistics and NLP.

Feel free to get in touch.
## [4][Do you run commands remotely using Rails?](https://www.reddit.com/r/rails/comments/huzqhf/do_you_run_commands_remotely_using_rails/)
- url: https://www.reddit.com/r/rails/comments/huzqhf/do_you_run_commands_remotely_using_rails/
---
Currently using Rails + Sidekiq + Net-Ssh gem to run commands, wait on responses, and perform additional steps. While this is extremely fast and convenient, it's also unreliable. Over the past several weeks, I've been absolutely exhausted trying to troubleshoot why my sidekiq workers are hanging at random, unpredictable times during the execution of sidekiq workers that may run for minutes at a time. 

Mike has been an absolutely amazing amount of help thus far but I find that I'm only able to make little bits of progress with researching and understanding the "behind the scenes" of threading, "thread-safe", sidekiq, and Rails 5. 

For example, I have a worker that runs every minute and lasts 3 seconds. Usually runs just fine. However, I have another worker that, when a task is executed based on sidekiq's perform_at schedule, it starts a worker that uses the net ssh gem. 

Sometimes this worker may last a few seconds and sometimes it may last 20 minutes (or longer). However, in very random and unreproducible circumstances, it'll cause other sidekiq workers to "back up" if you will. They just queue and I wake up with 700 workers in the queue that haven't processed because of one worker that's using this net ssh gem that may have not finished. 

Essentially this worker uses net ssh to run commands, wait on responses, runs other tasks, and then finishes. Depending on the size of my customers' network, the time just simply varies greatly.

Here are the important key things that I'm trying to accomplish:

- Run commands in parallel and remotely and know when they finish
- Use results of previously run commands to determine next command to run
- Currently using an ssh connection (obviously) and would prefer this channel if possible
- Scalability

Does anyone do anything very similar by chance?
## [5][Using Mocks/Stubs with Rspec](https://www.reddit.com/r/rails/comments/hv4cgn/using_mocksstubs_with_rspec/)
- url: https://www.reddit.com/r/rails/comments/hv4cgn/using_mocksstubs_with_rspec/
---
Hi, I need help with stubing and mocking with rspec.
So here is a very simplified version of my problem. I want to mock Balance class that is instantiated inside the SomethingSmart class.

```ruby
class SomethingSmart
  def run
    account_balance = Balance.new.fetch_balance

    if account_balance &lt; 1000
      # return something
    else
      # return something else
    end
  end
end
```

```ruby
describe SomethingSmart do
  subject { described_class.new.run }
  context 'when balance is under the limit' do
    before do
      # here I want to mock the Balance class and fetch_balance method
    end

    it 'does something' do
      expec(subject).to eq('something')
    end
  end
end
```

So how would I mock this. I tried to google the solution and to be honest I felt kinda lost. It seems that allow_any_instance_of method is deprecated and is not used anymore for some reason. Whatever I copy/paste and try is not working. And I'm not sure how to use doubles here since I'm not using dependency injection and passing Balance class as an argument.
## [6][is there a gem for linked form elements](https://www.reddit.com/r/rails/comments/hv2k96/is_there_a_gem_for_linked_form_elements/)
- url: https://www.reddit.com/r/rails/comments/hv2k96/is_there_a_gem_for_linked_form_elements/
---
before i go and start putting ideas to code, i wanted to see if it already exists. I’d imagine it’s an obscure gem.

i want to build server side rails forms but with inputs that are dependent upon other inputs in the form. for example, if you have two select tags, one for state and one for city, the city select tag should change its list of options depending on what the user selects for the state select tag. the catch is, i want to do this without javascript and with css but without actually writing all the css rules.

I understand that this will significantly increase the css and html on page render but i have some js fatigue and i’m in a legacy angular 1.x app. 

most of the angular was just binding the form elements to one another so that combinations of values were valid. so i started to wonder if i can create ruby-land form helpers.
## [7][eil5 how can I implement JWT with devise?](https://www.reddit.com/r/rails/comments/hv3yn0/eil5_how_can_i_implement_jwt_with_devise/)
- url: https://www.reddit.com/r/rails/comments/hv3yn0/eil5_how_can_i_implement_jwt_with_devise/
---
I searched a lot about this particular subject. I found a lot of blogs, video tutorials, etc. But most of them actually contradict each other and some of them are just about the implementation of the user by yourself. 

I really like devise, as it's a safe and secure way to implement user management system. So, I'm currently writing an API, and JWT authentication is really needed.
## [8][Question on "static" videos.](https://www.reddit.com/r/rails/comments/huxkqd/question_on_static_videos/)
- url: https://www.reddit.com/r/rails/comments/huxkqd/question_on_static_videos/
---
So I'm working on a project (turned out to be quite large based on my rookie skillset) and I sort of hit a big fat "?" on handling video. I know AWS offers nice features for uploading files (videos or photos) into the application and they're sort of the gold standard for handling them based on my current research. However, my Rich Text can already take in an image, and I don't really need such a robust service as AWS to host videos or anything like that. At least not considering this is just a personal project (extension to my portfolio). Wouldn't it be easier to just use embedded code for a YouTube video?
## [9][prefer stimulus reflex over js frameworks](https://www.reddit.com/r/rails/comments/hutotr/prefer_stimulus_reflex_over_js_frameworks/)
- url: https://www.reddit.com/r/rails/comments/hutotr/prefer_stimulus_reflex_over_js_frameworks/
---
For people who worked with ROR, do you prefer Stimulus Reflex over js frameworks like Vuejs or react?
## [10][Morph Modes in Stimulus Reflex](https://www.reddit.com/r/rails/comments/huvrav/morph_modes_in_stimulus_reflex/)
- url: https://www.reddit.com/r/rails/comments/huvrav/morph_modes_in_stimulus_reflex/
---
Stimulus Reflex v3.3.0 is adding different morph modes, and I think it could be a game changer. I put together a short demo if you're interested: [https://youtu.be/VoA86Id3vjg](https://youtu.be/VoA86Id3vjg)
## [11][Rails 6 Production Master.key](https://www.reddit.com/r/rails/comments/huuvym/rails_6_production_masterkey/)
- url: https://www.reddit.com/r/rails/comments/huuvym/rails_6_production_masterkey/
---
I am in the middle of setting up my production environment on my server in order to go live. I have come to something I am not familiar with enough to move forward. When I attempt to set up my `rake db:migrate` for production, I get the following error:

`ArgumentError: Missing \`secret_key_base\` for 'production' environment, set this string with \`rails credentials:edit\``

When I run `EDITOR="vim" bin/rails credentials:edit`

I get the following response: 

`Adding config/master.key to store the encryption key: 111111111111111111111111`

`Save this in a password manager your team can access.`

`If you lose the key, no one, including you, can access anything encrypted with it.`

`create  config/master.key`

`Couldn't decrypt config/credentials.yml.enc. Perhaps you passed the wrong key?`

&amp;#x200B;

I am not sure if I missed something in the development environment or if this something that I need to handle on the server itself.  

I am running Rails 6.0.2, Ruby 2.6.3, Capistrano 3.14, Apache2, Passenger.

I am happy to give any file code that I need to to figure this out. 

Thanks!!
## [12][[HELP] Does having hundreds of connections to hundreds of ActionCable channels make sense?](https://www.reddit.com/r/rails/comments/hus8qk/help_does_having_hundreds_of_connections_to/)
- url: https://www.reddit.com/r/rails/comments/hus8qk/help_does_having_hundreds_of_connections_to/
---
I have a fairly complex data model with numerous models that are all connected to a Person. I have a dashboard that displays, possibly, hundreds of Persons and each row of that dashboard can display information from different models that Person has. For example, it could pull their Car, it could pull their Address through their Home, etc. 

I also have other views where I am interacting with each Person one at a time. I would like to be able to live-reload all of these associated data points when any part of the Person is updated.

The biggest thing holding me back is the performance concern of possible hundreds of users making subscriptions to hundreds of channels each.

How does ActionCable handle multiple subscriptions in a single WebSocket connection?
