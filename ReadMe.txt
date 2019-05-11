Search Engine:

From a specific URL, below information is taken
1. Page title
2. Page Description
3. HTML body
4. Links present

Storage of the crawled information
- using Elastic Search
-- important points of Elastic Search
    a. highly scalable full text Search Engine
    b. uses inverted index algorithm. Link -  https://www.elastic.co/guide/en/elasticsearch/guide/current/inverted-index.html
    
What all we have in the code:
1. Queue: all the links that are found in the website are sent to the queue, workers will recieve and process all the links that are recieved in the queue
2. Worker pool of size 10. Each of the worker will visit varios website URLs, extract information and update the database. Workers will also send the links that they find and put it in to the queue.
3. Discard invalid links. Wont implement page rank.(try implementing)




What all the files have 

1. Scrape.go 

This is to define the scraper/crawler for each website which has functions which will take the information from the page like page title, Body and description and links. 

there will be 10 workers which will be running on seperate goroutines doing the above steps for 10 different websites at a time.
