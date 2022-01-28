### To Mike and Jake

Im not great at coding on the spot but I took some time to think this through and I think i found a good solution here. For our current universe it might be a little over kill but lets say that we had an interdimensional portal gun (rick and morty style). With millions of new elements and thousands of different character lengths for element names my algorithm would be able to process it in O(n) time given enough threads. Thats not really to my credit because Go is doing all the heavy lifting here using the same threading model as the Nginx load balancers that google and facebook use to direct millions of requests per second. 

Anyways I think that when operating at scale this will be a very easy way to maintain fast speeds and I hope you liked my solution. 