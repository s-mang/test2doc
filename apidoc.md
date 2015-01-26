FORMAT: 1A
HOST: http://gist.github.com

# Github Gists
The best place to host your code snippits!

# Group Gists
A Gist is a simple way to share snippets and pastes with others. All gists are Git repositories, so they are automatically versioned, forkable and usable from Git.

+ Request 

	+ Body

			{
				"Id": 2,
				"Title": "Some other Gist"
			}

+ Response 200 (application/json)

	+ Body

			{
				"Diffs": [
					{
						"Title": [
							"MyGist",
							"Some other Gist"
						]
					},
					{
						"Id": [
							1,
							2
						]
					}
				]
			}

+ Request 

+ Response 400 (application/json)

	+ Body

			{
				"Status": "Bad Request",
				"Reason": "No request body"
			}

+ Request 

+ Response 200 (application/json)

	+ Body

			{
				"Id": 1,
				"Title": "MyGist"
			}

+ Request 

+ Response 200 (application/json)

	+ Body

			[
				{
					"Id": 1,
					"Title": "MyGist"
				},
				{
					"Id": 2,
					"Title": "Some other Gist"
				}
			]

# Group Repos
A Repo is an on-disk data structure which stores metadata for a set of files and/or directory structure.

+ Request 

+ Response 200 (application/json)

	+ Body

			{
				"Id": 1,
				"Name": "my_repo",
				"Owner": "adams-sarah"
			}

+ Request 

+ Response 200 (application/json)

	+ Body

			[
				{
					"Id": 1,
					"Name": "my_repo",
					"Owner": "adams-sarah"
				},
				{
					"Id": 2,
					"Name": "my_other_repo",
					"Owner": "adams-sarah"
				},
				{
					"Id": 3,
					"Name": "someone_elses_repo",
					"Owner": "goody-twoshoes"
				}
			]

