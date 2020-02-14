# golang
## [1][Golang file manager for Linux with GTK-GUI](https://www.reddit.com/r/golang/comments/f3dny8/golang_file_manager_for_linux_with_gtkgui/)
- url: https://www.reddit.com/r/golang/comments/f3dny8/golang_file_manager_for_linux_with_gtkgui/
---
Hi. I created for myself some app by golang. I use it. Hope maybe you will like it too.

Or you can just add some stars on github.

[github - GopherFileManager](https://github.com/SilentGopherLnx/GopherFileManager)
## [2][cview - Terminal-based user interface toolkit](https://www.reddit.com/r/golang/comments/f3jky4/cview_terminalbased_user_interface_toolkit/)
- url: https://gitlab.com/tslocum/cview
---

## [3][Micro In Action, Part4：Pub/Sub](https://www.reddit.com/r/golang/comments/f3rqzv/micro_in_action_part4pubsub/)
- url: https://medium.com//micro-in-action-part4-pub-sub-564f3b054ecd?source=friends_link&amp;sk=45892c07676feae7086dae696fbf9a46
---

## [4][Structuring Web Applications in Go](https://www.reddit.com/r/golang/comments/f3rqbu/structuring_web_applications_in_go/)
- url: https://www.calhoun.io/structuring-web-applications-in-go
---

## [5][Joker - A commandline tool helps to check your dev environment](https://www.reddit.com/r/golang/comments/f3pl3u/joker_a_commandline_tool_helps_to_check_your_dev/)
- url: https://github.com/gingerhot/joker
---

## [6][Building Web Server with Go - Part 7](https://www.reddit.com/r/golang/comments/f3mpqm/building_web_server_with_go_part_7/)
- url: https://www.reddit.com/r/golang/comments/f3mpqm/building_web_server_with_go_part_7/
---
https://www.gophersumit.com/series/web-7/
## [7][Micro In Action, Part4：Pub/Sub](https://www.reddit.com/r/golang/comments/f3qx3g/micro_in_action_part4pubsub/)
- url: https://medium.com/@dche423/micro-in-action-part4-pub-sub-564f3b054ecd
---

## [8][XML implementation overhead and alternatives](https://www.reddit.com/r/golang/comments/f3q0lu/xml_implementation_overhead_and_alternatives/)
- url: https://www.reddit.com/r/golang/comments/f3q0lu/xml_implementation_overhead_and_alternatives/
---
Hi,
currently I'm using XML to allow the data exchange between two applications (no services just two separate applications that should be run one after another). There are some now cases where I want to store a map in XML, so I have to dynamically write and read the data for each component. Right now this seems like a lot of boilerplate code, but I'm arguing with myself that this also allows me to validate the fields during parsing. Should I generally never use XML if map-like structures are involved? TOML is not really an alternative, because I also store multiline base64 encoded data. Apart from that I also store XML output from an internal application, which is also why I started using XML in the first place (as wrapper).
## [9][How complicated is local network peer to peer in golang?](https://www.reddit.com/r/golang/comments/f3pkxo/how_complicated_is_local_network_peer_to_peer_in/)
- url: https://www.reddit.com/r/golang/comments/f3pkxo/how_complicated_is_local_network_peer_to_peer_in/
---

## [10][returning exception in postgres and get the exception message in go](https://www.reddit.com/r/golang/comments/f3p0ua/returning_exception_in_postgres_and_get_the/)
- url: https://www.reddit.com/r/golang/comments/f3p0ua/returning_exception_in_postgres_and_get_the/
---
hi, im trying to access the exception message in go, but without any success so far...

what ive got so far in Postgres: 

&gt;CREATE OR REPLACE FUNCTION data.insert\_data(num integer)  
&gt;  
&gt;  RETURNS text AS  
&gt;  
&gt;$BODY$  
&gt;  
&gt;  DECLARE errorMsg TEXT;  
&gt;  
&gt;BEGIN  
&gt;  
&gt;SET search\_path TO Data;  
&gt;  
&gt;COPY tbl FROM 'path' DELIMITER '	' CSV HEADER;  
&gt;  
&gt;errorMsg = 'im\_a\_return\_var';  
&gt;  
&gt;RETURN errorMsg;  
&gt;  
&gt;EXCEPTION  
&gt;  
&gt;   WHEN others THEN  
&gt;  
&gt;RAISE EXCEPTION  
&gt;  
&gt;USING ERRCODE = sqlstate  
&gt;  
&gt;,MESSAGE = 'Error: ' || sqlstate || '/ ' || sqlerrm;  
&gt;  
&gt;errorMsg = sqlstate  
&gt;  
&gt;,MESSAGE = 'Error: ' || sqlstate || '/ ' || sqlerrm  
&gt;  
&gt;RETURN errorMsg;  
&gt;  
&gt;END  
&gt;  
&gt;$BODY$  
&gt;  
&gt;  LANGUAGE plpgsql VOLATILE

on error i get, like i want it the errormessage and the sqlstate.

GoCode accessing the function: 

&gt;db.QueryRow("SELECT data.insert\_data(" + "10" + ")").Scan(&amp;errorMsg)

on success i get im\_a\_return\_var in go, but on error i just get no value at all.

Would anyone have a hint for me? Thanks!
