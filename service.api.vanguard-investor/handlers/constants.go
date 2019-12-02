package handlers


const IsaRequestJson = `
{
  "steps": [
	{
	  "function": "open",
	  "params": {
		"url": "https://secure.vanguardinvestor.co.uk"
	  }
	},
	{
	  "function": "input",
	  "params": {
		"selector": ".login-section-username input:first-of-type",
		"text": "%s"
	  }
	},
	{
	  "function": "input",
	  "params": {
		"selector": ".login-section-password input:first-of-type",
		"text": "%s"
	  }
	},
	{
	  "function": "submit",
	  "params": {
		"selector": ".panel-login-form button[type=submit]:first-of-type"
	  }
	},
	{
	  "function": "save-text",
	  "params": {
		"selector": ".value span",
		"response_key": "balance"
	  }
	}
  ]
}
`


