Risks of giving an agent a code-running tool:

Risk                          | Mitigation
Deletes files / runs rm -rf   | Limit tools — no raw shell access
Leaks secrets from env        | Never pass API keys into tool args
Calls wrong API endpoint      | Validate inputs, allowlist actions
Infinite tool loops           | Set max_iterations or step limit

Main idea: the more power you give an agent, the more damage it can do 
if it makes a mistake. Give it the least access needed — limited tools, 
validated inputs, and step limits.

My menu tools are safe: read-only, fixed data, no shell.