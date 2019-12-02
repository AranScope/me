package handlers


const MarcusBalanceRequestJson = `
{
  "steps": [
	{
	  "function": "open",
	  "params": {
		"url": "https://www.marcus.co.uk/uk/en/login"
	  }
	},
	{
	  "function": "input",
	  "params": {
		"selector": "input[type=email]:first-of-type",
		"text": "%s"
	  }
	},
	{
	  "function": "input",
	  "params": {
		"selector": "input[type=email]:first-of-type",
		"text": "%s"
	  }
	},
	{
	  "function": "submit",
	  "params": {
		"selector": ".PasswordEntryModule button[type=submit]:first-of-type"
	  }
	},
	{
	  "function": "save-text",
	  "params": {
		"selector": ".SavingsAccount__total-balance",
		"response_key": "balance"
	  }
	}
  ]
}
`


