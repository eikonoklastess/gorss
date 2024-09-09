# gorss
beautiful rss client for the terminal

spec:
 - vim keybinding
 - keyword search for channels and posts
 - page reader??? html to markdown? 

### database structure
**Channel**
 - id
 - title 
 - link 
 - desc 
 - pub_date 
 - category

 **Channels Image**
 - id
 - channel_id
 - url
 - title
 - link
 - w&h

 **Item**
 - id
 - channel_id
 - title
 - link
 - desc
 - author
 - category
 - comments
 - enclosure??
 - guid
 - pubdate

### Steps
**rss aggregator spec**
 - follow/unfollow channel -> database
 - when app is running fetch feeds
 - fetch new posts -> database
 - every 5 min because not always running

**rss viewer**
 - *CHARM POWERED*
 - channels viewer -> expand
 - posts viewer -> expand
 - singular post viewer -> link to site
 - mark post as read
 - vim bindings

 - post search with keywords???
 - webpage reader in terminal???
 - podcast player in terminal???

**litesql for embedded db**
 - how to setup litesql in project
 - db commmands (sqlc)

### TODO
    [ ] rss aggregator & DB
        [x] fetching feeds
        [x] insert them in db (follow command)
        [x] followed channel viewer
        [x] remove from db (unfollow command)
        [x] store post in db after following
        [x] view post by channel
        [ ] aggregating feeds in db concurently (worker)
        [ ] inserting post in db (after aggregating) (make sure logic schema is good)
    [ ] displaying
        [ ] following channels and following their sub? channels(comments)
        [ ] ergonomic way of displaying channels
        [ ] their posts
        [ ] notification
        [ ] with vim bindings
    [ ] Extra
        [ ] search engine for posts
            [ ]???
        [ ] post website fetcher & displayer
            [ ] gemini/gopher???









