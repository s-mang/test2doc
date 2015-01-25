FORMAT: 1A
HOST: http://www.google.com

# Google Search
The only search engine.

+ Request 

+ Response 200 (application/json)

	+ Body

			{
				"Message": "Fancy that!"
			}

+ Request 

	+ Body

			{
				"Foo": "Bars",
				"Something": "Funny"
			}

+ Response 200 (application/json)

	+ Body

			{
				"Message": "Fancy that!"
			}

