# ruby
## [1][date-formatter gem - date formatter by example; auto-builds the strftime format string from an example date](https://www.reddit.com/r/ruby/comments/f4pcn5/dateformatter_gem_date_formatter_by_example/)
- url: https://github.com/feedreader/pluto/tree/master/date-formatter
---

## [2][Linking a Webcam Directly to Rails' ActiveStorage](https://www.reddit.com/r/ruby/comments/f4ds3v/linking_a_webcam_directly_to_rails_activestorage/)
- url: https://emmanuelhayford.com/linking-a-webcam-directly-to-rails-activestorage/
---

## [3][What are the benefits and strengths of using Ruby in 2020?](https://www.reddit.com/r/ruby/comments/f3yuwz/what_are_the_benefits_and_strengths_of_using_ruby/)
- url: https://www.reddit.com/r/ruby/comments/f3yuwz/what_are_the_benefits_and_strengths_of_using_ruby/
---
Hello /r/ruby,

I am a basic data analyst who mostly plays around in R and Excel with some Python sprinkled in for browser automation. I wanted to try a language that was close in syntax to Python but different, as I'm not a fan of the python development environments.

I cam across Ruby and tried the little tutorial in the sidebar and found it fun. My question is: What are the benefits and strengths of using the Ruby language? What can I do with it that I cannot do with Python?

I hope this doesn't come across as adversarial or aggressive, I like the language but I want to use it for what it is good for, and learn it by playing to those strengths.

Thank you

Edit: Thank you for all the responses. It sounds like ruby is definitely worth learning. I bet you'll be seeing me around here relatively soon. :)
## [4][Rails adds ActiveRecord API for switching multiple database connections](https://www.reddit.com/r/ruby/comments/f3s13y/rails_adds_activerecord_api_for_switching/)
- url: https://blog.saeloun.com/2020/02/14/rails-6-multiple-database-support
---

## [5][Help with handling with rest-client exception](https://www.reddit.com/r/ruby/comments/f3wnvq/help_with_handling_with_restclient_exception/)
- url: https://www.reddit.com/r/ruby/comments/f3wnvq/help_with_handling_with_restclient_exception/
---
I need to scrape 10k urls from this website and someone of the urls are out of service (I think...its an error that does not return the json I'm looking for, so rest-client returns 500 Internal Server error in my program)

Error syntax: \`exception\_with\_response': 500 Internal Server Error (RestClient::InternalServerError)

To loop for the urls I'm using a range (1..30).each do |id|. I catenate the url with the current iteration of this range.

response = RestClient.get(url+id)

The problem is some times the url I'm storing in the response variable does not exist/the webpage return some error.

How could I protect my code so I can just pass through this problematic url and keep the scrapping?

Heres my current code:

(I put every code of the loop in a begin/rescue block, but I do not know how do write the code to do such thing)

    require 'nokogiri'
    require 'csv'
    require 'rest-client'
    require 'json'
    link = "https://webfec.org.br/Utils/GetCentrobyId?cod="
    CSV.open('data2.csv', 'ab') do |csv|
    csv &lt;&lt; ['Name', 'Street', 'Info', 'E-mail', 'Site']
        (1..30).each do |id|
            begin
            response = RestClient.get(link+id.to_s)
            json = JSON.parse(response)
            html = json["Data"]
            doc = Nokogiri::HTML.parse(html)
            name = doc.xpath("/html/body/table/tbody/tr[1]").text
            street = doc.xpath("/html/body/table/tbody/tr[2]").text.gsub(Regexp.union(REMOVER), " ")
            info = doc.xpath("/html/body/table/tbody/tr[3]").text.gsub(Regexp.union(REMOVER), " ")
            email = doc.xpath("/html/body/table/tbody/tr[4]").text.gsub(Regexp.union(REMOVER), " ")
            site = doc.xpath("/html/body/table/tbody/tr[5]").text.gsub(Regexp.union(REMOVER), " ")
            csv &lt;&lt; [nome, endereco, complemento, email, site]
            rescue
            end
       end
    end 

&amp;#x200B;
## [6][Nest.js for Ruby](https://www.reddit.com/r/ruby/comments/f3z470/nestjs_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/f3z470/nestjs_for_ruby/
---
Hi all,

Lately I've been using typescript and node.js at work and what i thought would be a nasty experience became a pleasant suprise. I've been reading nest.js documentation ( [https://docs.nestjs.com/](https://docs.nestjs.com/) which is more than great AFAIK) and it makes a lot of sense to me.

I mostly use ruby for my pet projects. I wonder if there's a book\\tutorial\\proper documentation that uses similar thoughts as these of nest.js or even .net core but applied to ruby ?  DI, AOP, modularity etc..
## [7][Ruby Memory Model](https://www.reddit.com/r/ruby/comments/f3r46i/ruby_memory_model/)
- url: https://www.reddit.com/r/ruby/comments/f3r46i/ruby_memory_model/
---
I found it interesting and good documentation on the ruby memory model, by concurrent ruby.  
Sharing for those who forgot to visit.  
[https://docs.google.com/document/d/1pVzU8w\_QF44YzUCCab990Q\_WZOdhpKolCIHaiXG-sPw/](https://docs.google.com/document/d/1pVzU8w_QF44YzUCCab990Q_WZOdhpKolCIHaiXG-sPw/edit#)
## [8][Can't open .rb file on IRB using the ruby command. Getting errors!](https://www.reddit.com/r/ruby/comments/f3o5lx/cant_open_rb_file_on_irb_using_the_ruby_command/)
- url: https://www.reddit.com/r/ruby/comments/f3o5lx/cant_open_rb_file_on_irb_using_the_ruby_command/
---
I've been trying to open files on IRB by using the ruby command (entering something like ruby text.rb to the terminal) but it only gives me errors. I'm trying it with a really basic new file (just one puts "hello") to be sure there are no typos, errors, or hidden characters, but it still doesn't work. I get this error:

ruby text.rb

SyntaxError: (irb):2: syntax error, unexpected tIDENTIFIER, expecting keyword\_do or '{' or '('

ruby text.rb

\^

from /usr/bin/irb:11:in \`&lt;main&gt;'

I made sure to cd into the location where the .rb file is at. No code errors in the file. I have no idea what's going on. It used to work fine months ago, but not anymore. I can load files using the load command, but I'm following some courses that require me to open files using the ruby command to be able to enter command line arguments. 

Does anyone know what's going on?
## [9][New ruby user here with questions about loops](https://www.reddit.com/r/ruby/comments/f3o6l7/new_ruby_user_here_with_questions_about_loops/)
- url: https://www.reddit.com/r/ruby/comments/f3o6l7/new_ruby_user_here_with_questions_about_loops/
---
I was wondering if someone can tutor me a bit on loops in ruby through private messages.
## [10][Sandi Metz Talk in Toronto - Sat Feb 22/2020](https://www.reddit.com/r/ruby/comments/f3c2x0/sandi_metz_talk_in_toronto_sat_feb_222020/)
- url: https://www.eventbrite.com/e/cover-presents-an-evening-with-sandi-metz-registration-94178429217
---

